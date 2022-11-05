// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: upload.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ConfigName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *ConfigName) Reset() {
	*x = ConfigName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigName) ProtoMessage() {}

func (x *ConfigName) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigName.ProtoReflect.Descriptor instead.
func (*ConfigName) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigName) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type ConfigBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConfigBody) Reset() {
	*x = ConfigBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigBody) ProtoMessage() {}

func (x *ConfigBody) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigBody.ProtoReflect.Descriptor instead.
func (*ConfigBody) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigBody) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service []byte            `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Data    map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetService() []byte {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Config) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type ConfigBodyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ConfigBodyResponse) Reset() {
	*x = ConfigBodyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigBodyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigBodyResponse) ProtoMessage() {}

func (x *ConfigBodyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigBodyResponse.ProtoReflect.Descriptor instead.
func (*ConfigBodyResponse) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigBodyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_upload_proto protoreflect.FileDescriptor

var file_upload_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0a,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x70, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a,
	0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x82, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2e, 0x0a, 0x12, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xcc, 0x04, 0x0a, 0x13,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x07, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x13, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x53, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3b, 0x12, 0x1c, 0x41, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x20, 0x6f, 0x66, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1b, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x20, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0b, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x1a, 0x0b, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x22,
	0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x7d, 0x92, 0x41, 0x46,
	0x12, 0x13, 0x47, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x20, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x2f, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x7b, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x40, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7d, 0x12, 0x8d, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x07, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0x13, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x32, 0x14, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3d, 0x12, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x20, 0x6f, 0x66, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x20, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0b, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x13, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x6f, 0x64,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x2a, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x7b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x7d, 0x92, 0x41, 0x40, 0x12, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x2f, 0x53, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20,
	0x7b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x40, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7d, 0x42, 0x82, 0x01, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x92, 0x41, 0x79, 0x12, 0x77, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x20, 0x53, 0x65, 0x74, 0x75, 0x70, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5a,
	0x0a, 0x19, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x20, 0x53, 0x65,
	0x74, 0x75, 0x70, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x68, 0x75, 0x73, 0x61, 0x69, 0x6e, 0x6e, 0x6f, 0x76, 0x2f, 0x73, 0x62, 0x65, 0x72, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x1a, 0x14, 0x6b, 0x68, 0x75, 0x73, 0x61, 0x69, 0x6e, 0x6e, 0x6f, 0x76,
	0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_upload_proto_rawDescOnce sync.Once
	file_upload_proto_rawDescData = file_upload_proto_rawDesc
)

func file_upload_proto_rawDescGZIP() []byte {
	file_upload_proto_rawDescOnce.Do(func() {
		file_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_upload_proto_rawDescData)
	})
	return file_upload_proto_rawDescData
}

var file_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_upload_proto_goTypes = []interface{}{
	(*ConfigName)(nil),         // 0: ConfigName
	(*ConfigBody)(nil),         // 1: ConfigBody
	(*Config)(nil),             // 2: Config
	(*ConfigBodyResponse)(nil), // 3: ConfigBodyResponse
	nil,                        // 4: ConfigBody.DataEntry
	nil,                        // 5: Config.DataEntry
}
var file_upload_proto_depIdxs = []int32{
	4, // 0: ConfigBody.data:type_name -> ConfigBody.DataEntry
	5, // 1: Config.data:type_name -> Config.DataEntry
	2, // 2: UploadConfigService.UploadConfig:input_type -> Config
	0, // 3: UploadConfigService.GetConfig:input_type -> ConfigName
	2, // 4: UploadConfigService.UpdateConfig:input_type -> Config
	0, // 5: UploadConfigService.DeleteConfig:input_type -> ConfigName
	3, // 6: UploadConfigService.UploadConfig:output_type -> ConfigBodyResponse
	1, // 7: UploadConfigService.GetConfig:output_type -> ConfigBody
	3, // 8: UploadConfigService.UpdateConfig:output_type -> ConfigBodyResponse
	3, // 9: UploadConfigService.DeleteConfig:output_type -> ConfigBodyResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_upload_proto_init() }
func file_upload_proto_init() {
	if File_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigName); i {
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
		file_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigBody); i {
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
		file_upload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_upload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigBodyResponse); i {
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
			RawDescriptor: file_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upload_proto_goTypes,
		DependencyIndexes: file_upload_proto_depIdxs,
		MessageInfos:      file_upload_proto_msgTypes,
	}.Build()
	File_upload_proto = out.File
	file_upload_proto_rawDesc = nil
	file_upload_proto_goTypes = nil
	file_upload_proto_depIdxs = nil
}
