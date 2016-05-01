// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	service.proto
	user.proto

It has these top-level messages:
	ServiceInfo
	ServiceInfos
	Notification
	Notifications
	NullMsg
	User
*/
package model

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
const _ = proto.ProtoPackageIsVersion1

type ServiceInfo struct {
	WorkerAddr string `protobuf:"bytes,1,opt,name=WorkerAddr,json=workerAddr" json:"WorkerAddr,omitempty"`
	Version    string `protobuf:"bytes,2,opt,name=Version,json=version" json:"Version,omitempty"`
}

func (m *ServiceInfo) Reset()                    { *m = ServiceInfo{} }
func (m *ServiceInfo) String() string            { return proto.CompactTextString(m) }
func (*ServiceInfo) ProtoMessage()               {}
func (*ServiceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServiceInfos struct {
	ServiceInfo []*ServiceInfo `protobuf:"bytes,1,rep,name=ServiceInfo,json=serviceInfo" json:"ServiceInfo,omitempty"`
}

func (m *ServiceInfos) Reset()                    { *m = ServiceInfos{} }
func (m *ServiceInfos) String() string            { return proto.CompactTextString(m) }
func (*ServiceInfos) ProtoMessage()               {}
func (*ServiceInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceInfos) GetServiceInfo() []*ServiceInfo {
	if m != nil {
		return m.ServiceInfo
	}
	return nil
}

type Notification struct {
	Error int32  `protobuf:"varint,1,opt,name=Error,json=error" json:"Error,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=Msg,json=msg" json:"Msg,omitempty"`
}

func (m *Notification) Reset()                    { *m = Notification{} }
func (m *Notification) String() string            { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()               {}
func (*Notification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Notifications struct {
	Notification []*Notification `protobuf:"bytes,1,rep,name=Notification,json=notification" json:"Notification,omitempty"`
}

func (m *Notifications) Reset()                    { *m = Notifications{} }
func (m *Notifications) String() string            { return proto.CompactTextString(m) }
func (*Notifications) ProtoMessage()               {}
func (*Notifications) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Notifications) GetNotification() []*Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type NullMsg struct {
}

func (m *NullMsg) Reset()                    { *m = NullMsg{} }
func (m *NullMsg) String() string            { return proto.CompactTextString(m) }
func (*NullMsg) ProtoMessage()               {}
func (*NullMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*ServiceInfo)(nil), "model.ServiceInfo")
	proto.RegisterType((*ServiceInfos)(nil), "model.ServiceInfos")
	proto.RegisterType((*Notification)(nil), "model.Notification")
	proto.RegisterType((*Notifications)(nil), "model.Notifications")
	proto.RegisterType((*NullMsg)(nil), "model.NullMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for RegistrationService service

type RegistrationServiceClient interface {
	// ListAllWorkerServices lists all available worker services.
	// Errors:
	ListAllWorkerServices(ctx context.Context, in *NullMsg, opts ...grpc.CallOption) (*ServiceInfos, error)
}

type registrationServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationServiceClient(cc *grpc.ClientConn) RegistrationServiceClient {
	return &registrationServiceClient{cc}
}

func (c *registrationServiceClient) ListAllWorkerServices(ctx context.Context, in *NullMsg, opts ...grpc.CallOption) (*ServiceInfos, error) {
	out := new(ServiceInfos)
	err := grpc.Invoke(ctx, "/model.RegistrationService/ListAllWorkerServices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RegistrationService service

type RegistrationServiceServer interface {
	// ListAllWorkerServices lists all available worker services.
	// Errors:
	ListAllWorkerServices(context.Context, *NullMsg) (*ServiceInfos, error)
}

func RegisterRegistrationServiceServer(s *grpc.Server, srv RegistrationServiceServer) {
	s.RegisterService(&_RegistrationService_serviceDesc, srv)
}

func _RegistrationService_ListAllWorkerServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(NullMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RegistrationServiceServer).ListAllWorkerServices(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _RegistrationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.RegistrationService",
	HandlerType: (*RegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAllWorkerServices",
			Handler:    _RegistrationService_ListAllWorkerServices_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for NotificationService service

type NotificationServiceClient interface {
	GetUnreadNotification(ctx context.Context, in *NullMsg, opts ...grpc.CallOption) (*Notifications, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) GetUnreadNotification(ctx context.Context, in *NullMsg, opts ...grpc.CallOption) (*Notifications, error) {
	out := new(Notifications)
	err := grpc.Invoke(ctx, "/model.NotificationService/GetUnreadNotification", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NotificationService service

type NotificationServiceServer interface {
	GetUnreadNotification(context.Context, *NullMsg) (*Notifications, error)
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_GetUnreadNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(NullMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NotificationServiceServer).GetUnreadNotification(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUnreadNotification",
			Handler:    _NotificationService_GetUnreadNotification_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xad, 0x21, 0x86, 0x4e, 0x52, 0x91, 0x49, 0x0b, 0xc1, 0x83, 0xc8, 0x9e, 0x7a, 0xca,
	0x21, 0x8a, 0xde, 0x94, 0x82, 0x52, 0x05, 0xed, 0x21, 0xd2, 0x7a, 0xae, 0xcd, 0x36, 0x04, 0xd3,
	0xac, 0xcc, 0xc4, 0xfa, 0xf7, 0x8d, 0xd9, 0xad, 0x6e, 0xd4, 0xe3, 0x7b, 0x8f, 0xfd, 0xde, 0xcc,
	0x2c, 0x0c, 0x58, 0xd2, 0xb6, 0x58, 0xc9, 0xf8, 0x8d, 0x54, 0xad, 0xd0, 0xdd, 0xa8, 0x4c, 0x96,
	0x62, 0x0a, 0xfe, 0x93, 0xf6, 0xef, 0xab, 0xb5, 0xc2, 0x13, 0x80, 0x67, 0x45, 0xaf, 0x92, 0x26,
	0x59, 0x46, 0x51, 0xef, 0xb4, 0x37, 0xee, 0xa7, 0xf0, 0xf1, 0xed, 0x60, 0x04, 0xde, 0x42, 0x12,
	0x17, 0xaa, 0x8a, 0xf6, 0xdb, 0xd0, 0xdb, 0x6a, 0x29, 0x6e, 0x20, 0xb0, 0x40, 0x8c, 0xe7, 0x1d,
	0x70, 0x83, 0x72, 0xc6, 0x7e, 0x82, 0x71, 0xdb, 0x1a, 0x5b, 0x49, 0xea, 0xf3, 0x8f, 0x10, 0x17,
	0x10, 0xcc, 0x54, 0x5d, 0xac, 0x8b, 0xd5, 0xb2, 0x6e, 0xa8, 0x38, 0x04, 0xf7, 0x96, 0x48, 0xe9,
	0x51, 0xdc, 0xd4, 0x95, 0x5f, 0x02, 0x8f, 0xc0, 0x79, 0xe4, 0xdc, 0x4c, 0xe0, 0x6c, 0x38, 0x17,
	0x77, 0x30, 0xb0, 0xdf, 0x31, 0x5e, 0x76, 0x41, 0xa6, 0x3f, 0x34, 0xfd, 0x76, 0x94, 0x06, 0x95,
	0xa5, 0x44, 0x1f, 0xbc, 0xd9, 0x7b, 0x59, 0x36, 0xfc, 0x64, 0x0e, 0x61, 0x2a, 0xf3, 0x82, 0x6b,
	0x6a, 0x23, 0x33, 0x34, 0x5e, 0xc1, 0xe8, 0xa1, 0x31, 0x27, 0x65, 0xa9, 0x4f, 0x65, 0x7c, 0xc6,
	0xc3, 0x1d, 0x5d, 0xbf, 0x3f, 0x0e, 0xff, 0x6e, 0xcb, 0x62, 0x2f, 0x59, 0x40, 0x68, 0xf7, 0xef,
	0xb0, 0xd7, 0x30, 0x9a, 0xca, 0x7a, 0x5e, 0x91, 0x5c, 0x66, 0x9d, 0x1b, 0xfc, 0xc6, 0x0e, 0xff,
	0x59, 0xa2, 0xe1, 0xbe, 0x1c, 0xb4, 0x1f, 0x7b, 0xf6, 0x19, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x72,
	0x57, 0x43, 0xe9, 0x01, 0x00, 0x00,
}