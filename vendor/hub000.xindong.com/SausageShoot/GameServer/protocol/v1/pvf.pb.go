// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pvf.proto

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
type PvfRoomStatus int32

const (
	PvfRoomStatus_waiting       PvfRoomStatus = 0
	PvfRoomStatus_gameSelection PvfRoomStatus = 1
	PvfRoomStatus_startGame     PvfRoomStatus = 2
	PvfRoomStatus_isContinue    PvfRoomStatus = 3
	PvfRoomStatus_leave         PvfRoomStatus = 4
	PvfRoomStatus_playAgain     PvfRoomStatus = 5
	PvfRoomStatus_cancel        PvfRoomStatus = 6
	PvfRoomStatus_refuse        PvfRoomStatus = 7
)

var PvfRoomStatus_name = map[int32]string{
	0: "waiting",
	1: "gameSelection",
	2: "startGame",
	3: "isContinue",
	4: "leave",
	5: "playAgain",
	6: "cancel",
	7: "refuse",
}

var PvfRoomStatus_value = map[string]int32{
	"waiting":       0,
	"gameSelection": 1,
	"startGame":     2,
	"isContinue":    3,
	"leave":         4,
	"playAgain":     5,
	"cancel":        6,
	"refuse":        7,
}

func (x PvfRoomStatus) String() string {
	return proto.EnumName(PvfRoomStatus_name, int32(x))
}

func (PvfRoomStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{0}
}

// 邀请好友
type PvfInviteReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	IsNew                bool     `protobuf:"varint,2,opt,name=isNew,proto3" json:"isNew,omitempty"`
	RoomID               string   `protobuf:"bytes,3,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfInviteReq) Reset()         { *m = PvfInviteReq{} }
func (m *PvfInviteReq) String() string { return proto.CompactTextString(m) }
func (*PvfInviteReq) ProtoMessage()    {}
func (*PvfInviteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{0}
}

func (m *PvfInviteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfInviteReq.Unmarshal(m, b)
}
func (m *PvfInviteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfInviteReq.Marshal(b, m, deterministic)
}
func (m *PvfInviteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfInviteReq.Merge(m, src)
}
func (m *PvfInviteReq) XXX_Size() int {
	return xxx_messageInfo_PvfInviteReq.Size(m)
}
func (m *PvfInviteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfInviteReq.DiscardUnknown(m)
}

var xxx_messageInfo_PvfInviteReq proto.InternalMessageInfo

func (m *PvfInviteReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PvfInviteReq) GetIsNew() bool {
	if m != nil {
		return m.IsNew
	}
	return false
}

func (m *PvfInviteReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

type PvfInviteResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	RoomID               string   `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	Link                 string   `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfInviteResp) Reset()         { *m = PvfInviteResp{} }
func (m *PvfInviteResp) String() string { return proto.CompactTextString(m) }
func (*PvfInviteResp) ProtoMessage()    {}
func (*PvfInviteResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{1}
}

func (m *PvfInviteResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfInviteResp.Unmarshal(m, b)
}
func (m *PvfInviteResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfInviteResp.Marshal(b, m, deterministic)
}
func (m *PvfInviteResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfInviteResp.Merge(m, src)
}
func (m *PvfInviteResp) XXX_Size() int {
	return xxx_messageInfo_PvfInviteResp.Size(m)
}
func (m *PvfInviteResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfInviteResp.DiscardUnknown(m)
}

var xxx_messageInfo_PvfInviteResp proto.InternalMessageInfo

func (m *PvfInviteResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func (m *PvfInviteResp) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func (m *PvfInviteResp) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

// 接受pvf邀请
type PvfAcceptReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoomID               string   `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfAcceptReq) Reset()         { *m = PvfAcceptReq{} }
func (m *PvfAcceptReq) String() string { return proto.CompactTextString(m) }
func (*PvfAcceptReq) ProtoMessage()    {}
func (*PvfAcceptReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{2}
}

