// Code generated by protoc-gen-go. DO NOT EDIT.
// source: EduData.proto

package protoTest

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EduData_SubjectType int32

const (
	EduData_KOREAN   EduData_SubjectType = 0
	EduData_ENGLISH  EduData_SubjectType = 1
	EduData_HISTORY  EduData_SubjectType = 2
	EduData_COMPUTER EduData_SubjectType = 3
	EduData_ETC      EduData_SubjectType = 4
)

var EduData_SubjectType_name = map[int32]string{
	0: "KOREAN",
	1: "ENGLISH",
	2: "HISTORY",
	3: "COMPUTER",
	4: "ETC",
}

var EduData_SubjectType_value = map[string]int32{
	"KOREAN":   0,
	"ENGLISH":  1,
	"HISTORY":  2,
	"COMPUTER": 3,
	"ETC":      4,
}

func (x EduData_SubjectType) String() string {
	return proto.EnumName(EduData_SubjectType_name, int32(x))
}

func (EduData_SubjectType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59dbc371a62aefe6, []int{0, 0}
}

type EduData struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string             `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string             `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Subjects             []*EduData_Subject `protobuf:"bytes,4,rep,name=subjects,proto3" json:"subjects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EduData) Reset()         { *m = EduData{} }
func (m *EduData) String() string { return proto.CompactTextString(m) }
func (*EduData) ProtoMessage()    {}
func (*EduData) Descriptor() ([]byte, []int) {
	return fileDescriptor_59dbc371a62aefe6, []int{0}
}

func (m *EduData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EduData.Unmarshal(m, b)
}
func (m *EduData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EduData.Marshal(b, m, deterministic)
}
func (m *EduData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EduData.Merge(m, src)
}
func (m *EduData) XXX_Size() int {
	return xxx_messageInfo_EduData.Size(m)
}
func (m *EduData) XXX_DiscardUnknown() {
	xxx_messageInfo_EduData.DiscardUnknown(m)
}

var xxx_messageInfo_EduData proto.InternalMessageInfo

func (m *EduData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EduData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EduData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EduData) GetSubjects() []*EduData_Subject {
	if m != nil {
		return m.Subjects
	}
	return nil
}

type EduData_Subject struct {
	SubjectName          string              `protobuf:"bytes,1,opt,name=SubjectName,proto3" json:"SubjectName,omitempty"`
	SubjectRank          string              `protobuf:"bytes,2,opt,name=SubjectRank,proto3" json:"SubjectRank,omitempty"`
	Type                 EduData_SubjectType `protobuf:"varint,3,opt,name=type,proto3,enum=protoTest.EduData_SubjectType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *EduData_Subject) Reset()         { *m = EduData_Subject{} }
func (m *EduData_Subject) String() string { return proto.CompactTextString(m) }
func (*EduData_Subject) ProtoMessage()    {}
func (*EduData_Subject) Descriptor() ([]byte, []int) {
	return fileDescriptor_59dbc371a62aefe6, []int{0, 0}
}

func (m *EduData_Subject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EduData_Subject.Unmarshal(m, b)
}
func (m *EduData_Subject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EduData_Subject.Marshal(b, m, deterministic)
}
func (m *EduData_Subject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EduData_Subject.Merge(m, src)
}
func (m *EduData_Subject) XXX_Size() int {
	return xxx_messageInfo_EduData_Subject.Size(m)
}
func (m *EduData_Subject) XXX_DiscardUnknown() {
	xxx_messageInfo_EduData_Subject.DiscardUnknown(m)
}

var xxx_messageInfo_EduData_Subject proto.InternalMessageInfo

func (m *EduData_Subject) GetSubjectName() string {
	if m != nil {
		return m.SubjectName
	}
	return ""
}

func (m *EduData_Subject) GetSubjectRank() string {
	if m != nil {
		return m.SubjectRank
	}
	return ""
}

func (m *EduData_Subject) GetType() EduData_SubjectType {
	if m != nil {
		return m.Type
	}
	return EduData_KOREAN
}

// Our address book file is just one of these.
type EduHistory struct {
	Data                 []*EduData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EduHistory) Reset()         { *m = EduHistory{} }
func (m *EduHistory) String() string { return proto.CompactTextString(m) }
func (*EduHistory) ProtoMessage()    {}
func (*EduHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_59dbc371a62aefe6, []int{1}
}

