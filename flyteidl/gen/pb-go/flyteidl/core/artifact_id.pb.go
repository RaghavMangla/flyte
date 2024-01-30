// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/artifact_id.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ArtifactKey struct {
	// Project and domain and suffix needs to be unique across a given artifact store.
	Project              string   `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Domain               string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactKey) Reset()         { *m = ArtifactKey{} }
func (m *ArtifactKey) String() string { return proto.CompactTextString(m) }
func (*ArtifactKey) ProtoMessage()    {}
func (*ArtifactKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_1079b0707e23f978, []int{0}
}

func (m *ArtifactKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactKey.Unmarshal(m, b)
}
func (m *ArtifactKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactKey.Marshal(b, m, deterministic)
}
func (m *ArtifactKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactKey.Merge(m, src)
}
func (m *ArtifactKey) XXX_Size() int {
	return xxx_messageInfo_ArtifactKey.Size(m)
}
func (m *ArtifactKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactKey.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactKey proto.InternalMessageInfo

func (m *ArtifactKey) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ArtifactKey) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ArtifactKey) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Only valid for triggers
type ArtifactBindingData struct {
	Index uint32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// These two fields are only relevant in the partition value case
	//
	// Types that are valid to be assigned to PartitionData:
	//	*ArtifactBindingData_PartitionKey
	//	*ArtifactBindingData_BindToTimePartition
	PartitionData isArtifactBindingData_PartitionData `protobuf_oneof:"partition_data"`
	// This is only relevant in the time partition case
	Transform            string   `protobuf:"bytes,4,opt,name=transform,proto3" json:"transform,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactBindingData) Reset()         { *m = ArtifactBindingData{} }
func (m *ArtifactBindingData) String() string { return proto.CompactTextString(m) }
func (*ArtifactBindingData) ProtoMessage()    {}
func (*ArtifactBindingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1079b0707e23f978, []int{1}
}

func (m *ArtifactBindingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactBindingData.Unmarshal(m, b)
}
func (m *ArtifactBindingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactBindingData.Marshal(b, m, deterministic)
}
func (m *ArtifactBindingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactBindingData.Merge(m, src)
}
func (m *ArtifactBindingData) XXX_Size() int {
	return xxx_messageInfo_ArtifactBindingData.Size(m)
}
func (m *ArtifactBindingData) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactBindingData.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactBindingData proto.InternalMessageInfo

func (m *ArtifactBindingData) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type isArtifactBindingData_PartitionData interface {
	isArtifactBindingData_PartitionData()
}

type ArtifactBindingData_PartitionKey struct {
	PartitionKey string `protobuf:"bytes,2,opt,name=partition_key,json=partitionKey,proto3,oneof"`
}

type ArtifactBindingData_BindToTimePartition struct {
	BindToTimePartition bool `protobuf:"varint,3,opt,name=bind_to_time_partition,json=bindToTimePartition,proto3,oneof"`
}

func (*ArtifactBindingData_PartitionKey) isArtifactBindingData_PartitionData() {}

func (*ArtifactBindingData_BindToTimePartition) isArtifactBindingData_PartitionData() {}

func (m *ArtifactBindingData) GetPartitionData() isArtifactBindingData_PartitionData {
	if m != nil {
		return m.PartitionData
	}
	return nil
}

func (m *ArtifactBindingData) GetPartitionKey() string {
	if x, ok := m.GetPartitionData().(*ArtifactBindingData_PartitionKey); ok {
		return x.PartitionKey
	}
	return ""
}

func (m *ArtifactBindingData) GetBindToTimePartition() bool {
	if x, ok := m.GetPartitionData().(*ArtifactBindingData_BindToTimePartition); ok {
		return x.BindToTimePartition
	}
	return false
}

func (m *ArtifactBindingData) GetTransform() string {
	if m != nil {
		return m.Transform
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ArtifactBindingData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ArtifactBindingData_PartitionKey)(nil),
		(*ArtifactBindingData_BindToTimePartition)(nil),
	}
}

