// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account/shared.proto

package account

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 策略
type Strategy int32

const (
	Strategy_NONE Strategy = 0
	Strategy_JWT  Strategy = 1
)

var Strategy_name = map[int32]string{
	0: "NONE",
	1: "JWT",
}

var Strategy_value = map[string]int32{
	"NONE": 0,
	"JWT":  1,
}

func (x Strategy) String() string {
	return proto.EnumName(Strategy_name, int32(x))
}

func (Strategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49158a9607562c1b, []int{0}
}

// 状态
type Status struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_49158a9607562c1b, []int{0}
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

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("account.Strategy", Strategy_name, Strategy_value)
	proto.RegisterType((*Status)(nil), "account.Status")
}

func init() {
	proto.RegisterFile("proto/account/shared.proto", fileDescriptor_49158a9607562c1b)
}

var fileDescriptor_49158a9607562c1b = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2f, 0xce, 0x48, 0x2c, 0x4a, 0x4d,
	0xd1, 0x03, 0x0b, 0x0a, 0xb1, 0x43, 0x45, 0x95, 0xcc, 0xb8, 0xd8, 0x82, 0x4b, 0x12, 0x4b, 0x4a,
	0x8b, 0x85, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83,
	0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x4b, 0x96, 0x8b, 0x23, 0xb8, 0xa4, 0x28, 0xb1, 0x24, 0x35,
	0xbd, 0x52, 0x88, 0x83, 0x8b, 0xc5, 0xcf, 0xdf, 0xcf, 0x55, 0x80, 0x41, 0x88, 0x9d, 0x8b, 0xd9,
	0x2b, 0x3c, 0x44, 0x80, 0x31, 0x89, 0x0d, 0x6c, 0x8d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc7,
	0x00, 0x96, 0x8b, 0x84, 0x00, 0x00, 0x00,
}
