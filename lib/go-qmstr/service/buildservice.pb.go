// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: buildservice.proto

package service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type BuildResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BuildResponse) Reset() {
	*x = BuildResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildResponse) ProtoMessage() {}

func (x *BuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buildservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildResponse.ProtoReflect.Descriptor instead.
func (*BuildResponse) Descriptor() ([]byte, []int) {
	return file_buildservice_proto_rawDescGZIP(), []int{0}
}

func (x *BuildResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PushFileMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PushFileMessage) Reset() {
	*x = PushFileMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushFileMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushFileMessage) ProtoMessage() {}

func (x *PushFileMessage) ProtoReflect() protoreflect.Message {
	mi := &file_buildservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushFileMessage.ProtoReflect.Descriptor instead.
func (*PushFileMessage) Descriptor() ([]byte, []int) {
	return file_buildservice_proto_rawDescGZIP(), []int{1}
}

func (x *PushFileMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PushFileMessage) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *PushFileMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PushFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *PushFileResponse) Reset() {
	*x = PushFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushFileResponse) ProtoMessage() {}

func (x *PushFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buildservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushFileResponse.ProtoReflect.Descriptor instead.
func (*PushFileResponse) Descriptor() ([]byte, []int) {
	return file_buildservice_proto_rawDescGZIP(), []int{2}
}

func (x *PushFileResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DeleteMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Edge string `protobuf:"bytes,2,opt,name=edge,proto3" json:"edge,omitempty"`
}

func (x *DeleteMessage) Reset() {
	*x = DeleteMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessage) ProtoMessage() {}

func (x *DeleteMessage) ProtoReflect() protoreflect.Message {
	mi := &file_buildservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessage.ProtoReflect.Descriptor instead.
func (*DeleteMessage) Descriptor() ([]byte, []int) {
	return file_buildservice_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteMessage) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *DeleteMessage) GetEdge() string {
	if x != nil {
		return x.Edge
	}
	return ""
}

type UpdatePackageNodeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *PackageNode `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	Targets []*FileNode  `protobuf:"bytes,2,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *UpdatePackageNodeMessage) Reset() {
	*x = UpdatePackageNodeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePackageNodeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePackageNodeMessage) ProtoMessage() {}

func (x *UpdatePackageNodeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_buildservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePackageNodeMessage.ProtoReflect.Descriptor instead.
func (*UpdatePackageNodeMessage) Descriptor() ([]byte, []int) {
	return file_buildservice_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePackageNodeMessage) GetPackage() *PackageNode {
	if x != nil {
		return x.Package
	}
	return nil
}

func (x *UpdatePackageNodeMessage) GetTargets() []*FileNode {
	if x != nil {
		return x.Targets
	}
	return nil
}

var File_buildservice_proto protoreflect.FileDescriptor

var file_buildservice_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0f, 0x64,
	0x61, 0x74, 0x61, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29,
	0x0a, 0x0d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4d, 0x0a, 0x0f, 0x50, 0x75, 0x73,
	0x68, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x35, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x64, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x65, 0x64, 0x67, 0x65, 0x22, 0x77, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73,
	0x32, 0xa5, 0x05, 0x0a, 0x0c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x36, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3d, 0x0a, 0x0e, 0x53, 0x65, 0x6e,
	0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x08, 0x50, 0x75, 0x73, 0x68,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x75, 0x73, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x64, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e,
	0x71, 0x6d, 0x73, 0x74, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x58, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buildservice_proto_rawDescOnce sync.Once
	file_buildservice_proto_rawDescData = file_buildservice_proto_rawDesc
)

func file_buildservice_proto_rawDescGZIP() []byte {
	file_buildservice_proto_rawDescOnce.Do(func() {
		file_buildservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_buildservice_proto_rawDescData)
	})
	return file_buildservice_proto_rawDescData
}

