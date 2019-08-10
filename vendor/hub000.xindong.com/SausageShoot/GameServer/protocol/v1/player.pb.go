// Code generated by protoc-gen-go. DO NOT EDIT.
// source: player.proto

package gamepb

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

// GetPlayerAttributeReq & GetPlayerAttributeResp is used to get basic player informations for ui
type GetPlayerAttributeReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerAttributeReq) Reset()         { *m = GetPlayerAttributeReq{} }
func (m *GetPlayerAttributeReq) String() string { return proto.CompactTextString(m) }
func (*GetPlayerAttributeReq) ProtoMessage()    {}
func (*GetPlayerAttributeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{0}
}

func (m *GetPlayerAttributeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerAttributeReq.Unmarshal(m, b)
}
func (m *GetPlayerAttributeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerAttributeReq.Marshal(b, m, deterministic)
}
func (m *GetPlayerAttributeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerAttributeReq.Merge(m, src)
}
func (m *GetPlayerAttributeReq) XXX_Size() int {
	return xxx_messageInfo_GetPlayerAttributeReq.Size(m)
}
func (m *GetPlayerAttributeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerAttributeReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerAttributeReq proto.InternalMessageInfo

func (m *GetPlayerAttributeReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetPlayerAttributeResp struct {
	Result               Result           `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	Attr                 *PlayerAttribute `protobuf:"bytes,2,opt,name=attr,proto3" json:"attr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetPlayerAttributeResp) Reset()         { *m = GetPlayerAttributeResp{} }
func (m *GetPlayerAttributeResp) String() string { return proto.CompactTextString(m) }
func (*GetPlayerAttributeResp) ProtoMessage()    {}
func (*GetPlayerAttributeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{1}
}

func (m *GetPlayerAttributeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerAttributeResp.Unmarshal(m, b)
}
func (m *GetPlayerAttributeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerAttributeResp.Marshal(b, m, deterministic)
}
func (m *GetPlayerAttributeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerAttributeResp.Merge(m, src)
}
func (m *GetPlayerAttributeResp) XXX_Size() int {
	return xxx_messageInfo_GetPlayerAttributeResp.Size(m)
}
func (m *GetPlayerAttributeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerAttributeResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerAttributeResp proto.InternalMessageInfo

func (m *GetPlayerAttributeResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func (m *GetPlayerAttributeResp) GetAttr() *PlayerAttribute {
	if m != nil {
		return m.Attr
	}
	return nil
}

type GetPlayerBufReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerBufReq) Reset()         { *m = GetPlayerBufReq{} }
func (m *GetPlayerBufReq) String() string { return proto.CompactTextString(m) }
func (*GetPlayerBufReq) ProtoMessage()    {}
func (*GetPlayerBufReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{2}
}

func (m *GetPlayerBufReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerBufReq.Unmarshal(m, b)
}
func (m *GetPlayerBufReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerBufReq.Marshal(b, m, deterministic)
}
func (m *GetPlayerBufReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerBufReq.Merge(m, src)
}
func (m *GetPlayerBufReq) XXX_Size() int {
	return xxx_messageInfo_GetPlayerBufReq.Size(m)
}
func (m *GetPlayerBufReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerBufReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerBufReq proto.InternalMessageInfo

func (m *GetPlayerBufReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetPlayerBufResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	Buffers              string   `protobuf:"bytes,2,opt,name=buffers,proto3" json:"buffers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerBufResp) Reset()         { *m = GetPlayerBufResp{} }
func (m *GetPlayerBufResp) String() string { return proto.CompactTextString(m) }
func (*GetPlayerBufResp) ProtoMessage()    {}
func (*GetPlayerBufResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{3}
}

func (m *GetPlayerBufResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerBufResp.Unmarshal(m, b)
}
func (m *GetPlayerBufResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerBufResp.Marshal(b, m, deterministic)
}
func (m *GetPlayerBufResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerBufResp.Merge(m, src)
}
func (m *GetPlayerBufResp) XXX_Size() int {
	return xxx_messageInfo_GetPlayerBufResp.Size(m)
}
func (m *GetPlayerBufResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerBufResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerBufResp proto.InternalMessageInfo

func (m *GetPlayerBufResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func (m *GetPlayerBufResp) GetBuffers() string {
	if m != nil {
		return m.Buffers
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPlayerAttributeReq)(nil), "Msg.GetPlayerAttributeReq")
	proto.RegisterType((*GetPlayerAttributeResp)(nil), "Msg.GetPlayerAttributeResp")
	proto.RegisterType((*GetPlayerBufReq)(nil), "Msg.GetPlayerBufReq")
	proto.RegisterType((*GetPlayerBufResp)(nil), "Msg.GetPlayerBufResp")
}

func init() { proto.RegisterFile("player.proto", fileDescriptor_41d803d1b635d5c6) }

var fileDescriptor_41d803d1b635d5c6 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x14, 0xc4, 0x89, 0x7f, 0x2a, 0x7d, 0x29, 0x2a, 0xa1, 0x4a, 0xe8, 0xa9, 0xc4, 0x83, 0xb9, 0xb8,
	0x89, 0xf5, 0x28, 0x22, 0xf6, 0xd2, 0x53, 0x41, 0x37, 0x37, 0x6f, 0xbb, 0xf1, 0x65, 0x1b, 0x4c,
	0xb2, 0x71, 0xf7, 0x6d, 0xd1, 0x6f, 0x2f, 0x6c, 0x54, 0x50, 0x14, 0x3c, 0xce, 0xec, 0x6f, 0x76,
	0x86, 0x07, 0x93, 0xbe, 0x11, 0x6f, 0x68, 0x58, 0x6f, 0x34, 0xe9, 0x68, 0x77, 0x6d, 0xd5, 0x6c,
	0x62, 0xd0, 0xba, 0x86, 0x06, 0x6b, 0x16, 0x4a, 0x61, 0xeb, 0x72, 0x10, 0xc9, 0x05, 0x9c, 0xac,
	0x90, 0xee, 0x7d, 0xe4, 0x8e, 0xc8, 0xd4, 0xd2, 0x11, 0x72, 0x7c, 0x89, 0xa6, 0xb0, 0x4f, 0xfa,
	0x19, 0xbb, 0x38, 0x98, 0x07, 0xe9, 0x98, 0x0f, 0x22, 0x51, 0x70, 0xfa, 0x1b, 0x6e, 0xfb, 0xe8,
	0x0c, 0x46, 0x43, 0x8b, 0x0f, 0x1c, 0x2e, 0x42, 0xb6, 0xb6, 0x8a, 0x71, 0x6f, 0xf1, 0x8f, 0xa7,
	0x28, 0x85, 0x3d, 0x41, 0x64, 0xe2, 0x9d, 0x79, 0x90, 0x86, 0x8b, 0xa9, 0x47, 0x7e, 0x7e, 0xe6,
	0x89, 0xe4, 0x1c, 0x8e, 0xbe, 0x8a, 0x96, 0xae, 0xfa, 0x7b, 0xd1, 0x03, 0x1c, 0x7f, 0x07, 0xff,
	0xbb, 0x25, 0x86, 0x03, 0xe9, 0xaa, 0x0a, 0x8d, 0xf5, 0x73, 0xc6, 0xfc, 0x53, 0x2e, 0x6f, 0x1f,
	0x6f, 0x36, 0x4e, 0xe6, 0x79, 0xce, 0x5e, 0xeb, 0xee, 0x49, 0x77, 0x8a, 0x95, 0xba, 0xcd, 0x0a,
	0xe1, 0xac, 0x50, 0x58, 0x6c, 0xb4, 0xa6, 0x6c, 0x25, 0x5a, 0x2c, 0xd0, 0x6c, 0xd1, 0x64, 0xfe,
	0x8e, 0xa5, 0x6e, 0xb2, 0xed, 0xe5, 0xb5, 0x12, 0x2d, 0xf6, 0x52, 0x8e, 0xbc, 0x77, 0xf5, 0x1e,
	0x00, 0x00, 0xff, 0xff, 0xe2, 0xb5, 0xc7, 0xf1, 0x8b, 0x01, 0x00, 0x00,
}