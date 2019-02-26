// Code generated by protoc-gen-go. DO NOT EDIT.
// source: controlservice.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogMessage struct {
	Msg                  []byte   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogMessage) Reset()         { *m = LogMessage{} }
func (m *LogMessage) String() string { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()    {}
func (*LogMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_controlservice_0a1a098ce6ace8f5, []int{0}
}
func (m *LogMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogMessage.Unmarshal(m, b)
}
func (m *LogMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogMessage.Marshal(b, m, deterministic)
}
func (dst *LogMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogMessage.Merge(dst, src)
}
func (m *LogMessage) XXX_Size() int {
	return xxx_messageInfo_LogMessage.Size(m)
}
func (m *LogMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LogMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LogMessage proto.InternalMessageInfo

func (m *LogMessage) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type LogResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogResponse) Reset()         { *m = LogResponse{} }
func (m *LogResponse) String() string { return proto.CompactTextString(m) }
func (*LogResponse) ProtoMessage()    {}
func (*LogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_controlservice_0a1a098ce6ace8f5, []int{1}
}
func (m *LogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogResponse.Unmarshal(m, b)
}
func (m *LogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogResponse.Marshal(b, m, deterministic)
}
func (dst *LogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogResponse.Merge(dst, src)
}
func (m *LogResponse) XXX_Size() int {
	return xxx_messageInfo_LogResponse.Size(m)
}
func (m *LogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogResponse proto.InternalMessageInfo

func (m *LogResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type QuitMessage struct {
	Kill                 bool     `protobuf:"varint,1,opt,name=kill,proto3" json:"kill,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuitMessage) Reset()         { *m = QuitMessage{} }
func (m *QuitMessage) String() string { return proto.CompactTextString(m) }
func (*QuitMessage) ProtoMessage()    {}
func (*QuitMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_controlservice_0a1a098ce6ace8f5, []int{2}
}
func (m *QuitMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuitMessage.Unmarshal(m, b)
}
func (m *QuitMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuitMessage.Marshal(b, m, deterministic)
}
func (dst *QuitMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuitMessage.Merge(dst, src)
}
func (m *QuitMessage) XXX_Size() int {
	return xxx_messageInfo_QuitMessage.Size(m)
}
func (m *QuitMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_QuitMessage.DiscardUnknown(m)
}

var xxx_messageInfo_QuitMessage proto.InternalMessageInfo

func (m *QuitMessage) GetKill() bool {
	if m != nil {
		return m.Kill
	}
	return false
}

type QuitResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuitResponse) Reset()         { *m = QuitResponse{} }
func (m *QuitResponse) String() string { return proto.CompactTextString(m) }
func (*QuitResponse) ProtoMessage()    {}
func (*QuitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_controlservice_0a1a098ce6ace8f5, []int{3}
}
func (m *QuitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuitResponse.Unmarshal(m, b)
}
func (m *QuitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuitResponse.Marshal(b, m, deterministic)
}
func (dst *QuitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuitResponse.Merge(dst, src)
}
func (m *QuitResponse) XXX_Size() int {
	return xxx_messageInfo_QuitResponse.Size(m)
}
func (m *QuitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuitResponse proto.InternalMessageInfo

func (m *QuitResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type SwitchPhaseMessage struct {
	Phase                Phase    `protobuf:"varint,1,opt,name=phase,proto3,enum=service.Phase" json:"phase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwitchPhaseMessage) Reset()         { *m = SwitchPhaseMessage{} }
func (m *SwitchPhaseMessage) String() string { return proto.CompactTextString(m) }
func (*SwitchPhaseMessage) ProtoMessage()    {}
func (*SwitchPhaseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_controlservice_0a1a098ce6ace8f5, []int{4}
}
func (m *SwitchPhaseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwitchPhaseMessage.Unmarshal(m, b)
}
func (m *SwitchPhaseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwitchPhaseMessage.Marshal(b, m, deterministic)
}
func (dst *SwitchPhaseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwitchPhaseMessage.Merge(dst, src)
}
func (m *SwitchPhaseMessage) XXX_Size() int {
	return xxx_messageInfo_SwitchPhaseMessage.Size(m)
}
func (m *SwitchPhaseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SwitchPhaseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SwitchPhaseMessage proto.InternalMessageInfo

func (m *SwitchPhaseMessage) GetPhase() Phase {
	if m != nil {
		return m.Phase
	}
	return Phase_INIT
}

type SwitchPhaseResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwitchPhaseResponse) Reset()         { *m = SwitchPhaseResponse{} }
func (m *SwitchPhaseResponse) String() string { return proto.CompactTextString(m) }
func (*SwitchPhaseResponse) ProtoMessage()    {}
func (*SwitchPhaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_controlservice_0a1a098ce6ace8f5, []int{5}
}
func (m *SwitchPhaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwitchPhaseResponse.Unmarshal(m, b)
}
func (m *SwitchPhaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwitchPhaseResponse.Marshal(b, m, deterministic)
}
func (dst *SwitchPhaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwitchPhaseResponse.Merge(dst, src)
}
func (m *SwitchPhaseResponse) XXX_Size() int {
	return xxx_messageInfo_SwitchPhaseResponse.Size(m)
}
func (m *SwitchPhaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SwitchPhaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SwitchPhaseResponse proto.InternalMessageInfo

func (m *SwitchPhaseResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SwitchPhaseResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type PackageRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PackageRequest) Reset()         { *m = PackageRequest{} }
func (m *PackageRequest) String() string { return proto.CompactTextString(m) }
func (*PackageRequest) ProtoMessage()    {}
func (*PackageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_controlservice_0a1a098ce6ace8f5, []int{6}
}
func (m *PackageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageRequest.Unmarshal(m, b)
}
func (m *PackageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageRequest.Marshal(b, m, deterministic)
}
func (dst *PackageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageRequest.Merge(dst, src)
}
func (m *PackageRequest) XXX_Size() int {
	return xxx_messageInfo_PackageRequest.Size(m)
}
func (m *PackageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PackageRequest proto.InternalMessageInfo

type StatusMessage struct {
	Phase                bool     `protobuf:"varint,1,opt,name=phase,proto3" json:"phase,omitempty"`
	Switch               bool     `protobuf:"varint,2,opt,name=switch,proto3" json:"switch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusMessage) Reset()         { *m = StatusMessage{} }
func (m *StatusMessage) String() string { return proto.CompactTextString(m) }
func (*StatusMessage) ProtoMessage()    {}
func (*StatusMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_controlservice_0a1a098ce6ace8f5, []int{7}
}
func (m *StatusMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusMessage.Unmarshal(m, b)
}
func (m *StatusMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusMessage.Marshal(b, m, deterministic)
}
func (dst *StatusMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusMessage.Merge(dst, src)
}
func (m *StatusMessage) XXX_Size() int {
	return xxx_messageInfo_StatusMessage.Size(m)
}
func (m *StatusMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StatusMessage proto.InternalMessageInfo

func (m *StatusMessage) GetPhase() bool {
	if m != nil {
		return m.Phase
	}
	return false
}

func (m *StatusMessage) GetSwitch() bool {
	if m != nil {
		return m.Switch
	}
	return false
}

type StatusResponse struct {
	Phase                string   `protobuf:"bytes,1,opt,name=phase,proto3" json:"phase,omitempty"`
	PhaseID              Phase    `protobuf:"varint,2,opt,name=phaseID,proto3,enum=service.Phase" json:"phaseID,omitempty"`
	Switching            bool     `protobuf:"varint,3,opt,name=switching,proto3" json:"switching,omitempty"`
	Error                string   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_controlservice_0a1a098ce6ace8f5, []int{8}
}
func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (dst *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(dst, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetPhase() string {
	if m != nil {
		return m.Phase
	}
	return ""
}

func (m *StatusResponse) GetPhaseID() Phase {
	if m != nil {
		return m.PhaseID
	}
	return Phase_INIT
}

func (m *StatusResponse) GetSwitching() bool {
	if m != nil {
		return m.Switching
	}
	return false
}

func (m *StatusResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type EventMessage struct {
	Class                EventClass `protobuf:"varint,1,opt,name=class,proto3,enum=service.EventClass" json:"class,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EventMessage) Reset()         { *m = EventMessage{} }
func (m *EventMessage) String() string { return proto.CompactTextString(m) }
func (*EventMessage) ProtoMessage()    {}
func (*EventMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_controlservice_0a1a098ce6ace8f5, []int{9}
}
func (m *EventMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventMessage.Unmarshal(m, b)
}
func (m *EventMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventMessage.Marshal(b, m, deterministic)
}
func (dst *EventMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMessage.Merge(dst, src)
}
func (m *EventMessage) XXX_Size() int {
	return xxx_messageInfo_EventMessage.Size(m)
}
func (m *EventMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EventMessage proto.InternalMessageInfo

func (m *EventMessage) GetClass() EventClass {
	if m != nil {
		return m.Class
	}
	return EventClass_ALL
}

type ExportRequest struct {
	Wait                 bool     `protobuf:"varint,1,opt,name=wait,proto3" json:"wait,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportRequest) Reset()         { *m = ExportRequest{} }
func (m *ExportRequest) String() string { return proto.CompactTextString(m) }
func (*ExportRequest) ProtoMessage()    {}
func (*ExportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_controlservice_0a1a098ce6ace8f5, []int{10}
}
func (m *ExportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportRequest.Unmarshal(m, b)
}
func (m *ExportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportRequest.Marshal(b, m, deterministic)
}
func (dst *ExportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportRequest.Merge(dst, src)
}
func (m *ExportRequest) XXX_Size() int {
	return xxx_messageInfo_ExportRequest.Size(m)
}
func (m *ExportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportRequest proto.InternalMessageInfo

func (m *ExportRequest) GetWait() bool {
	if m != nil {
		return m.Wait
	}
	return false
}

type ExportResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportResponse) Reset()         { *m = ExportResponse{} }
func (m *ExportResponse) String() string { return proto.CompactTextString(m) }
func (*ExportResponse) ProtoMessage()    {}
func (*ExportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_controlservice_0a1a098ce6ace8f5, []int{11}
}
func (m *ExportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportResponse.Unmarshal(m, b)
}
func (m *ExportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportResponse.Marshal(b, m, deterministic)
}
func (dst *ExportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportResponse.Merge(dst, src)
}
func (m *ExportResponse) XXX_Size() int {
	return xxx_messageInfo_ExportResponse.Size(m)
}
func (m *ExportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportResponse proto.InternalMessageInfo

func (m *ExportResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*LogMessage)(nil), "service.LogMessage")
	proto.RegisterType((*LogResponse)(nil), "service.LogResponse")
	proto.RegisterType((*QuitMessage)(nil), "service.QuitMessage")
	proto.RegisterType((*QuitResponse)(nil), "service.QuitResponse")
	proto.RegisterType((*SwitchPhaseMessage)(nil), "service.SwitchPhaseMessage")
	proto.RegisterType((*SwitchPhaseResponse)(nil), "service.SwitchPhaseResponse")
	proto.RegisterType((*PackageRequest)(nil), "service.PackageRequest")
	proto.RegisterType((*StatusMessage)(nil), "service.StatusMessage")
	proto.RegisterType((*StatusResponse)(nil), "service.StatusResponse")
	proto.RegisterType((*EventMessage)(nil), "service.EventMessage")
	proto.RegisterType((*ExportRequest)(nil), "service.ExportRequest")
	proto.RegisterType((*ExportResponse)(nil), "service.ExportResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ControlServiceClient is the client API for ControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControlServiceClient interface {
	Log(ctx context.Context, in *LogMessage, opts ...grpc.CallOption) (*LogResponse, error)
	Quit(ctx context.Context, in *QuitMessage, opts ...grpc.CallOption) (*QuitResponse, error)
	SwitchPhase(ctx context.Context, in *SwitchPhaseMessage, opts ...grpc.CallOption) (*SwitchPhaseResponse, error)
	GetPackageNode(ctx context.Context, in *PackageRequest, opts ...grpc.CallOption) (*PackageNode, error)
	GetFileNode(ctx context.Context, in *FileNode, opts ...grpc.CallOption) (ControlService_GetFileNodeClient, error)
	GetDiagnosticNode(ctx context.Context, in *DiagnosticNode, opts ...grpc.CallOption) (ControlService_GetDiagnosticNodeClient, error)
	Status(ctx context.Context, in *StatusMessage, opts ...grpc.CallOption) (*StatusResponse, error)
	SubscribeEvents(ctx context.Context, in *EventMessage, opts ...grpc.CallOption) (ControlService_SubscribeEventsClient, error)
	ExportSnapshot(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error)
}

type controlServiceClient struct {
	cc *grpc.ClientConn
}

func NewControlServiceClient(cc *grpc.ClientConn) ControlServiceClient {
	return &controlServiceClient{cc}
}

func (c *controlServiceClient) Log(ctx context.Context, in *LogMessage, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/service.ControlService/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) Quit(ctx context.Context, in *QuitMessage, opts ...grpc.CallOption) (*QuitResponse, error) {
	out := new(QuitResponse)
	err := c.cc.Invoke(ctx, "/service.ControlService/Quit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) SwitchPhase(ctx context.Context, in *SwitchPhaseMessage, opts ...grpc.CallOption) (*SwitchPhaseResponse, error) {
	out := new(SwitchPhaseResponse)
	err := c.cc.Invoke(ctx, "/service.ControlService/SwitchPhase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) GetPackageNode(ctx context.Context, in *PackageRequest, opts ...grpc.CallOption) (*PackageNode, error) {
	out := new(PackageNode)
	err := c.cc.Invoke(ctx, "/service.ControlService/GetPackageNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) GetFileNode(ctx context.Context, in *FileNode, opts ...grpc.CallOption) (ControlService_GetFileNodeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ControlService_serviceDesc.Streams[0], "/service.ControlService/GetFileNode", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlServiceGetFileNodeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ControlService_GetFileNodeClient interface {
	Recv() (*FileNode, error)
	grpc.ClientStream
}

type controlServiceGetFileNodeClient struct {
	grpc.ClientStream
}

func (x *controlServiceGetFileNodeClient) Recv() (*FileNode, error) {
	m := new(FileNode)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlServiceClient) GetDiagnosticNode(ctx context.Context, in *DiagnosticNode, opts ...grpc.CallOption) (ControlService_GetDiagnosticNodeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ControlService_serviceDesc.Streams[1], "/service.ControlService/GetDiagnosticNode", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlServiceGetDiagnosticNodeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ControlService_GetDiagnosticNodeClient interface {
	Recv() (*DiagnosticNode, error)
	grpc.ClientStream
}

type controlServiceGetDiagnosticNodeClient struct {
	grpc.ClientStream
}

func (x *controlServiceGetDiagnosticNodeClient) Recv() (*DiagnosticNode, error) {
	m := new(DiagnosticNode)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlServiceClient) Status(ctx context.Context, in *StatusMessage, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/service.ControlService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) SubscribeEvents(ctx context.Context, in *EventMessage, opts ...grpc.CallOption) (ControlService_SubscribeEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ControlService_serviceDesc.Streams[2], "/service.ControlService/SubscribeEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlServiceSubscribeEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ControlService_SubscribeEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type controlServiceSubscribeEventsClient struct {
	grpc.ClientStream
}

func (x *controlServiceSubscribeEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlServiceClient) ExportSnapshot(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error) {
	out := new(ExportResponse)
	err := c.cc.Invoke(ctx, "/service.ControlService/ExportSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServiceServer is the server API for ControlService service.
type ControlServiceServer interface {
	Log(context.Context, *LogMessage) (*LogResponse, error)
	Quit(context.Context, *QuitMessage) (*QuitResponse, error)
	SwitchPhase(context.Context, *SwitchPhaseMessage) (*SwitchPhaseResponse, error)
	GetPackageNode(context.Context, *PackageRequest) (*PackageNode, error)
	GetFileNode(*FileNode, ControlService_GetFileNodeServer) error
	GetDiagnosticNode(*DiagnosticNode, ControlService_GetDiagnosticNodeServer) error
	Status(context.Context, *StatusMessage) (*StatusResponse, error)
	SubscribeEvents(*EventMessage, ControlService_SubscribeEventsServer) error
	ExportSnapshot(context.Context, *ExportRequest) (*ExportResponse, error)
}

func RegisterControlServiceServer(s *grpc.Server, srv ControlServiceServer) {
	s.RegisterService(&_ControlService_serviceDesc, srv)
}

func _ControlService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ControlService/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).Log(ctx, req.(*LogMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_Quit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuitMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).Quit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ControlService/Quit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).Quit(ctx, req.(*QuitMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_SwitchPhase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwitchPhaseMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).SwitchPhase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ControlService/SwitchPhase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).SwitchPhase(ctx, req.(*SwitchPhaseMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_GetPackageNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).GetPackageNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ControlService/GetPackageNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).GetPackageNode(ctx, req.(*PackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_GetFileNode_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileNode)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControlServiceServer).GetFileNode(m, &controlServiceGetFileNodeServer{stream})
}

type ControlService_GetFileNodeServer interface {
	Send(*FileNode) error
	grpc.ServerStream
}

type controlServiceGetFileNodeServer struct {
	grpc.ServerStream
}

func (x *controlServiceGetFileNodeServer) Send(m *FileNode) error {
	return x.ServerStream.SendMsg(m)
}

func _ControlService_GetDiagnosticNode_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DiagnosticNode)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControlServiceServer).GetDiagnosticNode(m, &controlServiceGetDiagnosticNodeServer{stream})
}

type ControlService_GetDiagnosticNodeServer interface {
	Send(*DiagnosticNode) error
	grpc.ServerStream
}

type controlServiceGetDiagnosticNodeServer struct {
	grpc.ServerStream
}

func (x *controlServiceGetDiagnosticNodeServer) Send(m *DiagnosticNode) error {
	return x.ServerStream.SendMsg(m)
}

func _ControlService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ControlService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).Status(ctx, req.(*StatusMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_SubscribeEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControlServiceServer).SubscribeEvents(m, &controlServiceSubscribeEventsServer{stream})
}

type ControlService_SubscribeEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type controlServiceSubscribeEventsServer struct {
	grpc.ServerStream
}

func (x *controlServiceSubscribeEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _ControlService_ExportSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).ExportSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ControlService/ExportSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).ExportSnapshot(ctx, req.(*ExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.ControlService",
	HandlerType: (*ControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _ControlService_Log_Handler,
		},
		{
			MethodName: "Quit",
			Handler:    _ControlService_Quit_Handler,
		},
		{
			MethodName: "SwitchPhase",
			Handler:    _ControlService_SwitchPhase_Handler,
		},
		{
			MethodName: "GetPackageNode",
			Handler:    _ControlService_GetPackageNode_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _ControlService_Status_Handler,
		},
		{
			MethodName: "ExportSnapshot",
			Handler:    _ControlService_ExportSnapshot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFileNode",
			Handler:       _ControlService_GetFileNode_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetDiagnosticNode",
			Handler:       _ControlService_GetDiagnosticNode_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeEvents",
			Handler:       _ControlService_SubscribeEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "controlservice.proto",
}

func init() {
	proto.RegisterFile("controlservice.proto", fileDescriptor_controlservice_0a1a098ce6ace8f5)
}

var fileDescriptor_controlservice_0a1a098ce6ace8f5 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x75, 0xc8, 0xf7, 0x24, 0x75, 0xdb, 0x6d, 0x1a, 0x22, 0x53, 0xa1, 0xb2, 0x20, 0x11, 0x38,
	0x44, 0x28, 0x08, 0x21, 0xbe, 0x0e, 0x90, 0x86, 0xa8, 0x28, 0xa0, 0xe2, 0x5c, 0xb8, 0x6e, 0x9c,
	0x95, 0x63, 0xd5, 0xf1, 0xba, 0xbb, 0x9b, 0x96, 0x1f, 0xc0, 0xef, 0xe3, 0x37, 0x21, 0xaf, 0xbd,
	0xf6, 0x9a, 0x80, 0xda, 0xdb, 0xce, 0xcc, 0x9b, 0xf7, 0x66, 0x3c, 0x4f, 0x86, 0x9e, 0xc7, 0x22,
	0xc9, 0x59, 0x28, 0x28, 0xbf, 0x0e, 0x3c, 0x3a, 0x8a, 0x39, 0x93, 0x0c, 0x35, 0xb3, 0xd0, 0xd9,
	0x5f, 0x11, 0x49, 0x36, 0x6c, 0x45, 0xc3, 0xb4, 0x82, 0x1f, 0x02, 0xcc, 0x99, 0xff, 0x95, 0x0a,
	0x41, 0x7c, 0x8a, 0x0e, 0xa0, 0xba, 0x11, 0xfe, 0xa0, 0x72, 0x5a, 0x19, 0x76, 0xdd, 0xe4, 0x89,
	0x9f, 0x42, 0x67, 0xce, 0x7c, 0x97, 0x8a, 0x98, 0x45, 0x82, 0xa2, 0x01, 0x34, 0xc5, 0xd6, 0xf3,
	0xa8, 0x10, 0x0a, 0xd4, 0x72, 0x75, 0x88, 0x1f, 0x41, 0xe7, 0xfb, 0x36, 0x90, 0x9a, 0x09, 0x41,
	0xed, 0x32, 0x08, 0xc3, 0x0c, 0xa5, 0xde, 0x78, 0x08, 0xdd, 0x04, 0x72, 0x07, 0xb2, 0xb7, 0x80,
	0x16, 0x37, 0x81, 0xf4, 0xd6, 0x17, 0x6b, 0x22, 0xa8, 0xe6, 0x7c, 0x02, 0xf5, 0x38, 0x89, 0x15,
	0xda, 0x1e, 0xdb, 0x23, 0xbd, 0xa4, 0x42, 0xb9, 0x69, 0x11, 0x4f, 0xe1, 0xc8, 0xe8, 0xbd, 0x5d,
	0x0c, 0xf5, 0xa0, 0x4e, 0x39, 0x67, 0x7c, 0x70, 0xef, 0xb4, 0x32, 0x6c, 0xbb, 0x69, 0x80, 0x0f,
	0xc0, 0xbe, 0x20, 0xde, 0x25, 0xf1, 0xa9, 0x4b, 0xaf, 0xb6, 0x54, 0x48, 0xfc, 0x01, 0xf6, 0x16,
	0x92, 0xc8, 0xad, 0xd0, 0xf3, 0xf4, 0xcc, 0x79, 0x5a, 0x99, 0x3e, 0xea, 0x43, 0x43, 0x28, 0x7d,
	0xc5, 0xd7, 0x72, 0xb3, 0x08, 0xff, 0xaa, 0x80, 0x9d, 0xf6, 0xe7, 0x33, 0x95, 0x08, 0xda, 0x9a,
	0x60, 0x08, 0x4d, 0xf5, 0x38, 0x3f, 0x53, 0x0c, 0xbb, 0x8b, 0xea, 0x32, 0x3a, 0x81, 0x76, 0x4a,
	0x1e, 0x44, 0xfe, 0xa0, 0xaa, 0xd4, 0x8a, 0x44, 0xb1, 0x57, 0xcd, 0xdc, 0xeb, 0x0d, 0x74, 0xa7,
	0xd7, 0x34, 0xca, 0x0f, 0xf5, 0x0c, 0xea, 0x5e, 0x48, 0xb2, 0xaf, 0x62, 0x8f, 0x8f, 0x72, 0x2d,
	0x85, 0x9a, 0x24, 0x25, 0x37, 0x45, 0xe0, 0xc7, 0xb0, 0x37, 0xfd, 0x19, 0x33, 0x2e, 0xb3, 0x2f,
	0x92, 0x1c, 0xf9, 0x86, 0x04, 0x52, 0x1f, 0x39, 0x79, 0xe3, 0xe7, 0x60, 0x6b, 0xd0, 0x6d, 0x5f,
	0x7e, 0xfc, 0xbb, 0x06, 0xf6, 0x24, 0xf5, 0xeb, 0x22, 0x55, 0x45, 0x63, 0xa8, 0xce, 0x99, 0x8f,
	0x8a, 0x31, 0x0a, 0x77, 0x3a, 0x3d, 0x33, 0xa9, 0xe9, 0xb1, 0x85, 0x5e, 0x41, 0x2d, 0xf1, 0x15,
	0x2a, 0xea, 0x86, 0x13, 0x9d, 0xe3, 0x52, 0xd6, 0x68, 0xfb, 0x02, 0x1d, 0xc3, 0x28, 0xe8, 0x41,
	0x8e, 0xdb, 0xb5, 0x9e, 0x73, 0xf2, 0xaf, 0xa2, 0xc1, 0xf5, 0x11, 0xec, 0x19, 0x95, 0x99, 0x61,
	0xbe, 0xb1, 0x15, 0x45, 0xf7, 0x8b, 0xa3, 0x95, 0x6c, 0x64, 0x6c, 0x61, 0xc0, 0xb1, 0x85, 0x5e,
	0x43, 0x67, 0x46, 0xe5, 0xe7, 0x20, 0x4c, 0xfb, 0x0f, 0x73, 0x98, 0x4e, 0x39, 0xbb, 0x29, 0x6c,
	0xbd, 0xa8, 0xa0, 0x73, 0x38, 0x9c, 0x51, 0x79, 0x16, 0x10, 0x3f, 0x62, 0x42, 0x06, 0xde, 0x5f,
	0xf2, 0xe5, 0x82, 0xf3, 0xbf, 0x82, 0xa2, 0x7a, 0x07, 0x8d, 0xd4, 0xa2, 0xa8, 0x5f, 0x2c, 0x6c,
	0x7a, 0xde, 0x68, 0x2f, 0x7b, 0x19, 0x5b, 0xe8, 0x3d, 0xec, 0x2f, 0xb6, 0x4b, 0xe1, 0xf1, 0x60,
	0x49, 0x95, 0x79, 0x04, 0x3a, 0x2e, 0xbb, 0x49, 0x93, 0xd8, 0xe5, 0xb4, 0x92, 0x9e, 0x68, 0xdf,
	0x2c, 0x22, 0x12, 0x8b, 0x35, 0x93, 0xc6, 0x08, 0x25, 0xd7, 0x19, 0x23, 0x94, 0x8d, 0x86, 0xad,
	0x4f, 0x03, 0xe8, 0x33, 0xee, 0x8f, 0xae, 0x36, 0x42, 0xf2, 0x91, 0xcf, 0x63, 0x4f, 0x43, 0x7f,
	0x58, 0xcb, 0x86, 0xfa, 0xe1, 0xbd, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xd0, 0x1b, 0xe6,
	0x22, 0x05, 0x00, 0x00,
}