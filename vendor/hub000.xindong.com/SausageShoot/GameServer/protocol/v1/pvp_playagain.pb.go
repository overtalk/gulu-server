// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pvp_playagain.proto

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

// pvf房间状态
type PvpRoomStatus int32

const (
	PvpRoomStatus_pvp_waiting_invite PvpRoomStatus = 0
	PvpRoomStatus_pvp_waiting_accept PvpRoomStatus = 1
	PvpRoomStatus_pvp_cancel         PvpRoomStatus = 2
	PvpRoomStatus_pvp_leave          PvpRoomStatus = 3
	PvpRoomStatus_pvp_start          PvpRoomStatus = 4
	PvpRoomStatus_pvp_refuse         PvpRoomStatus = 5
)

var PvpRoomStatus_name = map[int32]string{
	0: "pvp_waiting_invite",
	1: "pvp_waiting_accept",
	2: "pvp_cancel",
	3: "pvp_leave",
	4: "pvp_start",
	5: "pvp_refuse",
}

var PvpRoomStatus_value = map[string]int32{
	"pvp_waiting_invite": 0,
	"pvp_waiting_accept": 1,
	"pvp_cancel":         2,
	"pvp_leave":          3,
	"pvp_start":          4,
	"pvp_refuse":         5,
}

func (x PvpRoomStatus) String() string {
	return proto.EnumName(PvpRoomStatus_name, int32(x))
}

func (PvpRoomStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0523d21ed0ac0a7d, []int{0}
}

// pvp再来一局房间状态轮询
type PvpPlayAgainStatusReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoomID               string   `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvpPlayAgainStatusReq) Reset()         { *m = PvpPlayAgainStatusReq{} }
func (m *PvpPlayAgainStatusReq) String() string { return proto.CompactTextString(m) }
func (*PvpPlayAgainStatusReq) ProtoMessage()    {}
func (*PvpPlayAgainStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0523d21ed0ac0a7d, []int{0}
}

func (m *PvpPlayAgainStatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvpPlayAgainStatusReq.Unmarshal(m, b)
}
func (m *PvpPlayAgainStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvpPlayAgainStatusReq.Marshal(b, m, deterministic)
}
func (m *PvpPlayAgainStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvpPlayAgainStatusReq.Merge(m, src)
}
func (m *PvpPlayAgainStatusReq) XXX_Size() int {
	return xxx_messageInfo_PvpPlayAgainStatusReq.Size(m)
}
func (m *PvpPlayAgainStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PvpPlayAgainStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_PvpPlayAgainStatusReq proto.InternalMessageInfo

func (m *PvpPlayAgainStatusReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PvpPlayAgainStatusReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

type PvpPlayAgainStatusResp struct {
	Result               Result        `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	RoomStatus           PvpRoomStatus `protobuf:"varint,2,opt,name=roomStatus,proto3,enum=Msg.PvpRoomStatus" json:"roomStatus,omitempty"`
	IsInviter            bool          `protobuf:"varint,3,opt,name=isInviter,proto3" json:"isInviter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PvpPlayAgainStatusResp) Reset()         { *m = PvpPlayAgainStatusResp{} }
func (m *PvpPlayAgainStatusResp) String() string { return proto.CompactTextString(m) }
func (*PvpPlayAgainStatusResp) ProtoMessage()    {}
func (*PvpPlayAgainStatusResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0523d21ed0ac0a7d, []int{1}
}

func (m *PvpPlayAgainStatusResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvpPlayAgainStatusResp.Unmarshal(m, b)
}
func (m *PvpPlayAgainStatusResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvpPlayAgainStatusResp.Marshal(b, m, deterministic)
}
func (m *PvpPlayAgainStatusResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvpPlayAgainStatusResp.Merge(m, src)
}
func (m *PvpPlayAgainStatusResp) XXX_Size() int {
	return xxx_messageInfo_PvpPlayAgainStatusResp.Size(m)
}
func (m *PvpPlayAgainStatusResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PvpPlayAgainStatusResp.DiscardUnknown(m)
}

var xxx_messageInfo_PvpPlayAgainStatusResp proto.InternalMessageInfo

func (m *PvpPlayAgainStatusResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func (m *PvpPlayAgainStatusResp) GetRoomStatus() PvpRoomStatus {
	if m != nil {
		return m.RoomStatus
	}
	return PvpRoomStatus_pvp_waiting_invite
}

func (m *PvpPlayAgainStatusResp) GetIsInviter() bool {
	if m != nil {
		return m.IsInviter
	}
	return false
}

// pvp再来一局邀请
type PvpPlayAgainInviteReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoomID               string   `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvpPlayAgainInviteReq) Reset()         { *m = PvpPlayAgainInviteReq{} }
func (m *PvpPlayAgainInviteReq) String() string { return proto.CompactTextString(m) }
func (*PvpPlayAgainInviteReq) ProtoMessage()    {}
func (*PvpPlayAgainInviteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0523d21ed0ac0a7d, []int{2}
}