type InputBindingData struct {
	Var                  string   `protobuf:"bytes,1,opt,name=var,proto3" json:"var,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputBindingData) Reset()         { *m = InputBindingData{} }
func (m *InputBindingData) String() string { return proto.CompactTextString(m) }
func (*InputBindingData) ProtoMessage()    {}
func (*InputBindingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1079b0707e23f978, []int{2}
}

func (m *InputBindingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputBindingData.Unmarshal(m, b)
}
func (m *InputBindingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputBindingData.Marshal(b, m, deterministic)
}
func (m *InputBindingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputBindingData.Merge(m, src)
}
func (m *InputBindingData) XXX_Size() int {
	return xxx_messageInfo_InputBindingData.Size(m)
}
func (m *InputBindingData) XXX_DiscardUnknown() {
	xxx_messageInfo_InputBindingData.DiscardUnknown(m)
}

var xxx_messageInfo_InputBindingData proto.InternalMessageInfo

func (m *InputBindingData) GetVar() string {
	if m != nil {
		return m.Var
	}
	return ""
}

type LabelValue struct {
	// Types that are valid to be assigned to Value:
	//	*LabelValue_StaticValue
	//	*LabelValue_TimeValue
	//	*LabelValue_TriggeredBinding
	//	*LabelValue_InputBinding
	Value                isLabelValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LabelValue) Reset()         { *m = LabelValue{} }
func (m *LabelValue) String() string { return proto.CompactTextString(m) }
func (*LabelValue) ProtoMessage()    {}
func (*LabelValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_1079b0707e23f978, []int{3}
}

func (m *LabelValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelValue.Unmarshal(m, b)
}
func (m *LabelValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelValue.Marshal(b, m, deterministic)
}
func (m *LabelValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelValue.Merge(m, src)
}
func (m *LabelValue) XXX_Size() int {
	return xxx_messageInfo_LabelValue.Size(m)
}
func (m *LabelValue) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelValue.DiscardUnknown(m)
}

var xxx_messageInfo_LabelValue proto.InternalMessageInfo

type isLabelValue_Value interface {
	isLabelValue_Value()
}

type LabelValue_StaticValue struct {
	StaticValue string `protobuf:"bytes,1,opt,name=static_value,json=staticValue,proto3,oneof"`
}

type LabelValue_TimeValue struct {
	TimeValue *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time_value,json=timeValue,proto3,oneof"`
}

type LabelValue_TriggeredBinding struct {
	TriggeredBinding *ArtifactBindingData `protobuf:"bytes,3,opt,name=triggered_binding,json=triggeredBinding,proto3,oneof"`
}

type LabelValue_InputBinding struct {
	InputBinding *InputBindingData `protobuf:"bytes,4,opt,name=input_binding,json=inputBinding,proto3,oneof"`
}

func (*LabelValue_StaticValue) isLabelValue_Value() {}

func (*LabelValue_TimeValue) isLabelValue_Value() {}

func (*LabelValue_TriggeredBinding) isLabelValue_Value() {}

func (*LabelValue_InputBinding) isLabelValue_Value() {}

func (m *LabelValue) GetValue() isLabelValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *LabelValue) GetStaticValue() string {
	if x, ok := m.GetValue().(*LabelValue_StaticValue); ok {
		return x.StaticValue
	}
	return ""
}

func (m *LabelValue) GetTimeValue() *timestamp.Timestamp {
	if x, ok := m.GetValue().(*LabelValue_TimeValue); ok {
		return x.TimeValue
	}
	return nil
}

func (m *LabelValue) GetTriggeredBinding() *ArtifactBindingData {
	if x, ok := m.GetValue().(*LabelValue_TriggeredBinding); ok {
		return x.TriggeredBinding
	}
	return nil
}

