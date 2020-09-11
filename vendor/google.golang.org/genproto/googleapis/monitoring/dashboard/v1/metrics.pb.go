// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/dashboard/v1/metrics.proto

package dashboard

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

// Defines the possible types of spark chart supported by the `Scorecard`.
type SparkChartType int32

const (
	// Not allowed in well-formed requests.
	SparkChartType_SPARK_CHART_TYPE_UNSPECIFIED SparkChartType = 0
	// The sparkline will be rendered as a small line chart.
	SparkChartType_SPARK_LINE SparkChartType = 1
	// The sparkbar will be rendered as a small bar chart.
	SparkChartType_SPARK_BAR SparkChartType = 2
)

var SparkChartType_name = map[int32]string{
	0: "SPARK_CHART_TYPE_UNSPECIFIED",
	1: "SPARK_LINE",
	2: "SPARK_BAR",
}

var SparkChartType_value = map[string]int32{
	"SPARK_CHART_TYPE_UNSPECIFIED": 0,
	"SPARK_LINE":                   1,
	"SPARK_BAR":                    2,
}

func (x SparkChartType) String() string {
	return proto.EnumName(SparkChartType_name, int32(x))
}

func (SparkChartType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6380832df6588178, []int{0}
}

// The color suggests an interpretation to the viewer when actual values cross
// the threshold. Comments on each color provide UX guidance on how users can
// be expected to interpret a given state color.
type Threshold_Color int32

const (
	// Color is unspecified. Not allowed in well-formed requests.
	Threshold_COLOR_UNSPECIFIED Threshold_Color = 0
	// Crossing the threshold is "concerning" behavior.
	Threshold_YELLOW Threshold_Color = 4
	// Crossing the threshold is "emergency" behavior.
	Threshold_RED Threshold_Color = 6
)

var Threshold_Color_name = map[int32]string{
	0: "COLOR_UNSPECIFIED",
	4: "YELLOW",
	6: "RED",
}

var Threshold_Color_value = map[string]int32{
	"COLOR_UNSPECIFIED": 0,
	"YELLOW":            4,
	"RED":               6,
}

func (x Threshold_Color) String() string {
	return proto.EnumName(Threshold_Color_name, int32(x))
}

func (Threshold_Color) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6380832df6588178, []int{3, 0}
}

// Whether the threshold is considered crossed by an actual value above or
// below its threshold value.
type Threshold_Direction int32

const (
	// Not allowed in well-formed requests.
	Threshold_DIRECTION_UNSPECIFIED Threshold_Direction = 0
	// The threshold will be considered crossed if the actual value is above
	// the threshold value.
	Threshold_ABOVE Threshold_Direction = 1
	// The threshold will be considered crossed if the actual value is below
	// the threshold value.
	Threshold_BELOW Threshold_Direction = 2
)

var Threshold_Direction_name = map[int32]string{
	0: "DIRECTION_UNSPECIFIED",
	1: "ABOVE",
	2: "BELOW",
}

var Threshold_Direction_value = map[string]int32{
	"DIRECTION_UNSPECIFIED": 0,
	"ABOVE":                 1,
	"BELOW":                 2,
}

func (x Threshold_Direction) String() string {
	return proto.EnumName(Threshold_Direction_name, int32(x))
}

func (Threshold_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6380832df6588178, []int{3, 1}
}

