// Code generated by protoc-gen-go. DO NOT EDIT.
// source: student.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Student struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Sex                  string   `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Student) Reset()         { *m = Student{} }
func (m *Student) String() string { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()    {}
func (*Student) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{0}
}

func (m *Student) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Student.Unmarshal(m, b)
}
func (m *Student) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Student.Marshal(b, m, deterministic)
}
func (m *Student) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Student.Merge(m, src)
}
func (m *Student) XXX_Size() int {
	return xxx_messageInfo_Student.Size(m)
}
func (m *Student) XXX_DiscardUnknown() {
	xxx_messageInfo_Student.DiscardUnknown(m)
}

var xxx_messageInfo_Student proto.InternalMessageInfo

func (m *Student) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Student) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Student) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Student) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

type Reply struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{1}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Reply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type AddStudentArgs struct {
	Stu                  *Student `protobuf:"bytes,1,opt,name=stu,proto3" json:"stu,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddStudentArgs) Reset()         { *m = AddStudentArgs{} }
func (m *AddStudentArgs) String() string { return proto.CompactTextString(m) }
func (*AddStudentArgs) ProtoMessage()    {}
func (*AddStudentArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{2}
}

func (m *AddStudentArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStudentArgs.Unmarshal(m, b)
}
func (m *AddStudentArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStudentArgs.Marshal(b, m, deterministic)
}
func (m *AddStudentArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStudentArgs.Merge(m, src)
}
func (m *AddStudentArgs) XXX_Size() int {
	return xxx_messageInfo_AddStudentArgs.Size(m)
}
func (m *AddStudentArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStudentArgs.DiscardUnknown(m)
}

var xxx_messageInfo_AddStudentArgs proto.InternalMessageInfo

func (m *AddStudentArgs) GetStu() *Student {
	if m != nil {
		return m.Stu
	}
	return nil
}

type AddStudentReply struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddStudentReply) Reset()         { *m = AddStudentReply{} }
func (m *AddStudentReply) String() string { return proto.CompactTextString(m) }
func (*AddStudentReply) ProtoMessage()    {}
func (*AddStudentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{3}
}

func (m *AddStudentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStudentReply.Unmarshal(m, b)
}
func (m *AddStudentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStudentReply.Marshal(b, m, deterministic)
}
func (m *AddStudentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStudentReply.Merge(m, src)
}
func (m *AddStudentReply) XXX_Size() int {
	return xxx_messageInfo_AddStudentReply.Size(m)
}
func (m *AddStudentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStudentReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddStudentReply proto.InternalMessageInfo

func (m *AddStudentReply) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type GetStudentArgs struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStudentArgs) Reset()         { *m = GetStudentArgs{} }
func (m *GetStudentArgs) String() string { return proto.CompactTextString(m) }
func (*GetStudentArgs) ProtoMessage()    {}
func (*GetStudentArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{4}
}

func (m *GetStudentArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStudentArgs.Unmarshal(m, b)
}
func (m *GetStudentArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStudentArgs.Marshal(b, m, deterministic)
}
func (m *GetStudentArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentArgs.Merge(m, src)
}
func (m *GetStudentArgs) XXX_Size() int {
	return xxx_messageInfo_GetStudentArgs.Size(m)
}
func (m *GetStudentArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentArgs proto.InternalMessageInfo

func (m *GetStudentArgs) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type GetStudentReply struct {
	Stu                  *Student `protobuf:"bytes,1,opt,name=stu,proto3" json:"stu,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStudentReply) Reset()         { *m = GetStudentReply{} }
func (m *GetStudentReply) String() string { return proto.CompactTextString(m) }
func (*GetStudentReply) ProtoMessage()    {}
func (*GetStudentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{5}
}

func (m *GetStudentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStudentReply.Unmarshal(m, b)
}
func (m *GetStudentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStudentReply.Marshal(b, m, deterministic)
}
func (m *GetStudentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentReply.Merge(m, src)
}
func (m *GetStudentReply) XXX_Size() int {
	return xxx_messageInfo_GetStudentReply.Size(m)
}
func (m *GetStudentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentReply proto.InternalMessageInfo

func (m *GetStudentReply) GetStu() *Student {
	if m != nil {
		return m.Stu
	}
	return nil
}

type GetAllStudentArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllStudentArgs) Reset()         { *m = GetAllStudentArgs{} }
func (m *GetAllStudentArgs) String() string { return proto.CompactTextString(m) }
func (*GetAllStudentArgs) ProtoMessage()    {}
func (*GetAllStudentArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{6}
}

func (m *GetAllStudentArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllStudentArgs.Unmarshal(m, b)
}
func (m *GetAllStudentArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllStudentArgs.Marshal(b, m, deterministic)
}
func (m *GetAllStudentArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllStudentArgs.Merge(m, src)
}
func (m *GetAllStudentArgs) XXX_Size() int {
	return xxx_messageInfo_GetAllStudentArgs.Size(m)
}
func (m *GetAllStudentArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllStudentArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllStudentArgs proto.InternalMessageInfo

type GetAllStudentReply struct {
	// 声明一个Student的数组
	Stus                 []*Student `protobuf:"bytes,1,rep,name=stus,proto3" json:"stus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAllStudentReply) Reset()         { *m = GetAllStudentReply{} }
func (m *GetAllStudentReply) String() string { return proto.CompactTextString(m) }
func (*GetAllStudentReply) ProtoMessage()    {}
func (*GetAllStudentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{7}
}

func (m *GetAllStudentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllStudentReply.Unmarshal(m, b)
}
func (m *GetAllStudentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllStudentReply.Marshal(b, m, deterministic)
}
func (m *GetAllStudentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllStudentReply.Merge(m, src)
}
func (m *GetAllStudentReply) XXX_Size() int {
	return xxx_messageInfo_GetAllStudentReply.Size(m)
}
func (m *GetAllStudentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllStudentReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllStudentReply proto.InternalMessageInfo

func (m *GetAllStudentReply) GetStus() []*Student {
	if m != nil {
		return m.Stus
	}
	return nil
}

type UpdateStudentArgs struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Stu                  *Student `protobuf:"bytes,2,opt,name=stu,proto3" json:"stu,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStudentArgs) Reset()         { *m = UpdateStudentArgs{} }
func (m *UpdateStudentArgs) String() string { return proto.CompactTextString(m) }
func (*UpdateStudentArgs) ProtoMessage()    {}
func (*UpdateStudentArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{8}
}

func (m *UpdateStudentArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStudentArgs.Unmarshal(m, b)
}
func (m *UpdateStudentArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStudentArgs.Marshal(b, m, deterministic)
}
func (m *UpdateStudentArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStudentArgs.Merge(m, src)
}
func (m *UpdateStudentArgs) XXX_Size() int {
	return xxx_messageInfo_UpdateStudentArgs.Size(m)
}
func (m *UpdateStudentArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStudentArgs.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStudentArgs proto.InternalMessageInfo

func (m *UpdateStudentArgs) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UpdateStudentArgs) GetStu() *Student {
	if m != nil {
		return m.Stu
	}
	return nil
}

type UpdateStudentReply struct {
	Stu                  *Student `protobuf:"bytes,1,opt,name=stu,proto3" json:"stu,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStudentReply) Reset()         { *m = UpdateStudentReply{} }
func (m *UpdateStudentReply) String() string { return proto.CompactTextString(m) }
func (*UpdateStudentReply) ProtoMessage()    {}
func (*UpdateStudentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{9}
}

func (m *UpdateStudentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStudentReply.Unmarshal(m, b)
}
func (m *UpdateStudentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStudentReply.Marshal(b, m, deterministic)
}
func (m *UpdateStudentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStudentReply.Merge(m, src)
}
func (m *UpdateStudentReply) XXX_Size() int {
	return xxx_messageInfo_UpdateStudentReply.Size(m)
}
func (m *UpdateStudentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStudentReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStudentReply proto.InternalMessageInfo

func (m *UpdateStudentReply) GetStu() *Student {
	if m != nil {
		return m.Stu
	}
	return nil
}

type DeleteStudentArgs struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteStudentArgs) Reset()         { *m = DeleteStudentArgs{} }
func (m *DeleteStudentArgs) String() string { return proto.CompactTextString(m) }
func (*DeleteStudentArgs) ProtoMessage()    {}
func (*DeleteStudentArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{10}
}

func (m *DeleteStudentArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStudentArgs.Unmarshal(m, b)
}
func (m *DeleteStudentArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStudentArgs.Marshal(b, m, deterministic)
}
func (m *DeleteStudentArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStudentArgs.Merge(m, src)
}
func (m *DeleteStudentArgs) XXX_Size() int {
	return xxx_messageInfo_DeleteStudentArgs.Size(m)
}
func (m *DeleteStudentArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStudentArgs.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStudentArgs proto.InternalMessageInfo

func (m *DeleteStudentArgs) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type DeleteStudentReply struct {
	Reply                *Reply   `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteStudentReply) Reset()         { *m = DeleteStudentReply{} }
func (m *DeleteStudentReply) String() string { return proto.CompactTextString(m) }
func (*DeleteStudentReply) ProtoMessage()    {}
func (*DeleteStudentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{11}
}

func (m *DeleteStudentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStudentReply.Unmarshal(m, b)
}
func (m *DeleteStudentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStudentReply.Marshal(b, m, deterministic)
}
func (m *DeleteStudentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStudentReply.Merge(m, src)
}
func (m *DeleteStudentReply) XXX_Size() int {
	return xxx_messageInfo_DeleteStudentReply.Size(m)
}
func (m *DeleteStudentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStudentReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStudentReply proto.InternalMessageInfo

func (m *DeleteStudentReply) GetReply() *Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func init() {
	proto.RegisterType((*Student)(nil), "api.Student")
	proto.RegisterType((*Reply)(nil), "api.Reply")
	proto.RegisterType((*AddStudentArgs)(nil), "api.AddStudentArgs")
	proto.RegisterType((*AddStudentReply)(nil), "api.AddStudentReply")
	proto.RegisterType((*GetStudentArgs)(nil), "api.GetStudentArgs")
	proto.RegisterType((*GetStudentReply)(nil), "api.GetStudentReply")
	proto.RegisterType((*GetAllStudentArgs)(nil), "api.GetAllStudentArgs")
	proto.RegisterType((*GetAllStudentReply)(nil), "api.GetAllStudentReply")
	proto.RegisterType((*UpdateStudentArgs)(nil), "api.UpdateStudentArgs")
	proto.RegisterType((*UpdateStudentReply)(nil), "api.UpdateStudentReply")
	proto.RegisterType((*DeleteStudentArgs)(nil), "api.DeleteStudentArgs")
	proto.RegisterType((*DeleteStudentReply)(nil), "api.DeleteStudentReply")
}

func init() { proto.RegisterFile("student.proto", fileDescriptor_94a1c1b032ad0c00) }

var fileDescriptor_94a1c1b032ad0c00 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x15, 0x3b, 0x29, 0xea, 0x14, 0xa7, 0xcd, 0x04, 0x15, 0x63, 0xa1, 0x2a, 0x5a, 0x84,
	0x54, 0x21, 0x1a, 0x93, 0x82, 0x38, 0x70, 0x8b, 0x04, 0xea, 0xb9, 0xae, 0x7a, 0xe2, 0xc2, 0x52,
	0xaf, 0x56, 0x2b, 0xb9, 0xb6, 0xe5, 0x5d, 0x17, 0x10, 0xe2, 0xc2, 0x2b, 0xf0, 0x5c, 0x9c, 0x78,
	0x05, 0x1e, 0x04, 0xed, 0x9f, 0x28, 0xde, 0xae, 0xa0, 0x9c, 0x32, 0xd9, 0xf9, 0xe6, 0x37, 0x33,
	0xfb, 0xd9, 0x86, 0x44, 0xaa, 0xbe, 0x64, 0xb5, 0x5a, 0xb6, 0x5d, 0xa3, 0x1a, 0x8c, 0x69, 0x2b,
	0xb2, 0xc7, 0xbc, 0x69, 0x78, 0xc5, 0x72, 0xda, 0x8a, 0x9c, 0xd6, 0x75, 0xa3, 0xa8, 0x12, 0x4d,
	0x2d, 0xad, 0x24, 0x7b, 0x6e, 0x7e, 0xae, 0x4e, 0x38, 0xab, 0x4f, 0xe4, 0x27, 0xca, 0x39, 0xeb,
	0xf2, 0xa6, 0x35, 0x8a, 0x50, 0x4d, 0xce, 0xe1, 0xde, 0x85, 0xed, 0x80, 0x53, 0x88, 0x44, 0x99,
	0x8e, 0x16, 0xa3, 0xe3, 0x49, 0x11, 0x89, 0x12, 0x11, 0xc6, 0x35, 0xbd, 0x66, 0x69, 0xb4, 0x18,
	0x1d, 0xef, 0x16, 0x26, 0xc6, 0x03, 0x88, 0x29, 0x67, 0x69, 0x6c, 0x44, 0x3a, 0xd4, 0x27, 0x92,
	0x7d, 0x4e, 0xc7, 0x46, 0xa4, 0x43, 0xb2, 0x82, 0x49, 0xc1, 0xda, 0xea, 0x0b, 0x1e, 0xc2, 0x8e,
	0x54, 0x54, 0xf5, 0xd2, 0x41, 0xdd, 0x3f, 0x5d, 0x72, 0x2d, 0xb9, 0xe3, 0xea, 0x90, 0xbc, 0x80,
	0xe9, 0xba, 0x2c, 0xdd, 0x20, 0xeb, 0x8e, 0x4b, 0x3c, 0x82, 0x58, 0xaa, 0xde, 0x14, 0xee, 0x9d,
	0xde, 0x5f, 0xd2, 0x56, 0x2c, 0x5d, 0xba, 0xd0, 0x09, 0xf2, 0x04, 0xf6, 0xb7, 0x15, 0xb6, 0xdd,
	0x01, 0xc4, 0xbd, 0x5b, 0x60, 0xb7, 0xd0, 0x21, 0x21, 0x30, 0x3d, 0x63, 0x6a, 0x88, 0x0d, 0x35,
	0x2b, 0xd8, 0xdf, 0x6a, 0x2c, 0xe8, 0xae, 0xde, 0x73, 0x98, 0x9d, 0x31, 0xb5, 0xae, 0xaa, 0x01,
	0x99, 0xbc, 0x06, 0xf4, 0x0e, 0x2d, 0x6a, 0x01, 0x63, 0x69, 0x2f, 0x20, 0x0e, 0x58, 0x26, 0x43,
	0xde, 0xc1, 0xec, 0xb2, 0x2d, 0xa9, 0x62, 0xff, 0x1c, 0x73, 0x33, 0x53, 0xf4, 0xb7, 0x99, 0x5e,
	0x01, 0x7a, 0x98, 0xff, 0xdb, 0xe4, 0x29, 0xcc, 0xde, 0xb2, 0x8a, 0xdd, 0xd1, 0x5c, 0xef, 0xe6,
	0xc9, 0x36, 0xbb, 0x4d, 0x3a, 0x1d, 0x38, 0x3c, 0x18, 0xbc, 0x49, 0x15, 0x36, 0x71, 0xfa, 0x33,
	0x86, 0xa9, 0x2b, 0xb9, 0x60, 0xdd, 0x8d, 0xb8, 0x62, 0x78, 0x0e, 0xb0, 0xf5, 0x0d, 0xe7, 0xa6,
	0xc6, 0xb7, 0x3e, 0x7b, 0x70, 0xeb, 0xd0, 0x20, 0x49, 0xfa, 0xfd, 0xd7, 0xef, 0x1f, 0x11, 0x92,
	0xbd, 0xfc, 0x66, 0x95, 0xbb, 0x97, 0xe2, 0x8d, 0x5e, 0x02, 0x0b, 0x80, 0xad, 0x83, 0x0e, 0xe9,
	0xdb, 0xee, 0x90, 0xb7, 0x7c, 0x26, 0x8f, 0x0c, 0x72, 0x8e, 0xb3, 0x01, 0x32, 0xff, 0xda, 0x8b,
	0xf2, 0x1b, 0x5e, 0x42, 0xe2, 0xb9, 0x89, 0x87, 0x1b, 0x82, 0x6f, 0x7b, 0xf6, 0x30, 0x3c, 0xb7,
	0xf0, 0xb9, 0x81, 0x27, 0x38, 0x9c, 0x17, 0x3f, 0x40, 0xe2, 0xb9, 0xe4, 0xb0, 0xc1, 0x03, 0xe0,
	0xb0, 0xa1, 0xa3, 0xe4, 0xc8, 0x60, 0xd3, 0x2c, 0x9c, 0xd9, 0x5e, 0xc6, 0x7b, 0x48, 0x3c, 0xab,
	0x5c, 0x87, 0xc0, 0x65, 0xd7, 0x21, 0xb4, 0x75, 0x73, 0x2b, 0xcf, 0xc2, 0x0e, 0x1f, 0x77, 0xcc,
	0x37, 0xe3, 0xe5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xe6, 0x9d, 0xc9, 0x95, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudentServiceClient interface {
	// 添加学生信息
	AddStudent(ctx context.Context, in *AddStudentArgs, opts ...grpc.CallOption) (*AddStudentReply, error)
	// 得到学生信息
	GetStudent(ctx context.Context, in *GetStudentArgs, opts ...grpc.CallOption) (*GetStudentReply, error)
	// 得到所有学生信息
	GetAllStudent(ctx context.Context, in *GetAllStudentArgs, opts ...grpc.CallOption) (*GetAllStudentReply, error)
	// 更新学生信息
	UpdateStudent(ctx context.Context, in *UpdateStudentArgs, opts ...grpc.CallOption) (*UpdateStudentReply, error)
	// 删除学生信息
	DeleteStudent(ctx context.Context, in *DeleteStudentArgs, opts ...grpc.CallOption) (*DeleteStudentReply, error)
}

type studentServiceClient struct {
	cc *grpc.ClientConn
}

func NewStudentServiceClient(cc *grpc.ClientConn) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) AddStudent(ctx context.Context, in *AddStudentArgs, opts ...grpc.CallOption) (*AddStudentReply, error) {
	out := new(AddStudentReply)
	err := c.cc.Invoke(ctx, "/api.StudentService/AddStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetStudent(ctx context.Context, in *GetStudentArgs, opts ...grpc.CallOption) (*GetStudentReply, error) {
	out := new(GetStudentReply)
	err := c.cc.Invoke(ctx, "/api.StudentService/GetStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetAllStudent(ctx context.Context, in *GetAllStudentArgs, opts ...grpc.CallOption) (*GetAllStudentReply, error) {
	out := new(GetAllStudentReply)
	err := c.cc.Invoke(ctx, "/api.StudentService/GetAllStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) UpdateStudent(ctx context.Context, in *UpdateStudentArgs, opts ...grpc.CallOption) (*UpdateStudentReply, error) {
	out := new(UpdateStudentReply)
	err := c.cc.Invoke(ctx, "/api.StudentService/UpdateStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) DeleteStudent(ctx context.Context, in *DeleteStudentArgs, opts ...grpc.CallOption) (*DeleteStudentReply, error) {
	out := new(DeleteStudentReply)
	err := c.cc.Invoke(ctx, "/api.StudentService/DeleteStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
type StudentServiceServer interface {
	// 添加学生信息
	AddStudent(context.Context, *AddStudentArgs) (*AddStudentReply, error)
	// 得到学生信息
	GetStudent(context.Context, *GetStudentArgs) (*GetStudentReply, error)
	// 得到所有学生信息
	GetAllStudent(context.Context, *GetAllStudentArgs) (*GetAllStudentReply, error)
	// 更新学生信息
	UpdateStudent(context.Context, *UpdateStudentArgs) (*UpdateStudentReply, error)
	// 删除学生信息
	DeleteStudent(context.Context, *DeleteStudentArgs) (*DeleteStudentReply, error)
}

// UnimplementedStudentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (*UnimplementedStudentServiceServer) AddStudent(ctx context.Context, req *AddStudentArgs) (*AddStudentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStudent not implemented")
}
func (*UnimplementedStudentServiceServer) GetStudent(ctx context.Context, req *GetStudentArgs) (*GetStudentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}
func (*UnimplementedStudentServiceServer) GetAllStudent(ctx context.Context, req *GetAllStudentArgs) (*GetAllStudentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllStudent not implemented")
}
func (*UnimplementedStudentServiceServer) UpdateStudent(ctx context.Context, req *UpdateStudentArgs) (*UpdateStudentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudent not implemented")
}
func (*UnimplementedStudentServiceServer) DeleteStudent(ctx context.Context, req *DeleteStudentArgs) (*DeleteStudentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudent not implemented")
}

func RegisterStudentServiceServer(s *grpc.Server, srv StudentServiceServer) {
	s.RegisterService(&_StudentService_serviceDesc, srv)
}

func _StudentService_AddStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStudentArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).AddStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StudentService/AddStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).AddStudent(ctx, req.(*AddStudentArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StudentService/GetStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetStudent(ctx, req.(*GetStudentArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetAllStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllStudentArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetAllStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StudentService/GetAllStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetAllStudent(ctx, req.(*GetAllStudentArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_UpdateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).UpdateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StudentService/UpdateStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).UpdateStudent(ctx, req.(*UpdateStudentArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_DeleteStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStudentArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).DeleteStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StudentService/DeleteStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).DeleteStudent(ctx, req.(*DeleteStudentArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _StudentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddStudent",
			Handler:    _StudentService_AddStudent_Handler,
		},
		{
			MethodName: "GetStudent",
			Handler:    _StudentService_GetStudent_Handler,
		},
		{
			MethodName: "GetAllStudent",
			Handler:    _StudentService_GetAllStudent_Handler,
		},
		{
			MethodName: "UpdateStudent",
			Handler:    _StudentService_UpdateStudent_Handler,
		},
		{
			MethodName: "DeleteStudent",
			Handler:    _StudentService_DeleteStudent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student.proto",
}
