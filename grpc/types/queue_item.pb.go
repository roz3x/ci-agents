// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/tinyci/ci-agents/grpc/types/queue_item.proto

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

// QueueItems are the subject sent to runners when runners are able to execute
// a job. Runners poll for these endless through the queuesvc.
type QueueItem struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Running              bool                 `protobuf:"varint,2,opt,name=running,proto3" json:"running,omitempty"`
	RunningOn            string               `protobuf:"bytes,3,opt,name=runningOn,proto3" json:"runningOn,omitempty"`
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	QueueName            string               `protobuf:"bytes,5,opt,name=queueName,proto3" json:"queueName,omitempty"`
	Run                  *Run                 `protobuf:"bytes,6,opt,name=run,proto3" json:"run,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueueItem) Reset()         { *m = QueueItem{} }
func (m *QueueItem) String() string { return proto.CompactTextString(m) }
func (*QueueItem) ProtoMessage()    {}
func (*QueueItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ac4659e424e803, []int{0}
}

func (m *QueueItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueItem.Unmarshal(m, b)
}
func (m *QueueItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueItem.Marshal(b, m, deterministic)
}
func (m *QueueItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueItem.Merge(m, src)
}
func (m *QueueItem) XXX_Size() int {
	return xxx_messageInfo_QueueItem.Size(m)
}
func (m *QueueItem) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueItem.DiscardUnknown(m)
}

var xxx_messageInfo_QueueItem proto.InternalMessageInfo

func (m *QueueItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QueueItem) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

func (m *QueueItem) GetRunningOn() string {
	if m != nil {
		return m.RunningOn
	}
	return ""
}

func (m *QueueItem) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *QueueItem) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *QueueItem) GetRun() *Run {
	if m != nil {
		return m.Run
	}
	return nil
}

// QueueRequest is issued by runners to the queuesvc.
type QueueRequest struct {
	QueueName            string   `protobuf:"bytes,1,opt,name=queueName,proto3" json:"queueName,omitempty"`
	RunningOn            string   `protobuf:"bytes,2,opt,name=runningOn,proto3" json:"runningOn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueRequest) Reset()         { *m = QueueRequest{} }
func (m *QueueRequest) String() string { return proto.CompactTextString(m) }
func (*QueueRequest) ProtoMessage()    {}
func (*QueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ac4659e424e803, []int{1}
}

func (m *QueueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueRequest.Unmarshal(m, b)
}
func (m *QueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueRequest.Marshal(b, m, deterministic)
}
func (m *QueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueRequest.Merge(m, src)
}
func (m *QueueRequest) XXX_Size() int {
	return xxx_messageInfo_QueueRequest.Size(m)
}
func (m *QueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueueRequest proto.InternalMessageInfo

func (m *QueueRequest) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *QueueRequest) GetRunningOn() string {
	if m != nil {
		return m.RunningOn
	}
	return ""
}

