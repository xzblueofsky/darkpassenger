// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	// Static info
	Id               int64  `protobuf:"varint,1,opt,name=Id,json=id" json:"Id,omitempty"`
	Email            string `protobuf:"bytes,2,opt,name=Email,json=email" json:"Email,omitempty"`
	PwdHash          string `protobuf:"bytes,3,opt,name=PwdHash,json=pwdHash" json:"PwdHash,omitempty"`
	MaxConn          int32  `protobuf:"varint,4,opt,name=MaxConn,json=maxConn" json:"MaxConn,omitempty"`
	CreatedTimestamp int64  `protobuf:"varint,5,opt,name=CreatedTimestamp,json=createdTimestamp" json:"CreatedTimestamp,omitempty"`
	ExpiredTimestamp int64  `protobuf:"varint,6,opt,name=ExpiredTimestamp,json=expiredTimestamp" json:"ExpiredTimestamp,omitempty"`
	// Dynamic info
	LoginTimestamp int64  `protobuf:"varint,128,opt,name=LoginTimestamp,json=loginTimestamp" json:"LoginTimestamp,omitempty"`
	Traffic        int64  `protobuf:"varint,129,opt,name=Traffic,json=traffic" json:"Traffic,omitempty"`
	IP             string `protobuf:"bytes,130,opt,name=IP,json=iP" json:"IP,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*User)(nil), "model.User")
}

var fileDescriptor1 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x40, 0xd9, 0xdc, 0xe5, 0x82, 0x53, 0xac, 0xc7, 0x60, 0x31, 0x76, 0x62, 0xa3, 0x58, 0xd8,
	0xf8, 0x09, 0xc7, 0x81, 0x82, 0xc2, 0xb2, 0x9c, 0x1f, 0x10, 0x2f, 0x39, 0x0d, 0x6c, 0x36, 0x21,
	0x1b, 0x71, 0x4b, 0xb5, 0xf7, 0x9f, 0xcd, 0x26, 0x85, 0xde, 0x96, 0xef, 0xcd, 0x83, 0x64, 0x06,
	0xe0, 0x7d, 0xd0, 0xe1, 0xd6, 0x07, 0x17, 0x1d, 0x72, 0xeb, 0x94, 0xee, 0x2e, 0x7f, 0x18, 0x2c,
	0x9f, 0x93, 0xc5, 0x1a, 0xd8, 0x83, 0xa2, 0xea, 0xa2, 0xba, 0x5e, 0xb4, 0xcc, 0x28, 0x3c, 0x03,
	0xbe, 0xb5, 0xd2, 0x74, 0xc4, 0x92, 0x3a, 0x69, 0xb9, 0x9e, 0x00, 0x09, 0x44, 0xf3, 0xa1, 0xee,
	0xe5, 0xf0, 0x46, 0x8b, 0xec, 0x85, 0x2f, 0x38, 0x4d, 0x9e, 0xe4, 0xb8, 0x71, 0x7d, 0x4f, 0xcb,
	0x34, 0xe1, 0xad, 0xb0, 0x05, 0xf1, 0x06, 0xd6, 0x9b, 0xa0, 0x65, 0xd4, 0x6a, 0x67, 0xac, 0x1e,
	0xa2, 0xb4, 0x9e, 0x78, 0x7e, 0x67, 0xbd, 0x9f, 0xf9, 0xa9, 0xdd, 0x8e, 0xde, 0x84, 0xff, 0xed,
	0xaa, 0xb4, 0x7a, 0xe6, 0xf1, 0x0a, 0xea, 0x47, 0xf7, 0x6a, 0xfa, 0xbf, 0xf2, 0xb3, 0x7c, 0xbf,
	0xee, 0x8e, 0x34, 0x9e, 0x83, 0xd8, 0x05, 0x79, 0x38, 0x98, 0x3d, 0x7d, 0x95, 0x42, 0xc4, 0xc2,
	0x78, 0x9a, 0xb6, 0x6e, 0xe8, 0xbb, 0xca, 0xbb, 0x30, 0xd3, 0xbc, 0xac, 0xf2, 0x75, 0xee, 0x7e,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x12, 0x3c, 0x49, 0x16, 0x2b, 0x01, 0x00, 0x00,
}