func (m *PvpPlayAgainInviteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvpPlayAgainInviteReq.Unmarshal(m, b)
}
func (m *PvpPlayAgainInviteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvpPlayAgainInviteReq.Marshal(b, m, deterministic)
}
func (m *PvpPlayAgainInviteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvpPlayAgainInviteReq.Merge(m, src)
}
func (m *PvpPlayAgainInviteReq) XXX_Size() int {
	return xxx_messageInfo_PvpPlayAgainInviteReq.Size(m)
}
func (m *PvpPlayAgainInviteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PvpPlayAgainInviteReq.DiscardUnknown(m)
}

var xxx_messageInfo_PvpPlayAgainInviteReq proto.InternalMessageInfo

func (m *PvpPlayAgainInviteReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PvpPlayAgainInviteReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

type PvpPlayAgainInviteResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvpPlayAgainInviteResp) Reset()         { *m = PvpPlayAgainInviteResp{} }
func (m *PvpPlayAgainInviteResp) String() string { return proto.CompactTextString(m) }
func (*PvpPlayAgainInviteResp) ProtoMessage()    {}
func (*PvpPlayAgainInviteResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0523d21ed0ac0a7d, []int{3}
}

func (m *PvpPlayAgainInviteResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvpPlayAgainInviteResp.Unmarshal(m, b)
}
func (m *PvpPlayAgainInviteResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvpPlayAgainInviteResp.Marshal(b, m, deterministic)
}
func (m *PvpPlayAgainInviteResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvpPlayAgainInviteResp.Merge(m, src)
}
func (m *PvpPlayAgainInviteResp) XXX_Size() int {
	return xxx_messageInfo_PvpPlayAgainInviteResp.Size(m)
}
func (m *PvpPlayAgainInviteResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PvpPlayAgainInviteResp.DiscardUnknown(m)
}

var xxx_messageInfo_PvpPlayAgainInviteResp proto.InternalMessageInfo

func (m *PvpPlayAgainInviteResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

// pvp接受邀请
type PvpPlayAgainAcceptReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoomID               string   `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvpPlayAgainAcceptReq) Reset()         { *m = PvpPlayAgainAcceptReq{} }
func (m *PvpPlayAgainAcceptReq) String() string { return proto.CompactTextString(m) }
func (*PvpPlayAgainAcceptReq) ProtoMessage()    {}
func (*PvpPlayAgainAcceptReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0523d21ed0ac0a7d, []int{4}
}

func (m *PvpPlayAgainAcceptReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvpPlayAgainAcceptReq.Unmarshal(m, b)
}
func (m *PvpPlayAgainAcceptReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvpPlayAgainAcceptReq.Marshal(b, m, deterministic)
}
func (m *PvpPlayAgainAcceptReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvpPlayAgainAcceptReq.Merge(m, src)
}
func (m *PvpPlayAgainAcceptReq) XXX_Size() int {
	return xxx_messageInfo_PvpPlayAgainAcceptReq.Size(m)
}
func (m *PvpPlayAgainAcceptReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PvpPlayAgainAcceptReq.DiscardUnknown(m)
}

var xxx_messageInfo_PvpPlayAgainAcceptReq proto.InternalMessageInfo

func (m *PvpPlayAgainAcceptReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PvpPlayAgainAcceptReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

type PvpPlayAgainAcceptResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvpPlayAgainAcceptResp) Reset()         { *m = PvpPlayAgainAcceptResp{} }
func (m *PvpPlayAgainAcceptResp) String() string { return proto.CompactTextString(m) }
func (*PvpPlayAgainAcceptResp) ProtoMessage()    {}
func (*PvpPlayAgainAcceptResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0523d21ed0ac0a7d, []int{5}
}

func (m *PvpPlayAgainAcceptResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvpPlayAgainAcceptResp.Unmarshal(m, b)
}
func (m *PvpPlayAgainAcceptResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvpPlayAgainAcceptResp.Marshal(b, m, deterministic)
}
func (m *PvpPlayAgainAcceptResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvpPlayAgainAcceptResp.Merge(m, src)
}
func (m *PvpPlayAgainAcceptResp) XXX_Size() int {
	return xxx_messageInfo_PvpPlayAgainAcceptResp.Size(m)
}
func (m *PvpPlayAgainAcceptResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PvpPlayAgainAcceptResp.DiscardUnknown(m)
}

var xxx_messageInfo_PvpPlayAgainAcceptResp proto.InternalMessageInfo

func (m *PvpPlayAgainAcceptResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

// 取消邀请
type PvpPlayAgainCancelInviteReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoomID               string   `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvpPlayAgainCancelInviteReq) Reset()         { *m = PvpPlayAgainCancelInviteReq{} }
func (m *PvpPlayAgainCancelInviteReq) String() string { return proto.CompactTextString(m) }
func (*PvpPlayAgainCancelInviteReq) ProtoMessage()    {}
func (*PvpPlayAgainCancelInviteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0523d21ed0ac0a7d, []int{6}
}

func (m *PvpPlayAgainCancelInviteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvpPlayAgainCancelInviteReq.Unmarshal(m, b)
}
func (m *PvpPlayAgainCancelInviteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvpPlayAgainCancelInviteReq.Marshal(b, m, deterministic)
}
func (m *PvpPlayAgainCancelInviteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvpPlayAgainCancelInviteReq.Merge(m, src)
}
func (m *PvpPlayAgainCancelInviteReq) XXX_Size() int {
	return xxx_messageInfo_PvpPlayAgainCancelInviteReq.Size(m)
}
func (m *PvpPlayAgainCancelInviteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PvpPlayAgainCancelInviteReq.DiscardUnknown(m)
}

var xxx_messageInfo_PvpPlayAgainCancelInviteReq proto.InternalMessageInfo

func (m *PvpPlayAgainCancelInviteReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PvpPlayAgainCancelInviteReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

type PvpPlayAgainCancelInviteResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvpPlayAgainCancelInviteResp) Reset()         { *m = PvpPlayAgainCancelInviteResp{} }
func (m *PvpPlayAgainCancelInviteResp) String() string { return proto.CompactTextString(m) }
func (*PvpPlayAgainCancelInviteResp) ProtoMessage()    {}
func (*PvpPlayAgainCancelInviteResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0523d21ed0ac0a7d, []int{7}
}

func (m *PvpPlayAgainCancelInviteResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvpPlayAgainCancelInviteResp.Unmarshal(m, b)
}
func (m *PvpPlayAgainCancelInviteResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvpPlayAgainCancelInviteResp.Marshal(b, m, deterministic)
}
func (m *PvpPlayAgainCancelInviteResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvpPlayAgainCancelInviteResp.Merge(m, src)
}
func (m *PvpPlayAgainCancelInviteResp) XXX_Size() int {
	return xxx_messageInfo_PvpPlayAgainCancelInviteResp.Size(m)
}
func (m *PvpPlayAgainCancelInviteResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PvpPlayAgainCancelInviteResp.DiscardUnknown(m)
}

var xxx_messageInfo_PvpPlayAgainCancelInviteResp proto.InternalMessageInfo

func (m *PvpPlayAgainCancelInviteResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func init() {
	proto.RegisterEnum("Msg.PvpRoomStatus", PvpRoomStatus_name, PvpRoomStatus_value)
	proto.RegisterType((*PvpPlayAgainStatusReq)(nil), "Msg.PvpPlayAgainStatusReq")
	proto.RegisterType((*PvpPlayAgainStatusResp)(nil), "Msg.PvpPlayAgainStatusResp")
	proto.RegisterType((*PvpPlayAgainInviteReq)(nil), "Msg.PvpPlayAgainInviteReq")
	proto.RegisterType((*PvpPlayAgainInviteResp)(nil), "Msg.PvpPlayAgainInviteResp")
	proto.RegisterType((*PvpPlayAgainAcceptReq)(nil), "Msg.PvpPlayAgainAcceptReq")
	proto.RegisterType((*PvpPlayAgainAcceptResp)(nil), "Msg.PvpPlayAgainAcceptResp")
	proto.RegisterType((*PvpPlayAgainCancelInviteReq)(nil), "Msg.PvpPlayAgainCancelInviteReq")
	proto.RegisterType((*PvpPlayAgainCancelInviteResp)(nil), "Msg.PvpPlayAgainCancelInviteResp")
}

func init() { proto.RegisterFile("pvp_playagain.proto", fileDescriptor_0523d21ed0ac0a7d) }

var fileDescriptor_0523d21ed0ac0a7d = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x4f, 0xe2, 0x40,
	0x14, 0xc7, 0xb7, 0xb0, 0x90, 0xe5, 0xed, 0x42, 0x9a, 0xd9, 0x5d, 0x42, 0x76, 0x39, 0x10, 0xf6,
	0x42, 0xf6, 0xd0, 0x22, 0x1e, 0x0d, 0x31, 0x88, 0xc6, 0x10, 0x43, 0x42, 0xda, 0x9b, 0x17, 0x32,
	0xd4, 0x67, 0x69, 0x6c, 0x3b, 0xe3, 0xcc, 0xb4, 0xca, 0xc1, 0xbf, 0xc1, 0x7f, 0xd9, 0x74, 0x4a,
	0x15, 0x88, 0x26, 0xa8, 0xc7, 0xf7, 0x99, 0x37, 0x9f, 0xf7, 0xbe, 0xfd, 0x01, 0x3f, 0x79, 0xca,
	0xe7, 0x3c, 0xa4, 0x2b, 0xea, 0xd3, 0x20, 0xb6, 0xb8, 0x60, 0x8a, 0x91, 0xf2, 0x54, 0xfa, 0x7f,
	0x7e, 0x08, 0x94, 0x49, 0xa8, 0x72, 0xd4, 0x3d, 0x83, 0xdf, 0xb3, 0x94, 0xcf, 0x42, 0xba, 0x1a,
	0x65, 0x8d, 0xae, 0xa2, 0x2a, 0x91, 0x0e, 0xde, 0x92, 0x5f, 0x50, 0x51, 0xec, 0x06, 0xe3, 0x96,
	0xd1, 0x31, 0x7a, 0x35, 0x27, 0x2f, 0x48, 0x13, 0xaa, 0x82, 0xb1, 0x68, 0x72, 0xda, 0x2a, 0x69,
	0xbc, 0xae, 0xba, 0x8f, 0x06, 0x34, 0x5f, 0xf3, 0x48, 0x4e, 0xfe, 0x41, 0x35, 0x9f, 0xa8, 0x4d,
	0x8d, 0xc1, 0x77, 0x6b, 0x2a, 0x7d, 0xcb, 0xd1, 0xc8, 0x59, 0x1f, 0x91, 0x01, 0x40, 0x66, 0xca,
	0xaf, 0x69, 0x77, 0x63, 0x40, 0x74, 0xe3, 0x2c, 0xe5, 0xce, 0xf3, 0x89, 0xb3, 0xd1, 0x45, 0xda,
	0x50, 0x0b, 0xe4, 0x24, 0x4e, 0x03, 0x85, 0xa2, 0x55, 0xee, 0x18, 0xbd, 0x6f, 0xce, 0x0b, 0xd8,
	0x0d, 0x96, 0xe3, 0xf7, 0x07, 0x1b, 0x6e, 0xe7, 0x2a, 0x34, 0x7b, 0xe6, 0xda, 0xdd, 0x62, 0xe4,
	0x79, 0xc8, 0xd5, 0xa7, 0xb7, 0x28, 0x34, 0xfb, 0x6e, 0x71, 0x01, 0x7f, 0x37, 0xaf, 0x8f, 0x69,
	0xec, 0x61, 0xf8, 0xd1, 0x27, 0x32, 0x86, 0xf6, 0xdb, 0xb2, 0x3d, 0x37, 0xfa, 0xff, 0x00, 0xf5,
	0xad, 0x17, 0x4b, 0x9a, 0x40, 0xb2, 0x2f, 0xf6, 0x8e, 0x06, 0x2a, 0x88, 0xfd, 0x79, 0xa0, 0x7d,
	0xe6, 0x97, 0x5d, 0x4e, 0x75, 0x72, 0xd3, 0x20, 0x0d, 0x80, 0x8c, 0x7b, 0x7a, 0xba, 0x59, 0x22,
	0x75, 0xa8, 0x65, 0x75, 0x88, 0x34, 0x45, 0xb3, 0x5c, 0x94, 0x52, 0x51, 0xa1, 0xcc, 0xaf, 0x45,
	0xb7, 0xc0, 0xeb, 0x44, 0xa2, 0x59, 0x39, 0x39, 0xbe, 0x1c, 0x2e, 0x93, 0x45, 0xbf, 0xdf, 0xb7,
	0xee, 0x83, 0xf8, 0x8a, 0xc5, 0xbe, 0xe5, 0xb1, 0xc8, 0x76, 0x69, 0x22, 0xa9, 0x8f, 0xee, 0x92,
	0x31, 0x65, 0x9f, 0xd3, 0x08, 0x5d, 0x14, 0x29, 0x0a, 0x5b, 0xff, 0x29, 0x1e, 0x0b, 0xed, 0xf4,
	0xe0, 0xc8, 0xa7, 0x11, 0xf2, 0xc5, 0xa2, 0xaa, 0xd9, 0xe1, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb0, 0xbc, 0xd5, 0x6b, 0x67, 0x03, 0x00, 0x00,
}