// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basic.proto

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

// PlayerAttribute used for ui
type PlayerAttribute struct {
	PlayerID             string          `protobuf:"bytes,20,opt,name=playerID,proto3" json:"playerID,omitempty"`
	Nickname             string          `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Url                  string          `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Diamond              int64           `protobuf:"varint,3,opt,name=diamond,proto3" json:"diamond,omitempty"`
	Gold                 int64           `protobuf:"varint,4,opt,name=gold,proto3" json:"gold,omitempty"`
	Exp                  int64           `protobuf:"varint,5,opt,name=exp,proto3" json:"exp,omitempty"`
	ExpLimit             int64           `protobuf:"varint,6,opt,name=expLimit,proto3" json:"expLimit,omitempty"`
	Level                int64           `protobuf:"varint,7,opt,name=level,proto3" json:"level,omitempty"`
	Strength             int64           `protobuf:"varint,8,opt,name=strength,proto3" json:"strength,omitempty"`
	StrengthLimit        int64           `protobuf:"varint,9,opt,name=strengthLimit,proto3" json:"strengthLimit,omitempty"`
	NextUpdate           int64           `protobuf:"varint,10,opt,name=nextUpdate,proto3" json:"nextUpdate,omitempty"`
	Id                   string          `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
	Cup                  int64           `protobuf:"varint,12,opt,name=cup,proto3" json:"cup,omitempty"`
	Star                 int64           `protobuf:"varint,16,opt,name=star,proto3" json:"star,omitempty"`
	Sex                  int64           `protobuf:"varint,17,opt,name=sex,proto3" json:"sex,omitempty"`
	Arena                int64           `protobuf:"varint,18,opt,name=arena,proto3" json:"arena,omitempty"`
	PveStage             int64           `protobuf:"varint,19,opt,name=pveStage,proto3" json:"pveStage,omitempty"`
	PvpAwardGold         int64           `protobuf:"varint,21,opt,name=pvpAwardGold,proto3" json:"pvpAwardGold,omitempty"`
	Fashion              []int64         `protobuf:"varint,13,rep,packed,name=fashion,proto3" json:"fashion,omitempty"`
	Weapon               []int64         `protobuf:"varint,14,rep,packed,name=weapon,proto3" json:"weapon,omitempty"`
	RedDot               map[string]bool `protobuf:"bytes,15,rep,name=redDot,proto3" json:"redDot,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PlayerAttribute) Reset()         { *m = PlayerAttribute{} }
func (m *PlayerAttribute) String() string { return proto.CompactTextString(m) }
func (*PlayerAttribute) ProtoMessage()    {}
func (*PlayerAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc45f6ac745dd88, []int{0}
}

func (m *PlayerAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerAttribute.Unmarshal(m, b)
}
func (m *PlayerAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerAttribute.Marshal(b, m, deterministic)
}
func (m *PlayerAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerAttribute.Merge(m, src)
}
func (m *PlayerAttribute) XXX_Size() int {
	return xxx_messageInfo_PlayerAttribute.Size(m)
}
func (m *PlayerAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerAttribute proto.InternalMessageInfo

func (m *PlayerAttribute) GetPlayerID() string {
	if m != nil {
		return m.PlayerID
	}
	return ""
}

func (m *PlayerAttribute) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *PlayerAttribute) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PlayerAttribute) GetDiamond() int64 {
	if m != nil {
		return m.Diamond
	}
	return 0
}

func (m *PlayerAttribute) GetGold() int64 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *PlayerAttribute) GetExp() int64 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *PlayerAttribute) GetExpLimit() int64 {
	if m != nil {
		return m.ExpLimit
	}
	return 0
}

func (m *PlayerAttribute) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *PlayerAttribute) GetStrength() int64 {
	if m != nil {
		return m.Strength
	}
	return 0
}

func (m *PlayerAttribute) GetStrengthLimit() int64 {
	if m != nil {
		return m.StrengthLimit
	}
	return 0
}

