// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/tinyci/ci-agents/ci-gen/grpc/types/session.proto

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
	return fileDescriptor_a2e55cd2a1053d6f, []int{0}
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
	proto.RegisterFile("github.com/tinyci/ci-agents/ci-gen/grpc/types/session.proto", fileDescriptor_a2e55cd2a1053d6f)
}

var fileDescriptor_a2e55cd2a1053d6f = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0xbf, 0x4b, 0xc5, 0x30,
	0x10, 0xc7, 0xa9, 0x0f, 0x9f, 0xbc, 0xb8, 0x48, 0x06, 0x29, 0x9d, 0x8a, 0x53, 0x97, 0xe6, 0x40,
	0x17, 0xc1, 0xcd, 0x7f, 0x40, 0xa8, 0x4e, 0x6e, 0x69, 0x38, 0xcf, 0x60, 0xf3, 0x83, 0x5e, 0x2a,
	0xf6, 0xbf, 0x97, 0xa6, 0x96, 0xba, 0xbe, 0xed, 0x93, 0x6f, 0xee, 0x3e, 0x5f, 0x4e, 0x3c, 0x91,
	0x4d, 0x9f, 0x53, 0xaf, 0x4c, 0x70, 0x90, 0xac, 0x9f, 0x8d, 0x05, 0x63, 0x5b, 0x4d, 0xe8, 0x13,
	0x2f, 0x44, 0xe8, 0x81, 0xc6, 0x68, 0x20, 0xcd, 0x11, 0x19, 0x18, 0x99, 0x6d, 0xf0, 0x2a, 0x8e,
	0x21, 0x05, 0x79, 0x99, 0xc3, 0xea, 0xbf, 0x83, 0xc2, 0xa0, 0x3d, 0x41, 0xfe, 0xef, 0xa7, 0x0f,
	0x88, 0xeb, 0x5e, 0xb2, 0x0e, 0x39, 0x69, 0x17, 0x77, 0x5a, 0x1d, 0x77, 0x4e, 0x5c, 0xbd, 0xae,
	0x52, 0x79, 0x23, 0x0e, 0x5f, 0x38, 0x97, 0x45, 0x5d, 0x34, 0xa7, 0x6e, 0x41, 0x79, 0x2b, 0x8e,
	0xdf, 0x7a, 0x98, 0x90, 0xcb, 0x8b, 0x1c, 0xfe, 0xbd, 0xe4, 0xa3, 0x38, 0xe1, 0x4f, 0xb4, 0x23,
	0xf2, 0x8b, 0x2f, 0x0f, 0x75, 0xd1, 0x5c, 0xdf, 0x57, 0x8a, 0x42, 0xa0, 0x01, 0xd5, 0x56, 0xad,
	0xde, 0xb6, 0xa6, 0x6e, 0x1f, 0x7e, 0x86, 0xf7, 0xf6, 0xac, 0x8b, 0xfb, 0x63, 0xf6, 0x3d, 0xfc,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x5a, 0x19, 0xd8, 0x29, 0x01, 0x00, 0x00,
}
