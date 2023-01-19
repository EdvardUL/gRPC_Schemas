// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: driver/grpc.proto

package grpcTaxiDriver

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

type RequestTaxi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Taxitype string `protobuf:"bytes,1,opt,name=taxitype,proto3" json:"taxitype,omitempty"`
}

func (x *RequestTaxi) Reset() {
	*x = RequestTaxi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestTaxi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestTaxi) ProtoMessage() {}

func (x *RequestTaxi) ProtoReflect() protoreflect.Message {
	mi := &file_driver_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestTaxi.ProtoReflect.Descriptor instead.
func (*RequestTaxi) Descriptor() ([]byte, []int) {
	return file_driver_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *RequestTaxi) GetTaxitype() string {
	if x != nil {
		return x.Taxitype
	}
	return ""
}

type Driver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Taxitype    string `protobuf:"bytes,1,opt,name=taxitype,proto3" json:"taxitype,omitempty"`
	DriverEmail string `protobuf:"bytes,2,opt,name=DriverEmail,proto3" json:"DriverEmail,omitempty"`
}

func (x *Driver) Reset() {
	*x = Driver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Driver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Driver) ProtoMessage() {}

func (x *Driver) ProtoReflect() protoreflect.Message {
	mi := &file_driver_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Driver.ProtoReflect.Descriptor instead.
func (*Driver) Descriptor() ([]byte, []int) {
	return file_driver_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *Driver) GetTaxitype() string {
	if x != nil {
		return x.Taxitype
	}
	return ""
}

func (x *Driver) GetDriverEmail() string {
	if x != nil {
		return x.DriverEmail
	}
	return ""
}

var File_driver_grpc_proto protoreflect.FileDescriptor

var file_driver_grpc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x29, 0x0a, 0x0b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x61, 0x78, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x78, 0x69,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x78, 0x69,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x46, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x61, 0x78, 0x69, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x61, 0x78, 0x69, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0x64, 0x0a, 0x0e,
	0x49, 0x6e, 0x6e, 0x6f, 0x54, 0x61, 0x78, 0x69, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x52,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x61,
	0x78, 0x69, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x45, 0x64, 0x76, 0x61, 0x72, 0x64, 0x55, 0x4c, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x5f, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_driver_grpc_proto_rawDescOnce sync.Once
	file_driver_grpc_proto_rawDescData = file_driver_grpc_proto_rawDesc
)

func file_driver_grpc_proto_rawDescGZIP() []byte {
	file_driver_grpc_proto_rawDescOnce.Do(func() {
		file_driver_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_driver_grpc_proto_rawDescData)
	})
	return file_driver_grpc_proto_rawDescData
}

var file_driver_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_driver_grpc_proto_goTypes = []interface{}{
	(*RequestTaxi)(nil), // 0: grpcTaxiDriverSchema.RequestTaxi
	(*Driver)(nil),      // 1: grpcTaxiDriverSchema.Driver
}
var file_driver_grpc_proto_depIdxs = []int32{
	0, // 0: grpcTaxiDriverSchema.InnoTaxiDriver.GetFreeDriver:input_type -> grpcTaxiDriverSchema.RequestTaxi
	1, // 1: grpcTaxiDriverSchema.InnoTaxiDriver.GetFreeDriver:output_type -> grpcTaxiDriverSchema.Driver
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_driver_grpc_proto_init() }
func file_driver_grpc_proto_init() {
	if File_driver_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_driver_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestTaxi); i {
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
		file_driver_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Driver); i {
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
			RawDescriptor: file_driver_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_driver_grpc_proto_goTypes,
		DependencyIndexes: file_driver_grpc_proto_depIdxs,
		MessageInfos:      file_driver_grpc_proto_msgTypes,
	}.Build()
	File_driver_grpc_proto = out.File
	file_driver_grpc_proto_rawDesc = nil
	file_driver_grpc_proto_goTypes = nil
	file_driver_grpc_proto_depIdxs = nil
}