var file_buildservice_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_buildservice_proto_goTypes = []interface{}{
	(*BuildResponse)(nil),            // 0: service.BuildResponse
	(*PushFileMessage)(nil),          // 1: service.PushFileMessage
	(*PushFileResponse)(nil),         // 2: service.PushFileResponse
	(*DeleteMessage)(nil),            // 3: service.DeleteMessage
	(*UpdatePackageNodeMessage)(nil), // 4: service.UpdatePackageNodeMessage
	(*PackageNode)(nil),              // 5: service.PackageNode
	(*FileNode)(nil),                 // 6: service.FileNode
	(*InfoNode)(nil),                 // 7: service.InfoNode
	(*ProjectNode)(nil),              // 8: service.ProjectNode
}
var file_buildservice_proto_depIdxs = []int32{
	5,  // 0: service.UpdatePackageNodeMessage.package:type_name -> service.PackageNode
	6,  // 1: service.UpdatePackageNodeMessage.targets:type_name -> service.FileNode
	6,  // 2: service.BuildService.Build:input_type -> service.FileNode
	7,  // 3: service.BuildService.SendBuildError:input_type -> service.InfoNode
	1,  // 4: service.BuildService.PushFile:input_type -> service.PushFileMessage
	4,  // 5: service.BuildService.UpdatePackageNode:input_type -> service.UpdatePackageNodeMessage
	5,  // 6: service.BuildService.CreatePackage:input_type -> service.PackageNode
	8,  // 7: service.BuildService.CreateProject:input_type -> service.ProjectNode
	5,  // 8: service.BuildService.UpdateProjectNode:input_type -> service.PackageNode
	8,  // 9: service.BuildService.GetProjectNode:input_type -> service.ProjectNode
	3,  // 10: service.BuildService.DeleteNode:input_type -> service.DeleteMessage
	3,  // 11: service.BuildService.DeleteEdge:input_type -> service.DeleteMessage
	0,  // 12: service.BuildService.Build:output_type -> service.BuildResponse
	0,  // 13: service.BuildService.SendBuildError:output_type -> service.BuildResponse
	2,  // 14: service.BuildService.PushFile:output_type -> service.PushFileResponse
	0,  // 15: service.BuildService.UpdatePackageNode:output_type -> service.BuildResponse
	0,  // 16: service.BuildService.CreatePackage:output_type -> service.BuildResponse
	0,  // 17: service.BuildService.CreateProject:output_type -> service.BuildResponse
	0,  // 18: service.BuildService.UpdateProjectNode:output_type -> service.BuildResponse
	8,  // 19: service.BuildService.GetProjectNode:output_type -> service.ProjectNode
	0,  // 20: service.BuildService.DeleteNode:output_type -> service.BuildResponse
	0,  // 21: service.BuildService.DeleteEdge:output_type -> service.BuildResponse
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_buildservice_proto_init() }
func file_buildservice_proto_init() {
	if File_buildservice_proto != nil {
		return
	}
	file_datamodel_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_buildservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildResponse); i {
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
		file_buildservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushFileMessage); i {
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
		file_buildservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushFileResponse); i {
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
		file_buildservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessage); i {
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
		file_buildservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePackageNodeMessage); i {
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
			RawDescriptor: file_buildservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buildservice_proto_goTypes,
		DependencyIndexes: file_buildservice_proto_depIdxs,
		MessageInfos:      file_buildservice_proto_msgTypes,
	}.Build()
	File_buildservice_proto = out.File
	file_buildservice_proto_rawDesc = nil
	file_buildservice_proto_goTypes = nil
	file_buildservice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BuildServiceClient is the client API for BuildService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildServiceClient interface {
	Build(ctx context.Context, opts ...grpc.CallOption) (BuildService_BuildClient, error)
	SendBuildError(ctx context.Context, in *InfoNode, opts ...grpc.CallOption) (*BuildResponse, error)
	PushFile(ctx context.Context, in *PushFileMessage, opts ...grpc.CallOption) (*PushFileResponse, error)
	UpdatePackageNode(ctx context.Context, in *UpdatePackageNodeMessage, opts ...grpc.CallOption) (*BuildResponse, error)
	CreatePackage(ctx context.Context, in *PackageNode, opts ...grpc.CallOption) (*BuildResponse, error)
	CreateProject(ctx context.Context, in *ProjectNode, opts ...grpc.CallOption) (*BuildResponse, error)
	UpdateProjectNode(ctx context.Context, opts ...grpc.CallOption) (BuildService_UpdateProjectNodeClient, error)
	GetProjectNode(ctx context.Context, in *ProjectNode, opts ...grpc.CallOption) (*ProjectNode, error)
	DeleteNode(ctx context.Context, opts ...grpc.CallOption) (BuildService_DeleteNodeClient, error)
	DeleteEdge(ctx context.Context, in *DeleteMessage, opts ...grpc.CallOption) (*BuildResponse, error)
}

type buildServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildServiceClient(cc grpc.ClientConnInterface) BuildServiceClient {
	return &buildServiceClient{cc}
}

func (c *buildServiceClient) Build(ctx context.Context, opts ...grpc.CallOption) (BuildService_BuildClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BuildService_serviceDesc.Streams[0], "/service.BuildService/Build", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildServiceBuildClient{stream}
	return x, nil
}