func (m *LabelValue) GetInputBinding() *InputBindingData {
	if x, ok := m.GetValue().(*LabelValue_InputBinding); ok {
		return x.InputBinding
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LabelValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LabelValue_StaticValue)(nil),
		(*LabelValue_TimeValue)(nil),
		(*LabelValue_TriggeredBinding)(nil),
		(*LabelValue_InputBinding)(nil),
	}
}

type Partitions struct {
	Value                map[string]*LabelValue `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Partitions) Reset()         { *m = Partitions{} }
func (m *Partitions) String() string { return proto.CompactTextString(m) }
func (*Partitions) ProtoMessage()    {}
func (*Partitions) Descriptor() ([]byte, []int) {
	return fileDescriptor_1079b0707e23f978, []int{4}
}

func (m *Partitions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Partitions.Unmarshal(m, b)
}
func (m *Partitions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Partitions.Marshal(b, m, deterministic)
}
func (m *Partitions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Partitions.Merge(m, src)
}
func (m *Partitions) XXX_Size() int {
	return xxx_messageInfo_Partitions.Size(m)
}
func (m *Partitions) XXX_DiscardUnknown() {
	xxx_messageInfo_Partitions.DiscardUnknown(m)
}

var xxx_messageInfo_Partitions proto.InternalMessageInfo

func (m *Partitions) GetValue() map[string]*LabelValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type TimePartition struct {
	Value                *LabelValue `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TimePartition) Reset()         { *m = TimePartition{} }
func (m *TimePartition) String() string { return proto.CompactTextString(m) }
func (*TimePartition) ProtoMessage()    {}
func (*TimePartition) Descriptor() ([]byte, []int) {
	return fileDescriptor_1079b0707e23f978, []int{5}
}

func (m *TimePartition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimePartition.Unmarshal(m, b)
}
func (m *TimePartition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimePartition.Marshal(b, m, deterministic)
}
func (m *TimePartition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimePartition.Merge(m, src)
}
func (m *TimePartition) XXX_Size() int {
	return xxx_messageInfo_TimePartition.Size(m)
}
func (m *TimePartition) XXX_DiscardUnknown() {
	xxx_messageInfo_TimePartition.DiscardUnknown(m)
}

var xxx_messageInfo_TimePartition proto.InternalMessageInfo

func (m *TimePartition) GetValue() *LabelValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type ArtifactID struct {
	ArtifactKey *ArtifactKey `protobuf:"bytes,1,opt,name=artifact_key,json=artifactKey,proto3" json:"artifact_key,omitempty"`
	Version     string       `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Think of a partition as a tag on an Artifact, except it's a key-value pair.
	// Different partitions naturally have different versions (execution ids).
	Partitions *Partitions `protobuf:"bytes,3,opt,name=partitions,proto3" json:"partitions,omitempty"`
	// There is no such thing as an empty time partition - if it's not set, then there is no time partition.
	TimePartition        *TimePartition `protobuf:"bytes,4,opt,name=time_partition,json=timePartition,proto3" json:"time_partition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ArtifactID) Reset()         { *m = ArtifactID{} }
func (m *ArtifactID) String() string { return proto.CompactTextString(m) }
func (*ArtifactID) ProtoMessage()    {}
func (*ArtifactID) Descriptor() ([]byte, []int) {
	return fileDescriptor_1079b0707e23f978, []int{6}
}

func (m *ArtifactID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactID.Unmarshal(m, b)
}
func (m *ArtifactID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactID.Marshal(b, m, deterministic)
}
func (m *ArtifactID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactID.Merge(m, src)
}
func (m *ArtifactID) XXX_Size() int {
	return xxx_messageInfo_ArtifactID.Size(m)
}
func (m *ArtifactID) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactID.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactID proto.InternalMessageInfo

func (m *ArtifactID) GetArtifactKey() *ArtifactKey {
	if m != nil {
		return m.ArtifactKey
	}
	return nil
}

func (m *ArtifactID) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ArtifactID) GetPartitions() *Partitions {
	if m != nil {
		return m.Partitions
	}
	return nil
}