// TimeSeriesQuery collects the set of supported methods for querying time
// series data from the Stackdriver metrics API.
type TimeSeriesQuery struct {
	// Parameters needed to obtain data for the chart.
	//
	// Types that are valid to be assigned to Source:
	//	*TimeSeriesQuery_TimeSeriesFilter
	//	*TimeSeriesQuery_TimeSeriesFilterRatio
	Source isTimeSeriesQuery_Source `protobuf_oneof:"source"`
	// The unit of data contained in fetched time series. If non-empty, this
	// unit will override any unit that accompanies fetched data. The format is
	// the same as the
	// [`unit`](/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors)
	// field in `MetricDescriptor`.
	UnitOverride         string   `protobuf:"bytes,5,opt,name=unit_override,json=unitOverride,proto3" json:"unit_override,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeSeriesQuery) Reset()         { *m = TimeSeriesQuery{} }
func (m *TimeSeriesQuery) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesQuery) ProtoMessage()    {}
func (*TimeSeriesQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_6380832df6588178, []int{0}
}

func (m *TimeSeriesQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeriesQuery.Unmarshal(m, b)
}
func (m *TimeSeriesQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeriesQuery.Marshal(b, m, deterministic)
}
func (m *TimeSeriesQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeriesQuery.Merge(m, src)
}
func (m *TimeSeriesQuery) XXX_Size() int {
	return xxx_messageInfo_TimeSeriesQuery.Size(m)
}
func (m *TimeSeriesQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeriesQuery.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeriesQuery proto.InternalMessageInfo

type isTimeSeriesQuery_Source interface {
	isTimeSeriesQuery_Source()
}

type TimeSeriesQuery_TimeSeriesFilter struct {
	TimeSeriesFilter *TimeSeriesFilter `protobuf:"bytes,1,opt,name=time_series_filter,json=timeSeriesFilter,proto3,oneof"`
}

type TimeSeriesQuery_TimeSeriesFilterRatio struct {
	TimeSeriesFilterRatio *TimeSeriesFilterRatio `protobuf:"bytes,2,opt,name=time_series_filter_ratio,json=timeSeriesFilterRatio,proto3,oneof"`
}

func (*TimeSeriesQuery_TimeSeriesFilter) isTimeSeriesQuery_Source() {}

func (*TimeSeriesQuery_TimeSeriesFilterRatio) isTimeSeriesQuery_Source() {}

func (m *TimeSeriesQuery) GetSource() isTimeSeriesQuery_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *TimeSeriesQuery) GetTimeSeriesFilter() *TimeSeriesFilter {
	if x, ok := m.GetSource().(*TimeSeriesQuery_TimeSeriesFilter); ok {
		return x.TimeSeriesFilter
	}
	return nil
}

func (m *TimeSeriesQuery) GetTimeSeriesFilterRatio() *TimeSeriesFilterRatio {
	if x, ok := m.GetSource().(*TimeSeriesQuery_TimeSeriesFilterRatio); ok {
		return x.TimeSeriesFilterRatio
	}
	return nil
}

func (m *TimeSeriesQuery) GetUnitOverride() string {
	if m != nil {
		return m.UnitOverride
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TimeSeriesQuery) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TimeSeriesQuery_TimeSeriesFilter)(nil),
		(*TimeSeriesQuery_TimeSeriesFilterRatio)(nil),
	}
}

// A filter that defines a subset of time series data that is displayed in a
// widget. Time series data is fetched using the
// [`ListTimeSeries`](/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list)
// method.
type TimeSeriesFilter struct {
	// Required. The [monitoring filter](/monitoring/api/v3/filters) that identifies the
	// metric types, resources, and projects to query.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// By default, the raw time series data is returned.
	// Use this field to combine multiple time series for different views of the
	// data.
	Aggregation *Aggregation `protobuf:"bytes,2,opt,name=aggregation,proto3" json:"aggregation,omitempty"`
	// Selects an optional time series filter.
	//
	// Types that are valid to be assigned to OutputFilter:
	//	*TimeSeriesFilter_PickTimeSeriesFilter
	//	*TimeSeriesFilter_StatisticalTimeSeriesFilter
	OutputFilter         isTimeSeriesFilter_OutputFilter `protobuf_oneof:"output_filter"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *TimeSeriesFilter) Reset()         { *m = TimeSeriesFilter{} }
func (m *TimeSeriesFilter) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesFilter) ProtoMessage()    {}
func (*TimeSeriesFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_6380832df6588178, []int{1}
}

func (m *TimeSeriesFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeriesFilter.Unmarshal(m, b)
}
func (m *TimeSeriesFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeriesFilter.Marshal(b, m, deterministic)
}
func (m *TimeSeriesFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeriesFilter.Merge(m, src)
}
func (m *TimeSeriesFilter) XXX_Size() int {
	return xxx_messageInfo_TimeSeriesFilter.Size(m)
}
func (m *TimeSeriesFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeriesFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeriesFilter proto.InternalMessageInfo

func (m *TimeSeriesFilter) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *TimeSeriesFilter) GetAggregation() *Aggregation {
	if m != nil {
		return m.Aggregation
	}
	return nil
}

type isTimeSeriesFilter_OutputFilter interface {
	isTimeSeriesFilter_OutputFilter()
}