func (m *EduHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EduHistory.Unmarshal(m, b)
}
func (m *EduHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EduHistory.Marshal(b, m, deterministic)
}
func (m *EduHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EduHistory.Merge(m, src)
}
func (m *EduHistory) XXX_Size() int {
	return xxx_messageInfo_EduHistory.Size(m)
}
func (m *EduHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_EduHistory.DiscardUnknown(m)
}

var xxx_messageInfo_EduHistory proto.InternalMessageInfo

func (m *EduHistory) GetData() []*EduData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("protoTest.EduData_SubjectType", EduData_SubjectType_name, EduData_SubjectType_value)
	proto.RegisterType((*EduData)(nil), "protoTest.EduData")
	proto.RegisterType((*EduData_Subject)(nil), "protoTest.EduData.Subject")
	proto.RegisterType((*EduHistory)(nil), "protoTest.EduHistory")
}

func init() { proto.RegisterFile("EduData.proto", fileDescriptor_59dbc371a62aefe6) }

var fileDescriptor_59dbc371a62aefe6 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4d, 0x4f, 0x83, 0x40,
	0x10, 0x86, 0xe5, 0xc3, 0xd2, 0x0e, 0xda, 0x90, 0x89, 0x07, 0xd2, 0x83, 0x21, 0x1c, 0x4c, 0x4f,
	0x1c, 0xd0, 0x78, 0x37, 0x75, 0x23, 0xf5, 0x03, 0xcc, 0xb2, 0x1e, 0x3c, 0x6e, 0xdd, 0x3d, 0xa0,
	0xb6, 0x90, 0xb2, 0x1c, 0x38, 0xfa, 0xfb, 0xfc, 0x53, 0xa6, 0x5b, 0x44, 0x12, 0xd3, 0x13, 0x33,
	0xcf, 0xfb, 0x86, 0x79, 0xb2, 0x70, 0x4a, 0x44, 0x73, 0xcb, 0x15, 0x8f, 0xaa, 0x6d, 0xa9, 0x4a,
	0x9c, 0xe8, 0x0f, 0x93, 0xb5, 0x0a, 0xbf, 0x4d, 0x70, 0xba, 0x10, 0x11, 0xec, 0x0d, 0x5f, 0x4b,
	0xdf, 0x08, 0x8c, 0xf9, 0x84, 0xea, 0x19, 0xa7, 0x60, 0x16, 0xc2, 0x37, 0x35, 0x31, 0x0b, 0x81,
	0x67, 0x70, 0x2c, 0xd7, 0xbc, 0xf8, 0xf4, 0x2d, 0x8d, 0xf6, 0x0b, 0x5e, 0xc3, 0xb8, 0x6e, 0x56,
	0xef, 0xf2, 0x4d, 0xd5, 0xbe, 0x1d, 0x58, 0x73, 0x37, 0x9e, 0x45, 0xfd, 0x8d, 0xe8, 0xf7, 0x78,
	0xbe, 0xaf, 0xd0, 0xbe, 0x3b, 0xfb, 0x32, 0xc0, 0xe9, 0x28, 0x06, 0xe0, 0x76, 0x63, 0xfa, 0x27,
	0x31, 0x44, 0x83, 0x06, 0xe5, 0x9b, 0x8f, 0x4e, 0x6a, 0x88, 0x30, 0x06, 0x5b, 0xb5, 0x95, 0xd4,
	0x72, 0xd3, 0xf8, 0xfc, 0xb0, 0x03, 0x6b, 0x2b, 0x49, 0x75, 0x37, 0xbc, 0xef, 0xff, 0xba, 0x83,
	0x08, 0x30, 0x7a, 0xc8, 0x28, 0xb9, 0x49, 0xbd, 0x23, 0x74, 0xc1, 0x21, 0xe9, 0xdd, 0xe3, 0x32,
	0x4f, 0x3c, 0x63, 0xb7, 0x24, 0xcb, 0x9c, 0x65, 0xf4, 0xd5, 0x33, 0xf1, 0x04, 0xc6, 0x8b, 0xec,
	0xe9, 0xf9, 0x85, 0x11, 0xea, 0x59, 0xe8, 0x80, 0x45, 0xd8, 0xc2, 0xb3, 0xc3, 0x2b, 0x00, 0x22,
	0x9a, 0xa4, 0xa8, 0x55, 0xb9, 0x6d, 0xf1, 0x02, 0x6c, 0xc1, 0x15, 0xf7, 0x0d, 0xfd, 0x22, 0xf8,
	0xdf, 0x86, 0xea, 0x7c, 0x35, 0xd2, 0xc1, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x6c,
	0x96, 0xf2, 0xa6, 0x01, 0x00, 0x00,
}