func (m *ArtifactID) GetTimePartition() *TimePartition {
	if m != nil {
		return m.TimePartition
	}
	return nil
}

type ArtifactTag struct {
	ArtifactKey          *ArtifactKey `protobuf:"bytes,1,opt,name=artifact_key,json=artifactKey,proto3" json:"artifact_key,omitempty"`
	Value                *LabelValue  `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ArtifactTag) Reset()         { *m = ArtifactTag{} }
func (m *ArtifactTag) String() string { return proto.CompactTextString(m) }
func (*ArtifactTag) ProtoMessage()    {}
func (*ArtifactTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_1079b0707e23f978, []int{7}
}

func (m *ArtifactTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactTag.Unmarshal(m, b)
}
func (m *ArtifactTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactTag.Marshal(b, m, deterministic)
}
func (m *ArtifactTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactTag.Merge(m, src)
}
func (m *ArtifactTag) XXX_Size() int {
	return xxx_messageInfo_ArtifactTag.Size(m)
}
func (m *ArtifactTag) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactTag.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactTag proto.InternalMessageInfo

func (m *ArtifactTag) GetArtifactKey() *ArtifactKey {
	if m != nil {
		return m.ArtifactKey
	}
	return nil
}

func (m *ArtifactTag) GetValue() *LabelValue {
	if m != nil {
		return m.Value
	}
	return nil
}

// Uniqueness constraints for Artifacts
//  - project, domain, name, version, partitions
// Option 2 (tags are standalone, point to an individual artifact id):
//  - project, domain, name, alias (points to one partition if partitioned)
//  - project, domain, name, partition key, partition value
type ArtifactQuery struct {
	// Types that are valid to be assigned to Identifier:
	//	*ArtifactQuery_ArtifactId
	//	*ArtifactQuery_ArtifactTag
	//	*ArtifactQuery_Uri
	//	*ArtifactQuery_Binding
	Identifier           isArtifactQuery_Identifier `protobuf_oneof:"identifier"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ArtifactQuery) Reset()         { *m = ArtifactQuery{} }
func (m *ArtifactQuery) String() string { return proto.CompactTextString(m) }
func (*ArtifactQuery) ProtoMessage()    {}
func (*ArtifactQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_1079b0707e23f978, []int{8}
}

func (m *ArtifactQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactQuery.Unmarshal(m, b)
}
func (m *ArtifactQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactQuery.Marshal(b, m, deterministic)
}
func (m *ArtifactQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactQuery.Merge(m, src)
}
func (m *ArtifactQuery) XXX_Size() int {
	return xxx_messageInfo_ArtifactQuery.Size(m)
}
func (m *ArtifactQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactQuery proto.InternalMessageInfo

type isArtifactQuery_Identifier interface {
	isArtifactQuery_Identifier()
}

type ArtifactQuery_ArtifactId struct {
	ArtifactId *ArtifactID `protobuf:"bytes,1,opt,name=artifact_id,json=artifactId,proto3,oneof"`
}

type ArtifactQuery_ArtifactTag struct {
	ArtifactTag *ArtifactTag `protobuf:"bytes,2,opt,name=artifact_tag,json=artifactTag,proto3,oneof"`
}

type ArtifactQuery_Uri struct {
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3,oneof"`
}

type ArtifactQuery_Binding struct {
	Binding *ArtifactBindingData `protobuf:"bytes,4,opt,name=binding,proto3,oneof"`
}

func (*ArtifactQuery_ArtifactId) isArtifactQuery_Identifier() {}

func (*ArtifactQuery_ArtifactTag) isArtifactQuery_Identifier() {}

func (*ArtifactQuery_Uri) isArtifactQuery_Identifier() {}

func (*ArtifactQuery_Binding) isArtifactQuery_Identifier() {}

func (m *ArtifactQuery) GetIdentifier() isArtifactQuery_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *ArtifactQuery) GetArtifactId() *ArtifactID {
	if x, ok := m.GetIdentifier().(*ArtifactQuery_ArtifactId); ok {
		return x.ArtifactId
	}
	return nil
}

