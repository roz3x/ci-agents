// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/tinyci/ci-agents/grpc/types/session.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Session is a session key/value pair as well as an expiry time for the
// session. All values are encrypted with the session key.
type Session struct {
	Key                  string               `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               string               `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
	ExpiresOn            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expiresOn,proto3" json:"expiresOn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c4dfb5203e7e055, []int{0}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Session) GetValues() string {
	if m != nil {
		return m.Values
	}
	return ""
}

func (m *Session) GetExpiresOn() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresOn
	}
	return nil
}

func init() {
	proto.RegisterType((*Session)(nil), "types.Session")
}

func init() {
	proto.RegisterFile("github.com/tinyci/ci-agents/grpc/types/session.proto", fileDescriptor_9c4dfb5203e7e055)
}

var fileDescriptor_9c4dfb5203e7e055 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xc9, 0xcc, 0xab, 0x4c, 0xce, 0xd4, 0x4f, 0xce,
	0xd4, 0x4d, 0x4c, 0x4f, 0xcd, 0x2b, 0x29, 0xd6, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0xa9, 0x2c,
	0x48, 0x2d, 0xd6, 0x2f, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x05, 0x0b, 0x4a, 0x59, 0x23, 0x69, 0x4e, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x07,
	0xcb, 0x27, 0x95, 0xa6, 0xe9, 0x17, 0x40, 0xf4, 0x95, 0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6,
	0x16, 0x20, 0x58, 0x10, 0x33, 0x94, 0x72, 0xb9, 0xd8, 0x83, 0x21, 0x86, 0x0a, 0x09, 0x70, 0x31,
	0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x62, 0x5c, 0x6c,
	0x65, 0x89, 0x39, 0xa5, 0xa9, 0xc5, 0x12, 0x4c, 0x60, 0x41, 0x28, 0x4f, 0xc8, 0x82, 0x8b, 0x33,
	0xb5, 0xa2, 0x20, 0xb3, 0x28, 0xb5, 0xd8, 0x3f, 0x4f, 0x82, 0x59, 0x81, 0x51, 0x83, 0xdb, 0x48,
	0x4a, 0x2f, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x0f, 0x66, 0xb5, 0x5e, 0x08, 0xcc, 0xa6, 0x20,
	0x84, 0x62, 0x27, 0x8d, 0x28, 0x35, 0xe2, 0xbc, 0x9a, 0xc4, 0x06, 0x36, 0xc8, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x1d, 0x96, 0x75, 0x65, 0x1b, 0x01, 0x00, 0x00,
}
