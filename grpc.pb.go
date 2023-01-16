// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: grpc.proto

package orderproto

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

// The request message containing the rating name.
type Rate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating float32 `protobuf:"fixed32,1,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *Rate) Reset() {
	*x = Rate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rate) ProtoMessage() {}

func (x *Rate) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rate.ProtoReflect.Descriptor instead.
func (*Rate) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Rate) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type OrderUserSide struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Taxitype string `protobuf:"bytes,1,opt,name=taxitype,proto3" json:"taxitype,omitempty"`
	Driver   string `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	From     string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To       string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *OrderUserSide) Reset() {
	*x = OrderUserSide{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUserSide) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUserSide) ProtoMessage() {}

func (x *OrderUserSide) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUserSide.ProtoReflect.Descriptor instead.
func (*OrderUserSide) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *OrderUserSide) GetTaxitype() string {
	if x != nil {
		return x.Taxitype
	}
	return ""
}

func (x *OrderUserSide) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *OrderUserSide) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *OrderUserSide) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type UserOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order []*OrderUserSide `protobuf:"bytes,1,rep,name=order,proto3" json:"order,omitempty"`
}

func (x *UserOrders) Reset() {
	*x = UserOrders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrders) ProtoMessage() {}

func (x *UserOrders) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOrders.ProtoReflect.Descriptor instead.
func (*UserOrders) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *UserOrders) GetOrder() []*OrderUserSide {
	if x != nil {
		return x.Order
	}
	return nil
}

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *Name) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RequestTaxi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Taxitype string `protobuf:"bytes,1,opt,name=taxitype,proto3" json:"taxitype,omitempty"`
	From     string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To       string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *RequestTaxi) Reset() {
	*x = RequestTaxi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestTaxi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestTaxi) ProtoMessage() {}

func (x *RequestTaxi) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[4]
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
	return file_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *RequestTaxi) GetTaxitype() string {
	if x != nil {
		return x.Taxitype
	}
	return ""
}

func (x *RequestTaxi) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *RequestTaxi) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type TaxiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Taxitype string `protobuf:"bytes,1,opt,name=taxitype,proto3" json:"taxitype,omitempty"`
	Driver   string `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	From     string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To       string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Message  string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TaxiResponse) Reset() {
	*x = TaxiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaxiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaxiResponse) ProtoMessage() {}

func (x *TaxiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaxiResponse.ProtoReflect.Descriptor instead.
func (*TaxiResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *TaxiResponse) GetTaxitype() string {
	if x != nil {
		return x.Taxitype
	}
	return ""
}

func (x *TaxiResponse) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *TaxiResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TaxiResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TaxiResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type OrderRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating float32 `protobuf:"fixed32,1,opt,name=rating,proto3" json:"rating,omitempty"`
	Email  string  `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *OrderRate) Reset() {
	*x = OrderRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRate) ProtoMessage() {}

func (x *OrderRate) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRate.ProtoReflect.Descriptor instead.
func (*OrderRate) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *OrderRate) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *OrderRate) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type OrderDriverSide struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Taxitype string `protobuf:"bytes,1,opt,name=taxitype,proto3" json:"taxitype,omitempty"`
	User     string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	From     string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To       string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *OrderDriverSide) Reset() {
	*x = OrderDriverSide{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDriverSide) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDriverSide) ProtoMessage() {}

func (x *OrderDriverSide) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDriverSide.ProtoReflect.Descriptor instead.
func (*OrderDriverSide) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *OrderDriverSide) GetTaxitype() string {
	if x != nil {
		return x.Taxitype
	}
	return ""
}

func (x *OrderDriverSide) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *OrderDriverSide) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *OrderDriverSide) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type DriverOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order []*OrderDriverSide `protobuf:"bytes,1,rep,name=order,proto3" json:"order,omitempty"`
}

func (x *DriverOrders) Reset() {
	*x = DriverOrders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverOrders) ProtoMessage() {}

func (x *DriverOrders) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverOrders.ProtoReflect.Descriptor instead.
func (*DriverOrders) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *DriverOrders) GetOrder() []*OrderDriverSide {
	if x != nil {
		return x.Order
	}
	return nil
}

type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{9}
}

func (x *Email) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_grpc_proto protoreflect.FileDescriptor

