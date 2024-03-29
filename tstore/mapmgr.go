package tstore

import (
	"context"
	"fmt"
	"github.com/tristan-club/kit/tstore/pb"
	"reflect"
)

type MapManger struct {
	key   string
	err   error
	value *pb.IValue
}

func GetMapMgr() *MapManger {
	return &MapManger{}
}

func (m *MapManger) Save(uid, path string, key string, value interface{}) error {

	if key != "" {
		path = fmt.Sprintf("%s.%s", path, key)
	}

	_, err := save(&pb.SaveParam{
		Uid:    uid,
		Path:   path,
		IValue: pb.NewIValue(value),
	})
	return err
}

func (m *MapManger) BatchSave(uid, path string, param map[string]interface{}) error {
	if len(param) == 0 {
		return nil
	}

	return m.Save(uid, path, "", param)
}

func (m *MapManger) Keys(uid, path string) (ks []string, err error) {
	cli := pb.NewTStoreServiceClient(conn)
	if conn == nil {
		return nil, fmt.Errorf("tstore svc not init")
	}
	var resp *pb.FetchResp
	resp, err = cli.Keys(context.Background(), &pb.FetchParam{Uid: uid, Path: path})
	if err != nil {
		return
	} else if resp.Code == 404 {
		return
	} else if resp.Code != CodeSuccess {
		return nil, fmt.Errorf("fetch keys error: %s", resp.Msg)
	}
	return resp.Keys, nil
}

func (m *MapManger) Fetch(uid, path string, key string) *MapManger {

	if key != "" {
		path = fmt.Sprintf("%s.%s", path, key)
		m.key = key
	}

	v, err := fetch(uid, path)
	if err != nil {
		m.err = err
	} else if v.Code == 404 {
	} else if v.Code != CodeSuccess {
		msg := v.Msg
		if msg == "" {
			msg = fmt.Sprintf("unknown fetch error, code: %d", v.Code)
		}
		m.err = fmt.Errorf(msg)
	} else {
		m.value = v.IValue
	}

	return m
}

func (m *MapManger) Scan(dest interface{}) *MapManger {

	if m.err != nil {
		return m
	} else if m.value == nil {
		return m
	}

	if dest == nil || (reflect.ValueOf(dest).Kind() == reflect.Ptr && reflect.ValueOf(dest).IsNil()) {
		m.err = fmt.Errorf("invalid scan dest")
		return m
	}

	switch dest.(type) {
	case *map[string]string:
		dm, _ := dest.(*map[string]string)
		dmv := *dm
		for k, v := range m.value.MapValue {
			dmv[k] = v.GetStrValue()
		}
	default:
		m.err = fmt.Errorf("scan dest not support")
	}
	return m
}

func (m *MapManger) GetStrValue() (string, error) {
	if m.key == "" {
		return "", fmt.Errorf("path not set")
	}
	if m.value == nil {
		return "", nil
	}
	return m.value.GetStrValue(), nil
}

func (m *MapManger) Error() error {
	return m.err
}
