package tstore

import (
	"context"
	"errors"
	"fmt"
	"github.com/tristan-club/kit/grpc/client"
	"github.com/tristan-club/kit/tstore/pb"
	"google.golang.org/grpc"
	"time"
)

var conn *grpc.ClientConn

func InitTStore(svc string) (err error) {
	conn, err = client.GetDefault(svc)
	return err
}

func PBSave(uid string, path string, value interface{}) error {

	_, err := save(&pb.SaveParam{
		Uid:    uid,
		Path:   path,
		IValue: pb.NewIValue(value),
	})
	return err
}

func PBSaveStr(uid string, path string, value string) error {

	_, err := save(&pb.SaveParam{
		Uid:    uid,
		Path:   path,
		IValue: pb.NewIValue(value),
	})
	return err
}

func PBGetStr(uid string, path string) (string, error) {
	v, err := fetch(uid, path)
	if err != nil {
		return "", err
	}
	if v.Code == 404 {
		return "", nil
	}
	if v.IValue.Itype != pb.IValue_str || v.Code != CodeSuccess {
		return "", errors.New("pb get error")
	}
	return v.IValue.StrValue, nil
}

func PBGetInt(uid string, path string) (int64, error) {
	v, err := fetch(uid, path)
	if err != nil {
		return 0, err
	}
	if v.Code == 404 {
		return 0, nil
	}
	if v.IValue.Itype != pb.IValue_int || v.Code != CodeSuccess {
		return 0, errors.New("pb get error")
	}
	return int64(v.IValue.IntValue), nil
}

func PBGet(uid string, path string) ([]byte, error) {
	v, err := fetch(uid, path)
	if err != nil {
		return nil, err
	}
	if v.Code == 404 {
		return nil, nil
	}

	if v.IValue.Itype != pb.IValue_any || v.Code != CodeSuccess {

		return []byte{}, errors.New("pb get error")
	}

	return v.IValue.AnyValue.Value, nil
}

func save(v *pb.SaveParam) (*pb.SaveResp, error) {
	// 设定请求超时时间 3s
	if conn == nil {
		return nil, fmt.Errorf("tstore svc not init")
	}
	cli := pb.NewTStoreServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return cli.Save(ctx, v)
}

func fetch(uid string, path string) (*pb.FetchResp, error) {
	cli := pb.NewTStoreServiceClient(conn)
	if conn == nil {
		return nil, fmt.Errorf("tstore svc not init")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return cli.Fetch(ctx, &pb.FetchParam{
		Uid:  uid,
		Path: path,
	})
}
