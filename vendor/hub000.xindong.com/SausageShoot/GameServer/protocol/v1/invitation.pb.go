// Code generated by protoc-gen-go. DO NOT EDIT.
// source: invitation.proto

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

type InvitationStatusReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvitationStatusReq) Reset()         { *m = InvitationStatusReq{} }
func (m *InvitationStatusReq) String() string { return proto.CompactTextString(m) }
func (*InvitationStatusReq) ProtoMessage()    {}
func (*InvitationStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a60183ae41ad10, []int{0}
}

func (m *InvitationStatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvitationStatusReq.Unmarshal(m, b)
}
func (m *InvitationStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvitationStatusReq.Marshal(b, m, deterministic)
}
func (m *InvitationStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvitationStatusReq.Merge(m, src)
}
func (m *InvitationStatusReq) XXX_Size() int {
	return xxx_messageInfo_InvitationStatusReq.Size(m)
}
func (m *InvitationStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InvitationStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_InvitationStatusReq proto.InternalMessageInfo

func (m *InvitationStatusReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type InvitationStatus struct {
	Receive              int64      `protobuf:"varint,1,opt,name=receive,proto3" json:"receive,omitempty"`
	Box                  []*BoxItem `protobuf:"bytes,2,rep,name=box,proto3" json:"box,omitempty"`
	FriendID             string     `protobuf:"bytes,3,opt,name=friendID,proto3" json:"friendID,omitempty"`
	Nickname             string     `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Url                  string     `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Time                 int64      `protobuf:"varint,6,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *InvitationStatus) Reset()         { *m = InvitationStatus{} }
func (m *InvitationStatus) String() string { return proto.CompactTextString(m) }
func (*InvitationStatus) ProtoMessage()    {}
func (*InvitationStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a60183ae41ad10, []int{1}
}

func (m *InvitationStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvitationStatus.Unmarshal(m, b)
}
func (m *InvitationStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvitationStatus.Marshal(b, m, deterministic)
}
func (m *InvitationStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvitationStatus.Merge(m, src)
}
func (m *InvitationStatus) XXX_Size() int {
	return xxx_messageInfo_InvitationStatus.Size(m)
}
func (m *InvitationStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_InvitationStatus.DiscardUnknown(m)
}

var xxx_messageInfo_InvitationStatus proto.InternalMessageInfo

func (m *InvitationStatus) GetReceive() int64 {
	if m != nil {
		return m.Receive
	}
	return 0
}

func (m *InvitationStatus) GetBox() []*BoxItem {
	if m != nil {
		return m.Box
	}
	return nil
}

func (m *InvitationStatus) GetFriendID() string {
	if m != nil {
		return m.FriendID
	}
	return ""
}

func (m *InvitationStatus) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *InvitationStatus) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *InvitationStatus) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type InvitationStatusResp struct {
	Result               Result              `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	Lists                []*InvitationStatus `protobuf:"bytes,2,rep,name=lists,proto3" json:"lists,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *InvitationStatusResp) Reset()         { *m = InvitationStatusResp{} }
func (m *InvitationStatusResp) String() string { return proto.CompactTextString(m) }
func (*InvitationStatusResp) ProtoMessage()    {}
func (*InvitationStatusResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a60183ae41ad10, []int{2}
}

func (m *InvitationStatusResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvitationStatusResp.Unmarshal(m, b)
}
func (m *InvitationStatusResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvitationStatusResp.Marshal(b, m, deterministic)
}
func (m *InvitationStatusResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvitationStatusResp.Merge(m, src)
}
func (m *InvitationStatusResp) XXX_Size() int {
	return xxx_messageInfo_InvitationStatusResp.Size(m)
}
func (m *InvitationStatusResp) XXX_DiscardUnknown() {
	xxx_messageInfo_InvitationStatusResp.DiscardUnknown(m)
}

var xxx_messageInfo_InvitationStatusResp proto.InternalMessageInfo

func (m *InvitationStatusResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func (m *InvitationStatusResp) GetLists() []*InvitationStatus {
	if m != nil {
		return m.Lists
	}
	return nil
}

type InvitationGotAwardReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	FriendID             string   `protobuf:"bytes,2,opt,name=friendID,proto3" json:"friendID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvitationGotAwardReq) Reset()         { *m = InvitationGotAwardReq{} }
func (m *InvitationGotAwardReq) String() string { return proto.CompactTextString(m) }
func (*InvitationGotAwardReq) ProtoMessage()    {}
func (*InvitationGotAwardReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a60183ae41ad10, []int{3}
}

func (m *InvitationGotAwardReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvitationGotAwardReq.Unmarshal(m, b)
}
func (m *InvitationGotAwardReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvitationGotAwardReq.Marshal(b, m, deterministic)
}
func (m *InvitationGotAwardReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvitationGotAwardReq.Merge(m, src)
}
func (m *InvitationGotAwardReq) XXX_Size() int {
	return xxx_messageInfo_InvitationGotAwardReq.Size(m)
}
func (m *InvitationGotAwardReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InvitationGotAwardReq.DiscardUnknown(m)
}

var xxx_messageInfo_InvitationGotAwardReq proto.InternalMessageInfo

func (m *InvitationGotAwardReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *InvitationGotAwardReq) GetFriendID() string {
	if m != nil {
		return m.FriendID
	}
	return ""
}

type InvitationGotAwardResp struct {
	Result               Result     `protobuf:"varint,1,opt,name=result,proto3,enum=Msg.Result" json:"result,omitempty"`
	Box                  []*BoxItem `protobuf:"bytes,2,rep,name=box,proto3" json:"box,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *InvitationGotAwardResp) Reset()         { *m = InvitationGotAwardResp{} }
func (m *InvitationGotAwardResp) String() string { return proto.CompactTextString(m) }
func (*InvitationGotAwardResp) ProtoMessage()    {}
func (*InvitationGotAwardResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a60183ae41ad10, []int{4}
}

func (m *InvitationGotAwardResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvitationGotAwardResp.Unmarshal(m, b)
}
func (m *InvitationGotAwardResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvitationGotAwardResp.Marshal(b, m, deterministic)
}
func (m *InvitationGotAwardResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvitationGotAwardResp.Merge(m, src)
}
func (m *InvitationGotAwardResp) XXX_Size() int {
	return xxx_messageInfo_InvitationGotAwardResp.Size(m)
}
func (m *InvitationGotAwardResp) XXX_DiscardUnknown() {
	xxx_messageInfo_InvitationGotAwardResp.DiscardUnknown(m)
}

var xxx_messageInfo_InvitationGotAwardResp proto.InternalMessageInfo

func (m *InvitationGotAwardResp) GetResult() Result {
	if m != nil {
		return m.Result
	}
	return Result_INTERNAL
}

func (m *InvitationGotAwardResp) GetBox() []*BoxItem {
	if m != nil {
		return m.Box
	}
	return nil
}

func init() {
	proto.RegisterType((*InvitationStatusReq)(nil), "Msg.InvitationStatusReq")
	proto.RegisterType((*InvitationStatus)(nil), "Msg.InvitationStatus")
	proto.RegisterType((*InvitationStatusResp)(nil), "Msg.InvitationStatusResp")
	proto.RegisterType((*InvitationGotAwardReq)(nil), "Msg.InvitationGotAwardReq")
	proto.RegisterType((*InvitationGotAwardResp)(nil), "Msg.InvitationGotAwardResp")
}

func init() { proto.RegisterFile("invitation.proto", fileDescriptor_92a60183ae41ad10) }

var fileDescriptor_92a60183ae41ad10 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x0d, 0x14, 0x50, 0x07, 0x62, 0xc8, 0x0a, 0xa6, 0xe1, 0x60, 0x48, 0xbd, 0x90, 0x90, 0xb4,
	0x88, 0x47, 0x63, 0x8c, 0xc4, 0x84, 0xf4, 0xc0, 0xa5, 0xbd, 0x99, 0x78, 0xd8, 0x96, 0xb1, 0x6c,
	0x68, 0x77, 0x71, 0x77, 0x5b, 0xf9, 0x4d, 0xfe, 0x4a, 0xc3, 0x16, 0x30, 0x88, 0x46, 0x6f, 0xf3,
	0x3e, 0x3a, 0x7d, 0xf3, 0xb2, 0xd0, 0x66, 0xbc, 0x60, 0x9a, 0x6a, 0x26, 0xb8, 0xbb, 0x92, 0x42,
	0x0b, 0x62, 0xcd, 0x54, 0xd2, 0x6b, 0x49, 0x54, 0x79, 0xaa, 0x4b, 0xaa, 0xd7, 0x8c, 0xa8, 0x62,
	0x71, 0x09, 0x9c, 0x21, 0x5c, 0xf8, 0xfb, 0x6f, 0x42, 0x4d, 0x75, 0xae, 0x02, 0x7c, 0x23, 0x1d,
	0xa8, 0x6b, 0xb1, 0x44, 0x6e, 0x57, 0xfa, 0x95, 0xc1, 0x59, 0x50, 0x02, 0xe7, 0xa3, 0x02, 0xed,
	0xef, 0x6e, 0x62, 0xc3, 0x89, 0xc4, 0x18, 0x59, 0x81, 0xc6, 0x6c, 0x05, 0x3b, 0x48, 0xae, 0xc0,
	0x8a, 0xc4, 0xda, 0xae, 0xf6, 0xad, 0x41, 0x73, 0xdc, 0x72, 0x67, 0x2a, 0x71, 0x27, 0x62, 0xed,
	0x6b, 0xcc, 0x82, 0x8d, 0x40, 0x7a, 0x70, 0xfa, 0x2a, 0x19, 0xf2, 0xb9, 0xff, 0x64, 0x5b, 0xe6,
	0x3f, 0x7b, 0xbc, 0xd1, 0x38, 0x8b, 0x97, 0x9c, 0x66, 0x68, 0xd7, 0x4a, 0x6d, 0x87, 0x49, 0x1b,
	0xac, 0x5c, 0xa6, 0x76, 0xdd, 0xd0, 0x9b, 0x91, 0x10, 0xa8, 0x69, 0x96, 0xa1, 0xdd, 0x30, 0x01,
	0xcc, 0xec, 0x2c, 0xa0, 0x73, 0x7c, 0x99, 0x5a, 0x91, 0x6b, 0x68, 0x94, 0x75, 0x98, 0xb8, 0xe7,
	0xe3, 0xa6, 0x09, 0x16, 0x18, 0x2a, 0xd8, 0x4a, 0x64, 0x08, 0xf5, 0x94, 0x29, 0xad, 0xb6, 0xe1,
	0xbb, 0xc6, 0x73, 0xb4, 0xae, 0xf4, 0x38, 0x3e, 0x74, 0xbf, 0xa4, 0xa9, 0xd0, 0x8f, 0xef, 0x54,
	0xce, 0x7f, 0x6d, 0xf1, 0xe0, 0xec, 0xea, 0xe1, 0xd9, 0xce, 0x0b, 0x5c, 0xfe, 0xb4, 0xea, 0xbf,
	0xb1, 0xff, 0x68, 0x7c, 0xf2, 0xf0, 0x7c, 0xbf, 0xc8, 0xa3, 0xd1, 0x68, 0xe4, 0xae, 0x19, 0x9f,
	0x0b, 0x9e, 0xb8, 0xb1, 0xc8, 0xbc, 0x90, 0xe6, 0x8a, 0x26, 0x18, 0x2e, 0x84, 0xd0, 0xde, 0x94,
	0x66, 0x18, 0xa2, 0x2c, 0x50, 0x7a, 0xe6, 0x85, 0xc4, 0x22, 0xf5, 0x8a, 0x9b, 0xbb, 0x84, 0x66,
	0xb8, 0x8a, 0xa2, 0x86, 0xe1, 0x6e, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x27, 0x41, 0x17, 0x42,
	0x69, 0x02, 0x00, 0x00,
}
