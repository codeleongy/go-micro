// Code generated by protoc-gen-go. DO NOT EDIT.
// source: runtime.proto

package go_micro_runtime

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Service struct {
	// name of the service
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// version of the service
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// git url of the source
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// local path of the source
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// command to execute
	Exec []string `protobuf:"bytes,5,rep,name=exec,proto3" json:"exec,omitempty"`
	// service metadata
	Metadata             map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Service) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Service) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Service) GetExec() []string {
	if m != nil {
		return m.Exec
	}
	return nil
}

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type CreateOptions struct {
	// command to pass in
	Command []string `protobuf:"bytes,1,rep,name=command,proto3" json:"command,omitempty"`
	// environment to pass in
	Env []string `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
	// output to send to
	Output               string   `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOptions) Reset()         { *m = CreateOptions{} }
func (m *CreateOptions) String() string { return proto.CompactTextString(m) }
func (*CreateOptions) ProtoMessage()    {}
func (*CreateOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{1}
}

func (m *CreateOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOptions.Unmarshal(m, b)
}
func (m *CreateOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOptions.Marshal(b, m, deterministic)
}
func (m *CreateOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOptions.Merge(m, src)
}
func (m *CreateOptions) XXX_Size() int {
	return xxx_messageInfo_CreateOptions.Size(m)
}
func (m *CreateOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOptions proto.InternalMessageInfo

func (m *CreateOptions) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *CreateOptions) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *CreateOptions) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

