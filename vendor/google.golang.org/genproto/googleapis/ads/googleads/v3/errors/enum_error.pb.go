// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/enum_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible enum errors.
type EnumErrorEnum_EnumError int32

const (
	// Enum unspecified.
	EnumErrorEnum_UNSPECIFIED EnumErrorEnum_EnumError = 0
	// The received error code is not known in this version.
	EnumErrorEnum_UNKNOWN EnumErrorEnum_EnumError = 1
	// The enum value is not permitted.
	EnumErrorEnum_ENUM_VALUE_NOT_PERMITTED EnumErrorEnum_EnumError = 3
)

var EnumErrorEnum_EnumError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "ENUM_VALUE_NOT_PERMITTED",
}

var EnumErrorEnum_EnumError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"ENUM_VALUE_NOT_PERMITTED": 3,
}

func (x EnumErrorEnum_EnumError) String() string {
	return proto.EnumName(EnumErrorEnum_EnumError_name, int32(x))
}

func (EnumErrorEnum_EnumError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b002669bf1a51dfb, []int{0, 0}
}

// Container for enum describing possible enum errors.
type EnumErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumErrorEnum) Reset()         { *m = EnumErrorEnum{} }
func (m *EnumErrorEnum) String() string { return proto.CompactTextString(m) }
func (*EnumErrorEnum) ProtoMessage()    {}
func (*EnumErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b002669bf1a51dfb, []int{0}
}

func (m *EnumErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumErrorEnum.Unmarshal(m, b)
}
func (m *EnumErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumErrorEnum.Marshal(b, m, deterministic)
}
func (m *EnumErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumErrorEnum.Merge(m, src)
}
func (m *EnumErrorEnum) XXX_Size() int {
	return xxx_messageInfo_EnumErrorEnum.Size(m)
}
func (m *EnumErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_EnumErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.EnumErrorEnum_EnumError", EnumErrorEnum_EnumError_name, EnumErrorEnum_EnumError_value)
	proto.RegisterType((*EnumErrorEnum)(nil), "google.ads.googleads.v3.errors.EnumErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/enum_error.proto", fileDescriptor_b002669bf1a51dfb)
}

var fileDescriptor_b002669bf1a51dfb = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4e, 0xb4, 0x30,
	0x18, 0x86, 0xff, 0x61, 0x92, 0xdf, 0xd8, 0x89, 0x4a, 0x58, 0x19, 0x33, 0x99, 0x05, 0x07, 0x68,
	0x17, 0xdd, 0xd5, 0x55, 0x47, 0x2a, 0x21, 0x3a, 0x1d, 0xa2, 0x80, 0xc6, 0x90, 0x10, 0x14, 0xd2,
	0x90, 0x0c, 0x2d, 0xa1, 0xcc, 0x1c, 0xc8, 0xa5, 0x47, 0xf1, 0x26, 0x7a, 0x0a, 0x03, 0x1d, 0xd8,
	0xe9, 0xaa, 0x4f, 0x9a, 0xe7, 0xfd, 0xbe, 0x37, 0x1f, 0x40, 0x42, 0x29, 0xb1, 0x2b, 0x51, 0x5e,
	0xe8, 0x23, 0xf6, 0x74, 0xc0, 0xa8, 0x6c, 0x5b, 0xd5, 0x6a, 0x54, 0xca, 0x7d, 0x9d, 0x0d, 0x0c,
	0x9b, 0x56, 0x75, 0xca, 0x59, 0x19, 0x0b, 0xe6, 0x85, 0x86, 0x53, 0x00, 0x1e, 0x30, 0x34, 0x81,
	0xab, 0xe5, 0x38, 0xb0, 0xa9, 0x50, 0x2e, 0xa5, 0xea, 0xf2, 0xae, 0x52, 0x52, 0x9b, 0xb4, 0xfb,
	0x0c, 0xce, 0x98, 0xdc, 0xd7, 0xac, 0x77, 0x7b, 0x70, 0x7d, 0x70, 0x3a, 0x7d, 0x38, 0x17, 0x60,
	0x11, 0xf3, 0xc7, 0x90, 0xdd, 0x04, 0xb7, 0x01, 0xf3, 0xec, 0x7f, 0xce, 0x02, 0x9c, 0xc4, 0xfc,
	0x8e, 0x6f, 0x9f, 0xb8, 0x3d, 0x73, 0x96, 0xe0, 0x92, 0xf1, 0x78, 0x93, 0x25, 0xf4, 0x3e, 0x66,
	0x19, 0xdf, 0x46, 0x59, 0xc8, 0x1e, 0x36, 0x41, 0x14, 0x31, 0xcf, 0x9e, 0xaf, 0xbf, 0x66, 0xc0,
	0x7d, 0x53, 0x35, 0xfc, 0xbb, 0xde, 0xfa, 0x7c, 0xda, 0x16, 0xf6, 0x85, 0xc2, 0xd9, 0x8b, 0x77,
	0x4c, 0x08, 0xb5, 0xcb, 0xa5, 0x80, 0xaa, 0x15, 0x48, 0x94, 0x72, 0xa8, 0x3b, 0x5e, 0xa4, 0xa9,
	0xf4, 0x6f, 0x07, 0xba, 0x36, 0xcf, 0xbb, 0x35, 0xf7, 0x29, 0xfd, 0xb0, 0x56, 0xbe, 0x19, 0x46,
	0x0b, 0x0d, 0x0d, 0xf6, 0x94, 0x60, 0x38, 0xac, 0xd4, 0x9f, 0xa3, 0x90, 0xd2, 0x42, 0xa7, 0x93,
	0x90, 0x26, 0x38, 0x35, 0xc2, 0xb7, 0xe5, 0x9a, 0x5f, 0x42, 0x68, 0xa1, 0x09, 0x99, 0x14, 0x42,
	0x12, 0x4c, 0x88, 0x91, 0x5e, 0xff, 0x0f, 0xed, 0xf0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfa,
	0x67, 0x63, 0xa3, 0xbd, 0x01, 0x00, 0x00,
}