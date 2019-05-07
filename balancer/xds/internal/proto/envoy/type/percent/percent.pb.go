// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/percent.proto

package envoy_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/panjjo/grpc-go/balancer/xds/internal/proto/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FractionalPercent_DenominatorType int32

const (
	FractionalPercent_HUNDRED      FractionalPercent_DenominatorType = 0
	FractionalPercent_TEN_THOUSAND FractionalPercent_DenominatorType = 1
	FractionalPercent_MILLION      FractionalPercent_DenominatorType = 2
)

var FractionalPercent_DenominatorType_name = map[int32]string{
	0: "HUNDRED",
	1: "TEN_THOUSAND",
	2: "MILLION",
}
var FractionalPercent_DenominatorType_value = map[string]int32{
	"HUNDRED":      0,
	"TEN_THOUSAND": 1,
	"MILLION":      2,
}

func (x FractionalPercent_DenominatorType) String() string {
	return proto.EnumName(FractionalPercent_DenominatorType_name, int32(x))
}
func (FractionalPercent_DenominatorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_percent_cd85ccaca181f641, []int{1, 0}
}

type Percent struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Percent) Reset()         { *m = Percent{} }
func (m *Percent) String() string { return proto.CompactTextString(m) }
func (*Percent) ProtoMessage()    {}
func (*Percent) Descriptor() ([]byte, []int) {
	return fileDescriptor_percent_cd85ccaca181f641, []int{0}
}
func (m *Percent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Percent.Unmarshal(m, b)
}
func (m *Percent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Percent.Marshal(b, m, deterministic)
}
func (dst *Percent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Percent.Merge(dst, src)
}
func (m *Percent) XXX_Size() int {
	return xxx_messageInfo_Percent.Size(m)
}
func (m *Percent) XXX_DiscardUnknown() {
	xxx_messageInfo_Percent.DiscardUnknown(m)
}

var xxx_messageInfo_Percent proto.InternalMessageInfo

func (m *Percent) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type FractionalPercent struct {
	Numerator            uint32                            `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	Denominator          FractionalPercent_DenominatorType `protobuf:"varint,2,opt,name=denominator,proto3,enum=envoy.type.FractionalPercent_DenominatorType" json:"denominator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *FractionalPercent) Reset()         { *m = FractionalPercent{} }
func (m *FractionalPercent) String() string { return proto.CompactTextString(m) }
func (*FractionalPercent) ProtoMessage()    {}
func (*FractionalPercent) Descriptor() ([]byte, []int) {
	return fileDescriptor_percent_cd85ccaca181f641, []int{1}
}
func (m *FractionalPercent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FractionalPercent.Unmarshal(m, b)
}
func (m *FractionalPercent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FractionalPercent.Marshal(b, m, deterministic)
}
func (dst *FractionalPercent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FractionalPercent.Merge(dst, src)
}
func (m *FractionalPercent) XXX_Size() int {
	return xxx_messageInfo_FractionalPercent.Size(m)
}
func (m *FractionalPercent) XXX_DiscardUnknown() {
	xxx_messageInfo_FractionalPercent.DiscardUnknown(m)
}

var xxx_messageInfo_FractionalPercent proto.InternalMessageInfo

func (m *FractionalPercent) GetNumerator() uint32 {
	if m != nil {
		return m.Numerator
	}
	return 0
}

func (m *FractionalPercent) GetDenominator() FractionalPercent_DenominatorType {
	if m != nil {
		return m.Denominator
	}
	return FractionalPercent_HUNDRED
}

func init() {
	proto.RegisterType((*Percent)(nil), "envoy.type.Percent")
	proto.RegisterType((*FractionalPercent)(nil), "envoy.type.FractionalPercent")
	proto.RegisterEnum("envoy.type.FractionalPercent_DenominatorType", FractionalPercent_DenominatorType_name, FractionalPercent_DenominatorType_value)
}

func init() { proto.RegisterFile("envoy/type/percent.proto", fileDescriptor_percent_cd85ccaca181f641) }

var fileDescriptor_percent_cd85ccaca181f641 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x48, 0x2d, 0x4a, 0x4e, 0xcd, 0x2b, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0xcb, 0xe8, 0x81, 0x64, 0xa4, 0xc4, 0xcb, 0x12, 0x73,
	0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0x22, 0x25, 0x2b, 0x2e, 0xf6, 0x00, 0x88,
	0x2e, 0x21, 0x7d, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x46,
	0x27, 0xc9, 0x5d, 0x2f, 0x0f, 0x30, 0x8b, 0x08, 0x09, 0x49, 0x32, 0x80, 0x41, 0xa4, 0x83, 0x26,
	0x03, 0x14, 0x04, 0x41, 0xd4, 0x29, 0x9d, 0x65, 0xe4, 0x12, 0x74, 0x2b, 0x4a, 0x4c, 0x2e, 0xc9,
	0xcc, 0xcf, 0x4b, 0xcc, 0x81, 0x19, 0x23, 0xc3, 0xc5, 0x99, 0x57, 0x9a, 0x9b, 0x5a, 0x94, 0x58,
	0x92, 0x5f, 0x04, 0x36, 0x8a, 0x37, 0x08, 0x21, 0x20, 0x14, 0xcd, 0xc5, 0x9d, 0x92, 0x9a, 0x97,
	0x9f, 0x9b, 0x99, 0x07, 0x96, 0x67, 0x52, 0x60, 0xd4, 0xe0, 0x33, 0xd2, 0xd5, 0x43, 0x38, 0x55,
	0x0f, 0xc3, 0x44, 0x3d, 0x17, 0x84, 0x86, 0x90, 0xca, 0x82, 0x54, 0x27, 0x2e, 0x90, 0xcb, 0x58,
	0x9b, 0x18, 0x99, 0x04, 0x18, 0x83, 0x90, 0x4d, 0x53, 0xb2, 0xe5, 0xe2, 0x47, 0x53, 0x2b, 0xc4,
	0xcd, 0xc5, 0xee, 0x11, 0xea, 0xe7, 0x12, 0xe4, 0xea, 0x22, 0xc0, 0x20, 0x24, 0xc0, 0xc5, 0x13,
	0xe2, 0xea, 0x17, 0x1f, 0xe2, 0xe1, 0x1f, 0x1a, 0xec, 0xe8, 0xe7, 0x22, 0xc0, 0x08, 0x92, 0xf6,
	0xf5, 0xf4, 0xf1, 0xf1, 0xf4, 0xf7, 0x13, 0x60, 0x72, 0xd2, 0xe2, 0x92, 0xc8, 0xcc, 0x87, 0x38,
	0xa5, 0xa0, 0x28, 0xbf, 0xa2, 0x12, 0xc9, 0x55, 0x4e, 0x3c, 0x50, 0xc7, 0x04, 0x80, 0x42, 0x2d,
	0x80, 0x31, 0x89, 0x0d, 0x1c, 0x7c, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x79, 0xd6,
	0xbb, 0x7f, 0x01, 0x00, 0x00,
}