type TimeSeriesFilter_PickTimeSeriesFilter struct {
	PickTimeSeriesFilter *PickTimeSeriesFilter `protobuf:"bytes,4,opt,name=pick_time_series_filter,json=pickTimeSeriesFilter,proto3,oneof"`
}

type TimeSeriesFilter_StatisticalTimeSeriesFilter struct {
	StatisticalTimeSeriesFilter *StatisticalTimeSeriesFilter `protobuf:"bytes,5,opt,name=statistical_time_series_filter,json=statisticalTimeSeriesFilter,proto3,oneof"`
}

func (*TimeSeriesFilter_PickTimeSeriesFilter) isTimeSeriesFilter_OutputFilter() {}

func (*TimeSeriesFilter_StatisticalTimeSeriesFilter) isTimeSeriesFilter_OutputFilter() {}

func (m *TimeSeriesFilter) GetOutputFilter() isTimeSeriesFilter_OutputFilter {
	if m != nil {
		return m.OutputFilter
	}
	return nil
}

func (m *TimeSeriesFilter) GetPickTimeSeriesFilter() *PickTimeSeriesFilter {
	if x, ok := m.GetOutputFilter().(*TimeSeriesFilter_PickTimeSeriesFilter); ok {
		return x.PickTimeSeriesFilter
	}
	return nil
}

func (m *TimeSeriesFilter) GetStatisticalTimeSeriesFilter() *StatisticalTimeSeriesFilter {
	if x, ok := m.GetOutputFilter().(*TimeSeriesFilter_StatisticalTimeSeriesFilter); ok {
		return x.StatisticalTimeSeriesFilter
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TimeSeriesFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TimeSeriesFilter_PickTimeSeriesFilter)(nil),
		(*TimeSeriesFilter_StatisticalTimeSeriesFilter)(nil),
	}
}

// A pair of time series filters that define a ratio computation. The output
// time series is the pair-wise division of each aligned element from the
// numerator and denominator time series.
type TimeSeriesFilterRatio struct {
	// The numerator of the ratio.
	Numerator *TimeSeriesFilterRatio_RatioPart `protobuf:"bytes,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	// The denominator of the ratio.
	Denominator *TimeSeriesFilterRatio_RatioPart `protobuf:"bytes,2,opt,name=denominator,proto3" json:"denominator,omitempty"`
	// Apply a second aggregation after the ratio is computed.
	SecondaryAggregation *Aggregation `protobuf:"bytes,3,opt,name=secondary_aggregation,json=secondaryAggregation,proto3" json:"secondary_aggregation,omitempty"`
	// Selects an optional filter that is applied to the time series after
	// computing the ratio.
	//
	// Types that are valid to be assigned to OutputFilter:
	//	*TimeSeriesFilterRatio_PickTimeSeriesFilter
	//	*TimeSeriesFilterRatio_StatisticalTimeSeriesFilter
	OutputFilter         isTimeSeriesFilterRatio_OutputFilter `protobuf_oneof:"output_filter"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *TimeSeriesFilterRatio) Reset()         { *m = TimeSeriesFilterRatio{} }
func (m *TimeSeriesFilterRatio) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesFilterRatio) ProtoMessage()    {}
func (*TimeSeriesFilterRatio) Descriptor() ([]byte, []int) {
	return fileDescriptor_6380832df6588178, []int{2}
}

func (m *TimeSeriesFilterRatio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeriesFilterRatio.Unmarshal(m, b)
}
func (m *TimeSeriesFilterRatio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeriesFilterRatio.Marshal(b, m, deterministic)
}
func (m *TimeSeriesFilterRatio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeriesFilterRatio.Merge(m, src)
}
func (m *TimeSeriesFilterRatio) XXX_Size() int {
	return xxx_messageInfo_TimeSeriesFilterRatio.Size(m)
}
func (m *TimeSeriesFilterRatio) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeriesFilterRatio.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeriesFilterRatio proto.InternalMessageInfo

func (m *TimeSeriesFilterRatio) GetNumerator() *TimeSeriesFilterRatio_RatioPart {
	if m != nil {
		return m.Numerator
	}
	return nil
}

func (m *TimeSeriesFilterRatio) GetDenominator() *TimeSeriesFilterRatio_RatioPart {
	if m != nil {
		return m.Denominator
	}
	return nil
}