func (m *ArtifactQuery) GetArtifactTag() *ArtifactTag {
	if x, ok := m.GetIdentifier().(*ArtifactQuery_ArtifactTag); ok {
		return x.ArtifactTag
	}
	return nil
}

func (m *ArtifactQuery) GetUri() string {
	if x, ok := m.GetIdentifier().(*ArtifactQuery_Uri); ok {
		return x.Uri
	}
	return ""
}

func (m *ArtifactQuery) GetBinding() *ArtifactBindingData {
	if x, ok := m.GetIdentifier().(*ArtifactQuery_Binding); ok {
		return x.Binding
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ArtifactQuery) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ArtifactQuery_ArtifactId)(nil),
		(*ArtifactQuery_ArtifactTag)(nil),
		(*ArtifactQuery_Uri)(nil),
		(*ArtifactQuery_Binding)(nil),
	}
}

func init() {
	proto.RegisterType((*ArtifactKey)(nil), "flyteidl.core.ArtifactKey")
	proto.RegisterType((*ArtifactBindingData)(nil), "flyteidl.core.ArtifactBindingData")
	proto.RegisterType((*InputBindingData)(nil), "flyteidl.core.InputBindingData")
	proto.RegisterType((*LabelValue)(nil), "flyteidl.core.LabelValue")
	proto.RegisterType((*Partitions)(nil), "flyteidl.core.Partitions")
	proto.RegisterMapType((map[string]*LabelValue)(nil), "flyteidl.core.Partitions.ValueEntry")
	proto.RegisterType((*TimePartition)(nil), "flyteidl.core.TimePartition")
	proto.RegisterType((*ArtifactID)(nil), "flyteidl.core.ArtifactID")
	proto.RegisterType((*ArtifactTag)(nil), "flyteidl.core.ArtifactTag")
	proto.RegisterType((*ArtifactQuery)(nil), "flyteidl.core.ArtifactQuery")
}

func init() { proto.RegisterFile("flyteidl/core/artifact_id.proto", fileDescriptor_1079b0707e23f978) }

