// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account/profile.proto

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

type QueryProfileRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryProfileRequest) Reset()         { *m = QueryProfileRequest{} }
func (m *QueryProfileRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProfileRequest) ProtoMessage()    {}
func (*QueryProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe3d38111f3626c4, []int{0}
}

func (m *QueryProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryProfileRequest.Unmarshal(m, b)
}
func (m *QueryProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryProfileRequest.Marshal(b, m, deterministic)
}
func (m *QueryProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProfileRequest.Merge(m, src)
}
func (m *QueryProfileRequest) XXX_Size() int {
	return xxx_messageInfo_QueryProfileRequest.Size(m)
}
func (m *QueryProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProfileRequest proto.InternalMessageInfo

func (m *QueryProfileRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type QueryProfileResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Profile              string   `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryProfileResponse) Reset()         { *m = QueryProfileResponse{} }
func (m *QueryProfileResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProfileResponse) ProtoMessage()    {}
func (*QueryProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe3d38111f3626c4, []int{1}
}

func (m *QueryProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryProfileResponse.Unmarshal(m, b)
}
func (m *QueryProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryProfileResponse.Marshal(b, m, deterministic)
}
func (m *QueryProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProfileResponse.Merge(m, src)
}
func (m *QueryProfileResponse) XXX_Size() int {
	return xxx_messageInfo_QueryProfileResponse.Size(m)
}
func (m *QueryProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProfileResponse proto.InternalMessageInfo

func (m *QueryProfileResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *QueryProfileResponse) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

type UpdateProfileRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	Profile              string   `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProfileRequest) Reset()         { *m = UpdateProfileRequest{} }
func (m *UpdateProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateProfileRequest) ProtoMessage()    {}
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe3d38111f3626c4, []int{2}
}

func (m *UpdateProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProfileRequest.Unmarshal(m, b)
}
func (m *UpdateProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProfileRequest.Merge(m, src)
}
func (m *UpdateProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateProfileRequest.Size(m)
}
func (m *UpdateProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProfileRequest proto.InternalMessageInfo

func (m *UpdateProfileRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *UpdateProfileRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

type UpdateProfileResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProfileResponse) Reset()         { *m = UpdateProfileResponse{} }
func (m *UpdateProfileResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateProfileResponse) ProtoMessage()    {}
func (*UpdateProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe3d38111f3626c4, []int{3}
}

func (m *UpdateProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProfileResponse.Unmarshal(m, b)
}
func (m *UpdateProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProfileResponse.Marshal(b, m, deterministic)
}
func (m *UpdateProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProfileResponse.Merge(m, src)
}
func (m *UpdateProfileResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateProfileResponse.Size(m)
}
func (m *UpdateProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProfileResponse proto.InternalMessageInfo

func (m *UpdateProfileResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryProfileRequest)(nil), "account.QueryProfileRequest")
	proto.RegisterType((*QueryProfileResponse)(nil), "account.QueryProfileResponse")
	proto.RegisterType((*UpdateProfileRequest)(nil), "account.UpdateProfileRequest")
	proto.RegisterType((*UpdateProfileResponse)(nil), "account.UpdateProfileResponse")
}

func init() {
	proto.RegisterFile("proto/account/profile.proto", fileDescriptor_fe3d38111f3626c4)
}

var fileDescriptor_fe3d38111f3626c4 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2f, 0x28, 0xca, 0x4f, 0xcb, 0xcc,
	0x49, 0xd5, 0x03, 0x8b, 0x0a, 0xb1, 0x43, 0x85, 0xa5, 0xa4, 0x50, 0x55, 0x15, 0x67, 0x24, 0x16,
	0xa5, 0xa6, 0x40, 0x14, 0x29, 0x99, 0x73, 0x09, 0x07, 0x96, 0xa6, 0x16, 0x55, 0x06, 0x40, 0xb4,
	0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x29, 0x70, 0x71, 0x27, 0x26, 0x27, 0xa7, 0x16,
	0x17, 0x87, 0xe4, 0x67, 0xa7, 0xe6, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x0b, 0x29,
	0x45, 0x72, 0x89, 0xa0, 0x6a, 0x2c, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x52, 0xe7, 0x62, 0x2b,
	0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x06, 0x6b, 0xe2, 0x36, 0xe2, 0xd7, 0x83, 0xda, 0xab, 0x17, 0x0c,
	0x16, 0x0e, 0x82, 0x4a, 0x0b, 0x49, 0x70, 0xb1, 0x43, 0xdd, 0x2b, 0xc1, 0x04, 0x36, 0x1e, 0xc6,
	0x55, 0x0a, 0xe2, 0x12, 0x09, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x25, 0xd5, 0x51, 0x78, 0xcc, 0x74,
	0xe0, 0x12, 0x45, 0x33, 0x93, 0x44, 0xf7, 0x1a, 0xcd, 0x61, 0xe4, 0x62, 0x87, 0x6a, 0x16, 0x72,
	0xe3, 0x62, 0x05, 0x7b, 0x5e, 0x48, 0x06, 0xae, 0x1a, 0x4b, 0x28, 0x4a, 0xc9, 0xe2, 0x90, 0x85,
	0x58, 0xad, 0xc4, 0x20, 0xe4, 0xc9, 0xc5, 0x06, 0x71, 0x95, 0x10, 0x42, 0x29, 0x36, 0xaf, 0x4b,
	0xc9, 0xe1, 0x92, 0x86, 0x19, 0x95, 0xc4, 0x06, 0x8e, 0x4f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe2, 0x83, 0x4a, 0xf0, 0x13, 0x02, 0x00, 0x00,
}