func (m *TimeSeriesFilterRatio) GetSecondaryAggregation() *Aggregation {
	if m != nil {
		return m.SecondaryAggregation
	}
	return nil
}

type isTimeSeriesFilterRatio_OutputFilter interface {
	isTimeSeriesFilterRatio_OutputFilter()
}

type TimeSeriesFilterRatio_PickTimeSeriesFilter struct {
	PickTimeSeriesFilter *PickTimeSeriesFilter `protobuf:"bytes,4,opt,name=pick_time_series_filter,json=pickTimeSeriesFilter,proto3,oneof"`
}

type TimeSeriesFilterRatio_StatisticalTimeSeriesFilter struct {
	StatisticalTimeSeriesFilter *StatisticalTimeSeriesFilter `protobuf:"bytes,5,opt,name=statistical_time_series_filter,json=statisticalTimeSeriesFilter,proto3,oneof"`
}

func (*TimeSeriesFilterRatio_PickTimeSeriesFilter) isTimeSeriesFilterRatio_OutputFilter() {}

func (*TimeSeriesFilterRatio_StatisticalTimeSeriesFilter) isTimeSeriesFilterRatio_OutputFilter() {}

func (m *TimeSeriesFilterRatio) GetOutputFilter() isTimeSeriesFilterRatio_OutputFilter {
	if m != nil {
		return m.OutputFilter
	}
	return nil
}

func (m *TimeSeriesFilterRatio) GetPickTimeSeriesFilter() *PickTimeSeriesFilter {
	if x, ok := m.GetOutputFilter().(*TimeSeriesFilterRatio_PickTimeSeriesFilter); ok {
		return x.PickTimeSeriesFilter
	}
	return nil
}

func (m *TimeSeriesFilterRatio) GetStatisticalTimeSeriesFilter() *StatisticalTimeSeriesFilter {
	if x, ok := m.GetOutputFilter().(*TimeSeriesFilterRatio_StatisticalTimeSeriesFilter); ok {
		return x.StatisticalTimeSeriesFilter
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TimeSeriesFilterRatio) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TimeSeriesFilterRatio_PickTimeSeriesFilter)(nil),
		(*TimeSeriesFilterRatio_StatisticalTimeSeriesFilter)(nil),
	}
}

// Describes a query to build the numerator or denominator of a
// TimeSeriesFilterRatio.
type TimeSeriesFilterRatio_RatioPart struct {
	// Required. The [monitoring filter](/monitoring/api/v3/filters) that identifies the
	// metric types, resources, and projects to query.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// By default, the raw time series data is returned.
	// Use this field to combine multiple time series for different views of the
	// data.
	Aggregation          *Aggregation `protobuf:"bytes,2,opt,name=aggregation,proto3" json:"aggregation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TimeSeriesFilterRatio_RatioPart) Reset()         { *m = TimeSeriesFilterRatio_RatioPart{} }
func (m *TimeSeriesFilterRatio_RatioPart) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesFilterRatio_RatioPart) ProtoMessage()    {}
func (*TimeSeriesFilterRatio_RatioPart) Descriptor() ([]byte, []int) {
	return fileDescriptor_6380832df6588178, []int{2, 0}
}

func (m *TimeSeriesFilterRatio_RatioPart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeriesFilterRatio_RatioPart.Unmarshal(m, b)
}
func (m *TimeSeriesFilterRatio_RatioPart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeriesFilterRatio_RatioPart.Marshal(b, m, deterministic)
}
func (m *TimeSeriesFilterRatio_RatioPart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeriesFilterRatio_RatioPart.Merge(m, src)
}
func (m *TimeSeriesFilterRatio_RatioPart) XXX_Size() int {
	return xxx_messageInfo_TimeSeriesFilterRatio_RatioPart.Size(m)
}
func (m *TimeSeriesFilterRatio_RatioPart) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeriesFilterRatio_RatioPart.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeriesFilterRatio_RatioPart proto.InternalMessageInfo

func (m *TimeSeriesFilterRatio_RatioPart) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *TimeSeriesFilterRatio_RatioPart) GetAggregation() *Aggregation {
	if m != nil {
		return m.Aggregation
	}
	return nil
}

// Defines a threshold for categorizing time series values.
type Threshold struct {
	// A label for the threshold.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// The value of the threshold. The value should be defined in the native scale
	// of the metric.
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	// The state color for this threshold. Color is not allowed in a XyChart.
	Color Threshold_Color `protobuf:"varint,3,opt,name=color,proto3,enum=google.monitoring.dashboard.v1.Threshold_Color" json:"color,omitempty"`
	// The direction for the current threshold. Direction is not allowed in a
	// XyChart.
	Direction            Threshold_Direction `protobuf:"varint,4,opt,name=direction,proto3,enum=google.monitoring.dashboard.v1.Threshold_Direction" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Threshold) Reset()         { *m = Threshold{} }