var fileDescriptor_1079b0707e23f978 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xb6, 0x9b, 0xfe, 0xd0, 0x71, 0x5c, 0x85, 0x2d, 0xaa, 0xd2, 0xa8, 0xa2, 0x95, 0x29, 0x52,
	0x2f, 0xd8, 0x52, 0x11, 0x12, 0x94, 0xf2, 0x17, 0x0a, 0x72, 0x55, 0x0e, 0xd4, 0x8d, 0x38, 0x70,
	0xb1, 0xd6, 0xf1, 0xc6, 0x2c, 0x8d, 0xbd, 0xd1, 0x66, 0x53, 0xe1, 0x03, 0x2f, 0xc1, 0x9d, 0x07,
	0xe1, 0x81, 0xb8, 0xf3, 0x08, 0x68, 0x77, 0xfd, 0x13, 0x47, 0xad, 0x54, 0xc4, 0x6d, 0x67, 0x76,
	0xe6, 0xdb, 0x99, 0xef, 0xdb, 0x19, 0xd8, 0x1d, 0x8d, 0x73, 0x41, 0x68, 0x3c, 0xf6, 0x86, 0x8c,
	0x13, 0x0f, 0x73, 0x41, 0x47, 0x78, 0x28, 0x42, 0x1a, 0xbb, 0x13, 0xce, 0x04, 0x43, 0x76, 0x19,
	0xe0, 0xca, 0x80, 0xde, 0x6e, 0xc2, 0x58, 0x32, 0x26, 0x9e, 0xba, 0x8c, 0x66, 0x23, 0x4f, 0xd0,
	0x94, 0x4c, 0x05, 0x4e, 0x27, 0x3a, 0xbe, 0x77, 0xbf, 0x09, 0x48, 0x63, 0x92, 0x09, 0x3a, 0xa2,
	0x84, 0xeb, 0x7b, 0xe7, 0x02, 0xac, 0x37, 0xc5, 0x23, 0x67, 0x24, 0x47, 0x5d, 0x58, 0x9b, 0x70,
	0xf6, 0x95, 0x0c, 0x45, 0xd7, 0xdc, 0x33, 0x0f, 0xd6, 0x83, 0xd2, 0x44, 0x5b, 0xb0, 0x1a, 0xb3,
	0x14, 0xd3, 0xac, 0xbb, 0xa4, 0x2e, 0x0a, 0x0b, 0x21, 0x58, 0xce, 0x70, 0x4a, 0xba, 0x2d, 0xe5,
	0x55, 0x67, 0xe7, 0x97, 0x09, 0x9b, 0x25, 0x6a, 0x9f, 0x66, 0x31, 0xcd, 0x92, 0x13, 0x2c, 0x30,
	0xba, 0x07, 0x2b, 0x34, 0x8b, 0xc9, 0x37, 0x85, 0x6d, 0x07, 0xda, 0x40, 0x0f, 0xc1, 0x9e, 0xc8,
	0x46, 0x05, 0x65, 0x59, 0x78, 0x49, 0x72, 0xfd, 0x80, 0x6f, 0x04, 0xed, 0xca, 0x2d, 0x4b, 0x7b,
	0x02, 0x5b, 0x11, 0xcd, 0xe2, 0x50, 0xb0, 0x50, 0x36, 0x19, 0x56, 0x97, 0xea, 0xe9, 0x3b, 0xbe,
	0x11, 0x6c, 0xca, 0xfb, 0x01, 0x1b, 0xd0, 0x94, 0x7c, 0x2c, 0x2f, 0xd1, 0x0e, 0xac, 0x0b, 0x8e,
	0xb3, 0xe9, 0x88, 0xf1, 0xb4, 0xbb, 0xac, 0x8a, 0xac, 0x1d, 0xfd, 0x0e, 0x6c, 0xd4, 0x6f, 0xc7,
	0x58, 0x60, 0x67, 0x1f, 0x3a, 0xa7, 0xd9, 0x64, 0xd6, 0xa8, 0xbb, 0x03, 0xad, 0x2b, 0xcc, 0x0b,
	0x46, 0xe4, 0xd1, 0xf9, 0xb1, 0x04, 0xf0, 0x01, 0x47, 0x64, 0xfc, 0x09, 0x8f, 0x67, 0x04, 0x3d,
	0x80, 0xf6, 0x54, 0x60, 0x41, 0x87, 0xe1, 0x95, 0xb4, 0x75, 0xa4, 0x6f, 0x04, 0x96, 0xf6, 0xea,
	0xa0, 0xe7, 0x00, 0xaa, 0x70, 0x1d, 0x22, 0x9b, 0xb4, 0x0e, 0x7b, 0xae, 0x16, 0xd0, 0x2d, 0x05,
	0x74, 0x07, 0xa5, 0x80, 0xbe, 0x11, 0xac, 0xcb, 0x78, 0x9d, 0x7c, 0x0e, 0x77, 0x05, 0xa7, 0x49,
	0x42, 0x38, 0x89, 0xc3, 0x48, 0xd7, 0xa6, 0x1a, 0xb7, 0x0e, 0x1d, 0xb7, 0xf1, 0x27, 0xdc, 0x6b,
	0x98, 0xf7, 0x8d, 0xa0, 0x53, 0xa5, 0x17, 0x7e, 0xf4, 0x1e, 0x6c, 0x2a, 0x3b, 0xad, 0xe0, 0x96,
	0x15, 0xdc, 0xee, 0x02, 0xdc, 0x22, 0x1b, 0x52, 0x18, 0x3a, 0xe7, 0xeb, 0xaf, 0xc1, 0x8a, 0x6a,
	0xc9, 0xf9, 0x69, 0x02, 0x54, 0xc4, 0x4f, 0xd1, 0x51, 0xe1, 0xef, 0x9a, 0x7b, 0xad, 0x03, 0xeb,
	0x70, 0x7f, 0x01, 0xb7, 0x8e, 0x74, 0x55, 0x8b, 0xef, 0x32, 0xc1, 0xf3, 0x40, 0xa7, 0xf4, 0x2e,
	0x00, 0x6a, 0xa7, 0xe4, 0x5f, 0xfe, 0x8b, 0x82, 0xff, 0x4b, 0x92, 0x23, 0xaf, 0xc4, 0xd6, 0x34,
	0x6e, 0x2f, 0x60, 0xd7, 0xd2, 0x14, 0x80, 0x47, 0x4b, 0x4f, 0x4d, 0xe7, 0x35, 0xd8, 0xcd, 0xbf,
	0xe1, 0xd5, 0x15, 0xde, 0x0a, 0xc5, 0xf9, 0x6d, 0x02, 0x94, 0xf4, 0x9e, 0x9e, 0xa0, 0x17, 0xd0,
	0xae, 0x26, 0xb4, 0x2c, 0x50, 0x6a, 0x7a, 0xbd, 0x1e, 0x67, 0x24, 0x0f, 0x2c, 0xdc, 0x1c, 0xb6,
	0x2b, 0xc2, 0xa7, 0xf2, 0x0b, 0xeb, 0x99, 0x2a, 0x4d, 0xf4, 0x0c, 0xa0, 0xfa, 0x96, 0xd3, 0x42,
	0xe6, 0xed, 0x1b, 0xf9, 0x0b, 0xe6, 0x82, 0xd1, 0x5b, 0xd8, 0x58, 0x18, 0x0f, 0x2d, 0xeb, 0xce,
	0x42, 0x7a, 0x83, 0x89, 0xc0, 0x16, 0xf3, 0xa6, 0xf3, 0xbd, 0xde, 0x0a, 0x03, 0x9c, 0xfc, 0x6f,
	0x9f, 0xff, 0x2a, 0x96, 0xf3, 0xc7, 0x04, 0xbb, 0x44, 0x3b, 0x9f, 0x11, 0x9e, 0xa3, 0x63, 0xb0,
	0xe6, 0x76, 0xe1, 0x0d, 0x7a, 0xd5, 0xca, 0xf8, 0x46, 0x00, 0x65, 0xfc, 0x69, 0x8c, 0x5e, 0xcd,
	0xd5, 0x2f, 0x70, 0x52, 0xcd, 0xde, 0xf5, 0xe9, 0x03, 0x9c, 0xc8, 0xd1, 0xc5, 0x73, 0x04, 0x20,
	0x68, 0xcd, 0x38, 0xd5, 0x3b, 0xce, 0x37, 0x02, 0x69, 0xa0, 0x97, 0xb0, 0xd6, 0x1c, 0x9c, 0xdb,
	0xcd, 0x61, 0x99, 0xd4, 0x6f, 0x03, 0xd4, 0xdb, 0xb8, 0x7f, 0xfc, 0xf9, 0x28, 0xa1, 0xe2, 0xcb,
	0x2c, 0x72, 0x87, 0x2c, 0xf5, 0x14, 0x10, 0xe3, 0x89, 0x3e, 0x78, 0xd5, 0x0e, 0x4f, 0x48, 0xe6,
	0x4d, 0xa2, 0x47, 0x09, 0xf3, 0x1a, 0x6b, 0x3d, 0x5a, 0x55, 0xeb, 0xe3, 0xf1, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xdf, 0x09, 0x3c, 0xc7, 0x3f, 0x06, 0x00, 0x00,
}
