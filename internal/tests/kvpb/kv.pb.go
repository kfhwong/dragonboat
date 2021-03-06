// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kv.proto

package kvpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PBKV struct {
	Key                  *string  `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Val                  *string  `protobuf:"bytes,2,opt,name=Val" json:"Val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBKV) Reset()         { *m = PBKV{} }
func (m *PBKV) String() string { return proto.CompactTextString(m) }
func (*PBKV) ProtoMessage()    {}
func (*PBKV) Descriptor() ([]byte, []int) {
	return fileDescriptor_kv_3f9814a31249a87b, []int{0}
}
func (m *PBKV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBKV.Unmarshal(m, b)
}
func (m *PBKV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBKV.Marshal(b, m, deterministic)
}
func (dst *PBKV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBKV.Merge(dst, src)
}
func (m *PBKV) XXX_Size() int {
	return xxx_messageInfo_PBKV.Size(m)
}
func (m *PBKV) XXX_DiscardUnknown() {
	xxx_messageInfo_PBKV.DiscardUnknown(m)
}

var xxx_messageInfo_PBKV proto.InternalMessageInfo

func (m *PBKV) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *PBKV) GetVal() string {
	if m != nil && m.Val != nil {
		return *m.Val
	}
	return ""
}

func init() {
	proto.RegisterType((*PBKV)(nil), "kvpb.PBKV")
}

func init() { proto.RegisterFile("kv.proto", fileDescriptor_kv_3f9814a31249a87b) }

var fileDescriptor_kv_3f9814a31249a87b = []byte{
	// 76 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x2e, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x2e, 0x2b, 0x48, 0x52, 0xd2, 0xe2, 0x62, 0x09, 0x70,
	0xf2, 0x0e, 0x13, 0x12, 0xe0, 0x62, 0xf6, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0x31, 0x41, 0x22, 0x61, 0x89, 0x39, 0x12, 0x4c, 0x10, 0x91, 0xb0, 0xc4, 0x1c, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xce, 0xd5, 0xec, 0x2b, 0x3c, 0x00, 0x00, 0x00,
}
