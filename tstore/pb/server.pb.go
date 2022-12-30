// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: server.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IValue_IType int32

const (
	IValue_str IValue_IType = 0
	IValue_int IValue_IType = 1
	IValue_map IValue_IType = 2
	IValue_nil IValue_IType = 3
	IValue_any IValue_IType = 4
	IValue_arr IValue_IType = 5
)

// Enum value maps for IValue_IType.
var (
	IValue_IType_name = map[int32]string{
		0: "str",
		1: "int",
		2: "map",
		3: "nil",
		4: "any",
		5: "arr",
	}
	IValue_IType_value = map[string]int32{
		"str": 0,
		"int": 1,
		"map": 2,
		"nil": 3,
		"any": 4,
		"arr": 5,
	}
)

func (x IValue_IType) Enum() *IValue_IType {
	p := new(IValue_IType)
	*p = x
	return p
}

func (x IValue_IType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IValue_IType) Descriptor() protoreflect.EnumDescriptor {
	return file_server_proto_enumTypes[0].Descriptor()
}

func (IValue_IType) Type() protoreflect.EnumType {
	return &file_server_proto_enumTypes[0]
}

func (x IValue_IType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IValue_IType.Descriptor instead.
func (IValue_IType) EnumDescriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0, 0}
}

type IValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StrValue string             `protobuf:"bytes,1,opt,name=str_value,json=strValue,proto3" json:"str_value,omitempty"`
	IntValue int32              `protobuf:"varint,2,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	MapValue map[string]*IValue `protobuf:"bytes,3,rep,name=map_value,json=mapValue,proto3" json:"map_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AnyValue *anypb.Any         `protobuf:"bytes,4,opt,name=any_value,json=anyValue,proto3" json:"any_value,omitempty"`
	Itype    IValue_IType       `protobuf:"varint,5,opt,name=itype,proto3,enum=pb.IValue_IType" json:"itype,omitempty"`
	ArrValue []*IValue          `protobuf:"bytes,6,rep,name=arr_value,json=arrValue,proto3" json:"arr_value,omitempty"`
}

func (x *IValue) Reset() {
	*x = IValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IValue) ProtoMessage() {}

func (x *IValue) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IValue.ProtoReflect.Descriptor instead.
func (*IValue) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *IValue) GetStrValue() string {
	if x != nil {
		return x.StrValue
	}
	return ""
}

func (x *IValue) GetIntValue() int32 {
	if x != nil {
		return x.IntValue
	}
	return 0
}

func (x *IValue) GetMapValue() map[string]*IValue {
	if x != nil {
		return x.MapValue
	}
	return nil
}

func (x *IValue) GetAnyValue() *anypb.Any {
	if x != nil {
		return x.AnyValue
	}
	return nil
}

func (x *IValue) GetItype() IValue_IType {
	if x != nil {
		return x.Itype
	}
	return IValue_str
}

func (x *IValue) GetArrValue() []*IValue {
	if x != nil {
		return x.ArrValue
	}
	return nil
}

type FetchParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FetchParam) Reset() {
	*x = FetchParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchParam) ProtoMessage() {}

func (x *FetchParam) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchParam.ProtoReflect.Descriptor instead.
func (*FetchParam) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *FetchParam) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *FetchParam) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type FetchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg    string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	IValue *IValue  `protobuf:"bytes,6,opt,name=iValue,proto3" json:"iValue,omitempty"`
	Keys   []string `protobuf:"bytes,7,rep,name=keys,proto3" json:"keys,omitempty"`
	ArrLen int32    `protobuf:"varint,8,opt,name=arrLen,proto3" json:"arrLen,omitempty"`
}

func (x *FetchResp) Reset() {
	*x = FetchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResp) ProtoMessage() {}

func (x *FetchResp) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResp.ProtoReflect.Descriptor instead.
func (*FetchResp) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (x *FetchResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FetchResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *FetchResp) GetIValue() *IValue {
	if x != nil {
		return x.IValue
	}
	return nil
}

func (x *FetchResp) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *FetchResp) GetArrLen() int32 {
	if x != nil {
		return x.ArrLen
	}
	return 0
}

type SaveParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string  `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Path   string  `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	IValue *IValue `protobuf:"bytes,6,opt,name=iValue,proto3" json:"iValue,omitempty"`
}

func (x *SaveParam) Reset() {
	*x = SaveParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveParam) ProtoMessage() {}

func (x *SaveParam) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveParam.ProtoReflect.Descriptor instead.
func (*SaveParam) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