func (m *Threshold) String() string { return proto.CompactTextString(m) }
func (*Threshold) ProtoMessage()    {}
func (*Threshold) Descriptor() ([]byte, []int) {
	return fileDescriptor_6380832df6588178, []int{3}
}

func (m *Threshold) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Threshold.Unmarshal(m, b)
}
func (m *Threshold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Threshold.Marshal(b, m, deterministic)
}
func (m *Threshold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Threshold.Merge(m, src)
}
func (m *Threshold) XXX_Size() int {
	return xxx_messageInfo_Threshold.Size(m)
}
func (m *Threshold) XXX_DiscardUnknown() {
	xxx_messageInfo_Threshold.DiscardUnknown(m)
}

var xxx_messageInfo_Threshold proto.InternalMessageInfo

func (m *Threshold) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Threshold) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Threshold) GetColor() Threshold_Color {
	if m != nil {
		return m.Color
	}
	return Threshold_COLOR_UNSPECIFIED
}

func (m *Threshold) GetDirection() Threshold_Direction {
	if m != nil {
		return m.Direction
	}
	return Threshold_DIRECTION_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("google.monitoring.dashboard.v1.SparkChartType", SparkChartType_name, SparkChartType_value)
	proto.RegisterEnum("google.monitoring.dashboard.v1.Threshold_Color", Threshold_Color_name, Threshold_Color_value)
	proto.RegisterEnum("google.monitoring.dashboard.v1.Threshold_Direction", Threshold_Direction_name, Threshold_Direction_value)
	proto.RegisterType((*TimeSeriesQuery)(nil), "google.monitoring.dashboard.v1.TimeSeriesQuery")
	proto.RegisterType((*TimeSeriesFilter)(nil), "google.monitoring.dashboard.v1.TimeSeriesFilter")
	proto.RegisterType((*TimeSeriesFilterRatio)(nil), "google.monitoring.dashboard.v1.TimeSeriesFilterRatio")
	proto.RegisterType((*TimeSeriesFilterRatio_RatioPart)(nil), "google.monitoring.dashboard.v1.TimeSeriesFilterRatio.RatioPart")
	proto.RegisterType((*Threshold)(nil), "google.monitoring.dashboard.v1.Threshold")
}

func init() {
	proto.RegisterFile("google/monitoring/dashboard/v1/metrics.proto", fileDescriptor_6380832df6588178)
}