type BuildService_BuildClient interface {
	Send(*FileNode) error
	CloseAndRecv() (*BuildResponse, error)
	grpc.ClientStream
}

type buildServiceBuildClient struct {
	grpc.ClientStream
}

func (x *buildServiceBuildClient) Send(m *FileNode) error {
	return x.ClientStream.SendMsg(m)
}

func (x *buildServiceBuildClient) CloseAndRecv() (*BuildResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BuildResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *buildServiceClient) SendBuildError(ctx context.Context, in *InfoNode, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, "/service.BuildService/SendBuildError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) PushFile(ctx context.Context, in *PushFileMessage, opts ...grpc.CallOption) (*PushFileResponse, error) {
	out := new(PushFileResponse)
	err := c.cc.Invoke(ctx, "/service.BuildService/PushFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) UpdatePackageNode(ctx context.Context, in *UpdatePackageNodeMessage, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, "/service.BuildService/UpdatePackageNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) CreatePackage(ctx context.Context, in *PackageNode, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, "/service.BuildService/CreatePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) CreateProject(ctx context.Context, in *ProjectNode, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, "/service.BuildService/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) UpdateProjectNode(ctx context.Context, opts ...grpc.CallOption) (BuildService_UpdateProjectNodeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BuildService_serviceDesc.Streams[1], "/service.BuildService/UpdateProjectNode", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildServiceUpdateProjectNodeClient{stream}
	return x, nil
}

type BuildService_UpdateProjectNodeClient interface {
	Send(*PackageNode) error
	CloseAndRecv() (*BuildResponse, error)
	grpc.ClientStream
}

type buildServiceUpdateProjectNodeClient struct {
	grpc.ClientStream
}

func (x *buildServiceUpdateProjectNodeClient) Send(m *PackageNode) error {
	return x.ClientStream.SendMsg(m)
}

func (x *buildServiceUpdateProjectNodeClient) CloseAndRecv() (*BuildResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BuildResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *buildServiceClient) GetProjectNode(ctx context.Context, in *ProjectNode, opts ...grpc.CallOption) (*ProjectNode, error) {
	out := new(ProjectNode)
	err := c.cc.Invoke(ctx, "/service.BuildService/GetProjectNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) DeleteNode(ctx context.Context, opts ...grpc.CallOption) (BuildService_DeleteNodeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BuildService_serviceDesc.Streams[2], "/service.BuildService/DeleteNode", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildServiceDeleteNodeClient{stream}
	return x, nil
}

type BuildService_DeleteNodeClient interface {
	Send(*DeleteMessage) error
	CloseAndRecv() (*BuildResponse, error)
	grpc.ClientStream
}

type buildServiceDeleteNodeClient struct {
	grpc.ClientStream
}

func (x *buildServiceDeleteNodeClient) Send(m *DeleteMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *buildServiceDeleteNodeClient) CloseAndRecv() (*BuildResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BuildResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *buildServiceClient) DeleteEdge(ctx context.Context, in *DeleteMessage, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, "/service.BuildService/DeleteEdge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildServiceServer is the server API for BuildService service.
type BuildServiceServer interface {
	Build(BuildService_BuildServer) error
	SendBuildError(context.Context, *InfoNode) (*BuildResponse, error)
	PushFile(context.Context, *PushFileMessage) (*PushFileResponse, error)
	UpdatePackageNode(context.Context, *UpdatePackageNodeMessage) (*BuildResponse, error)
	CreatePackage(context.Context, *PackageNode) (*BuildResponse, error)
	CreateProject(context.Context, *ProjectNode) (*BuildResponse, error)
	UpdateProjectNode(BuildService_UpdateProjectNodeServer) error
	GetProjectNode(context.Context, *ProjectNode) (*ProjectNode, error)
	DeleteNode(BuildService_DeleteNodeServer) error
	DeleteEdge(context.Context, *DeleteMessage) (*BuildResponse, error)
}

// UnimplementedBuildServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBuildServiceServer struct {
}

func (*UnimplementedBuildServiceServer) Build(BuildService_BuildServer) error {
	return status.Errorf(codes.Unimplemented, "method Build not implemented")
}
func (*UnimplementedBuildServiceServer) SendBuildError(context.Context, *InfoNode) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBuildError not implemented")
}
func (*UnimplementedBuildServiceServer) PushFile(context.Context, *PushFileMessage) (*PushFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushFile not implemented")
}
func (*UnimplementedBuildServiceServer) UpdatePackageNode(context.Context, *UpdatePackageNodeMessage) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePackageNode not implemented")
}
func (*UnimplementedBuildServiceServer) CreatePackage(context.Context, *PackageNode) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePackage not implemented")
}
func (*UnimplementedBuildServiceServer) CreateProject(context.Context, *ProjectNode) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (*UnimplementedBuildServiceServer) UpdateProjectNode(BuildService_UpdateProjectNodeServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateProjectNode not implemented")
}
func (*UnimplementedBuildServiceServer) GetProjectNode(context.Context, *ProjectNode) (*ProjectNode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectNode not implemented")
}
func (*UnimplementedBuildServiceServer) DeleteNode(BuildService_DeleteNodeServer) error {
	return status.Errorf(codes.Unimplemented, "method DeleteNode not implemented")
}
func (*UnimplementedBuildServiceServer) DeleteEdge(context.Context, *DeleteMessage) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEdge not implemented")
}

