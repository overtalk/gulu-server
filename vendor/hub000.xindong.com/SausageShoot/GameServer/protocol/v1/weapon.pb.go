// Code generated by protoc-gen-go. DO NOT EDIT.
// source: weapon.proto

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

type GetPlayerWeaponReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerWeaponReq) Reset()         { *m = GetPlayerWeaponReq{} }
func (m *GetPlayerWeaponReq) String() string { return proto.CompactTextString(m) }
func (*GetPlayerWeaponReq) ProtoMessage()    {}
func (*GetPlayerWeaponReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e07c0075c06b2ab6, []int{0}
}

func (m *GetPlayerWeaponReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerWeaponReq.Unmarshal(m, b)
}
func (m *GetPlayerWeaponReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerWeaponReq.Marshal(b, m, deterministic)
}
func (m *GetPlayerWeaponReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerWeaponReq.Merge(m, src)
}
func (m *GetPlayerWeaponReq) XXX_Size() int {
	return xxx_messageInfo_GetPlayerWeaponReq.Size(m)
}
func (m *GetPlayerWeaponReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerWeaponReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerWeaponReq proto.InternalMessageInfo

func (m *GetPlayerWeaponReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetPlayerWeaponResp struct {
	Result               Result         `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	Weapons              string         `protobuf:"bytes,2,opt,name=weapons,proto3" json:"weapons,omitempty"`
	IsOld                map[int64]bool `protobuf:"bytes,3,rep,name=isOld,proto3" json:"isOld,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetPlayerWeaponResp) Reset()         { *m = GetPlayerWeaponResp{} }
func (m *GetPlayerWeaponResp) String() string { return proto.CompactTextString(m) }
func (*GetPlayerWeaponResp) ProtoMessage()    {}
func (*GetPlayerWeaponResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e07c0075c06b2ab6, []int{1}
}

func (m *GetPlayerWeaponResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerWeaponResp.Unmarshal(m, b)
}
func (m *GetPlayerWeaponResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerWeaponResp.Marshal(b, m, deterministic)
}
func (m *GetPlayerWeaponResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerWeaponResp.Merge(m, src)
}
func (m *GetPlayerWeaponResp) XXX_Size() int {
	return xxx_messageInfo_GetPlayerWeaponResp.Size(m)
}
func (m *GetPlayerWeaponResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerWeaponResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerWeaponResp proto.InternalMessageInfo

func (m *GetPlayerWeaponResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func (m *GetPlayerWeaponResp) GetWeapons() string {
	if m != nil {
		return m.Weapons
	}
	return ""
}

func (m *GetPlayerWeaponResp) GetIsOld() map[int64]bool {
	if m != nil {
		return m.IsOld
	}
	return nil
}

type CheckWeaponReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	WeaponID             int64    `protobuf:"varint,2,opt,name=weaponID,proto3" json:"weaponID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckWeaponReq) Reset()         { *m = CheckWeaponReq{} }
func (m *CheckWeaponReq) String() string { return proto.CompactTextString(m) }
func (*CheckWeaponReq) ProtoMessage()    {}
func (*CheckWeaponReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e07c0075c06b2ab6, []int{2}
}

func (m *CheckWeaponReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckWeaponReq.Unmarshal(m, b)
}
func (m *CheckWeaponReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckWeaponReq.Marshal(b, m, deterministic)
}
func (m *CheckWeaponReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckWeaponReq.Merge(m, src)
}
func (m *CheckWeaponReq) XXX_Size() int {
	return xxx_messageInfo_CheckWeaponReq.Size(m)
}
func (m *CheckWeaponReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckWeaponReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckWeaponReq proto.InternalMessageInfo

func (m *CheckWeaponReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CheckWeaponReq) GetWeaponID() int64 {
	if m != nil {
		return m.WeaponID
	}
	return 0
}

type CheckWeaponResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	WeaponID             int64    `protobuf:"varint,2,opt,name=weaponID,proto3" json:"weaponID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckWeaponResp) Reset()         { *m = CheckWeaponResp{} }
func (m *CheckWeaponResp) String() string { return proto.CompactTextString(m) }
func (*CheckWeaponResp) ProtoMessage()    {}
func (*CheckWeaponResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e07c0075c06b2ab6, []int{3}
}

func (m *CheckWeaponResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckWeaponResp.Unmarshal(m, b)
}
func (m *CheckWeaponResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckWeaponResp.Marshal(b, m, deterministic)
}
func (m *CheckWeaponResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckWeaponResp.Merge(m, src)
}
func (m *CheckWeaponResp) XXX_Size() int {
	return xxx_messageInfo_CheckWeaponResp.Size(m)
}
func (m *CheckWeaponResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckWeaponResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckWeaponResp proto.InternalMessageInfo

func (m *CheckWeaponResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func (m *CheckWeaponResp) GetWeaponID() int64 {
	if m != nil {
		return m.WeaponID
	}
	return 0
}

type UpdateWeaponReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Type                 int64    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	WeaponID             int64    `protobuf:"varint,3,opt,name=weaponID,proto3" json:"weaponID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateWeaponReq) Reset()         { *m = UpdateWeaponReq{} }
func (m *UpdateWeaponReq) String() string { return proto.CompactTextString(m) }
func (*UpdateWeaponReq) ProtoMessage()    {}
func (*UpdateWeaponReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e07c0075c06b2ab6, []int{4}
}

func (m *UpdateWeaponReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWeaponReq.Unmarshal(m, b)
}
func (m *UpdateWeaponReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWeaponReq.Marshal(b, m, deterministic)
}
func (m *UpdateWeaponReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWeaponReq.Merge(m, src)
}
func (m *UpdateWeaponReq) XXX_Size() int {
	return xxx_messageInfo_UpdateWeaponReq.Size(m)
}
func (m *UpdateWeaponReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWeaponReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWeaponReq proto.InternalMessageInfo

func (m *UpdateWeaponReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateWeaponReq) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *UpdateWeaponReq) GetWeaponID() int64 {
	if m != nil {
		return m.WeaponID
	}
	return 0
}

type UpdateWeaponResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateWeaponResp) Reset()         { *m = UpdateWeaponResp{} }
func (m *UpdateWeaponResp) String() string { return proto.CompactTextString(m) }
func (*UpdateWeaponResp) ProtoMessage()    {}
func (*UpdateWeaponResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e07c0075c06b2ab6, []int{5}
}

func (m *UpdateWeaponResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWeaponResp.Unmarshal(m, b)
}
func (m *UpdateWeaponResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWeaponResp.Marshal(b, m, deterministic)
}
func (m *UpdateWeaponResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWeaponResp.Merge(m, src)
}
func (m *UpdateWeaponResp) XXX_Size() int {
	return xxx_messageInfo_UpdateWeaponResp.Size(m)
}
func (m *UpdateWeaponResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWeaponResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWeaponResp proto.InternalMessageInfo

func (m *UpdateWeaponResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

type UpdateBuyWeaponReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	WeaponID             int64    `protobuf:"varint,2,opt,name=weaponID,proto3" json:"weaponID,omitempty"`
	Type                 int64    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBuyWeaponReq) Reset()         { *m = UpdateBuyWeaponReq{} }
func (m *UpdateBuyWeaponReq) String() string { return proto.CompactTextString(m) }
func (*UpdateBuyWeaponReq) ProtoMessage()    {}
func (*UpdateBuyWeaponReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e07c0075c06b2ab6, []int{6}
}

func (m *UpdateBuyWeaponReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBuyWeaponReq.Unmarshal(m, b)
}
func (m *UpdateBuyWeaponReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBuyWeaponReq.Marshal(b, m, deterministic)
}
func (m *UpdateBuyWeaponReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBuyWeaponReq.Merge(m, src)
}
func (m *UpdateBuyWeaponReq) XXX_Size() int {
	return xxx_messageInfo_UpdateBuyWeaponReq.Size(m)
}
func (m *UpdateBuyWeaponReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBuyWeaponReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBuyWeaponReq proto.InternalMessageInfo

func (m *UpdateBuyWeaponReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateBuyWeaponReq) GetWeaponID() int64 {
	if m != nil {
		return m.WeaponID
	}
	return 0
}

func (m *UpdateBuyWeaponReq) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

type UpdateBuyWeaponResp struct {
	Result               Result   `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBuyWeaponResp) Reset()         { *m = UpdateBuyWeaponResp{} }
func (m *UpdateBuyWeaponResp) String() string { return proto.CompactTextString(m) }
func (*UpdateBuyWeaponResp) ProtoMessage()    {}
func (*UpdateBuyWeaponResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e07c0075c06b2ab6, []int{7}
}

func (m *UpdateBuyWeaponResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBuyWeaponResp.Unmarshal(m, b)
}
func (m *UpdateBuyWeaponResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBuyWeaponResp.Marshal(b, m, deterministic)
}
func (m *UpdateBuyWeaponResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBuyWeaponResp.Merge(m, src)
}
func (m *UpdateBuyWeaponResp) XXX_Size() int {
	return xxx_messageInfo_UpdateBuyWeaponResp.Size(m)
}
func (m *UpdateBuyWeaponResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBuyWeaponResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBuyWeaponResp proto.InternalMessageInfo

func (m *UpdateBuyWeaponResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func init() {
	proto.RegisterType((*GetPlayerWeaponReq)(nil), "Msg.GetPlayerWeaponReq")
	proto.RegisterType((*GetPlayerWeaponResp)(nil), "Msg.GetPlayerWeaponResp")
	proto.RegisterMapType((map[int64]bool)(nil), "Msg.GetPlayerWeaponResp.IsOldEntry")
	proto.RegisterType((*CheckWeaponReq)(nil), "Msg.CheckWeaponReq")
	proto.RegisterType((*CheckWeaponResp)(nil), "Msg.CheckWeaponResp")
	proto.RegisterType((*UpdateWeaponReq)(nil), "Msg.UpdateWeaponReq")
	proto.RegisterType((*UpdateWeaponResp)(nil), "Msg.UpdateWeaponResp")
	proto.RegisterType((*UpdateBuyWeaponReq)(nil), "Msg.UpdateBuyWeaponReq")
	proto.RegisterType((*UpdateBuyWeaponResp)(nil), "Msg.UpdateBuyWeaponResp")
}

func init() { proto.RegisterFile("weapon.proto", fileDescriptor_e07c0075c06b2ab6) }

var fileDescriptor_e07c0075c06b2ab6 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x6b, 0xea, 0x40,
	0x10, 0xc7, 0x89, 0xfb, 0xf4, 0xf9, 0x46, 0x51, 0x59, 0xdf, 0x21, 0x78, 0x92, 0x78, 0x91, 0x1e,
	0x12, 0x6b, 0x0f, 0xb5, 0x96, 0x52, 0xb0, 0x2d, 0xe2, 0x41, 0x5a, 0x56, 0x8a, 0xe0, 0x6d, 0xd5,
	0x21, 0x8a, 0x31, 0x9b, 0x26, 0x1b, 0xdb, 0x7c, 0xc4, 0x7e, 0xab, 0xe2, 0xae, 0x55, 0xac, 0xa5,
	0x0a, 0xbd, 0xed, 0x4c, 0xfe, 0xf3, 0xdb, 0xdf, 0x0e, 0x81, 0xfc, 0x2b, 0xf2, 0x40, 0xf8, 0x76,
	0x10, 0x0a, 0x29, 0x28, 0xe9, 0x47, 0x6e, 0x25, 0x1f, 0x62, 0x14, 0x7b, 0x52, 0xb7, 0xac, 0x33,
	0xa0, 0x5d, 0x94, 0x4f, 0x1e, 0x4f, 0x30, 0x1c, 0xaa, 0x2c, 0xc3, 0x17, 0xfa, 0x1f, 0xd2, 0x52,
	0x2c, 0xd0, 0x37, 0x8d, 0xaa, 0x51, 0xff, 0xc7, 0x74, 0x61, 0xbd, 0x1b, 0x50, 0x3e, 0x08, 0x47,
	0x01, 0xad, 0x41, 0x46, 0x33, 0x55, 0xbc, 0xd0, 0xcc, 0xd9, 0xfd, 0xc8, 0xb5, 0x99, 0x6a, 0xb1,
	0xcd, 0x27, 0x6a, 0xc2, 0x5f, 0xed, 0x12, 0x99, 0x29, 0x05, 0xfd, 0x2c, 0xe9, 0x15, 0xa4, 0xe7,
	0xd1, 0xa3, 0x37, 0x35, 0x49, 0x95, 0xd4, 0x73, 0xcd, 0x9a, 0x9a, 0xfe, 0xe6, 0x1e, 0xbb, 0xb7,
	0x4e, 0x3d, 0xf8, 0x32, 0x4c, 0x98, 0x9e, 0xa8, 0xb4, 0x00, 0x76, 0x4d, 0x5a, 0x02, 0xb2, 0xc0,
	0x44, 0x49, 0x10, 0xb6, 0x3e, 0xae, 0xdf, 0xb1, 0xe2, 0x5e, 0x8c, 0xea, 0xca, 0x2c, 0xd3, 0x45,
	0x3b, 0xd5, 0x32, 0xac, 0x0e, 0x14, 0xee, 0x66, 0x38, 0x59, 0x1c, 0x79, 0x33, 0xad, 0x40, 0x56,
	0x7b, 0xf6, 0xee, 0x15, 0x84, 0xb0, 0x6d, 0x6d, 0x31, 0x28, 0xee, 0x31, 0x4e, 0x5d, 0xc5, 0x4f,
	0xcc, 0x21, 0x14, 0x9f, 0x83, 0x29, 0x97, 0x78, 0x4c, 0x8c, 0xc2, 0x1f, 0x99, 0x04, 0xb8, 0x01,
	0xa8, 0xf3, 0x1e, 0x98, 0x7c, 0x01, 0x5f, 0x42, 0x69, 0x1f, 0x7c, 0xa2, 0xad, 0x35, 0x02, 0xaa,
	0x07, 0x3b, 0x71, 0xf2, 0x8b, 0x6d, 0x6d, 0x85, 0xc9, 0x4e, 0xd8, 0x6a, 0x43, 0xf9, 0x80, 0x7d,
	0xa2, 0x57, 0xe7, 0x76, 0x74, 0x33, 0x8b, 0xc7, 0x8d, 0x46, 0xc3, 0x7e, 0x9b, 0xfb, 0x53, 0xe1,
	0xbb, 0xf6, 0x44, 0x2c, 0x9d, 0x01, 0x8f, 0x23, 0xee, 0xe2, 0x60, 0x26, 0x84, 0x74, 0xba, 0x7c,
	0x89, 0x03, 0x0c, 0x57, 0x18, 0x3a, 0xea, 0x6f, 0x9f, 0x08, 0xcf, 0x59, 0x9d, 0x5f, 0xbb, 0x7c,
	0x89, 0xc1, 0x78, 0x9c, 0x51, 0xbd, 0x8b, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x83, 0x5a,
	0x23, 0x24, 0x03, 0x00, 0x00,
}