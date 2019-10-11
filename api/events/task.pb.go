// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	codec "github.com/kcarretto/paragon/api/codec"
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

type TaskQueued struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              string               `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Filter               *codec.AgentMetadata `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskQueued) Reset()         { *m = TaskQueued{} }
func (m *TaskQueued) String() string { return proto.CompactTextString(m) }
func (*TaskQueued) ProtoMessage()    {}
func (*TaskQueued) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *TaskQueued) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskQueued.Unmarshal(m, b)
}
func (m *TaskQueued) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskQueued.Marshal(b, m, deterministic)
}
func (m *TaskQueued) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskQueued.Merge(m, src)
}
func (m *TaskQueued) XXX_Size() int {
	return xxx_messageInfo_TaskQueued.Size(m)
}
func (m *TaskQueued) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskQueued.DiscardUnknown(m)
}

var xxx_messageInfo_TaskQueued proto.InternalMessageInfo

func (m *TaskQueued) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskQueued) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TaskQueued) GetFilter() *codec.AgentMetadata {
	if m != nil {
		return m.Filter
	}
	return nil
}

type TaskClaimed struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Agent                *codec.AgentMetadata `protobuf:"bytes,2,opt,name=agent,proto3" json:"agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskClaimed) Reset()         { *m = TaskClaimed{} }
func (m *TaskClaimed) String() string { return proto.CompactTextString(m) }
func (*TaskClaimed) ProtoMessage()    {}
func (*TaskClaimed) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{1}
}

func (m *TaskClaimed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskClaimed.Unmarshal(m, b)
}
func (m *TaskClaimed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskClaimed.Marshal(b, m, deterministic)
}
func (m *TaskClaimed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskClaimed.Merge(m, src)
}
func (m *TaskClaimed) XXX_Size() int {
	return xxx_messageInfo_TaskClaimed.Size(m)
}
func (m *TaskClaimed) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskClaimed.DiscardUnknown(m)
}

var xxx_messageInfo_TaskClaimed proto.InternalMessageInfo

func (m *TaskClaimed) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskClaimed) GetAgent() *codec.AgentMetadata {
	if m != nil {
		return m.Agent
	}
	return nil
}

type TaskExecuted struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Output               string               `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Error                string               `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	ExecStartTime        int64                `protobuf:"varint,4,opt,name=execStartTime,proto3" json:"execStartTime,omitempty"`
	ExecStopTime         int64                `protobuf:"varint,5,opt,name=execStopTime,proto3" json:"execStopTime,omitempty"`
	RecvTime             int64                `protobuf:"varint,6,opt,name=recvTime,proto3" json:"recvTime,omitempty"`
	Agent                *codec.AgentMetadata `protobuf:"bytes,7,opt,name=agent,proto3" json:"agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskExecuted) Reset()         { *m = TaskExecuted{} }
func (m *TaskExecuted) String() string { return proto.CompactTextString(m) }
func (*TaskExecuted) ProtoMessage()    {}
func (*TaskExecuted) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{2}
}

func (m *TaskExecuted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecuted.Unmarshal(m, b)
}
func (m *TaskExecuted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecuted.Marshal(b, m, deterministic)
}
func (m *TaskExecuted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecuted.Merge(m, src)
}
func (m *TaskExecuted) XXX_Size() int {
	return xxx_messageInfo_TaskExecuted.Size(m)
}
func (m *TaskExecuted) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecuted.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecuted proto.InternalMessageInfo

func (m *TaskExecuted) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskExecuted) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *TaskExecuted) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TaskExecuted) GetExecStartTime() int64 {
	if m != nil {
		return m.ExecStartTime
	}
	return 0
}

func (m *TaskExecuted) GetExecStopTime() int64 {
	if m != nil {
		return m.ExecStopTime
	}
	return 0
}

func (m *TaskExecuted) GetRecvTime() int64 {
	if m != nil {
		return m.RecvTime
	}
	return 0
}

