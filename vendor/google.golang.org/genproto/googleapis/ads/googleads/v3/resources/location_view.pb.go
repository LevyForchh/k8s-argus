// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/location_view.proto

package resources

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

// A location view summarizes the performance of campaigns by
// Location criteria.
type LocationView struct {
	// The resource name of the location view.
	// Location view resource names have the form:
	//
	// `customers/{customer_id}/locationViews/{campaign_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationView) Reset()         { *m = LocationView{} }
func (m *LocationView) String() string { return proto.CompactTextString(m) }
func (*LocationView) ProtoMessage()    {}
func (*LocationView) Descriptor() ([]byte, []int) {
	return fileDescriptor_8294fbe2fa5282fe, []int{0}
}

func (m *LocationView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationView.Unmarshal(m, b)
}
func (m *LocationView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationView.Marshal(b, m, deterministic)
}
func (m *LocationView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationView.Merge(m, src)
}
func (m *LocationView) XXX_Size() int {
	return xxx_messageInfo_LocationView.Size(m)
}
func (m *LocationView) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationView.DiscardUnknown(m)
}

var xxx_messageInfo_LocationView proto.InternalMessageInfo

func (m *LocationView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*LocationView)(nil), "google.ads.googleads.v3.resources.LocationView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/location_view.proto", fileDescriptor_8294fbe2fa5282fe)
}

var fileDescriptor_8294fbe2fa5282fe = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x49, 0x05, 0xc1, 0x50, 0x17, 0x76, 0xa5, 0xc5, 0x85, 0x55, 0x0a, 0xae, 0x66, 0xc0,
	0xc1, 0xcd, 0xb8, 0x9a, 0x6e, 0x0a, 0x22, 0x52, 0xba, 0xc8, 0x42, 0x83, 0x65, 0x4c, 0x86, 0x21,
	0xd0, 0xcc, 0x2b, 0x79, 0x69, 0xba, 0x28, 0x3d, 0x85, 0x37, 0x70, 0xe9, 0x51, 0x3c, 0x4a, 0x0f,
	0x21, 0x92, 0x4c, 0x66, 0xda, 0x6e, 0x74, 0xf7, 0x33, 0xef, 0xfb, 0xff, 0xf7, 0x7e, 0x26, 0xbc,
	0xd7, 0x00, 0x7a, 0xae, 0xa8, 0x4c, 0x91, 0x5a, 0x59, 0xab, 0x8a, 0xd1, 0x42, 0x21, 0x2c, 0x8b,
	0x44, 0x21, 0x9d, 0x43, 0x22, 0xcb, 0x0c, 0xcc, 0xac, 0xca, 0xd4, 0x8a, 0x2c, 0x0a, 0x28, 0xa1,
	0x37, 0xb0, 0x2c, 0x91, 0x29, 0x12, 0x6f, 0x23, 0x15, 0x23, 0xde, 0xd6, 0xbf, 0x70, 0xc9, 0x8b,
	0xcc, 0x87, 0x59, 0x77, 0xff, 0x72, 0x6f, 0x24, 0x8d, 0x81, 0xb2, 0xc9, 0x47, 0x3b, 0xbd, 0xfe,
	0x08, 0xc2, 0xee, 0x53, 0xbb, 0x33, 0xca, 0xd4, 0xaa, 0x77, 0x13, 0x9e, 0xba, 0x80, 0x99, 0x91,
	0xb9, 0x3a, 0x0f, 0xae, 0x82, 0xdb, 0x93, 0x69, 0xd7, 0x3d, 0x3e, 0xcb, 0x5c, 0xf1, 0xb7, 0xad,
	0x78, 0x0d, 0x87, 0xbb, 0x5b, 0x5a, 0xb5, 0xc8, 0x90, 0x24, 0x90, 0xd3, 0x83, 0xc0, 0xbb, 0x64,
	0x89, 0x25, 0xe4, 0xaa, 0x40, 0xba, 0x76, 0x72, 0xe3, 0x7b, 0xd6, 0x08, 0xd2, 0xf5, 0x41, 0xed,
	0xcd, 0xe8, 0x27, 0x08, 0x87, 0x09, 0xe4, 0xe4, 0xdf, 0xe2, 0xa3, 0xb3, 0xfd, 0x5d, 0x93, 0xba,
	0xd2, 0x24, 0x78, 0x79, 0x6c, 0x7d, 0x1a, 0xe6, 0xd2, 0x68, 0x02, 0x85, 0xa6, 0x5a, 0x99, 0xa6,
	0x30, 0xdd, 0x9d, 0xfa, 0xc7, 0x37, 0x3c, 0x78, 0xf5, 0xd9, 0x39, 0x1a, 0x0b, 0xf1, 0xd5, 0x19,
	0x8c, 0x6d, 0xa4, 0x48, 0x91, 0x58, 0x59, 0xab, 0x88, 0x91, 0xa9, 0x23, 0xbf, 0x1d, 0x13, 0x8b,
	0x14, 0x63, 0xcf, 0xc4, 0x11, 0x8b, 0x3d, 0xb3, 0xed, 0x0c, 0xed, 0x80, 0x73, 0x91, 0x22, 0xe7,
	0x9e, 0xe2, 0x3c, 0x62, 0x9c, 0x7b, 0xee, 0xfd, 0xb8, 0x39, 0x96, 0xfd, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0x0a, 0xb2, 0x64, 0x32, 0x02, 0x00, 0x00,
}