// Status is reported to the queuesvc on completion of a run.
type Status struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               bool     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	AdditionalMessage    string   `protobuf:"bytes,3,opt,name=additionalMessage,proto3" json:"additionalMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ac4659e424e803, []int{2}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Status) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Status) GetAdditionalMessage() string {
	if m != nil {
		return m.AdditionalMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*QueueItem)(nil), "types.QueueItem")
	proto.RegisterType((*QueueRequest)(nil), "types.QueueRequest")
	proto.RegisterType((*Status)(nil), "types.Status")
}

func init() {
	proto.RegisterFile("github.com/tinyci/ci-agents/grpc/types/queue_item.proto", fileDescriptor_83ac4659e424e803)
}

var fileDescriptor_83ac4659e424e803 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4f, 0xeb, 0x30,
	0x10, 0xc6, 0x95, 0xf4, 0x35, 0xef, 0xc5, 0xef, 0xe9, 0x49, 0x64, 0x40, 0x56, 0xd5, 0x21, 0xea,
	0x80, 0x32, 0x80, 0x8d, 0x60, 0x00, 0x89, 0x09, 0x36, 0x90, 0x00, 0x61, 0x98, 0x18, 0x40, 0x6e,
	0x72, 0x18, 0x4b, 0x8d, 0x9d, 0xc6, 0xe7, 0xa1, 0x7f, 0x22, 0xff, 0x15, 0xaa, 0x93, 0xd2, 0xaa,
	0x5d, 0xba, 0xdd, 0x7d, 0x97, 0xfb, 0xe5, 0xbe, 0xcf, 0xe4, 0x42, 0x69, 0xfc, 0xf4, 0x53, 0x56,
	0xda, 0x9a, 0xa3, 0x36, 0x8b, 0x52, 0xf3, 0x52, 0x9f, 0x48, 0x05, 0x06, 0x1d, 0x57, 0x6d, 0x53,
	0x72, 0x5c, 0x34, 0xe0, 0xf8, 0xdc, 0x83, 0x87, 0x77, 0x8d, 0x50, 0xb3, 0xa6, 0xb5, 0x68, 0xb3,
	0x61, 0xd0, 0x47, 0x57, 0x1b, 0xfb, 0xca, 0xce, 0xa4, 0x51, 0x3c, 0xcc, 0xa7, 0xfe, 0x83, 0x37,
	0xdd, 0x2a, 0xea, 0x1a, 0x1c, 0xca, 0xba, 0x59, 0x57, 0x1d, 0x63, 0x74, 0xba, 0xe7, 0xcf, 0x5b,
	0x6f, 0xba, 0x8d, 0xc9, 0x57, 0x44, 0xd2, 0xa7, 0xe5, 0x29, 0xb7, 0x08, 0x75, 0xf6, 0x9f, 0xc4,
	0xba, 0xa2, 0x51, 0x1e, 0x15, 0x03, 0x11, 0xeb, 0x2a, 0xa3, 0xe4, 0x77, 0xeb, 0x8d, 0xd1, 0x46,
	0xd1, 0x38, 0x8f, 0x8a, 0x3f, 0x62, 0xd5, 0x66, 0x63, 0x92, 0xf6, 0xe5, 0xa3, 0xa1, 0x83, 0x3c,
	0x2a, 0x52, 0xb1, 0x16, 0xb2, 0x4b, 0x92, 0x3a, 0x94, 0x2d, 0x42, 0x75, 0x8d, 0xf4, 0x57, 0x1e,
	0x15, 0x7f, 0xcf, 0x46, 0x4c, 0x59, 0xab, 0x66, 0xc0, 0x56, 0x6e, 0xd8, 0xcb, 0xea, 0x78, 0xb1,
	0xfe, 0x78, 0xc9, 0x0d, 0xc9, 0x3c, 0xc8, 0x1a, 0xe8, 0xb0, 0xe3, 0xfe, 0x08, 0xd9, 0x98, 0x0c,
	0x5a, 0x6f, 0x68, 0x12, 0x88, 0x84, 0x05, 0x33, 0x4c, 0x78, 0x23, 0x96, 0xf2, 0xe4, 0x8e, 0xfc,
	0x0b, 0x56, 0x04, 0xcc, 0x3d, 0xb8, 0x2d, 0x56, 0xb4, 0xcb, 0xda, 0x70, 0x10, 0x6f, 0x39, 0x98,
	0xbc, 0x91, 0xe4, 0x19, 0x25, 0x7a, 0xb7, 0x93, 0xc9, 0x21, 0x49, 0x5c, 0x98, 0xf4, 0x91, 0xf4,
	0x5d, 0x76, 0x4c, 0x0e, 0x64, 0x55, 0x69, 0xd4, 0xd6, 0xc8, 0xd9, 0x3d, 0x38, 0x27, 0x15, 0xf4,
	0xc9, 0xec, 0x0e, 0x6e, 0x8a, 0xd7, 0xa3, 0xfd, 0xde, 0x6a, 0x9a, 0x84, 0xc0, 0xce, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xbc, 0x08, 0x94, 0xaa, 0x59, 0x02, 0x00, 0x00,
}