var fileDescriptor_6380832df6588178 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x95, 0xdd, 0x6e, 0x12, 0x4f,
	0x18, 0xc6, 0x0b, 0x14, 0xfe, 0xff, 0x7d, 0xfb, 0x85, 0x93, 0x12, 0xb1, 0x35, 0xb5, 0xc1, 0x93,
	0xc6, 0x9a, 0x5d, 0xdb, 0xea, 0x51, 0x4d, 0x0c, 0x0b, 0x5b, 0x4b, 0xa4, 0x85, 0x0e, 0xa8, 0xa9,
	0x89, 0xd9, 0x0e, 0xbb, 0xd3, 0x65, 0xd2, 0xdd, 0x9d, 0xcd, 0xec, 0x80, 0xe9, 0x81, 0x27, 0x7a,
	0x0b, 0xde, 0x8c, 0xe7, 0x5e, 0x88, 0x97, 0x62, 0x76, 0x97, 0x02, 0xa1, 0x54, 0xac, 0x31, 0x9e,
	0x78, 0x42, 0x98, 0x77, 0xde, 0xe7, 0xf9, 0xcd, 0xc7, 0xb3, 0x19, 0x78, 0xec, 0x70, 0xee, 0xb8,
	0x54, 0xf3, 0xb8, 0xcf, 0x24, 0x17, 0xcc, 0x77, 0x34, 0x9b, 0x84, 0xdd, 0x0e, 0x27, 0xc2, 0xd6,
	0xfa, 0x3b, 0x9a, 0x47, 0xa5, 0x60, 0x56, 0xa8, 0x06, 0x82, 0x4b, 0x8e, 0x36, 0x92, 0x6e, 0x75,
	0xd4, 0xad, 0x0e, 0xbb, 0xd5, 0xfe, 0xce, 0xda, 0x83, 0x81, 0x1b, 0x09, 0x98, 0x76, 0xce, 0xa8,
	0x6b, 0x9b, 0x1d, 0xda, 0x25, 0x7d, 0xc6, 0x45, 0x62, 0xb0, 0xb6, 0x3d, 0x03, 0x67, 0x71, 0xcf,
	0xe3, 0x7e, 0xd2, 0x5c, 0xfa, 0x92, 0x86, 0x95, 0x36, 0xf3, 0x68, 0x8b, 0x0a, 0x46, 0xc3, 0x93,
	0x1e, 0x15, 0x97, 0xe8, 0x0c, 0x90, 0x64, 0x1e, 0x35, 0xc3, 0xb8, 0x66, 0x9e, 0x33, 0x57, 0x52,
	0x51, 0x4c, 0x6d, 0xa6, 0xb6, 0x16, 0x76, 0x9f, 0xa8, 0x3f, 0x5f, 0x9e, 0x3a, 0x32, 0x3b, 0x88,
	0x75, 0x87, 0x73, 0x38, 0x2f, 0x27, 0x6a, 0x28, 0x80, 0xe2, 0x75, 0x82, 0x29, 0x88, 0x64, 0xbc,
	0x98, 0x8e, 0x39, 0xcf, 0x6e, 0xcb, 0xc1, 0x91, 0xf8, 0x70, 0x0e, 0x17, 0xe4, 0xb4, 0x09, 0xf4,
	0x10, 0x96, 0x7a, 0x3e, 0x93, 0x26, 0xef, 0x53, 0x21, 0x98, 0x4d, 0x8b, 0xd9, 0xcd, 0xd4, 0x96,
	0x82, 0x17, 0xa3, 0x62, 0x63, 0x50, 0xd3, 0xff, 0x87, 0x5c, 0xc8, 0x7b, 0xc2, 0xa2, 0xa5, 0xcf,
	0x19, 0xc8, 0x4f, 0x12, 0xd0, 0x3a, 0xe4, 0xc6, 0xce, 0x42, 0xd1, 0x33, 0xdf, 0xcb, 0x69, 0x3c,
	0x28, 0xa1, 0x23, 0x58, 0x20, 0x8e, 0x23, 0xa8, 0x13, 0xe1, 0xfc, 0xc1, 0x2e, 0xb6, 0x67, 0xed,
	0xa2, 0x3c, 0x92, 0xe0, 0x71, 0x3d, 0xf2, 0xe0, 0x6e, 0xc0, 0xac, 0x0b, 0x73, 0xca, 0x45, 0xcc,
	0xc7, 0xd6, 0x4f, 0x67, 0x59, 0x37, 0x99, 0x75, 0x31, 0xe5, 0x32, 0x56, 0x83, 0x29, 0x75, 0xf4,
	0x29, 0x05, 0x1b, 0xa1, 0x24, 0x92, 0x85, 0x92, 0x59, 0xc4, 0x9d, 0x86, 0xcd, 0xc6, 0xd8, 0xfd,
	0x59, 0xd8, 0xd6, 0xc8, 0x65, 0x0a, 0x7d, 0x3d, 0xbc, 0x79, 0x5a, 0x5f, 0x81, 0x25, 0xde, 0x93,
	0x41, 0x4f, 0x0e, 0x90, 0xa5, 0xaf, 0x59, 0x28, 0x4c, 0xbd, 0x67, 0xf4, 0x1e, 0x14, 0xbf, 0xe7,
	0x51, 0x41, 0x24, 0xbf, 0x4a, 0xe6, 0x8b, 0xdf, 0x4a, 0x8c, 0x1a, 0xff, 0x36, 0x89, 0x90, 0x78,
	0xe4, 0x88, 0x08, 0x2c, 0xd8, 0xd4, 0xe7, 0x1e, 0xf3, 0x63, 0x40, 0xfa, 0xcf, 0x00, 0xc6, 0x3d,
	0xd1, 0x19, 0x14, 0x42, 0x6a, 0x71, 0xdf, 0x26, 0xe2, 0xd2, 0x1c, 0x4f, 0x4e, 0xe6, 0xf6, 0xc9,
	0x59, 0x1d, 0x3a, 0x95, 0xff, 0xed, 0x08, 0xad, 0x7d, 0x00, 0x65, 0x78, 0xde, 0x7f, 0xf3, 0x7b,
	0xbd, 0x9e, 0xdd, 0x6f, 0x69, 0x50, 0xda, 0x5d, 0x41, 0xc3, 0x2e, 0x77, 0x6d, 0xb4, 0x0a, 0x59,
	0x97, 0x74, 0xa8, 0x9b, 0xac, 0x04, 0x27, 0x83, 0xa8, 0xda, 0x27, 0x6e, 0x8f, 0xc6, 0xf4, 0x14,
	0x4e, 0x06, 0xc8, 0x80, 0xac, 0xc5, 0x5d, 0x2e, 0xe2, 0x24, 0x2c, 0xef, 0x6a, 0x33, 0x63, 0x77,
	0x45, 0x51, 0x2b, 0x91, 0x0c, 0x27, 0x6a, 0x74, 0x02, 0x8a, 0xcd, 0x04, 0xb5, 0xe2, 0xed, 0xcd,
	0xc7, 0x56, 0x7b, 0xbf, 0x6e, 0x55, 0xbd, 0x92, 0xe2, 0x91, 0x4b, 0x69, 0x0f, 0xb2, 0x31, 0x02,
	0x15, 0xe0, 0x4e, 0xa5, 0x51, 0x6f, 0x60, 0xf3, 0xf5, 0x71, 0xab, 0x69, 0x54, 0x6a, 0x07, 0x35,
	0xa3, 0x9a, 0x9f, 0x43, 0x00, 0xb9, 0x53, 0xa3, 0x5e, 0x6f, 0xbc, 0xcd, 0xcf, 0xa3, 0xff, 0x20,
	0x83, 0x8d, 0x6a, 0x3e, 0x57, 0x7a, 0x0e, 0xca, 0xd0, 0x0c, 0xdd, 0x83, 0x42, 0xb5, 0x86, 0x8d,
	0x4a, 0xbb, 0xd6, 0x38, 0x9e, 0x10, 0x2b, 0x90, 0x2d, 0xeb, 0x8d, 0x37, 0x46, 0x3e, 0x15, 0xfd,
	0xd5, 0x8d, 0xc8, 0x26, 0xfd, 0xe8, 0x04, 0x96, 0x5b, 0x01, 0x11, 0x17, 0x95, 0x2e, 0x11, 0xb2,
	0x7d, 0x19, 0x50, 0xb4, 0x09, 0xf7, 0x5b, 0xcd, 0x32, 0x7e, 0x65, 0x56, 0x0e, 0xcb, 0xb8, 0x6d,
	0xb6, 0x4f, 0x9b, 0xc6, 0x84, 0xd3, 0x32, 0x40, 0xd2, 0x51, 0xaf, 0x1d, 0x47, 0x76, 0x4b, 0xa0,
	0x24, 0x63, 0xbd, 0x8c, 0xf3, 0x69, 0xfd, 0x23, 0x94, 0x2c, 0xee, 0xcd, 0x38, 0x0a, 0x7d, 0xf1,
	0x28, 0x79, 0x95, 0x9b, 0xd1, 0x33, 0xd9, 0x4c, 0xbd, 0x7b, 0x39, 0xe8, 0x77, 0xb8, 0x4b, 0x7c,
	0x47, 0xe5, 0xc2, 0xd1, 0x1c, 0xea, 0xc7, 0x8f, 0xa8, 0x96, 0x4c, 0x91, 0x80, 0x85, 0x37, 0x3d,
	0xba, 0xfb, 0xc3, 0x41, 0x27, 0x17, 0x6b, 0xf6, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x4d,
	0x94, 0xe9, 0x16, 0x08, 0x00, 0x00,
}