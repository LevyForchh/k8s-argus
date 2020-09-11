// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/dsa_page_feed_criterion_field.proto

package enums

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

// Possible values for Dynamic Search Ad Page Feed criterion fields.
type DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField int32

const (
	// Not specified.
	DsaPageFeedCriterionFieldEnum_UNSPECIFIED DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField = 0
	// Used for return value only. Represents value unknown in this version.
	DsaPageFeedCriterionFieldEnum_UNKNOWN DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField = 1
	// Data Type: URL or URL_LIST. URL of the web page you want to target.
	DsaPageFeedCriterionFieldEnum_PAGE_URL DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField = 2
	// Data Type: STRING_LIST. The labels that will help you target ads within
	// your page feed.
	DsaPageFeedCriterionFieldEnum_LABEL DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField = 3
)

var DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PAGE_URL",
	3: "LABEL",
}

var DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"PAGE_URL":    2,
	"LABEL":       3,
}

func (x DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField) String() string {
	return proto.EnumName(DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField_name, int32(x))
}

func (DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_084fbedb240e0f43, []int{0, 0}
}

// Values for Dynamic Search Ad Page Feed criterion fields.
type DsaPageFeedCriterionFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DsaPageFeedCriterionFieldEnum) Reset()         { *m = DsaPageFeedCriterionFieldEnum{} }
func (m *DsaPageFeedCriterionFieldEnum) String() string { return proto.CompactTextString(m) }
func (*DsaPageFeedCriterionFieldEnum) ProtoMessage()    {}
func (*DsaPageFeedCriterionFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_084fbedb240e0f43, []int{0}
}

func (m *DsaPageFeedCriterionFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DsaPageFeedCriterionFieldEnum.Unmarshal(m, b)
}
func (m *DsaPageFeedCriterionFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DsaPageFeedCriterionFieldEnum.Marshal(b, m, deterministic)
}
func (m *DsaPageFeedCriterionFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DsaPageFeedCriterionFieldEnum.Merge(m, src)
}
func (m *DsaPageFeedCriterionFieldEnum) XXX_Size() int {
	return xxx_messageInfo_DsaPageFeedCriterionFieldEnum.Size(m)
}
func (m *DsaPageFeedCriterionFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DsaPageFeedCriterionFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DsaPageFeedCriterionFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField", DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField_name, DsaPageFeedCriterionFieldEnum_DsaPageFeedCriterionField_value)
	proto.RegisterType((*DsaPageFeedCriterionFieldEnum)(nil), "google.ads.googleads.v3.enums.DsaPageFeedCriterionFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/dsa_page_feed_criterion_field.proto", fileDescriptor_084fbedb240e0f43)
}

var fileDescriptor_084fbedb240e0f43 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4d, 0x6a, 0xeb, 0x30,
	0x18, 0x7c, 0x71, 0x78, 0xfd, 0x51, 0x0a, 0x35, 0xde, 0xb5, 0x34, 0x85, 0xe4, 0x00, 0xf2, 0xc2,
	0x3b, 0x75, 0x25, 0x27, 0x4e, 0x08, 0x0d, 0xae, 0x49, 0x49, 0x0a, 0xc5, 0x60, 0xd4, 0x48, 0x11,
	0x82, 0x44, 0x32, 0xfe, 0x9c, 0x1c, 0xa8, 0xcb, 0x1e, 0xa5, 0x47, 0xe9, 0xb6, 0x17, 0x28, 0x96,
	0xea, 0xec, 0xd2, 0x8d, 0x19, 0x3c, 0xf3, 0xcd, 0x8c, 0x06, 0x51, 0x69, 0x8c, 0xdc, 0x8a, 0x90,
	0x71, 0x08, 0x1d, 0x6c, 0xd0, 0x21, 0x0a, 0x85, 0xde, 0xef, 0x20, 0xe4, 0xc0, 0x8a, 0x92, 0x49,
	0x51, 0x6c, 0x84, 0xe0, 0xc5, 0xba, 0x52, 0xb5, 0xa8, 0x94, 0xd1, 0xc5, 0x46, 0x89, 0x2d, 0xc7,
	0x65, 0x65, 0x6a, 0x13, 0xf4, 0xdd, 0x1d, 0x66, 0x1c, 0xf0, 0xd1, 0x02, 0x1f, 0x22, 0x6c, 0x2d,
	0x6e, 0xef, 0xda, 0x84, 0x52, 0x85, 0x4c, 0x6b, 0x53, 0xb3, 0x5a, 0x19, 0x0d, 0xee, 0x78, 0x08,
	0xa8, 0x3f, 0x06, 0x96, 0x31, 0x29, 0x26, 0x42, 0xf0, 0x51, 0x1b, 0x30, 0x69, 0xfc, 0x13, 0xbd,
	0xdf, 0x0d, 0x17, 0xe8, 0xe6, 0xa4, 0x20, 0xb8, 0x46, 0xbd, 0x65, 0xfa, 0x9c, 0x25, 0xa3, 0xd9,
	0x64, 0x96, 0x8c, 0xfd, 0x7f, 0x41, 0x0f, 0x9d, 0x2f, 0xd3, 0xc7, 0xf4, 0xe9, 0x25, 0xf5, 0x3b,
	0xc1, 0x15, 0xba, 0xc8, 0xe8, 0x34, 0x29, 0x96, 0x8b, 0xb9, 0xef, 0x05, 0x97, 0xe8, 0xff, 0x9c,
	0xc6, 0xc9, 0xdc, 0xef, 0xc6, 0xdf, 0x1d, 0x34, 0x58, 0x9b, 0x1d, 0xfe, 0xb3, 0x78, 0x7c, 0x7f,
	0x32, 0x37, 0x6b, 0xaa, 0x67, 0x9d, 0xd7, 0xf8, 0xd7, 0x40, 0x9a, 0x2d, 0xd3, 0x12, 0x9b, 0x4a,
	0x86, 0x52, 0x68, 0xfb, 0xb0, 0x76, 0xcc, 0x52, 0xc1, 0x89, 0x6d, 0x1f, 0xec, 0xf7, 0xdd, 0xeb,
	0x4e, 0x29, 0xfd, 0xf0, 0xfa, 0x53, 0x67, 0x45, 0x39, 0x60, 0x07, 0x1b, 0xb4, 0x8a, 0x70, 0x33,
	0x02, 0x7c, 0xb6, 0x7c, 0x4e, 0x39, 0xe4, 0x47, 0x3e, 0x5f, 0x45, 0xb9, 0xe5, 0xbf, 0xbc, 0x81,
	0xfb, 0x49, 0x08, 0xe5, 0x40, 0xc8, 0x51, 0x41, 0xc8, 0x2a, 0x22, 0xc4, 0x6a, 0xde, 0xce, 0x6c,
	0xb1, 0xe8, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xc0, 0x0d, 0xd4, 0xf3, 0x01, 0x00, 0x00,
}