func RegisterBuildServiceServer(s *grpc.Server, srv BuildServiceServer) {
	s.RegisterService(&_BuildService_serviceDesc, srv)
}

func _BuildService_Build_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BuildServiceServer).Build(&buildServiceBuildServer{stream})
}

type BuildService_BuildServer interface {
	SendAndClose(*BuildResponse) error
	Recv() (*FileNode, error)
	grpc.ServerStream
}

type buildServiceBuildServer struct {
	grpc.ServerStream
}

func (x *buildServiceBuildServer) SendAndClose(m *BuildResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *buildServiceBuildServer) Recv() (*FileNode, error) {
	m := new(FileNode)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BuildService_SendBuildError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).SendBuildError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.BuildService/SendBuildError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).SendBuildError(ctx, req.(*InfoNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildService_PushFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushFileMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).PushFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.BuildService/PushFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).PushFile(ctx, req.(*PushFileMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildService_UpdatePackageNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePackageNodeMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).UpdatePackageNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.BuildService/UpdatePackageNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).UpdatePackageNode(ctx, req.(*UpdatePackageNodeMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildService_CreatePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).CreatePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.BuildService/CreatePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).CreatePackage(ctx, req.(*PackageNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildService_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.BuildService/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).CreateProject(ctx, req.(*ProjectNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildService_UpdateProjectNode_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BuildServiceServer).UpdateProjectNode(&buildServiceUpdateProjectNodeServer{stream})
}

type BuildService_UpdateProjectNodeServer interface {
	SendAndClose(*BuildResponse) error
	Recv() (*PackageNode, error)
	grpc.ServerStream
}

type buildServiceUpdateProjectNodeServer struct {
	grpc.ServerStream
}

func (x *buildServiceUpdateProjectNodeServer) SendAndClose(m *BuildResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *buildServiceUpdateProjectNodeServer) Recv() (*PackageNode, error) {
	m := new(PackageNode)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BuildService_GetProjectNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).GetProjectNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.BuildService/GetProjectNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).GetProjectNode(ctx, req.(*ProjectNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildService_DeleteNode_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BuildServiceServer).DeleteNode(&buildServiceDeleteNodeServer{stream})
}

type BuildService_DeleteNodeServer interface {
	SendAndClose(*BuildResponse) error
	Recv() (*DeleteMessage, error)
	grpc.ServerStream
}

type buildServiceDeleteNodeServer struct {
	grpc.ServerStream
}

func (x *buildServiceDeleteNodeServer) SendAndClose(m *BuildResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *buildServiceDeleteNodeServer) Recv() (*DeleteMessage, error) {
	m := new(DeleteMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BuildService_DeleteEdge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).DeleteEdge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.BuildService/DeleteEdge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).DeleteEdge(ctx, req.(*DeleteMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _BuildService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.BuildService",
	HandlerType: (*BuildServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendBuildError",
			Handler:    _BuildService_SendBuildError_Handler,
		},
		{
			MethodName: "PushFile",
			Handler:    _BuildService_PushFile_Handler,
		},
		{
			MethodName: "UpdatePackageNode",
			Handler:    _BuildService_UpdatePackageNode_Handler,
		},
		{
			MethodName: "CreatePackage",
			Handler:    _BuildService_CreatePackage_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _BuildService_CreateProject_Handler,
		},
		{
			MethodName: "GetProjectNode",
			Handler:    _BuildService_GetProjectNode_Handler,
		},
		{
			MethodName: "DeleteEdge",
			Handler:    _BuildService_DeleteEdge_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Build",
			Handler:       _BuildService_Build_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateProjectNode",
			Handler:       _BuildService_UpdateProjectNode_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DeleteNode",
			Handler:       _BuildService_DeleteNode_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "buildservice.proto",
}
