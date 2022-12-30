package pb

import (
	"github.com/golang/protobuf/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/anypb"
)

func NewStrValue(str string) *IValue {
	return &IValue{
		StrValue: str,
		Itype:    IValue_str,
	}
}

func NewIntValue(i int32) *IValue {
	return &IValue{
		IntValue: i,
		Itype:    IValue_int,
	}
}

func NewMapValue(m map[string]interface{}) *IValue {
	v := map[string]*IValue{}

	for key, value := range m {
		v[key] = NewIValue(value)
	}

	return &IValue{
		MapValue: v,
		Itype:    IValue_map,
	}
}

func NewArrValue(m []interface{}) *IValue {
	var v []*IValue

	for _, value := range m {
		if iv, ok := value.(*IValue); ok {
			v = append(v, iv)
		} else {
			v = append(v, NewIValue(value))
		}
	}

	return &IValue{
		ArrValue: v,
		Itype:    IValue_arr,
	}
}

func NewAnyValue(m proto.Message, typeName string) *IValue {
	bytes, err := proto.Marshal(m)
	if err != nil {
		log.Error().Msgf("NewAnyValue error, %s", err)
		typeName = "error"
		bytes = []byte{}
	}

	return &IValue{
		AnyValue: &anypb.Any{
			TypeUrl: typeName,
			Value:   bytes,
		},
		Itype: IValue_any,
	}
}

func NewIValue(value interface{}) *IValue {
	if value == nil {
		return &IValue{Itype: IValue_nil}
	}

	switch value.(type) {
	case int:
		return NewIntValue(int32(value.(int)))
	case int32:
		return NewIntValue(value.(int32))
	case string:
		return NewStrValue(value.(string))
	case map[string]interface{}:
		return NewMapValue(value.(map[string]interface{}))
	case proto.Message:
		return NewAnyValue(value.(proto.Message), "")
	default:
		log.Error().Msgf("new ivalue, unsupported type, value = %v", value)
	}

	return &IValue{
		Itype: IValue_nil,
	}
}

func (x *IValue) MapSet(key string, v *IValue) {
	if x == nil {
		return
	}
	if x.Itype != IValue_map {
		log.Error().Msgf("MapSet error, self is not a map value, key = %s, value = %v", key, v)
		return
	}
	x.MapValue[key] = v
}

func (x *IValue) MapDelete(key string) {
	if x == nil {
		return
	}
	if x.Itype != IValue_map {
		log.Error().Msgf("MapDelete error, self is not a map value, key = %s, value = %v", key)
		return
	}
	delete(x.MapValue, key)
}

func (x *IValue) MapGet(key string) (*IValue, bool) {
	if x == nil {
		return nil, false
	}
	if x.Itype != IValue_map {
		log.Error().Msgf("MapSet error, self is not a map value")
		return &IValue{Itype: IValue_nil}, false
	}

	v, ok := x.MapValue[key]

	return v, ok
}

func (x *IValue) ArrPush(v *IValue) int {
	if x == nil {
		return -1
	}

	if x.Itype != IValue_arr {
		log.Error().Msgf("ArrPush error, self is not a array value")
		return -1
	}

	x.ArrValue = append(x.ArrValue, v)
	return len(x.ArrValue)
}
