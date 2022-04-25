// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.5.0
// source: ServiceRegistrationDiscovery.proto

package ServiceRegistrationDiscovery

import (
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

type ResponseCode int32

const (
	ResponseCode_SUCCUSS ResponseCode = 0
	ResponseCode_FAIL    ResponseCode = 1
)

// Enum value maps for ResponseCode.
var (
	ResponseCode_name = map[int32]string{
		0: "SUCCUSS",
		1: "FAIL",
	}
	ResponseCode_value = map[string]int32{
		"SUCCUSS": 0,
		"FAIL":    1,
	}
)

func (x ResponseCode) Enum() *ResponseCode {
	p := new(ResponseCode)
	*p = x
	return p
}

func (x ResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_ServiceRegistrationDiscovery_proto_enumTypes[0].Descriptor()
}

func (ResponseCode) Type() protoreflect.EnumType {
	return &file_ServiceRegistrationDiscovery_proto_enumTypes[0]
}

func (x ResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseCode.Descriptor instead.
func (ResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_ServiceRegistrationDiscovery_proto_rawDescGZIP(), []int{0}
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber   int32 `protobuf:"varint,1,opt,name=firstNumber,proto3" json:"firstNumber,omitempty"`
	SecondNumber  int32 `protobuf:"varint,2,opt,name=secondNumber,proto3" json:"secondNumber,omitempty"`
	ServiceNumber int32 `protobuf:"varint,3,opt,name=serviceNumber,proto3" json:"serviceNumber,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ServiceRegistrationDiscovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_ServiceRegistrationDiscovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_ServiceRegistrationDiscovery_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *Operation) GetSecondNumber() int32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

func (x *Operation) GetServiceNumber() int32 {
	if x != nil {
		return x.ServiceNumber
	}
	return 0
}

type ServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation []*Operation `protobuf:"bytes,1,rep,name=operation,proto3" json:"operation,omitempty"`
}

func (x *ServiceRequest) Reset() {
	*x = ServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ServiceRegistrationDiscovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRequest) ProtoMessage() {}

func (x *ServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ServiceRegistrationDiscovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceRequest.ProtoReflect.Descriptor instead.
func (*ServiceRequest) Descriptor() ([]byte, []int) {
	return file_ServiceRegistrationDiscovery_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceRequest) GetOperation() []*Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type ServiceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Result []int32 `protobuf:"varint,3,rep,packed,name=result,proto3" json:"result,omitempty"`
}

func (x *ServiceReply) Reset() {
	*x = ServiceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ServiceRegistrationDiscovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceReply) ProtoMessage() {}

func (x *ServiceReply) ProtoReflect() protoreflect.Message {
	mi := &file_ServiceRegistrationDiscovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceReply.ProtoReflect.Descriptor instead.
func (*ServiceReply) Descriptor() ([]byte, []int) {
	return file_ServiceRegistrationDiscovery_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ServiceReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ServiceReply) GetResult() []int32 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_ServiceRegistrationDiscovery_proto protoreflect.FileDescriptor

var file_ServiceRegistrationDiscovery_proto_rawDesc = []byte{
	0x0a, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x22, 0x77, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x57, 0x0a, 0x0e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2a, 0x25, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x55, 0x53, 0x53, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x32, 0x76, 0x0a, 0x07, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x65, 0x72, 0x12, 0x6b, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x2e, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x3b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ServiceRegistrationDiscovery_proto_rawDescOnce sync.Once
	file_ServiceRegistrationDiscovery_proto_rawDescData = file_ServiceRegistrationDiscovery_proto_rawDesc
)

func file_ServiceRegistrationDiscovery_proto_rawDescGZIP() []byte {
	file_ServiceRegistrationDiscovery_proto_rawDescOnce.Do(func() {
		file_ServiceRegistrationDiscovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_ServiceRegistrationDiscovery_proto_rawDescData)
	})
	return file_ServiceRegistrationDiscovery_proto_rawDescData
}

var file_ServiceRegistrationDiscovery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ServiceRegistrationDiscovery_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ServiceRegistrationDiscovery_proto_goTypes = []interface{}{
	(ResponseCode)(0),      // 0: ServiceRegistrationDiscovery.ResponseCode
	(*Operation)(nil),      // 1: ServiceRegistrationDiscovery.Operation
	(*ServiceRequest)(nil), // 2: ServiceRegistrationDiscovery.ServiceRequest
	(*ServiceReply)(nil),   // 3: ServiceRegistrationDiscovery.ServiceReply
}
var file_ServiceRegistrationDiscovery_proto_depIdxs = []int32{
	1, // 0: ServiceRegistrationDiscovery.ServiceRequest.operation:type_name -> ServiceRegistrationDiscovery.Operation
	2, // 1: ServiceRegistrationDiscovery.Greeter.ServiceMethod:input_type -> ServiceRegistrationDiscovery.ServiceRequest
	3, // 2: ServiceRegistrationDiscovery.Greeter.ServiceMethod:output_type -> ServiceRegistrationDiscovery.ServiceReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ServiceRegistrationDiscovery_proto_init() }
func file_ServiceRegistrationDiscovery_proto_init() {
	if File_ServiceRegistrationDiscovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ServiceRegistrationDiscovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
		file_ServiceRegistrationDiscovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceRequest); i {
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
		file_ServiceRegistrationDiscovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceReply); i {
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
			RawDescriptor: file_ServiceRegistrationDiscovery_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ServiceRegistrationDiscovery_proto_goTypes,
		DependencyIndexes: file_ServiceRegistrationDiscovery_proto_depIdxs,
		EnumInfos:         file_ServiceRegistrationDiscovery_proto_enumTypes,
		MessageInfos:      file_ServiceRegistrationDiscovery_proto_msgTypes,
	}.Build()
	File_ServiceRegistrationDiscovery_proto = out.File
	file_ServiceRegistrationDiscovery_proto_rawDesc = nil
	file_ServiceRegistrationDiscovery_proto_goTypes = nil
	file_ServiceRegistrationDiscovery_proto_depIdxs = nil
}