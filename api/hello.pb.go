// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/hello.proto

package hello

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// Specifies the int64 start and end of the range using half-open interval semantics [start,
// end).
type Strings struct {
	// start of the range (inclusive)
	Start string `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  string   `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Strings) Reset()         { *m = Strings{} }
func (m *Strings) String() string { return proto.CompactTextString(m) }
func (*Strings) ProtoMessage()    {}
func (*Strings) Descriptor() ([]byte, []int) {
	return fileDescriptor_860c41a0e7a04411, []int{0}
}

func (m *Strings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Strings.Unmarshal(m, b)
}
func (m *Strings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Strings.Marshal(b, m, deterministic)
}
func (m *Strings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Strings.Merge(m, src)
}
func (m *Strings) XXX_Size() int {
	return xxx_messageInfo_Strings.Size(m)
}
func (m *Strings) XXX_DiscardUnknown() {
	xxx_messageInfo_Strings.DiscardUnknown(m)
}

var xxx_messageInfo_Strings proto.InternalMessageInfo

func (m *Strings) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Strings) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

// Specifies the double start and end of the range using half-open interval semantics [start,
// end).
type DoubleRange struct {
	// start of the range (inclusive)
	Start float64 `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  float64  `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleRange) Reset()         { *m = DoubleRange{} }
func (m *DoubleRange) String() string { return proto.CompactTextString(m) }
func (*DoubleRange) ProtoMessage()    {}
func (*DoubleRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_860c41a0e7a04411, []int{1}
}

func (m *DoubleRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleRange.Unmarshal(m, b)
}
func (m *DoubleRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleRange.Marshal(b, m, deterministic)
}
func (m *DoubleRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleRange.Merge(m, src)
}
func (m *DoubleRange) XXX_Size() int {
	return xxx_messageInfo_DoubleRange.Size(m)
}
func (m *DoubleRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleRange.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleRange proto.InternalMessageInfo

func (m *DoubleRange) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *DoubleRange) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

type Nested struct {
	DoubleRange          *DoubleRange    `protobuf:"bytes,1,opt,name=double_range,json=doubleRange,proto3" json:"double_range,omitempty"`
	Strings              *Strings        `protobuf:"bytes,2,opt,name=strings,proto3" json:"strings,omitempty"`
	Details              *_struct.Struct `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
	Hello                []string        `protobuf:"bytes,4,rep,name=hello,proto3" json:"hello,omitempty"`
	X                    []*Strings      `protobuf:"bytes,7,rep,name=x,proto3" json:"x,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Nested) Reset()         { *m = Nested{} }
func (m *Nested) String() string { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()    {}
func (*Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_860c41a0e7a04411, []int{2}
}

func (m *Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nested.Unmarshal(m, b)
}
func (m *Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nested.Marshal(b, m, deterministic)
}
func (m *Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested.Merge(m, src)
}
func (m *Nested) XXX_Size() int {
	return xxx_messageInfo_Nested.Size(m)
}
func (m *Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_Nested proto.InternalMessageInfo

func (m *Nested) GetDoubleRange() *DoubleRange {
	if m != nil {
		return m.DoubleRange
	}
	return nil
}

func (m *Nested) GetStrings() *Strings {
	if m != nil {
		return m.Strings
	}
	return nil
}

func (m *Nested) GetDetails() *_struct.Struct {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Nested) GetHello() []string {
	if m != nil {
		return m.Hello
	}
	return nil
}

func (m *Nested) GetX() []*Strings {
	if m != nil {
		return m.X
	}
	return nil
}

func init() {
	proto.RegisterType((*Strings)(nil), "envoy.type.Strings")
	proto.RegisterType((*DoubleRange)(nil), "envoy.type.DoubleRange")
	proto.RegisterType((*Nested)(nil), "envoy.type.Nested")
}

func init() { proto.RegisterFile("api/hello.proto", fileDescriptor_860c41a0e7a04411) }

var fileDescriptor_860c41a0e7a04411 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x89, 0xeb, 0x6e, 0x71, 0x2a, 0x28, 0x51, 0x30, 0x88, 0x87, 0xba, 0x27, 0x2f, 0xa6,
	0xec, 0x8a, 0x17, 0x8f, 0xe2, 0xd9, 0x43, 0xbc, 0x79, 0x91, 0x76, 0x33, 0xd6, 0x42, 0x68, 0x4a,
	0x32, 0x95, 0xdd, 0x5f, 0xeb, 0x5f, 0x91, 0x24, 0x5b, 0xac, 0xa0, 0xb7, 0x4e, 0xf3, 0xbe, 0xf7,
	0x66, 0x1e, 0x9c, 0x54, 0x7d, 0x5b, 0x7e, 0xa0, 0x31, 0x56, 0xf6, 0xce, 0x92, 0xe5, 0x80, 0xdd,
	0xa7, 0xdd, 0x49, 0xda, 0xf5, 0x78, 0x79, 0xd5, 0x58, 0xdb, 0x18, 0x2c, 0xe3, 0x4b, 0x3d, 0xbc,
	0x97, 0x9e, 0xdc, 0xb0, 0xa1, 0xa4, 0x5c, 0xae, 0x20, 0x7b, 0x21, 0xd7, 0x76, 0x8d, 0xe7, 0xe7,
	0x30, 0xf7, 0x54, 0x39, 0x12, 0xac, 0x60, 0x37, 0x47, 0x2a, 0x0d, 0xfc, 0x14, 0x66, 0xd8, 0x69,
	0x71, 0x10, 0xff, 0x85, 0xcf, 0xe5, 0x3d, 0xe4, 0x4f, 0x76, 0xa8, 0x0d, 0xaa, 0xaa, 0x6b, 0xf0,
	0x37, 0xc6, 0xfe, 0xc0, 0x58, 0xc2, 0xbe, 0x18, 0x2c, 0x9e, 0xd1, 0x13, 0x6a, 0xfe, 0x00, 0xc7,
	0x3a, 0x3a, 0xbc, 0xb9, 0x60, 0x11, 0xc9, 0x7c, 0x7d, 0x21, 0x7f, 0xb6, 0x96, 0x93, 0x04, 0x95,
	0xeb, 0x49, 0xdc, 0x2d, 0x64, 0x3e, 0x2d, 0x1c, 0xcd, 0xf3, 0xf5, 0xd9, 0x14, 0xdb, 0xdf, 0xa2,
	0x46, 0x0d, 0x5f, 0x41, 0xa6, 0x91, 0xaa, 0xd6, 0x78, 0x31, 0xdf, 0xa7, 0xa4, 0x3e, 0xe4, 0xd8,
	0x47, 0x60, 0x86, 0x0d, 0xa9, 0x51, 0x17, 0x0e, 0x8a, 0x5d, 0x8a, 0xc3, 0x62, 0x16, 0x7a, 0x88,
	0x03, 0xbf, 0x06, 0xb6, 0x15, 0x59, 0x31, 0xfb, 0x2f, 0x91, 0x6d, 0x1f, 0xb3, 0xd7, 0xa4, 0xad,
	0x17, 0xd1, 0xfb, 0xee, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xc6, 0x0d, 0x90, 0x98, 0x01, 0x00,
	0x00,
}