// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/tinyci/ci-agents/ci-gen/grpc/types/submission.proto

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

type Submission struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	HeadRef              *Ref                 `protobuf:"bytes,2,opt,name=headRef,proto3" json:"headRef,omitempty"`
	BaseRef              *Ref                 `protobuf:"bytes,3,opt,name=baseRef,proto3" json:"baseRef,omitempty"`
	User                 *User                `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Submission) Reset()         { *m = Submission{} }
func (m *Submission) String() string { return proto.CompactTextString(m) }
func (*Submission) ProtoMessage()    {}
func (*Submission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f00c03eb97056e8, []int{0}
}

func (m *Submission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Submission.Unmarshal(m, b)
}
func (m *Submission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Submission.Marshal(b, m, deterministic)
}
func (m *Submission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Submission.Merge(m, src)
}
func (m *Submission) XXX_Size() int {
	return xxx_messageInfo_Submission.Size(m)
}
func (m *Submission) XXX_DiscardUnknown() {
	xxx_messageInfo_Submission.DiscardUnknown(m)
}

var xxx_messageInfo_Submission proto.InternalMessageInfo

func (m *Submission) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Submission) GetHeadRef() *Ref {
	if m != nil {
		return m.HeadRef
	}
	return nil
}

func (m *Submission) GetBaseRef() *Ref {
	if m != nil {
		return m.BaseRef
	}
	return nil
}

func (m *Submission) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Submission) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Submission)(nil), "types.Submission")
}

func init() {
	proto.RegisterFile("github.com/tinyci/ci-agents/ci-gen/grpc/types/submission.proto", fileDescriptor_9f00c03eb97056e8)
}

var fileDescriptor_9f00c03eb97056e8 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x69, 0xb7, 0x29, 0x66, 0xe0, 0xa1, 0xa7, 0xd2, 0x8b, 0x43, 0x3c, 0xec, 0xb2, 0x04,
	0xf4, 0xe0, 0x40, 0x10, 0xf4, 0x23, 0x44, 0xbd, 0x78, 0x4b, 0xd2, 0x7f, 0xb3, 0xc0, 0x9a, 0x84,
	0xe4, 0xdf, 0xc3, 0x3e, 0x9f, 0x5f, 0x4c, 0x9a, 0xb6, 0x4e, 0xd8, 0xa9, 0xb7, 0x47, 0xde, 0xef,
	0xbd, 0x24, 0x8f, 0xbc, 0x6a, 0x83, 0x87, 0x4e, 0x52, 0xe5, 0x5a, 0x86, 0xc6, 0x9e, 0x94, 0x61,
	0xca, 0xec, 0x84, 0x06, 0x8b, 0xb1, 0x57, 0x1a, 0x2c, 0xd3, 0xc1, 0x2b, 0x86, 0x27, 0x0f, 0x91,
	0xc5, 0x4e, 0xb6, 0x26, 0x46, 0xe3, 0x2c, 0xf5, 0xc1, 0xa1, 0x2b, 0x56, 0xe9, 0xbc, 0x7a, 0xf9,
	0x57, 0xa3, 0xdd, 0x51, 0x58, 0xcd, 0x92, 0x2f, 0xbb, 0x86, 0xf9, 0x21, 0x8a, 0xa6, 0x85, 0x88,
	0xa2, 0xf5, 0x67, 0x35, 0x74, 0x54, 0xcf, 0xf3, 0xde, 0x10, 0xa0, 0x19, 0x83, 0xfb, 0x79, 0xc1,
	0x2e, 0x42, 0x18, 0x92, 0xf7, 0x3f, 0x19, 0x21, 0x1f, 0x7f, 0x7f, 0x29, 0x6e, 0x49, 0x6e, 0xea,
	0x32, 0xdb, 0x64, 0xdb, 0x05, 0xcf, 0x4d, 0x5d, 0x3c, 0x90, 0xeb, 0x03, 0x88, 0x9a, 0x43, 0x53,
	0xe6, 0x9b, 0x6c, 0xbb, 0x7e, 0x24, 0x34, 0x55, 0x50, 0x0e, 0x0d, 0x9f, 0xac, 0x9e, 0x92, 0x22,
	0x42, 0x4f, 0x2d, 0x2e, 0xa9, 0xd1, 0x2a, 0xee, 0xc8, 0xb2, 0xbf, 0xb8, 0x5c, 0x26, 0x64, 0x3d,
	0x22, 0x5f, 0x11, 0x02, 0x4f, 0x46, 0xb1, 0x27, 0x37, 0x2a, 0x80, 0x40, 0xa8, 0xdf, 0xb0, 0x5c,
	0x25, 0xaa, 0xa2, 0xda, 0x39, 0x7d, 0x04, 0x3a, 0x8d, 0x48, 0x3f, 0xa7, 0xcd, 0xf8, 0x19, 0x7e,
	0x67, 0xdf, 0xbb, 0x59, 0x0b, 0xc8, 0xab, 0xd4, 0xf7, 0xf4, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x90,
	0xed, 0x3f, 0xe1, 0xf6, 0x01, 0x00, 0x00,
}