var file_grpc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x72,
	0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x22, 0x1e, 0x0a, 0x04, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x67, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x78, 0x69, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x78, 0x69, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x22,
	0x3b, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x69, 0x64, 0x65, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x1a, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x54, 0x61, 0x78, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x78, 0x69, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x78, 0x69, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x78, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x78, 0x69,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x78, 0x69,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x39, 0x0a, 0x09, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x65, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x78, 0x69,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x78, 0x69,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0c,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x1d, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0xfc, 0x02, 0x0a,
	0x08, 0x49, 0x6e, 0x6e, 0x6f, 0x54, 0x61, 0x78, 0x69, 0x12, 0x33, 0x0a, 0x0f, 0x52, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x1a, 0x0e, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x2e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x54, 0x61, 0x78, 0x69, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x54,
	0x61, 0x78, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x61, 0x78, 0x69, 0x1a,
	0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x2e, 0x54, 0x61, 0x78, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x11, 0x52, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x13,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x61, 0x74, 0x65, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x54, 0x61, 0x78,
	0x69, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x54, 0x61,
	0x78, 0x69, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x19,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x54, 0x61, 0x78, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x64, 0x76, 0x61, 0x72, 0x64,
	0x55, 0x4c, 0x2f, 0x49, 0x6e, 0x6e, 0x6f, 0x54, 0x61, 0x78, 0x69, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_rawDescData = file_grpc_proto_rawDesc
)

func file_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_rawDescData)
	})
	return file_grpc_proto_rawDescData
}

var file_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_grpc_proto_goTypes = []interface{}{
	(*Rate)(nil),            // 0: grpcTaxi.Rate
	(*OrderUserSide)(nil),   // 1: grpcTaxi.OrderUserSide
	(*UserOrders)(nil),      // 2: grpcTaxi.UserOrders
	(*Name)(nil),            // 3: grpcTaxi.Name
	(*RequestTaxi)(nil),     // 4: grpcTaxi.RequestTaxi
	(*TaxiResponse)(nil),    // 5: grpcTaxi.TaxiResponse
	(*OrderRate)(nil),       // 6: grpcTaxi.OrderRate
	(*OrderDriverSide)(nil), // 7: grpcTaxi.OrderDriverSide
	(*DriverOrders)(nil),    // 8: grpcTaxi.DriverOrders
	(*Email)(nil),           // 9: grpcTaxi.Email
}
var file_grpc_proto_depIdxs = []int32{
	1, // 0: grpcTaxi.UserOrders.order:type_name -> grpcTaxi.OrderUserSide
	7, // 1: grpcTaxi.DriverOrders.order:type_name -> grpcTaxi.OrderDriverSide
	0, // 2: grpcTaxi.InnoTaxi.RateOrderByUser:input_type -> grpcTaxi.Rate
	9, // 3: grpcTaxi.InnoTaxi.GetAllUserOrders:input_type -> grpcTaxi.Email
	4, // 4: grpcTaxi.InnoTaxi.OrderTaxi:input_type -> grpcTaxi.RequestTaxi
	6, // 5: grpcTaxi.InnoTaxi.RateOrderByDriver:input_type -> grpcTaxi.OrderRate
	9, // 6: grpcTaxi.InnoTaxi.GetAllDriverOrders:input_type -> grpcTaxi.Email
	9, // 7: grpcTaxi.InnoTaxi.EndOrder:input_type -> grpcTaxi.Email
	0, // 8: grpcTaxi.InnoTaxi.RateOrderByUser:output_type -> grpcTaxi.Rate
	2, // 9: grpcTaxi.InnoTaxi.GetAllUserOrders:output_type -> grpcTaxi.UserOrders
	5, // 10: grpcTaxi.InnoTaxi.OrderTaxi:output_type -> grpcTaxi.TaxiResponse
	7, // 11: grpcTaxi.InnoTaxi.RateOrderByDriver:output_type -> grpcTaxi.OrderDriverSide
	8, // 12: grpcTaxi.InnoTaxi.GetAllDriverOrders:output_type -> grpcTaxi.DriverOrders
	7, // 13: grpcTaxi.InnoTaxi.EndOrder:output_type -> grpcTaxi.OrderDriverSide
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_proto_init() }
func file_grpc_proto_init() {
	if File_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rate); i {
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
		file_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderUserSide); i {
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
		file_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOrders); i {
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
		file_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaxiResponse); i {
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
		file_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRate); i {
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
		file_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDriverSide); i {
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
		file_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverOrders); i {
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
		file_grpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
			RawDescriptor: file_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto = out.File
	file_grpc_proto_rawDesc = nil
	file_grpc_proto_goTypes = nil
	file_grpc_proto_depIdxs = nil
}