func (m *TaskExecuted) GetAgent() *codec.AgentMetadata {
	if m != nil {
		return m.Agent
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskQueued)(nil), "events.TaskQueued")
	proto.RegisterType((*TaskClaimed)(nil), "events.TaskClaimed")
	proto.RegisterType((*TaskExecuted)(nil), "events.TaskExecuted")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xc1, 0x4e, 0xc2, 0x40,
	0x14, 0x4c, 0x41, 0x8a, 0x3c, 0xd0, 0xc3, 0x06, 0x4d, 0xc3, 0x89, 0x34, 0x26, 0xa2, 0x31, 0x6d,
	0xa2, 0x5f, 0xa0, 0xc6, 0x83, 0x07, 0x0f, 0x56, 0x4e, 0xde, 0x1e, 0xbb, 0x4f, 0xdc, 0x00, 0xdd,
	0x66, 0xfb, 0x96, 0xf0, 0xb5, 0x7e, 0x8b, 0xe9, 0x2e, 0x68, 0x48, 0x30, 0x1e, 0x67, 0xde, 0x64,
	0x66, 0x5e, 0x06, 0x80, 0xb1, 0x5e, 0x64, 0x95, 0x35, 0x6c, 0x44, 0x4c, 0x6b, 0x2a, 0xb9, 0x1e,
	0x9d, 0x61, 0xa5, 0x73, 0x69, 0x14, 0xc9, 0x1c, 0xe7, 0x54, 0x72, 0x38, 0xa7, 0x0a, 0x60, 0x8a,
	0xf5, 0xe2, 0xd5, 0x91, 0x23, 0x25, 0x4e, 0xa1, 0xa5, 0x55, 0x12, 0x8d, 0xa3, 0x49, 0xaf, 0x68,
	0x69, 0x25, 0x12, 0xe8, 0x4a, 0x53, 0x32, 0x95, 0x9c, 0xb4, 0x3c, 0xb9, 0x83, 0xe2, 0x06, 0xe2,
	0x0f, 0xbd, 0x64, 0xb2, 0x49, 0x7b, 0x1c, 0x4d, 0xfa, 0xb7, 0xc3, 0xcc, 0x7b, 0x67, 0xf7, 0x8d,
	0xf7, 0x0b, 0x31, 0x2a, 0x64, 0x2c, 0xb6, 0x9a, 0xf4, 0x19, 0xfa, 0x4d, 0xca, 0xe3, 0x12, 0xf5,
	0xea, 0x40, 0xcc, 0x35, 0x74, 0x7c, 0x27, 0x1f, 0xf2, 0x97, 0x57, 0x90, 0xa4, 0x5f, 0x11, 0x0c,
	0x1a, 0xaf, 0xa7, 0x0d, 0x49, 0xc7, 0x07, 0xcc, 0xce, 0x21, 0x36, 0x8e, 0x2b, 0xb7, 0xab, 0xbc,
	0x45, 0x62, 0x08, 0x1d, 0xb2, 0xd6, 0x84, 0xc2, 0xbd, 0x22, 0x00, 0x71, 0x01, 0x27, 0xb4, 0x21,
	0xf9, 0xc6, 0x68, 0x79, 0xaa, 0x57, 0x94, 0x1c, 0x8d, 0xa3, 0x49, 0xbb, 0xd8, 0x27, 0x45, 0x0a,
	0x83, 0x40, 0x98, 0xca, 0x8b, 0x3a, 0x5e, 0xb4, 0xc7, 0x89, 0x11, 0x1c, 0x5b, 0x92, 0x6b, 0x7f,
	0x8f, 0xfd, 0xfd, 0x07, 0xff, 0x3e, 0xd8, 0xfd, 0xf7, 0xc1, 0x87, 0xab, 0xf7, 0xcb, 0xb9, 0xe6,
	0x4f, 0x37, 0xcb, 0xa4, 0x59, 0xe5, 0x0b, 0x89, 0xd6, 0x12, 0xb3, 0xc9, 0x2b, 0xb4, 0x38, 0x37,
	0x65, 0xde, 0xec, 0x18, 0x36, 0x9d, 0xc5, 0x7e, 0xc3, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5e, 0x6f, 0x75, 0xbf, 0xf0, 0x01, 0x00, 0x00,
}