func (x *SaveParam) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *SaveParam) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SaveParam) GetIValue() *IValue {
	if x != nil {
		return x.IValue
	}
	return nil
}

type SaveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	ArrLen int32  `protobuf:"varint,3,opt,name=arrLen,proto3" json:"arrLen,omitempty"`
}

func (x *SaveResp) Reset() {
	*x = SaveResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResp) ProtoMessage() {}

func (x *SaveResp) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResp.ProtoReflect.Descriptor instead.
func (*SaveResp) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{4}
}

func (x *SaveResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SaveResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SaveResp) GetArrLen() int32 {
	if x != nil {
		return x.ArrLen
	}
	return 0
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x03,
	0x0a, 0x06, 0x49, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x2e, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x6d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x61, 0x6e, 0x79,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x08, 0x61, 0x6e, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a, 0x05,
	0x69, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x49, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x49, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x69,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x61, 0x72, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x08, 0x61, 0x72, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x47, 0x0a,
	0x0d, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x05, 0x49, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x07, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x6e, 0x69,
	0x6c, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x61, 0x6e, 0x79, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03,
	0x61, 0x72, 0x72, 0x10, 0x05, 0x22, 0x32, 0x0a, 0x0a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x81, 0x01, 0x0a, 0x09, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x22, 0x0a,
	0x06, 0x69, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x49, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x69, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x72, 0x4c, 0x65, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x72, 0x72, 0x4c, 0x65, 0x6e, 0x22, 0x55, 0x0a,
	0x09, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x22, 0x0a, 0x06, 0x69, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x69, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x48, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x72, 0x4c, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x72, 0x72, 0x4c, 0x65, 0x6e, 0x32, 0x85,
	0x02, 0x0a, 0x0d, 0x54, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x25, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x29, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0d, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x04,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0c, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x05,
	0x53, 0x68, 0x69, 0x66, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x62, 0x2f, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_server_proto_goTypes = []interface{}{
	(IValue_IType)(0),  // 0: pb.IValue.IType
	(*IValue)(nil),     // 1: pb.IValue
	(*FetchParam)(nil), // 2: pb.FetchParam
	(*FetchResp)(nil),  // 3: pb.FetchResp
	(*SaveParam)(nil),  // 4: pb.SaveParam
	(*SaveResp)(nil),   // 5: pb.SaveResp
	nil,                // 6: pb.IValue.MapValueEntry
	(*anypb.Any)(nil),  // 7: google.protobuf.Any
}
var file_server_proto_depIdxs = []int32{
	6,  // 0: pb.IValue.map_value:type_name -> pb.IValue.MapValueEntry
	7,  // 1: pb.IValue.any_value:type_name -> google.protobuf.Any
	0,  // 2: pb.IValue.itype:type_name -> pb.IValue.IType
	1,  // 3: pb.IValue.arr_value:type_name -> pb.IValue
	1,  // 4: pb.FetchResp.iValue:type_name -> pb.IValue
	1,  // 5: pb.SaveParam.iValue:type_name -> pb.IValue
	1,  // 6: pb.IValue.MapValueEntry.value:type_name -> pb.IValue
	4,  // 7: pb.TStoreService.Save:input_type -> pb.SaveParam
	2,  // 8: pb.TStoreService.Fetch:input_type -> pb.FetchParam
	2,  // 9: pb.TStoreService.Delete:input_type -> pb.FetchParam
	2,  // 10: pb.TStoreService.Keys:input_type -> pb.FetchParam
	4,  // 11: pb.TStoreService.Push:input_type -> pb.SaveParam
	2,  // 12: pb.TStoreService.Shift:input_type -> pb.FetchParam
	5,  // 13: pb.TStoreService.Save:output_type -> pb.SaveResp
	3,  // 14: pb.TStoreService.Fetch:output_type -> pb.FetchResp
	3,  // 15: pb.TStoreService.Delete:output_type -> pb.FetchResp
	3,  // 16: pb.TStoreService.Keys:output_type -> pb.FetchResp
	5,  // 17: pb.TStoreService.Push:output_type -> pb.SaveResp
	3,  // 18: pb.TStoreService.Shift:output_type -> pb.FetchResp
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		EnumInfos:         file_server_proto_enumTypes,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TStoreServiceClient is the client API for TStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TStoreServiceClient interface {
	Save(ctx context.Context, in *SaveParam, opts ...grpc.CallOption) (*SaveResp, error)
	Fetch(ctx context.Context, in *FetchParam, opts ...grpc.CallOption) (*FetchResp, error)
	Delete(ctx context.Context, in *FetchParam, opts ...grpc.CallOption) (*FetchResp, error)
	Keys(ctx context.Context, in *FetchParam, opts ...grpc.CallOption) (*FetchResp, error)
	Push(ctx context.Context, in *SaveParam, opts ...grpc.CallOption) (*SaveResp, error)
	Shift(ctx context.Context, in *FetchParam, opts ...grpc.CallOption) (*FetchResp, error)
}

type tStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTStoreServiceClient(cc grpc.ClientConnInterface) TStoreServiceClient {
	return &tStoreServiceClient{cc}
}

func (c *tStoreServiceClient) Save(ctx context.Context, in *SaveParam, opts ...grpc.CallOption) (*SaveResp, error) {
	out := new(SaveResp)
	err := c.cc.Invoke(ctx, "/pb.TStoreService/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tStoreServiceClient) Fetch(ctx context.Context, in *FetchParam, opts ...grpc.CallOption) (*FetchResp, error) {
	out := new(FetchResp)
	err := c.cc.Invoke(ctx, "/pb.TStoreService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tStoreServiceClient) Delete(ctx context.Context, in *FetchParam, opts ...grpc.CallOption) (*FetchResp, error) {
	out := new(FetchResp)
	err := c.cc.Invoke(ctx, "/pb.TStoreService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tStoreServiceClient) Keys(ctx context.Context, in *FetchParam, opts ...grpc.CallOption) (*FetchResp, error) {
	out := new(FetchResp)
	err := c.cc.Invoke(ctx, "/pb.TStoreService/Keys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tStoreServiceClient) Push(ctx context.Context, in *SaveParam, opts ...grpc.CallOption) (*SaveResp, error) {
	out := new(SaveResp)
	err := c.cc.Invoke(ctx, "/pb.TStoreService/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tStoreServiceClient) Shift(ctx context.Context, in *FetchParam, opts ...grpc.CallOption) (*FetchResp, error) {
	out := new(FetchResp)
	err := c.cc.Invoke(ctx, "/pb.TStoreService/Shift", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TStoreServiceServer is the server API for TStoreService service.
type TStoreServiceServer interface {
	Save(context.Context, *SaveParam) (*SaveResp, error)
	Fetch(context.Context, *FetchParam) (*FetchResp, error)
	Delete(context.Context, *FetchParam) (*FetchResp, error)
	Keys(context.Context, *FetchParam) (*FetchResp, error)
	Push(context.Context, *SaveParam) (*SaveResp, error)
	Shift(context.Context, *FetchParam) (*FetchResp, error)
}

// UnimplementedTStoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTStoreServiceServer struct {
}

func (*UnimplementedTStoreServiceServer) Save(context.Context, *SaveParam) (*SaveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (*UnimplementedTStoreServiceServer) Fetch(context.Context, *FetchParam) (*FetchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (*UnimplementedTStoreServiceServer) Delete(context.Context, *FetchParam) (*FetchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedTStoreServiceServer) Keys(context.Context, *FetchParam) (*FetchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Keys not implemented")
}
func (*UnimplementedTStoreServiceServer) Push(context.Context, *SaveParam) (*SaveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (*UnimplementedTStoreServiceServer) Shift(context.Context, *FetchParam) (*FetchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shift not implemented")
}

func RegisterTStoreServiceServer(s *grpc.Server, srv TStoreServiceServer) {
	s.RegisterService(&_TStoreService_serviceDesc, srv)
}

func _TStoreService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TStoreServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TStoreService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TStoreServiceServer).Save(ctx, req.(*SaveParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _TStoreService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TStoreServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TStoreService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TStoreServiceServer).Fetch(ctx, req.(*FetchParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _TStoreService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TStoreServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TStoreService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TStoreServiceServer).Delete(ctx, req.(*FetchParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _TStoreService_Keys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TStoreServiceServer).Keys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TStoreService/Keys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TStoreServiceServer).Keys(ctx, req.(*FetchParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _TStoreService_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TStoreServiceServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TStoreService/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TStoreServiceServer).Push(ctx, req.(*SaveParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _TStoreService_Shift_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TStoreServiceServer).Shift(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TStoreService/Shift",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TStoreServiceServer).Shift(ctx, req.(*FetchParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _TStoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TStoreService",
	HandlerType: (*TStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _TStoreService_Save_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _TStoreService_Fetch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TStoreService_Delete_Handler,
		},
		{
			MethodName: "Keys",
			Handler:    _TStoreService_Keys_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _TStoreService_Push_Handler,
		},
		{
			MethodName: "Shift",
			Handler:    _TStoreService_Shift_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
