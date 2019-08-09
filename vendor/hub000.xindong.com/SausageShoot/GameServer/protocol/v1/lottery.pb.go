// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lottery.proto

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

type GetLotteryInfoReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLotteryInfoReq) Reset()         { *m = GetLotteryInfoReq{} }
func (m *GetLotteryInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetLotteryInfoReq) ProtoMessage()    {}
func (*GetLotteryInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cce7afd61783b10, []int{0}
}

func (m *GetLotteryInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLotteryInfoReq.Unmarshal(m, b)
}
func (m *GetLotteryInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLotteryInfoReq.Marshal(b, m, deterministic)
}
func (m *GetLotteryInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLotteryInfoReq.Merge(m, src)
}
func (m *GetLotteryInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetLotteryInfoReq.Size(m)
}
func (m *GetLotteryInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLotteryInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetLotteryInfoReq proto.InternalMessageInfo

func (m *GetLotteryInfoReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetLotteryInfoResp struct {
	Result               Result     `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	Box                  []*BoxItem `protobuf:"bytes,2,rep,name=box,proto3" json:"box,omitempty"`
	NextTime             int64      `protobuf:"varint,3,opt,name=nextTime,proto3" json:"nextTime,omitempty"`
	Price                int64      `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	RefreshTime          int64      `protobuf:"varint,5,opt,name=refreshTime,proto3" json:"refreshTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetLotteryInfoResp) Reset()         { *m = GetLotteryInfoResp{} }
func (m *GetLotteryInfoResp) String() string { return proto.CompactTextString(m) }
func (*GetLotteryInfoResp) ProtoMessage()    {}
func (*GetLotteryInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cce7afd61783b10, []int{1}
}

func (m *GetLotteryInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLotteryInfoResp.Unmarshal(m, b)
}
func (m *GetLotteryInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLotteryInfoResp.Marshal(b, m, deterministic)
}
func (m *GetLotteryInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLotteryInfoResp.Merge(m, src)
}
func (m *GetLotteryInfoResp) XXX_Size() int {
	return xxx_messageInfo_GetLotteryInfoResp.Size(m)
}
func (m *GetLotteryInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLotteryInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetLotteryInfoResp proto.InternalMessageInfo

func (m *GetLotteryInfoResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func (m *GetLotteryInfoResp) GetBox() []*BoxItem {
	if m != nil {
		return m.Box
	}
	return nil
}

func (m *GetLotteryInfoResp) GetNextTime() int64 {
	if m != nil {
		return m.NextTime
	}
	return 0
}

func (m *GetLotteryInfoResp) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *GetLotteryInfoResp) GetRefreshTime() int64 {
	if m != nil {
		return m.RefreshTime
	}
	return 0
}

type GetLotteryAwardReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	LotteryType          int64    `protobuf:"varint,2,opt,name=lotteryType,proto3" json:"lotteryType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLotteryAwardReq) Reset()         { *m = GetLotteryAwardReq{} }
func (m *GetLotteryAwardReq) String() string { return proto.CompactTextString(m) }
func (*GetLotteryAwardReq) ProtoMessage()    {}
func (*GetLotteryAwardReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cce7afd61783b10, []int{2}
}

func (m *GetLotteryAwardReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLotteryAwardReq.Unmarshal(m, b)
}
func (m *GetLotteryAwardReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLotteryAwardReq.Marshal(b, m, deterministic)
}
func (m *GetLotteryAwardReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLotteryAwardReq.Merge(m, src)
}
func (m *GetLotteryAwardReq) XXX_Size() int {
	return xxx_messageInfo_GetLotteryAwardReq.Size(m)
}
func (m *GetLotteryAwardReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLotteryAwardReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetLotteryAwardReq proto.InternalMessageInfo

func (m *GetLotteryAwardReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GetLotteryAwardReq) GetLotteryType() int64 {
	if m != nil {
		return m.LotteryType
	}
	return 0
}

type GetLotteryAwardResp struct {
	Result               Result     `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	Box                  []*BoxItem `protobuf:"bytes,2,rep,name=box,proto3" json:"box,omitempty"`
	List                 []*BoxItem `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetLotteryAwardResp) Reset()         { *m = GetLotteryAwardResp{} }
func (m *GetLotteryAwardResp) String() string { return proto.CompactTextString(m) }
func (*GetLotteryAwardResp) ProtoMessage()    {}
func (*GetLotteryAwardResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cce7afd61783b10, []int{3}
}

func (m *GetLotteryAwardResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLotteryAwardResp.Unmarshal(m, b)
}
func (m *GetLotteryAwardResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLotteryAwardResp.Marshal(b, m, deterministic)
}
func (m *GetLotteryAwardResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLotteryAwardResp.Merge(m, src)
}
func (m *GetLotteryAwardResp) XXX_Size() int {
	return xxx_messageInfo_GetLotteryAwardResp.Size(m)
}
func (m *GetLotteryAwardResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLotteryAwardResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetLotteryAwardResp proto.InternalMessageInfo

func (m *GetLotteryAwardResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func (m *GetLotteryAwardResp) GetBox() []*BoxItem {
	if m != nil {
		return m.Box
	}
	return nil
}

func (m *GetLotteryAwardResp) GetList() []*BoxItem {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*GetLotteryInfoReq)(nil), "Msg.GetLotteryInfoReq")
	proto.RegisterType((*GetLotteryInfoResp)(nil), "Msg.GetLotteryInfoResp")
	proto.RegisterType((*GetLotteryAwardReq)(nil), "Msg.GetLotteryAwardReq")
	proto.RegisterType((*GetLotteryAwardResp)(nil), "Msg.GetLotteryAwardResp")
}

func init() { proto.RegisterFile("lottery.proto", fileDescriptor_2cce7afd61783b10) }

var fileDescriptor_2cce7afd61783b10 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x53, 0x0a, 0xe4, 0xfb, 0xa6, 0x68, 0xe2, 0xe8, 0x62, 0xc2, 0xc2, 0x34, 0x75, 0x83,
	0x9b, 0x16, 0x71, 0x69, 0x8c, 0x91, 0x0d, 0x21, 0x81, 0xcd, 0xc0, 0xca, 0xdd, 0xb4, 0x5c, 0x4a,
	0x63, 0xdb, 0x5b, 0x67, 0x06, 0x2c, 0x89, 0xaf, 0xe4, 0x3b, 0x1a, 0x66, 0x88, 0x69, 0xfc, 0xb3,
	0x73, 0x79, 0xce, 0xef, 0xdc, 0x93, 0x7b, 0x67, 0xc8, 0x49, 0x8e, 0x5a, 0x83, 0xdc, 0x87, 0x95,
	0x44, 0x8d, 0xd4, 0x9d, 0xab, 0xb4, 0xdf, 0x93, 0xa0, 0xb6, 0xb9, 0xb6, 0x56, 0xdf, 0x8b, 0x85,
	0xca, 0x12, 0x2b, 0x82, 0x6b, 0x72, 0x36, 0x01, 0x3d, 0xb3, 0x33, 0xd3, 0x72, 0x8d, 0x1c, 0x5e,
	0xe8, 0x05, 0xe9, 0x68, 0x7c, 0x86, 0x92, 0x39, 0xbe, 0x33, 0xf8, 0xcf, 0xad, 0x08, 0xde, 0x1d,
	0x42, 0xbf, 0x66, 0x55, 0x45, 0xaf, 0x48, 0xd7, 0xd6, 0x9b, 0xf4, 0xe9, 0xc8, 0x0b, 0xe7, 0x2a,
	0x0d, 0xb9, 0xb1, 0xf8, 0x11, 0xd1, 0x4b, 0xe2, 0xc6, 0x58, 0xb3, 0x96, 0xef, 0x0e, 0xbc, 0x51,
	0xcf, 0x24, 0xc6, 0x58, 0x4f, 0x35, 0x14, 0xfc, 0x00, 0x68, 0x9f, 0xfc, 0x2b, 0xa1, 0xd6, 0xcb,
	0xac, 0x00, 0xe6, 0xfa, 0xce, 0xc0, 0xe5, 0x9f, 0xfa, 0xb0, 0x4d, 0x25, 0xb3, 0x04, 0x58, 0xdb,
	0x00, 0x2b, 0xa8, 0x4f, 0x3c, 0x09, 0x6b, 0x09, 0x6a, 0x63, 0x86, 0x3a, 0x86, 0x35, 0xad, 0x60,
	0xd6, 0x5c, 0xf7, 0xf1, 0x55, 0xc8, 0xd5, 0xaf, 0xb7, 0x1d, 0xda, 0x8e, 0xef, 0xb6, 0xdc, 0x57,
	0xc0, 0x5a, 0xb6, 0xad, 0x61, 0x05, 0x6f, 0xe4, 0xfc, 0x5b, 0xdb, 0x5f, 0x5d, 0xef, 0x93, 0x76,
	0x9e, 0x29, 0xcd, 0xdc, 0x1f, 0x02, 0x86, 0x8c, 0x1f, 0x9e, 0xee, 0x37, 0xdb, 0x78, 0x38, 0x1c,
	0x86, 0x75, 0x56, 0xae, 0xb0, 0x4c, 0xc3, 0x04, 0x8b, 0x68, 0x21, 0xb6, 0x4a, 0xa4, 0xb0, 0xd8,
	0x20, 0xea, 0x68, 0x22, 0x0a, 0x58, 0x80, 0xdc, 0x81, 0x8c, 0xcc, 0xd7, 0x26, 0x98, 0x47, 0xbb,
	0x9b, 0xbb, 0x54, 0x14, 0x50, 0xc5, 0x71, 0xd7, 0x78, 0xb7, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xeb, 0xce, 0xea, 0x66, 0x1f, 0x02, 0x00, 0x00,
}
