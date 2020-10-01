// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.11.4
// source: data/LoginUserSession.proto

package dataImpl

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LoginUserSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//用户id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//类型
	Type uint64 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	//用户昵称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	//是否登陆
	IsLogin bool `protobuf:"varint,4,opt,name=isLogin,proto3" json:"isLogin,omitempty"`
}

func (x *LoginUserSession) Reset() {
	*x = LoginUserSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_LoginUserSession_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserSession) ProtoMessage() {}

func (x *LoginUserSession) ProtoReflect() protoreflect.Message {
	mi := &file_data_LoginUserSession_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserSession.ProtoReflect.Descriptor instead.
func (*LoginUserSession) Descriptor() ([]byte, []int) {
	return file_data_LoginUserSession_proto_rawDescGZIP(), []int{0}
}

func (x *LoginUserSession) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoginUserSession) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *LoginUserSession) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginUserSession) GetIsLogin() bool {
	if x != nil {
		return x.IsLogin
	}
	return false
}

var File_data_LoginUserSession_proto protoreflect.FileDescriptor

var file_data_LoginUserSession_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a,
	0x10, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x42, 0x0f, 0x5a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x49, 0x6d, 0x70, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_LoginUserSession_proto_rawDescOnce sync.Once
	file_data_LoginUserSession_proto_rawDescData = file_data_LoginUserSession_proto_rawDesc
)

func file_data_LoginUserSession_proto_rawDescGZIP() []byte {
	file_data_LoginUserSession_proto_rawDescOnce.Do(func() {
		file_data_LoginUserSession_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_LoginUserSession_proto_rawDescData)
	})
	return file_data_LoginUserSession_proto_rawDescData
}

var file_data_LoginUserSession_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_data_LoginUserSession_proto_goTypes = []interface{}{
	(*LoginUserSession)(nil), // 0: LoginUserSession
}
var file_data_LoginUserSession_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_LoginUserSession_proto_init() }
func file_data_LoginUserSession_proto_init() {
	if File_data_LoginUserSession_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_LoginUserSession_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserSession); i {
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
			RawDescriptor: file_data_LoginUserSession_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_LoginUserSession_proto_goTypes,
		DependencyIndexes: file_data_LoginUserSession_proto_depIdxs,
		MessageInfos:      file_data_LoginUserSession_proto_msgTypes,
	}.Build()
	File_data_LoginUserSession_proto = out.File
	file_data_LoginUserSession_proto_rawDesc = nil
	file_data_LoginUserSession_proto_goTypes = nil
	file_data_LoginUserSession_proto_depIdxs = nil
}