// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foo_test.proto

package codec

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

// 测试
type Foo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_foo_test_b51d4077268954bd, []int{0}
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (dst *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(dst, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Foo)(nil), "codec.Foo")
}

func init() { proto.RegisterFile("foo_test.proto", fileDescriptor_foo_test_b51d4077268954bd) }

var fileDescriptor_foo_test_b51d4077268954bd = []byte{
	// 76 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcb, 0xcf, 0x8f,
	0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0xce, 0x4f, 0x49,
	0x4d, 0x56, 0x92, 0xe4, 0x62, 0x76, 0xcb, 0xcf, 0x17, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x93, 0xd8, 0xc0, 0x0a, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x0c, 0x00, 0x59, 0xbb, 0x3a, 0x00, 0x00, 0x00,
}