type CreateRequest struct {
	Service              *Service       `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Options              *CreateOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{2}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *CreateRequest) GetOptions() *CreateOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type GetOptions struct {
	// version of the service
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOptions) Reset()         { *m = GetOptions{} }
func (m *GetOptions) String() string { return proto.CompactTextString(m) }
func (*GetOptions) ProtoMessage()    {}
func (*GetOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{4}
}

func (m *GetOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOptions.Unmarshal(m, b)
}
func (m *GetOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOptions.Marshal(b, m, deterministic)
}
func (m *GetOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOptions.Merge(m, src)
}
func (m *GetOptions) XXX_Size() int {
	return xxx_messageInfo_GetOptions.Size(m)
}
func (m *GetOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GetOptions proto.InternalMessageInfo

func (m *GetOptions) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type GetRequest struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Options              *GetOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{5}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetRequest) GetOptions() *GetOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type GetResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{6}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type DeleteRequest struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type UpdateRequest struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{9}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{10}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{11}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{12}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "go.micro.runtime.Service")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.runtime.Service.MetadataEntry")
	proto.RegisterType((*CreateOptions)(nil), "go.micro.runtime.CreateOptions")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.runtime.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.runtime.CreateResponse")
	proto.RegisterType((*GetOptions)(nil), "go.micro.runtime.GetOptions")
	proto.RegisterType((*GetRequest)(nil), "go.micro.runtime.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "go.micro.runtime.GetResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.runtime.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.runtime.DeleteResponse")
	proto.RegisterType((*UpdateRequest)(nil), "go.micro.runtime.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.runtime.UpdateResponse")
	proto.RegisterType((*ListRequest)(nil), "go.micro.runtime.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.runtime.ListResponse")
}

func init() { proto.RegisterFile("runtime.proto", fileDescriptor_86e2dd377c869464) }

var fileDescriptor_86e2dd377c869464 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcb, 0x6f, 0x94, 0x40,
	0x18, 0x2f, 0xcb, 0x16, 0xea, 0x87, 0x98, 0xcd, 0xc4, 0x98, 0x71, 0xe3, 0x63, 0xc3, 0x41, 0x7b,
	0xe2, 0xb0, 0x8d, 0xc6, 0xc7, 0xb1, 0x34, 0x7b, 0x71, 0x63, 0x42, 0x63, 0xe2, 0x75, 0x64, 0xbf,
	0x28, 0xb1, 0x30, 0xc8, 0x0c, 0x1b, 0x7b, 0xf2, 0xe4, 0x1f, 0xed, 0xcd, 0xcc, 0x6b, 0xdb, 0xa5,
	0xd0, 0xcb, 0xde, 0xbe, 0x17, 0x3f, 0x7e, 0x0f, 0x02, 0xc4, 0x6d, 0x57, 0xcb, 0xb2, 0xc2, 0xb4,
	0x69, 0xb9, 0xe4, 0x64, 0xf6, 0x9d, 0xa7, 0x55, 0x59, 0xb4, 0x3c, 0xb5, 0xf3, 0xe4, 0x9f, 0x07,
	0xe1, 0x25, 0xb6, 0xdb, 0xb2, 0x40, 0x42, 0x60, 0x5a, 0xb3, 0x0a, 0xa9, 0xb7, 0xf0, 0x4e, 0x1f,
	0xe4, 0xba, 0x26, 0x14, 0xc2, 0x2d, 0xb6, 0xa2, 0xe4, 0x35, 0x9d, 0xe8, 0xb1, 0x6b, 0xc9, 0x13,
	0x08, 0x04, 0xef, 0xda, 0x02, 0xa9, 0xaf, 0x17, 0xb6, 0x53, 0x28, 0x0d, 0x93, 0x3f, 0xe8, 0xd4,
	0xa0, 0xa8, 0x5a, 0xcd, 0xf0, 0x37, 0x16, 0xf4, 0x78, 0xe1, 0xab, 0x99, 0xaa, 0xc9, 0x39, 0x9c,
	0x54, 0x28, 0xd9, 0x86, 0x49, 0x46, 0x83, 0x85, 0x7f, 0x1a, 0x2d, 0x5f, 0xa7, 0x7d, 0x7a, 0xa9,
	0xa5, 0x96, 0xae, 0xed, 0xe5, 0x45, 0x2d, 0xdb, 0xeb, 0x7c, 0xf7, 0xe0, 0xfc, 0x23, 0xc4, 0x7b,
	0x2b, 0x32, 0x03, 0xff, 0x27, 0x5e, 0x5b, 0x09, 0xaa, 0x24, 0x8f, 0xe1, 0x78, 0xcb, 0xae, 0x3a,
	0xb4, 0xfc, 0x4d, 0xf3, 0x61, 0xf2, 0xce, 0x4b, 0x2e, 0x21, 0x3e, 0x6f, 0x91, 0x49, 0xfc, 0xdc,
	0xc8, 0x92, 0xd7, 0x42, 0x89, 0x2d, 0x78, 0x55, 0xb1, 0x7a, 0x43, 0x3d, 0xcd, 0xd4, 0xb5, 0x0a,
	0x16, 0xeb, 0x2d, 0x9d, 0xe8, 0xa9, 0x2a, 0x95, 0x7c, 0xde, 0xc9, 0xa6, 0x93, 0x4e, 0xbe, 0xe9,
	0x92, 0x3f, 0x0e, 0x34, 0xc7, 0x5f, 0x1d, 0x0a, 0x49, 0xce, 0x20, 0x14, 0x46, 0x85, 0x66, 0x15,
	0x2d, 0x9f, 0x8e, 0xca, 0xcc, 0xdd, 0x25, 0x79, 0x0f, 0x21, 0x37, 0xa4, 0x34, 0xed, 0x68, 0xf9,
	0xf2, 0xee, 0x43, 0x7b, 0xdc, 0x73, 0x77, 0x9f, 0xcc, 0xe0, 0x91, 0x23, 0x20, 0x1a, 0x5e, 0x0b,
	0x4c, 0x5e, 0x01, 0xac, 0x50, 0xde, 0x12, 0x39, 0x9c, 0x68, 0xf2, 0x55, 0xdf, 0x39, 0xde, 0x43,
	0x5f, 0xc3, 0xdb, 0x3e, 0xad, 0x67, 0x77, 0x69, 0xdd, 0xbc, 0xea, 0x86, 0x53, 0x06, 0x91, 0x46,
	0x36, 0x84, 0xc8, 0x1b, 0x38, 0xb1, 0x42, 0x85, 0x36, 0xfa, 0x5e, 0x4f, 0x76, 0xa7, 0x49, 0x06,
	0x71, 0x86, 0x57, 0x78, 0x98, 0xb5, 0xca, 0x1f, 0x87, 0x62, 0xfd, 0xc9, 0x20, 0xfe, 0xd2, 0x6c,
	0xd8, 0xe1, 0xb8, 0x0e, 0xc5, 0xe2, 0xc6, 0x10, 0x7d, 0x2a, 0x85, 0x33, 0x34, 0xb9, 0x80, 0x87,
	0xa6, 0x3d, 0xc8, 0x85, 0xe5, 0x5f, 0x1f, 0xc2, 0xdc, 0x6c, 0xc9, 0x1a, 0x02, 0x93, 0x35, 0x19,
	0xfd, 0x3e, 0xec, 0xdb, 0xe7, 0x8b, 0xf1, 0x03, 0x4b, 0xf7, 0x88, 0x64, 0xe0, 0xaf, 0x50, 0x92,
	0xe1, 0x50, 0x1d, 0xd0, 0xf3, 0x91, 0xed, 0x0e, 0x65, 0x0d, 0x81, 0x31, 0x78, 0x88, 0xd4, 0x5e,
	0x80, 0x43, 0xa4, 0x7a, 0xd9, 0x68, 0x38, 0xe3, 0xeb, 0x10, 0xdc, 0x5e, 0x6e, 0x43, 0x70, 0xbd,
	0x48, 0x8e, 0xc8, 0x0a, 0xa6, 0x2a, 0x05, 0x32, 0x20, 0xe3, 0x56, 0x58, 0xf3, 0x17, 0x63, 0x6b,
	0x07, 0xf4, 0x2d, 0xd0, 0xbf, 0xd4, 0xb3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xce, 0x6e,
	0x31, 0x63, 0x05, 0x00, 0x00,
}