func (m *PlayerAttribute) GetNextUpdate() int64 {
	if m != nil {
		return m.NextUpdate
	}
	return 0
}

func (m *PlayerAttribute) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PlayerAttribute) GetCup() int64 {
	if m != nil {
		return m.Cup
	}
	return 0
}

func (m *PlayerAttribute) GetStar() int64 {
	if m != nil {
		return m.Star
	}
	return 0
}

func (m *PlayerAttribute) GetSex() int64 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *PlayerAttribute) GetArena() int64 {
	if m != nil {
		return m.Arena
	}
	return 0
}

func (m *PlayerAttribute) GetPveStage() int64 {
	if m != nil {
		return m.PveStage
	}
	return 0
}

func (m *PlayerAttribute) GetPvpAwardGold() int64 {
	if m != nil {
		return m.PvpAwardGold
	}
	return 0
}

func (m *PlayerAttribute) GetFashion() []int64 {
	if m != nil {
		return m.Fashion
	}
	return nil
}

func (m *PlayerAttribute) GetWeapon() []int64 {
	if m != nil {
		return m.Weapon
	}
	return nil
}

func (m *PlayerAttribute) GetRedDot() map[string]bool {
	if m != nil {
		return m.RedDot
	}
	return nil
}

type Boxs struct {
	Box                  []*BoxItem `protobuf:"bytes,1,rep,name=box,proto3" json:"box,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Boxs) Reset()         { *m = Boxs{} }
func (m *Boxs) String() string { return proto.CompactTextString(m) }
func (*Boxs) ProtoMessage()    {}
func (*Boxs) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc45f6ac745dd88, []int{1}
}

func (m *Boxs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Boxs.Unmarshal(m, b)
}
func (m *Boxs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Boxs.Marshal(b, m, deterministic)
}
func (m *Boxs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Boxs.Merge(m, src)
}
func (m *Boxs) XXX_Size() int {
	return xxx_messageInfo_Boxs.Size(m)
}
func (m *Boxs) XXX_DiscardUnknown() {
	xxx_messageInfo_Boxs.DiscardUnknown(m)
}

var xxx_messageInfo_Boxs proto.InternalMessageInfo

func (m *Boxs) GetBox() []*BoxItem {
	if m != nil {
		return m.Box
	}
	return nil
}

type BoxItem struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Number               int64    `protobuf:"varint,2,opt,name=Number,proto3" json:"Number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoxItem) Reset()         { *m = BoxItem{} }
func (m *BoxItem) String() string { return proto.CompactTextString(m) }
func (*BoxItem) ProtoMessage()    {}
func (*BoxItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc45f6ac745dd88, []int{2}
}

func (m *BoxItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoxItem.Unmarshal(m, b)
}
func (m *BoxItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoxItem.Marshal(b, m, deterministic)
}
func (m *BoxItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoxItem.Merge(m, src)
}
func (m *BoxItem) XXX_Size() int {
	return xxx_messageInfo_BoxItem.Size(m)
}
func (m *BoxItem) XXX_DiscardUnknown() {
	xxx_messageInfo_BoxItem.DiscardUnknown(m)
}

var xxx_messageInfo_BoxItem proto.InternalMessageInfo

func (m *BoxItem) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *BoxItem) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

