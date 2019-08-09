// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dailysign.proto

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

type PlayerDailySignInfoReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerDailySignInfoReq) Reset()         { *m = PlayerDailySignInfoReq{} }
func (m *PlayerDailySignInfoReq) String() string { return proto.CompactTextString(m) }
func (*PlayerDailySignInfoReq) ProtoMessage()    {}
func (*PlayerDailySignInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d876fc97f1bbfe, []int{0}
}

func (m *PlayerDailySignInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerDailySignInfoReq.Unmarshal(m, b)
}
func (m *PlayerDailySignInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerDailySignInfoReq.Marshal(b, m, deterministic)
}
func (m *PlayerDailySignInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerDailySignInfoReq.Merge(m, src)
}
func (m *PlayerDailySignInfoReq) XXX_Size() int {
	return xxx_messageInfo_PlayerDailySignInfoReq.Size(m)
}
func (m *PlayerDailySignInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerDailySignInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerDailySignInfoReq proto.InternalMessageInfo

func (m *PlayerDailySignInfoReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type PlayerDailySignInfoResp struct {
	Data                 uint32     `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	LastSignSt           int64      `protobuf:"varint,2,opt,name=lastSignSt,proto3" json:"lastSignSt,omitempty"`
	RestDay              uint32     `protobuf:"varint,3,opt,name=restDay,proto3" json:"restDay,omitempty"`
	TodaySigned          uint32     `protobuf:"varint,4,opt,name=todaySigned,proto3" json:"todaySigned,omitempty"`
	Box                  []*BoxItem `protobuf:"bytes,5,rep,name=box,proto3" json:"box,omitempty"`
	Result               Result     `protobuf:"varint,6,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	Cycle                int64      `protobuf:"varint,7,opt,name=cycle,proto3" json:"cycle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PlayerDailySignInfoResp) Reset()         { *m = PlayerDailySignInfoResp{} }
func (m *PlayerDailySignInfoResp) String() string { return proto.CompactTextString(m) }
func (*PlayerDailySignInfoResp) ProtoMessage()    {}
func (*PlayerDailySignInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d876fc97f1bbfe, []int{1}
}

func (m *PlayerDailySignInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerDailySignInfoResp.Unmarshal(m, b)
}
func (m *PlayerDailySignInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerDailySignInfoResp.Marshal(b, m, deterministic)
}
func (m *PlayerDailySignInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerDailySignInfoResp.Merge(m, src)
}
func (m *PlayerDailySignInfoResp) XXX_Size() int {
	return xxx_messageInfo_PlayerDailySignInfoResp.Size(m)
}
func (m *PlayerDailySignInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerDailySignInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerDailySignInfoResp proto.InternalMessageInfo

func (m *PlayerDailySignInfoResp) GetData() uint32 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *PlayerDailySignInfoResp) GetLastSignSt() int64 {
	if m != nil {
		return m.LastSignSt
	}
	return 0
}

func (m *PlayerDailySignInfoResp) GetRestDay() uint32 {
	if m != nil {
		return m.RestDay
	}
	return 0
}

func (m *PlayerDailySignInfoResp) GetTodaySigned() uint32 {
	if m != nil {
		return m.TodaySigned
	}
	return 0
}

func (m *PlayerDailySignInfoResp) GetBox() []*BoxItem {
	if m != nil {
		return m.Box
	}
	return nil
}

func (m *PlayerDailySignInfoResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func (m *PlayerDailySignInfoResp) GetCycle() int64 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

type PlayerDailySignReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	FixDay               uint32   `protobuf:"varint,2,opt,name=fixDay,proto3" json:"fixDay,omitempty"`
	IsShare              bool     `protobuf:"varint,3,opt,name=isShare,proto3" json:"isShare,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerDailySignReq) Reset()         { *m = PlayerDailySignReq{} }
func (m *PlayerDailySignReq) String() string { return proto.CompactTextString(m) }
func (*PlayerDailySignReq) ProtoMessage()    {}
func (*PlayerDailySignReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d876fc97f1bbfe, []int{2}
}

func (m *PlayerDailySignReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerDailySignReq.Unmarshal(m, b)
}
func (m *PlayerDailySignReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerDailySignReq.Marshal(b, m, deterministic)
}
func (m *PlayerDailySignReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerDailySignReq.Merge(m, src)
}
func (m *PlayerDailySignReq) XXX_Size() int {
	return xxx_messageInfo_PlayerDailySignReq.Size(m)
}
func (m *PlayerDailySignReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerDailySignReq.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerDailySignReq proto.InternalMessageInfo

func (m *PlayerDailySignReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PlayerDailySignReq) GetFixDay() uint32 {
	if m != nil {
		return m.FixDay
	}
	return 0
}

func (m *PlayerDailySignReq) GetIsShare() bool {
	if m != nil {
		return m.IsShare
	}
	return false
}

type PlayerDailySignResp struct {
	Data                 uint32     `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	LastSignSt           int64      `protobuf:"varint,2,opt,name=lastSignSt,proto3" json:"lastSignSt,omitempty"`
	RestDay              uint32     `protobuf:"varint,3,opt,name=restDay,proto3" json:"restDay,omitempty"`
	TodaySigned          uint32     `protobuf:"varint,4,opt,name=todaySigned,proto3" json:"todaySigned,omitempty"`
	Box                  []*BoxItem `protobuf:"bytes,5,rep,name=box,proto3" json:"box,omitempty"`
	ExtraBox             []*BoxItem `protobuf:"bytes,7,rep,name=extraBox,proto3" json:"extraBox,omitempty"`
	Result               Result     `protobuf:"varint,6,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PlayerDailySignResp) Reset()         { *m = PlayerDailySignResp{} }
func (m *PlayerDailySignResp) String() string { return proto.CompactTextString(m) }
func (*PlayerDailySignResp) ProtoMessage()    {}
func (*PlayerDailySignResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d876fc97f1bbfe, []int{3}
}

func (m *PlayerDailySignResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerDailySignResp.Unmarshal(m, b)
}
func (m *PlayerDailySignResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerDailySignResp.Marshal(b, m, deterministic)
}
func (m *PlayerDailySignResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerDailySignResp.Merge(m, src)
}
func (m *PlayerDailySignResp) XXX_Size() int {
	return xxx_messageInfo_PlayerDailySignResp.Size(m)
}
func (m *PlayerDailySignResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerDailySignResp.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerDailySignResp proto.InternalMessageInfo

func (m *PlayerDailySignResp) GetData() uint32 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *PlayerDailySignResp) GetLastSignSt() int64 {
	if m != nil {
		return m.LastSignSt
	}
	return 0
}

func (m *PlayerDailySignResp) GetRestDay() uint32 {
	if m != nil {
		return m.RestDay
	}
	return 0
}

func (m *PlayerDailySignResp) GetTodaySigned() uint32 {
	if m != nil {
		return m.TodaySigned
	}
	return 0
}

func (m *PlayerDailySignResp) GetBox() []*BoxItem {
	if m != nil {
		return m.Box
	}
	return nil
}

func (m *PlayerDailySignResp) GetExtraBox() []*BoxItem {
	if m != nil {
		return m.ExtraBox
	}
	return nil
}

func (m *PlayerDailySignResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func init() {
	proto.RegisterType((*PlayerDailySignInfoReq)(nil), "Msg.PlayerDailySignInfoReq")
	proto.RegisterType((*PlayerDailySignInfoResp)(nil), "Msg.PlayerDailySignInfoResp")
	proto.RegisterType((*PlayerDailySignReq)(nil), "Msg.PlayerDailySignReq")
	proto.RegisterType((*PlayerDailySignResp)(nil), "Msg.PlayerDailySignResp")
}

func init() { proto.RegisterFile("dailysign.proto", fileDescriptor_58d876fc97f1bbfe) }

var fileDescriptor_58d876fc97f1bbfe = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x52, 0xc1, 0x6a, 0xe3, 0x30,
	0x14, 0xc4, 0x71, 0xe2, 0x64, 0xe5, 0xec, 0x2e, 0x68, 0x97, 0xac, 0xc8, 0x21, 0x18, 0xef, 0xc5,
	0x27, 0x3b, 0x4d, 0x8f, 0xa5, 0x14, 0x42, 0xa0, 0xe4, 0x10, 0x28, 0xf2, 0xad, 0xf4, 0x22, 0xdb,
	0x2f, 0x8e, 0xa9, 0x6d, 0xa5, 0x92, 0x12, 0xec, 0xff, 0xed, 0x27, 0xf4, 0x03, 0x8a, 0xe5, 0xa4,
	0x84, 0x36, 0x85, 0x1e, 0x7b, 0xd3, 0xcc, 0xbc, 0x91, 0xde, 0x0c, 0x42, 0xbf, 0x13, 0x96, 0xe5,
	0xb5, 0xcc, 0xd2, 0xd2, 0xdf, 0x0a, 0xae, 0x38, 0x36, 0x57, 0x32, 0x1d, 0x0f, 0x05, 0xc8, 0x5d,
	0xae, 0x5a, 0x6a, 0x6c, 0x47, 0x4c, 0x66, 0x71, 0x0b, 0x5c, 0x1f, 0x8d, 0xee, 0x72, 0x56, 0x83,
	0x58, 0x34, 0xc6, 0x30, 0x4b, 0xcb, 0x65, 0xb9, 0xe6, 0x14, 0x9e, 0xf0, 0x5f, 0xd4, 0x53, 0xfc,
	0x11, 0x4a, 0x62, 0x38, 0x86, 0xf7, 0x83, 0xb6, 0xc0, 0x7d, 0x36, 0xd0, 0xbf, 0xb3, 0x06, 0xb9,
	0xc5, 0x18, 0x75, 0x13, 0xa6, 0x98, 0x36, 0xfc, 0xa4, 0xfa, 0x8c, 0x27, 0x08, 0xe5, 0x4c, 0xaa,
	0x66, 0x2e, 0x54, 0xa4, 0xe3, 0x18, 0x9e, 0x49, 0x4f, 0x18, 0x4c, 0x50, 0x5f, 0x80, 0x54, 0x0b,
	0x56, 0x13, 0x53, 0xdb, 0x8e, 0x10, 0x3b, 0xc8, 0x56, 0x3c, 0x61, 0xfa, 0x09, 0x48, 0x48, 0x57,
	0xab, 0xa7, 0x14, 0x9e, 0x20, 0x33, 0xe2, 0x15, 0xe9, 0x39, 0xa6, 0x67, 0xcf, 0x86, 0xfe, 0x4a,
	0xa6, 0xfe, 0x9c, 0x57, 0x4b, 0x05, 0x05, 0x6d, 0x04, 0xfc, 0x1f, 0x59, 0x6d, 0x70, 0x62, 0x39,
	0x86, 0xf7, 0x6b, 0x66, 0xeb, 0x11, 0xaa, 0x29, 0x7a, 0x90, 0x9a, 0x98, 0x71, 0x1d, 0xe7, 0x40,
	0xfa, 0x7a, 0xb7, 0x16, 0xb8, 0x0f, 0x08, 0xbf, 0x4b, 0xf9, 0x69, 0x25, 0x78, 0x84, 0xac, 0x75,
	0x56, 0x35, 0x09, 0x3a, 0x7a, 0xc7, 0x03, 0x6a, 0xa2, 0x65, 0x32, 0xdc, 0x30, 0x01, 0x3a, 0xda,
	0x80, 0x1e, 0xa1, 0xfb, 0x62, 0xa0, 0x3f, 0x1f, 0xae, 0xff, 0x86, 0x05, 0x7a, 0x68, 0x00, 0x95,
	0x12, 0x6c, 0xce, 0x2b, 0xd2, 0x3f, 0x33, 0xf4, 0xa6, 0x7e, 0xa9, 0xea, 0xf9, 0xcd, 0xfd, 0xf5,
	0x66, 0x17, 0x4d, 0xa7, 0x53, 0xbf, 0xca, 0xca, 0x84, 0x97, 0xa9, 0x1f, 0xf3, 0x22, 0x08, 0xd9,
	0x4e, 0xb2, 0x14, 0xc2, 0x0d, 0xe7, 0x2a, 0xb8, 0x65, 0x05, 0x84, 0x20, 0xf6, 0x20, 0x02, 0xfd,
	0x3f, 0x63, 0x9e, 0x07, 0xfb, 0x8b, 0xab, 0x94, 0x15, 0xb0, 0x8d, 0x22, 0x4b, 0x73, 0x97, 0xaf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x45, 0x80, 0x54, 0xe6, 0x02, 0x00, 0x00,
}