func (m *PvfAcceptReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfAcceptReq.Unmarshal(m, b)
}
func (m *PvfAcceptReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfAcceptReq.Marshal(b, m, deterministic)
}
func (m *PvfAcceptReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfAcceptReq.Merge(m, src)
}
func (m *PvfAcceptReq) XXX_Size() int {
	return xxx_messageInfo_PvfAcceptReq.Size(m)
}
func (m *PvfAcceptReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfAcceptReq.DiscardUnknown(m)
}

var xxx_messageInfo_PvfAcceptReq proto.InternalMessageInfo

func (m *PvfAcceptReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PvfAcceptReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

type PvfAcceptResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfAcceptResp) Reset()         { *m = PvfAcceptResp{} }
func (m *PvfAcceptResp) String() string { return proto.CompactTextString(m) }
func (*PvfAcceptResp) ProtoMessage()    {}
func (*PvfAcceptResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{3}
}

func (m *PvfAcceptResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfAcceptResp.Unmarshal(m, b)
}
func (m *PvfAcceptResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfAcceptResp.Marshal(b, m, deterministic)
}
func (m *PvfAcceptResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfAcceptResp.Merge(m, src)
}
func (m *PvfAcceptResp) XXX_Size() int {
	return xxx_messageInfo_PvfAcceptResp.Size(m)
}
func (m *PvfAcceptResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfAcceptResp.DiscardUnknown(m)
}

var xxx_messageInfo_PvfAcceptResp proto.InternalMessageInfo

