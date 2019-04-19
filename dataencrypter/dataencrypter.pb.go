// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataencrypter/dataencrypter.proto

package dataencrypter

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type EncryptRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptRequest) Reset()         { *m = EncryptRequest{} }
func (m *EncryptRequest) String() string { return proto.CompactTextString(m) }
func (*EncryptRequest) ProtoMessage()    {}
func (*EncryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d50c45b9e93e7828, []int{0}
}

func (m *EncryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptRequest.Unmarshal(m, b)
}
func (m *EncryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptRequest.Marshal(b, m, deterministic)
}
func (m *EncryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptRequest.Merge(m, src)
}
func (m *EncryptRequest) XXX_Size() int {
	return xxx_messageInfo_EncryptRequest.Size(m)
}
func (m *EncryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptRequest proto.InternalMessageInfo

func (m *EncryptRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type EncryptResponse struct {
	EncryptedData        string   `protobuf:"bytes,1,opt,name=encrypted_data,json=encryptedData,proto3" json:"encrypted_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptResponse) Reset()         { *m = EncryptResponse{} }
func (m *EncryptResponse) String() string { return proto.CompactTextString(m) }
func (*EncryptResponse) ProtoMessage()    {}
func (*EncryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d50c45b9e93e7828, []int{1}
}

func (m *EncryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptResponse.Unmarshal(m, b)
}
func (m *EncryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptResponse.Marshal(b, m, deterministic)
}
func (m *EncryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptResponse.Merge(m, src)
}
func (m *EncryptResponse) XXX_Size() int {
	return xxx_messageInfo_EncryptResponse.Size(m)
}
func (m *EncryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptResponse proto.InternalMessageInfo

func (m *EncryptResponse) GetEncryptedData() string {
	if m != nil {
		return m.EncryptedData
	}
	return ""
}

type DecryptRequest struct {
	EncryptedData        string   `protobuf:"bytes,1,opt,name=encrypted_data,json=encryptedData,proto3" json:"encrypted_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptRequest) Reset()         { *m = DecryptRequest{} }
func (m *DecryptRequest) String() string { return proto.CompactTextString(m) }
func (*DecryptRequest) ProtoMessage()    {}
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d50c45b9e93e7828, []int{2}
}

func (m *DecryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptRequest.Unmarshal(m, b)
}
func (m *DecryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptRequest.Marshal(b, m, deterministic)
}
func (m *DecryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptRequest.Merge(m, src)
}
func (m *DecryptRequest) XXX_Size() int {
	return xxx_messageInfo_DecryptRequest.Size(m)
}
func (m *DecryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptRequest proto.InternalMessageInfo

func (m *DecryptRequest) GetEncryptedData() string {
	if m != nil {
		return m.EncryptedData
	}
	return ""
}

type DecryptResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptResponse) Reset()         { *m = DecryptResponse{} }
func (m *DecryptResponse) String() string { return proto.CompactTextString(m) }
func (*DecryptResponse) ProtoMessage()    {}
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d50c45b9e93e7828, []int{3}
}

func (m *DecryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptResponse.Unmarshal(m, b)
}
func (m *DecryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptResponse.Marshal(b, m, deterministic)
}
func (m *DecryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptResponse.Merge(m, src)
}
func (m *DecryptResponse) XXX_Size() int {
	return xxx_messageInfo_DecryptResponse.Size(m)
}
func (m *DecryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptResponse proto.InternalMessageInfo

func (m *DecryptResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*EncryptRequest)(nil), "EncryptRequest")
	proto.RegisterType((*EncryptResponse)(nil), "EncryptResponse")
	proto.RegisterType((*DecryptRequest)(nil), "DecryptRequest")
	proto.RegisterType((*DecryptResponse)(nil), "DecryptResponse")
}

func init() { proto.RegisterFile("dataencrypter/dataencrypter.proto", fileDescriptor_d50c45b9e93e7828) }

var fileDescriptor_d50c45b9e93e7828 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x49, 0x2c, 0x49,
	0x4c, 0xcd, 0x4b, 0x2e, 0xaa, 0x2c, 0x28, 0x49, 0x2d, 0xd2, 0x47, 0xe1, 0xe9, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x2b, 0xa9, 0x70, 0xf1, 0xb9, 0x42, 0x84, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x84, 0x84, 0xb8, 0x58, 0x40, 0x0a, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25,
	0x0b, 0x2e, 0x7e, 0xb8, 0xaa, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x55, 0x2e, 0x3e, 0x98,
	0x59, 0x29, 0xf1, 0x48, 0x1a, 0x78, 0xe1, 0xa2, 0x2e, 0x20, 0x9d, 0xe6, 0x5c, 0x7c, 0x2e, 0xa9,
	0x28, 0xe6, 0x13, 0xa9, 0x51, 0x95, 0x8b, 0x1f, 0xae, 0x11, 0x6a, 0x25, 0x16, 0x97, 0x19, 0xe5,
	0x73, 0xf1, 0x82, 0x94, 0xbb, 0xc2, 0xbc, 0x25, 0xa4, 0xc7, 0xc5, 0x0e, 0xe5, 0x08, 0xf1, 0xeb,
	0xa1, 0x7a, 0x4d, 0x4a, 0x40, 0x0f, 0xcd, 0x17, 0x4a, 0x0c, 0x20, 0xf5, 0x50, 0x7b, 0x84, 0xf8,
	0xf5, 0x50, 0x9d, 0x2a, 0x25, 0xa0, 0x87, 0xe6, 0x04, 0x25, 0x86, 0x24, 0x36, 0x70, 0xb8, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xc6, 0xd2, 0x12, 0x5c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataEncrypterClient is the client API for DataEncrypter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataEncrypterClient interface {
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
}

type dataEncrypterClient struct {
	cc *grpc.ClientConn
}

func NewDataEncrypterClient(cc *grpc.ClientConn) DataEncrypterClient {
	return &dataEncrypterClient{cc}
}

func (c *dataEncrypterClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := c.cc.Invoke(ctx, "/DataEncrypter/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataEncrypterClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, "/DataEncrypter/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataEncrypterServer is the server API for DataEncrypter service.
type DataEncrypterServer interface {
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
}

// UnimplementedDataEncrypterServer can be embedded to have forward compatible implementations.
type UnimplementedDataEncrypterServer struct {
}

func (*UnimplementedDataEncrypterServer) Encrypt(ctx context.Context, req *EncryptRequest) (*EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (*UnimplementedDataEncrypterServer) Decrypt(ctx context.Context, req *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}

func RegisterDataEncrypterServer(s *grpc.Server, srv DataEncrypterServer) {
	s.RegisterService(&_DataEncrypter_serviceDesc, srv)
}

func _DataEncrypter_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataEncrypterServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataEncrypter/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataEncrypterServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataEncrypter_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataEncrypterServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataEncrypter/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataEncrypterServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataEncrypter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DataEncrypter",
	HandlerType: (*DataEncrypterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _DataEncrypter_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _DataEncrypter_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataencrypter/dataencrypter.proto",
}