// ShootCollection is for player achievememt
type ShootCollection struct {
	Bullseye             int64    `protobuf:"varint,1,opt,name=bullseye,proto3" json:"bullseye,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShootCollection) Reset()         { *m = ShootCollection{} }
func (m *ShootCollection) String() string { return proto.CompactTextString(m) }
func (*ShootCollection) ProtoMessage()    {}
func (*ShootCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc45f6ac745dd88, []int{3}
}

func (m *ShootCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShootCollection.Unmarshal(m, b)
}
func (m *ShootCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShootCollection.Marshal(b, m, deterministic)
}
func (m *ShootCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShootCollection.Merge(m, src)
}
func (m *ShootCollection) XXX_Size() int {
	return xxx_messageInfo_ShootCollection.Size(m)
}
func (m *ShootCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_ShootCollection.DiscardUnknown(m)
}

var xxx_messageInfo_ShootCollection proto.InternalMessageInfo

func (m *ShootCollection) GetBullseye() int64 {
	if m != nil {
		return m.Bullseye
	}
	return 0
}

func init() {
	proto.RegisterType((*PlayerAttribute)(nil), "Msg.PlayerAttribute")
	proto.RegisterMapType((map[string]bool)(nil), "Msg.PlayerAttribute.RedDotEntry")
	proto.RegisterType((*Boxs)(nil), "Msg.Boxs")
	proto.RegisterType((*BoxItem)(nil), "Msg.BoxItem")
	proto.RegisterType((*ShootCollection)(nil), "Msg.ShootCollection")
}

func init() { proto.RegisterFile("basic.proto", fileDescriptor_0cc45f6ac745dd88) }

var fileDescriptor_0cc45f6ac745dd88 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0xd5, 0xa6, 0xeb, 0x3a, 0x77, 0xbf, 0x30, 0x03, 0x59, 0x7d, 0x98, 0xaa, 0x0a, 0xa1,
	0xbe, 0x90, 0x16, 0x10, 0xd2, 0x00, 0x21, 0xb4, 0x32, 0x34, 0x4d, 0x62, 0x08, 0xa5, 0xf0, 0xc2,
	0x9b, 0x93, 0x1c, 0x69, 0xb4, 0xc4, 0xb6, 0x6c, 0x27, 0x4b, 0xff, 0x6f, 0xfe, 0x00, 0xe4, 0x73,
	0x52, 0x6d, 0xbc, 0xdd, 0xe7, 0x7b, 0x17, 0xdf, 0x7d, 0x4f, 0x17, 0x32, 0x8e, 0xb9, 0xc9, 0x93,
	0x50, 0x69, 0x69, 0x25, 0x0d, 0x6e, 0x4d, 0x36, 0xfb, 0x3b, 0x20, 0x27, 0x3f, 0x0a, 0xbe, 0x05,
	0x7d, 0x69, 0xad, 0xce, 0xe3, 0xca, 0x02, 0x9d, 0x90, 0x91, 0x42, 0xe9, 0xe6, 0x8a, 0x9d, 0x4d,
	0x7b, 0xf3, 0x83, 0x68, 0xc7, 0x2e, 0x27, 0xf2, 0xe4, 0x4e, 0xf0, 0x12, 0x58, 0xcf, 0xe7, 0x3a,
	0xa6, 0xa7, 0x24, 0xa8, 0x74, 0xc1, 0xfa, 0x28, 0xbb, 0x90, 0x32, 0xb2, 0x9f, 0xe6, 0xbc, 0x94,
	0x22, 0x65, 0xc1, 0xb4, 0x37, 0x0f, 0xa2, 0x0e, 0x29, 0x25, 0x83, 0x4c, 0x16, 0x29, 0x1b, 0xa0,
	0x8c, 0xb1, 0xfb, 0x1e, 0x1a, 0xc5, 0xf6, 0x50, 0x72, 0xa1, 0xeb, 0x06, 0x8d, 0xfa, 0x96, 0x97,
	0xb9, 0x65, 0x43, 0x94, 0x77, 0x4c, 0xcf, 0xc8, 0x5e, 0x01, 0x35, 0x14, 0x6c, 0x1f, 0x13, 0x1e,
	0xdc, 0x17, 0xc6, 0x6a, 0x10, 0x99, 0xdd, 0xb0, 0x91, 0xff, 0xa2, 0x63, 0xfa, 0x82, 0x1c, 0x75,
	0xb1, 0x7f, 0xf2, 0x00, 0x0b, 0x1e, 0x8b, 0xf4, 0x9c, 0x10, 0x01, 0x8d, 0xfd, 0xa5, 0x52, 0x6e,
	0x81, 0x11, 0x2c, 0x79, 0xa0, 0xd0, 0x63, 0xd2, 0xcf, 0x53, 0x36, 0x46, 0x93, 0xfd, 0x1c, 0xa7,
	0x4e, 0x2a, 0xc5, 0x0e, 0xfd, 0xd4, 0x49, 0xa5, 0x9c, 0x37, 0x63, 0xb9, 0x66, 0xa7, 0xde, 0x9b,
	0x8b, 0x5d, 0x95, 0x81, 0x86, 0x3d, 0xf1, 0x55, 0x06, 0x1a, 0x37, 0x3f, 0xd7, 0x20, 0x38, 0xa3,
	0x7e, 0x7e, 0x04, 0xdc, 0x7d, 0x0d, 0x6b, 0xcb, 0x33, 0x60, 0x4f, 0xfd, 0xfc, 0x1d, 0xd3, 0x19,
	0x39, 0x54, 0xb5, 0xba, 0xbc, 0xe7, 0x3a, 0xbd, 0x76, 0xbb, 0x7b, 0x86, 0xf9, 0x47, 0x9a, 0xdb,
	0xf8, 0x1f, 0x6e, 0x36, 0xb9, 0x14, 0xec, 0x68, 0x1a, 0xb8, 0x8d, 0xb7, 0x48, 0x9f, 0x93, 0xe1,
	0x3d, 0x70, 0x25, 0x05, 0x3b, 0xc6, 0x44, 0x4b, 0xf4, 0x82, 0x0c, 0x35, 0xa4, 0x57, 0xd2, 0xb2,
	0x93, 0x69, 0x30, 0x1f, 0xbf, 0x99, 0x86, 0xb7, 0x26, 0x0b, 0xff, 0xbb, 0x89, 0x30, 0xc2, 0x92,
	0xaf, 0xc2, 0xea, 0x6d, 0xd4, 0xd6, 0x4f, 0xde, 0x93, 0xf1, 0x03, 0xd9, 0x59, 0xbc, 0x83, 0x6d,
	0x7b, 0x15, 0x2e, 0x74, 0x16, 0x6b, 0x5e, 0x54, 0x80, 0x27, 0x31, 0x8a, 0x3c, 0x7c, 0xe8, 0x5f,
	0xf4, 0x66, 0x2f, 0xc9, 0x60, 0x25, 0x1b, 0x43, 0xcf, 0x49, 0x10, 0xcb, 0x86, 0xf5, 0xb0, 0xf3,
	0x21, 0x76, 0x5e, 0xc9, 0xe6, 0xc6, 0x42, 0x19, 0xb9, 0xc4, 0xec, 0x1d, 0xd9, 0x6f, 0xd9, 0x6d,
	0xf5, 0xe7, 0x56, 0x75, 0x57, 0x87, 0xb1, 0xf3, 0xf4, 0xbd, 0x2a, 0x63, 0xd0, 0xd8, 0x21, 0x88,
	0x5a, 0x9a, 0xbd, 0x22, 0x27, 0xeb, 0x8d, 0x94, 0xf6, 0x8b, 0x2c, 0x0a, 0x48, 0xac, 0xb3, 0x3f,
	0x21, 0xa3, 0xb8, 0x2a, 0x0a, 0x03, 0x5b, 0xff, 0x44, 0x10, 0xed, 0x78, 0xf5, 0xf9, 0xf7, 0xa7,
	0x4d, 0x15, 0x2f, 0x97, 0xcb, 0xb0, 0xc9, 0x45, 0x2a, 0x45, 0x16, 0x26, 0xb2, 0x5c, 0xac, 0x79,
	0x65, 0x78, 0x06, 0xf8, 0xd0, 0xe2, 0x9a, 0x97, 0xb0, 0x06, 0x5d, 0x83, 0x5e, 0xe0, 0xdf, 0x93,
	0xc8, 0x62, 0x51, 0xbf, 0xfe, 0x98, 0xf1, 0x12, 0x54, 0x1c, 0x0f, 0x51, 0x7b, 0xfb, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0x6c, 0xa3, 0x3e, 0x68, 0x60, 0x03, 0x00, 0x00,
}