func (m *PvfAcceptResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

// pvf玩法选择
type PvfSelectReq struct {
	Token                string     `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoomID               string     `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	PlayType             PlayType   `protobuf:"varint,3,opt,name=playType,proto3,enum=Msg.PlayType" json:"playType,omitempty"`
	TargetType           TargetType `protobuf:"varint,4,opt,name=targetType,proto3,enum=Msg.TargetType" json:"targetType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PvfSelectReq) Reset()         { *m = PvfSelectReq{} }
func (m *PvfSelectReq) String() string { return proto.CompactTextString(m) }
func (*PvfSelectReq) ProtoMessage()    {}
func (*PvfSelectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{4}
}

func (m *PvfSelectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfSelectReq.Unmarshal(m, b)
}
func (m *PvfSelectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfSelectReq.Marshal(b, m, deterministic)
}
func (m *PvfSelectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfSelectReq.Merge(m, src)
}
func (m *PvfSelectReq) XXX_Size() int {
	return xxx_messageInfo_PvfSelectReq.Size(m)
}
func (m *PvfSelectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfSelectReq.DiscardUnknown(m)
}

var xxx_messageInfo_PvfSelectReq proto.InternalMessageInfo

func (m *PvfSelectReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PvfSelectReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func (m *PvfSelectReq) GetPlayType() PlayType {
	if m != nil {
		return m.PlayType
	}
	return PlayType_Time
}

func (m *PvfSelectReq) GetTargetType() TargetType {
	if m != nil {
		return m.TargetType
	}
	return TargetType_Standard
}

type PvfSelectResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfSelectResp) Reset()         { *m = PvfSelectResp{} }
func (m *PvfSelectResp) String() string { return proto.CompactTextString(m) }
func (*PvfSelectResp) ProtoMessage()    {}
func (*PvfSelectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{5}
}

func (m *PvfSelectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfSelectResp.Unmarshal(m, b)
}
func (m *PvfSelectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfSelectResp.Marshal(b, m, deterministic)
}
func (m *PvfSelectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfSelectResp.Merge(m, src)
}
func (m *PvfSelectResp) XXX_Size() int {
	return xxx_messageInfo_PvfSelectResp.Size(m)
}
func (m *PvfSelectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfSelectResp.DiscardUnknown(m)
}

var xxx_messageInfo_PvfSelectResp proto.InternalMessageInfo

func (m *PvfSelectResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

// pvf状态
type PvfStatusReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoomID               string   `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfStatusReq) Reset()         { *m = PvfStatusReq{} }
func (m *PvfStatusReq) String() string { return proto.CompactTextString(m) }
func (*PvfStatusReq) ProtoMessage()    {}
func (*PvfStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{6}
}

func (m *PvfStatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfStatusReq.Unmarshal(m, b)
}
func (m *PvfStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfStatusReq.Marshal(b, m, deterministic)
}
func (m *PvfStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfStatusReq.Merge(m, src)
}
func (m *PvfStatusReq) XXX_Size() int {
	return xxx_messageInfo_PvfStatusReq.Size(m)
}
func (m *PvfStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_PvfStatusReq proto.InternalMessageInfo

func (m *PvfStatusReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PvfStatusReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

type PvfStatusResp struct {
	Result               Result           `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	Status               PvfRoomStatus    `protobuf:"varint,2,opt,name=status,proto3,enum=Msg.PvfRoomStatus" json:"status,omitempty"`
	IsOwner              bool             `protobuf:"varint,3,opt,name=isOwner,proto3" json:"isOwner,omitempty"`
	Opponent             *PlayerAttribute `protobuf:"bytes,4,opt,name=opponent,proto3" json:"opponent,omitempty"`
	Joining              bool             `protobuf:"varint,5,opt,name=joining,proto3" json:"joining,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PvfStatusResp) Reset()         { *m = PvfStatusResp{} }
func (m *PvfStatusResp) String() string { return proto.CompactTextString(m) }
func (*PvfStatusResp) ProtoMessage()    {}
func (*PvfStatusResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{7}
}

func (m *PvfStatusResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfStatusResp.Unmarshal(m, b)
}
func (m *PvfStatusResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfStatusResp.Marshal(b, m, deterministic)
}
func (m *PvfStatusResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfStatusResp.Merge(m, src)
}
func (m *PvfStatusResp) XXX_Size() int {
	return xxx_messageInfo_PvfStatusResp.Size(m)
}
func (m *PvfStatusResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfStatusResp.DiscardUnknown(m)
}

var xxx_messageInfo_PvfStatusResp proto.InternalMessageInfo

func (m *PvfStatusResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func (m *PvfStatusResp) GetStatus() PvfRoomStatus {
	if m != nil {
		return m.Status
	}
	return PvfRoomStatus_waiting
}

func (m *PvfStatusResp) GetIsOwner() bool {
	if m != nil {
		return m.IsOwner
	}
	return false
}

func (m *PvfStatusResp) GetOpponent() *PlayerAttribute {
	if m != nil {
		return m.Opponent
	}
	return nil
}

func (m *PvfStatusResp) GetJoining() bool {
	if m != nil {
		return m.Joining
	}
	return false
}

// 再来一局
type PvfPlayAgainReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoomID               string   `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfPlayAgainReq) Reset()         { *m = PvfPlayAgainReq{} }
func (m *PvfPlayAgainReq) String() string { return proto.CompactTextString(m) }
func (*PvfPlayAgainReq) ProtoMessage()    {}
func (*PvfPlayAgainReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{8}
}

func (m *PvfPlayAgainReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfPlayAgainReq.Unmarshal(m, b)
}
func (m *PvfPlayAgainReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfPlayAgainReq.Marshal(b, m, deterministic)
}
func (m *PvfPlayAgainReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfPlayAgainReq.Merge(m, src)
}
func (m *PvfPlayAgainReq) XXX_Size() int {
	return xxx_messageInfo_PvfPlayAgainReq.Size(m)
}
func (m *PvfPlayAgainReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfPlayAgainReq.DiscardUnknown(m)
}

var xxx_messageInfo_PvfPlayAgainReq proto.InternalMessageInfo

func (m *PvfPlayAgainReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PvfPlayAgainReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

type PvfPlayAgainResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfPlayAgainResp) Reset()         { *m = PvfPlayAgainResp{} }
func (m *PvfPlayAgainResp) String() string { return proto.CompactTextString(m) }
func (*PvfPlayAgainResp) ProtoMessage()    {}
func (*PvfPlayAgainResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{9}
}

func (m *PvfPlayAgainResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfPlayAgainResp.Unmarshal(m, b)
}
func (m *PvfPlayAgainResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfPlayAgainResp.Marshal(b, m, deterministic)
}
func (m *PvfPlayAgainResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfPlayAgainResp.Merge(m, src)
}
func (m *PvfPlayAgainResp) XXX_Size() int {
	return xxx_messageInfo_PvfPlayAgainResp.Size(m)
}
func (m *PvfPlayAgainResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfPlayAgainResp.DiscardUnknown(m)
}

var xxx_messageInfo_PvfPlayAgainResp proto.InternalMessageInfo

func (m *PvfPlayAgainResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

// 取消再来一局
type PvfCacnelPlayAgainReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoomID               string   `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfCacnelPlayAgainReq) Reset()         { *m = PvfCacnelPlayAgainReq{} }
func (m *PvfCacnelPlayAgainReq) String() string { return proto.CompactTextString(m) }
func (*PvfCacnelPlayAgainReq) ProtoMessage()    {}
func (*PvfCacnelPlayAgainReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{10}
}

func (m *PvfCacnelPlayAgainReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfCacnelPlayAgainReq.Unmarshal(m, b)
}
func (m *PvfCacnelPlayAgainReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfCacnelPlayAgainReq.Marshal(b, m, deterministic)
}
func (m *PvfCacnelPlayAgainReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfCacnelPlayAgainReq.Merge(m, src)
}
func (m *PvfCacnelPlayAgainReq) XXX_Size() int {
	return xxx_messageInfo_PvfCacnelPlayAgainReq.Size(m)
}
func (m *PvfCacnelPlayAgainReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfCacnelPlayAgainReq.DiscardUnknown(m)
}

var xxx_messageInfo_PvfCacnelPlayAgainReq proto.InternalMessageInfo

func (m *PvfCacnelPlayAgainReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PvfCacnelPlayAgainReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

type PvfCacnelPlayAgainResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfCacnelPlayAgainResp) Reset()         { *m = PvfCacnelPlayAgainResp{} }
func (m *PvfCacnelPlayAgainResp) String() string { return proto.CompactTextString(m) }
func (*PvfCacnelPlayAgainResp) ProtoMessage()    {}
func (*PvfCacnelPlayAgainResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{11}
}

func (m *PvfCacnelPlayAgainResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfCacnelPlayAgainResp.Unmarshal(m, b)
}
func (m *PvfCacnelPlayAgainResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfCacnelPlayAgainResp.Marshal(b, m, deterministic)
}
func (m *PvfCacnelPlayAgainResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfCacnelPlayAgainResp.Merge(m, src)
}
func (m *PvfCacnelPlayAgainResp) XXX_Size() int {
	return xxx_messageInfo_PvfCacnelPlayAgainResp.Size(m)
}
func (m *PvfCacnelPlayAgainResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfCacnelPlayAgainResp.DiscardUnknown(m)
}

var xxx_messageInfo_PvfCacnelPlayAgainResp proto.InternalMessageInfo

func (m *PvfCacnelPlayAgainResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

// 再来一局回复
type PvfPlayAgainResponseReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoomID               string   `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	IsNext               bool     `protobuf:"varint,3,opt,name=isNext,proto3" json:"isNext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfPlayAgainResponseReq) Reset()         { *m = PvfPlayAgainResponseReq{} }
func (m *PvfPlayAgainResponseReq) String() string { return proto.CompactTextString(m) }
func (*PvfPlayAgainResponseReq) ProtoMessage()    {}
func (*PvfPlayAgainResponseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{12}
}

func (m *PvfPlayAgainResponseReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfPlayAgainResponseReq.Unmarshal(m, b)
}
func (m *PvfPlayAgainResponseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfPlayAgainResponseReq.Marshal(b, m, deterministic)
}
func (m *PvfPlayAgainResponseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfPlayAgainResponseReq.Merge(m, src)
}
func (m *PvfPlayAgainResponseReq) XXX_Size() int {
	return xxx_messageInfo_PvfPlayAgainResponseReq.Size(m)
}
func (m *PvfPlayAgainResponseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfPlayAgainResponseReq.DiscardUnknown(m)
}

var xxx_messageInfo_PvfPlayAgainResponseReq proto.InternalMessageInfo

func (m *PvfPlayAgainResponseReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PvfPlayAgainResponseReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func (m *PvfPlayAgainResponseReq) GetIsNext() bool {
	if m != nil {
		return m.IsNext
	}
	return false
}

type PvfPlayAgainResponseResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfPlayAgainResponseResp) Reset()         { *m = PvfPlayAgainResponseResp{} }
func (m *PvfPlayAgainResponseResp) String() string { return proto.CompactTextString(m) }
func (*PvfPlayAgainResponseResp) ProtoMessage()    {}
func (*PvfPlayAgainResponseResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{13}
}

func (m *PvfPlayAgainResponseResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfPlayAgainResponseResp.Unmarshal(m, b)
}
func (m *PvfPlayAgainResponseResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfPlayAgainResponseResp.Marshal(b, m, deterministic)
}
func (m *PvfPlayAgainResponseResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfPlayAgainResponseResp.Merge(m, src)
}
func (m *PvfPlayAgainResponseResp) XXX_Size() int {
	return xxx_messageInfo_PvfPlayAgainResponseResp.Size(m)
}
func (m *PvfPlayAgainResponseResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfPlayAgainResponseResp.DiscardUnknown(m)
}

var xxx_messageInfo_PvfPlayAgainResponseResp proto.InternalMessageInfo

func (m *PvfPlayAgainResponseResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

// 离开结算页面请求
type LeaveCountPageReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoomID               string   `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	IsPvf                bool     `protobuf:"varint,3,opt,name=isPvf,proto3" json:"isPvf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveCountPageReq) Reset()         { *m = LeaveCountPageReq{} }
func (m *LeaveCountPageReq) String() string { return proto.CompactTextString(m) }
func (*LeaveCountPageReq) ProtoMessage()    {}
func (*LeaveCountPageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{14}
}

func (m *LeaveCountPageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveCountPageReq.Unmarshal(m, b)
}
func (m *LeaveCountPageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveCountPageReq.Marshal(b, m, deterministic)
}
func (m *LeaveCountPageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveCountPageReq.Merge(m, src)
}
func (m *LeaveCountPageReq) XXX_Size() int {
	return xxx_messageInfo_LeaveCountPageReq.Size(m)
}
func (m *LeaveCountPageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveCountPageReq.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveCountPageReq proto.InternalMessageInfo

func (m *LeaveCountPageReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LeaveCountPageReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func (m *LeaveCountPageReq) GetIsPvf() bool {
	if m != nil {
		return m.IsPvf
	}
	return false
}

type LeaveCountPageResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveCountPageResp) Reset()         { *m = LeaveCountPageResp{} }
func (m *LeaveCountPageResp) String() string { return proto.CompactTextString(m) }
func (*LeaveCountPageResp) ProtoMessage()    {}
func (*LeaveCountPageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{15}
}

func (m *LeaveCountPageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveCountPageResp.Unmarshal(m, b)
}
func (m *LeaveCountPageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveCountPageResp.Marshal(b, m, deterministic)
}
func (m *LeaveCountPageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveCountPageResp.Merge(m, src)
}
func (m *LeaveCountPageResp) XXX_Size() int {
	return xxx_messageInfo_LeaveCountPageResp.Size(m)
}
func (m *LeaveCountPageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveCountPageResp.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveCountPageResp proto.InternalMessageInfo

func (m *LeaveCountPageResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

type PvfEnterGameReq struct {
	RoomID               string   `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfEnterGameReq) Reset()         { *m = PvfEnterGameReq{} }
func (m *PvfEnterGameReq) String() string { return proto.CompactTextString(m) }
func (*PvfEnterGameReq) ProtoMessage()    {}
func (*PvfEnterGameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{16}
}

func (m *PvfEnterGameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfEnterGameReq.Unmarshal(m, b)
}
func (m *PvfEnterGameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfEnterGameReq.Marshal(b, m, deterministic)
}
func (m *PvfEnterGameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfEnterGameReq.Merge(m, src)
}
func (m *PvfEnterGameReq) XXX_Size() int {
	return xxx_messageInfo_PvfEnterGameReq.Size(m)
}
func (m *PvfEnterGameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfEnterGameReq.DiscardUnknown(m)
}

var xxx_messageInfo_PvfEnterGameReq proto.InternalMessageInfo

func (m *PvfEnterGameReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

type PvfEnterGameResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PvfEnterGameResp) Reset()         { *m = PvfEnterGameResp{} }
func (m *PvfEnterGameResp) String() string { return proto.CompactTextString(m) }
func (*PvfEnterGameResp) ProtoMessage()    {}
func (*PvfEnterGameResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5069751322eedbc0, []int{17}
}

func (m *PvfEnterGameResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PvfEnterGameResp.Unmarshal(m, b)
}
func (m *PvfEnterGameResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PvfEnterGameResp.Marshal(b, m, deterministic)
}
func (m *PvfEnterGameResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PvfEnterGameResp.Merge(m, src)
}
func (m *PvfEnterGameResp) XXX_Size() int {
	return xxx_messageInfo_PvfEnterGameResp.Size(m)
}
func (m *PvfEnterGameResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PvfEnterGameResp.DiscardUnknown(m)
}

var xxx_messageInfo_PvfEnterGameResp proto.InternalMessageInfo

func (m *PvfEnterGameResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func init() {
	proto.RegisterEnum("Msg.PvfRoomStatus", PvfRoomStatus_name, PvfRoomStatus_value)
	proto.RegisterType((*PvfInviteReq)(nil), "Msg.PvfInviteReq")
	proto.RegisterType((*PvfInviteResp)(nil), "Msg.PvfInviteResp")
	proto.RegisterType((*PvfAcceptReq)(nil), "Msg.PvfAcceptReq")
	proto.RegisterType((*PvfAcceptResp)(nil), "Msg.PvfAcceptResp")
	proto.RegisterType((*PvfSelectReq)(nil), "Msg.PvfSelectReq")
	proto.RegisterType((*PvfSelectResp)(nil), "Msg.PvfSelectResp")
	proto.RegisterType((*PvfStatusReq)(nil), "Msg.PvfStatusReq")
	proto.RegisterType((*PvfStatusResp)(nil), "Msg.PvfStatusResp")
	proto.RegisterType((*PvfPlayAgainReq)(nil), "Msg.PvfPlayAgainReq")
	proto.RegisterType((*PvfPlayAgainResp)(nil), "Msg.PvfPlayAgainResp")
	proto.RegisterType((*PvfCacnelPlayAgainReq)(nil), "Msg.PvfCacnelPlayAgainReq")
	proto.RegisterType((*PvfCacnelPlayAgainResp)(nil), "Msg.PvfCacnelPlayAgainResp")
	proto.RegisterType((*PvfPlayAgainResponseReq)(nil), "Msg.PvfPlayAgainResponseReq")
	proto.RegisterType((*PvfPlayAgainResponseResp)(nil), "Msg.PvfPlayAgainResponseResp")
	proto.RegisterType((*LeaveCountPageReq)(nil), "Msg.LeaveCountPageReq")
	proto.RegisterType((*LeaveCountPageResp)(nil), "Msg.LeaveCountPageResp")
	proto.RegisterType((*PvfEnterGameReq)(nil), "Msg.PvfEnterGameReq")
	proto.RegisterType((*PvfEnterGameResp)(nil), "Msg.PvfEnterGameResp")
}

func init() { proto.RegisterFile("pvf.proto", fileDescriptor_5069751322eedbc0) }

var fileDescriptor_5069751322eedbc0 = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5d, 0x4f, 0xdb, 0x30,
	0x14, 0x5d, 0x80, 0x86, 0xf6, 0x96, 0x96, 0x60, 0x31, 0x56, 0xf1, 0x84, 0xb2, 0x17, 0xe0, 0xa1,
	0xed, 0xba, 0x49, 0xd3, 0xb4, 0x21, 0xc4, 0x18, 0x42, 0x48, 0x63, 0x8b, 0x52, 0xa4, 0x49, 0x7b,
	0xd9, 0xdc, 0xec, 0x26, 0x78, 0xa4, 0x76, 0x16, 0x3b, 0x01, 0xde, 0xf6, 0x33, 0xf6, 0x8b, 0xf6,
	0xbb, 0xa6, 0x38, 0x6e, 0xe9, 0x50, 0x91, 0xd2, 0xed, 0xcd, 0xe7, 0xe6, 0xf8, 0x9c, 0xe3, 0xeb,
	0x8f, 0x40, 0x23, 0xc9, 0xc3, 0x6e, 0x92, 0x0a, 0x25, 0xc8, 0xf2, 0xb9, 0x8c, 0xb6, 0xd7, 0x52,
	0x94, 0x59, 0xac, 0xca, 0xd2, 0x76, 0x23, 0xc9, 0x13, 0x33, 0x6c, 0x8e, 0xa8, 0x64, 0x41, 0x09,
	0x5c, 0x1f, 0xd6, 0xbc, 0x3c, 0x3c, 0xe3, 0x39, 0x53, 0xe8, 0xe3, 0x0f, 0xb2, 0x09, 0x35, 0x25,
	0xae, 0x90, 0x77, 0xac, 0x1d, 0x6b, 0xb7, 0xe1, 0x97, 0xa0, 0xa8, 0x32, 0xf9, 0x01, 0xaf, 0x3b,
	0x4b, 0x3b, 0xd6, 0x6e, 0xdd, 0x2f, 0x01, 0xd9, 0x02, 0x3b, 0x15, 0x62, 0x7c, 0xf6, 0xae, 0xb3,
	0xac, 0xc9, 0x06, 0xb9, 0x5f, 0xa1, 0x35, 0xa3, 0x29, 0x13, 0xf2, 0x14, 0xec, 0x32, 0x8c, 0x56,
	0x6d, 0x0f, 0x9a, 0xdd, 0x73, 0x19, 0x75, 0x7d, 0x5d, 0xf2, 0xcd, 0xa7, 0x19, 0xb5, 0xa5, 0x59,
	0x35, 0x42, 0x60, 0x25, 0x66, 0xfc, 0xca, 0x78, 0xe8, 0xb1, 0xfb, 0x46, 0xa7, 0x3e, 0x0a, 0x02,
	0x4c, 0xd4, 0xc3, 0xa9, 0x1f, 0x50, 0x74, 0x5f, 0xe8, 0x7c, 0x93, 0xd9, 0x15, 0xf3, 0xb9, 0xbf,
	0x2c, 0x6d, 0x3a, 0xc4, 0x18, 0x83, 0xc5, 0x4d, 0xc9, 0x1e, 0xd4, 0x93, 0x98, 0xde, 0x5e, 0xdc,
	0x26, 0xa8, 0x97, 0xd2, 0x1e, 0xb4, 0xb4, 0x8b, 0x67, 0x8a, 0xfe, 0xf4, 0x33, 0xe9, 0x01, 0x28,
	0x9a, 0x46, 0xa8, 0x34, 0x79, 0x45, 0x93, 0xd7, 0x35, 0xf9, 0x62, 0x5a, 0xf6, 0x67, 0x28, 0x66,
	0x41, 0x93, 0x64, 0x55, 0x17, 0x54, 0x36, 0x71, 0xa8, 0xa8, 0xca, 0xe4, 0xe2, 0x4d, 0xfc, 0x6d,
	0x95, 0xa6, 0x66, 0x7a, 0xd5, 0x5d, 0xde, 0x07, 0x5b, 0xea, 0x29, 0x5a, 0xae, 0x3d, 0x20, 0x65,
	0x13, 0xf2, 0xd0, 0x17, 0x62, 0x6c, 0xc4, 0x0c, 0x83, 0x74, 0x60, 0x95, 0xc9, 0x8f, 0xd7, 0x1c,
	0x53, 0xdd, 0xb1, 0xba, 0x3f, 0x81, 0xa4, 0x0f, 0x75, 0x91, 0x24, 0x82, 0x23, 0x57, 0xba, 0x3f,
	0xcd, 0xc1, 0xe6, 0xb4, 0x99, 0x98, 0x1e, 0x29, 0x95, 0xb2, 0x51, 0xa6, 0xd0, 0x9f, 0xb2, 0x0a,
	0xad, 0xef, 0x82, 0x71, 0xc6, 0xa3, 0x4e, 0xad, 0xd4, 0x32, 0xd0, 0x3d, 0x84, 0x75, 0x2f, 0x0f,
	0x8b, 0x99, 0x47, 0x11, 0x65, 0x7c, 0xf1, 0x4e, 0xbc, 0x04, 0xe7, 0x6f, 0x81, 0xaa, 0x1b, 0x70,
	0x02, 0x8f, 0xbd, 0x3c, 0x3c, 0xa6, 0x01, 0xc7, 0xf8, 0x3f, 0xfc, 0x0f, 0x60, 0x6b, 0x9e, 0x4c,
	0xd5, 0x14, 0x5f, 0xe0, 0xc9, 0xfd, 0xf8, 0x82, 0x4b, 0x5c, 0xfc, 0x84, 0x6f, 0x81, 0x5d, 0xbc,
	0x0b, 0x37, 0xca, 0xec, 0x96, 0x41, 0xee, 0x21, 0x74, 0xe6, 0x1b, 0x54, 0x4d, 0xf8, 0x09, 0x36,
	0xde, 0x23, 0xcd, 0xf1, 0x58, 0x64, 0x5c, 0x79, 0x34, 0xfa, 0x87, 0x6c, 0xfa, 0x01, 0xf3, 0xf2,
	0xd0, 0x44, 0x2b, 0x81, 0xfb, 0x0a, 0xc8, 0x7d, 0xe1, 0xaa, 0x99, 0xf6, 0xf4, 0xa9, 0x39, 0xe1,
	0x0a, 0xd3, 0x53, 0x3a, 0xd6, 0x89, 0xee, 0xbc, 0xad, 0x39, 0xe7, 0x63, 0x86, 0x5a, 0xd1, 0x63,
	0xff, 0x67, 0x79, 0xc5, 0xee, 0x6e, 0x06, 0x69, 0xc2, 0xea, 0x35, 0x65, 0x8a, 0xf1, 0xc8, 0x79,
	0x44, 0x36, 0xa0, 0x15, 0xd1, 0x31, 0x96, 0xd7, 0x9e, 0x09, 0xee, 0x58, 0xa4, 0x05, 0x0d, 0xa9,
	0x68, 0xaa, 0x0a, 0x1f, 0x67, 0x89, 0xb4, 0x01, 0x98, 0x3c, 0x16, 0x5c, 0x31, 0x9e, 0xa1, 0xb3,
	0x4c, 0x1a, 0x50, 0x8b, 0x8b, 0xf5, 0x3a, 0x2b, 0x05, 0x33, 0x99, 0xec, 0x88, 0x53, 0x23, 0x00,
	0x76, 0x40, 0x79, 0x80, 0xb1, 0x63, 0x17, 0xe3, 0x14, 0xc3, 0x4c, 0xa2, 0xb3, 0xfa, 0xf6, 0xf0,
	0xf3, 0xc1, 0x65, 0x36, 0xea, 0xf7, 0xfb, 0xdd, 0x1b, 0xc6, 0xbf, 0x09, 0x1e, 0x75, 0x03, 0x31,
	0xee, 0x0d, 0x69, 0x26, 0x69, 0x84, 0xc3, 0x4b, 0x21, 0x54, 0xef, 0x54, 0x67, 0x48, 0x73, 0x4c,
	0x7b, 0xfa, 0x97, 0x12, 0x88, 0xb8, 0x97, 0x3f, 0x7b, 0x5d, 0x44, 0x4b, 0x46, 0x23, 0x5b, 0xd7,
	0x9e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xda, 0x05, 0xa6, 0xbc, 0x9e, 0x06, 0x00, 0x00,
}