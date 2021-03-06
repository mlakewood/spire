// Code generated by protoc-gen-go. DO NOT EDIT.
// source: upstreamca.proto

package upstreamca

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	plugin "github.com/spiffe/spire/proto/spire/common/plugin"
	grpc "google.golang.org/grpc"
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

type SignedCertificate struct {
	// Contains ASN.1 encoded certificates representing the signed certificate
	// along with any intermediates necessary to chain the certificate back to
	// a certificate present in the upstream_trust_bundle.
	CertChain []byte `protobuf:"bytes,1,opt,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	// The upstream trust bundle.
	Bundle               []byte   `protobuf:"bytes,2,opt,name=bundle,proto3" json:"bundle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedCertificate) Reset()         { *m = SignedCertificate{} }
func (m *SignedCertificate) String() string { return proto.CompactTextString(m) }
func (*SignedCertificate) ProtoMessage()    {}
func (*SignedCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de8507909f29eef, []int{0}
}

func (m *SignedCertificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedCertificate.Unmarshal(m, b)
}
func (m *SignedCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedCertificate.Marshal(b, m, deterministic)
}
func (m *SignedCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedCertificate.Merge(m, src)
}
func (m *SignedCertificate) XXX_Size() int {
	return xxx_messageInfo_SignedCertificate.Size(m)
}
func (m *SignedCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_SignedCertificate proto.InternalMessageInfo

func (m *SignedCertificate) GetCertChain() []byte {
	if m != nil {
		return m.CertChain
	}
	return nil
}

func (m *SignedCertificate) GetBundle() []byte {
	if m != nil {
		return m.Bundle
	}
	return nil
}

type SubmitCSRRequest struct {
	// Certificate signing request
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	// Preferred TTL is the TTL preferred by SPIRE server for signed CA. If
	// zero, the plugin should determine its own TTL value.  Upstream CA
	// plugins are free to ignore this and use their own policies around TTLs.
	PreferredTtl         int32    `protobuf:"varint,2,opt,name=preferred_ttl,json=preferredTtl,proto3" json:"preferred_ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitCSRRequest) Reset()         { *m = SubmitCSRRequest{} }
func (m *SubmitCSRRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitCSRRequest) ProtoMessage()    {}
func (*SubmitCSRRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de8507909f29eef, []int{1}
}

func (m *SubmitCSRRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitCSRRequest.Unmarshal(m, b)
}
func (m *SubmitCSRRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitCSRRequest.Marshal(b, m, deterministic)
}
func (m *SubmitCSRRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitCSRRequest.Merge(m, src)
}
func (m *SubmitCSRRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitCSRRequest.Size(m)
}
func (m *SubmitCSRRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitCSRRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitCSRRequest proto.InternalMessageInfo

func (m *SubmitCSRRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

func (m *SubmitCSRRequest) GetPreferredTtl() int32 {
	if m != nil {
		return m.PreferredTtl
	}
	return 0
}

type SubmitCSRResponse struct {
	// Signed certificate
	SignedCertificate    *SignedCertificate `protobuf:"bytes,3,opt,name=signed_certificate,json=signedCertificate,proto3" json:"signed_certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SubmitCSRResponse) Reset()         { *m = SubmitCSRResponse{} }
func (m *SubmitCSRResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitCSRResponse) ProtoMessage()    {}
func (*SubmitCSRResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de8507909f29eef, []int{2}
}

func (m *SubmitCSRResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitCSRResponse.Unmarshal(m, b)
}
func (m *SubmitCSRResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitCSRResponse.Marshal(b, m, deterministic)
}
func (m *SubmitCSRResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitCSRResponse.Merge(m, src)
}
func (m *SubmitCSRResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitCSRResponse.Size(m)
}
func (m *SubmitCSRResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitCSRResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitCSRResponse proto.InternalMessageInfo

func (m *SubmitCSRResponse) GetSignedCertificate() *SignedCertificate {
	if m != nil {
		return m.SignedCertificate
	}
	return nil
}

func init() {
	proto.RegisterType((*SignedCertificate)(nil), "spire.server.upstreamca.SignedCertificate")
	proto.RegisterType((*SubmitCSRRequest)(nil), "spire.server.upstreamca.SubmitCSRRequest")
	proto.RegisterType((*SubmitCSRResponse)(nil), "spire.server.upstreamca.SubmitCSRResponse")
}

func init() { proto.RegisterFile("upstreamca.proto", fileDescriptor_9de8507909f29eef) }

var fileDescriptor_9de8507909f29eef = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0x4d, 0x0b, 0x8f, 0xc0, 0x7d, 0x90, 0x94, 0x59, 0xbc, 0x47, 0x48, 0x5e, 0x42, 0x78, 0xd1,
	0x20, 0x8b, 0x36, 0xc1, 0x18, 0xd7, 0xda, 0x85, 0x81, 0x95, 0x29, 0xba, 0x90, 0x0d, 0x69, 0xcb,
	0x6d, 0x99, 0xa4, 0x5f, 0xce, 0x87, 0x0b, 0x7f, 0x97, 0x3f, 0xd0, 0x74, 0x3a, 0x14, 0x45, 0x8c,
	0xac, 0x3a, 0xbd, 0xe7, 0xdc, 0x73, 0xee, 0x3d, 0x33, 0x60, 0xc9, 0x82, 0x0b, 0x86, 0x7e, 0x1a,
	0xfa, 0x76, 0xc1, 0x72, 0x91, 0x93, 0xbf, 0xbc, 0xa0, 0x0c, 0x6d, 0x8e, 0xec, 0x05, 0x99, 0xbd,
	0x87, 0x87, 0x23, 0x05, 0x38, 0x61, 0x9e, 0xa6, 0x79, 0xe6, 0x14, 0x89, 0x8c, 0xe9, 0xee, 0x53,
	0xb5, 0x8e, 0x17, 0xd0, 0x5f, 0xd2, 0x38, 0xc3, 0x8d, 0x8b, 0x4c, 0xd0, 0x88, 0x86, 0xbe, 0x40,
	0xf2, 0x0f, 0x20, 0x44, 0x26, 0xd6, 0xe1, 0xd6, 0xa7, 0xd9, 0xc0, 0x18, 0x19, 0x93, 0xae, 0xd7,
	0x29, 0x2b, 0x6e, 0x59, 0x20, 0x7f, 0xa0, 0x15, 0xc8, 0x6c, 0x93, 0xe0, 0xc0, 0x54, 0x90, 0xfe,
	0x1b, 0xcf, 0xc1, 0x5a, 0xca, 0x20, 0xa5, 0xc2, 0x5d, 0x7a, 0x1e, 0x3e, 0x4b, 0xe4, 0x82, 0x58,
	0xd0, 0x08, 0x39, 0xd3, 0x1a, 0xe5, 0x91, 0xfc, 0x87, 0x5e, 0xc1, 0x30, 0x42, 0xc6, 0x70, 0xb3,
	0x16, 0x22, 0x51, 0x22, 0xbf, 0xbc, 0x6e, 0x5d, 0x7c, 0x10, 0xc9, 0xf8, 0x15, 0xfa, 0x1f, 0xa4,
	0x78, 0x91, 0x67, 0x1c, 0xc9, 0x13, 0x10, 0xae, 0x66, 0x5d, 0x87, 0xfb, 0x61, 0x07, 0x8d, 0x91,
	0x31, 0xf9, 0x3d, 0x9b, 0xda, 0xdf, 0x64, 0x60, 0x7f, 0x59, 0xcf, 0xeb, 0xf3, 0xc3, 0xd2, 0xa2,
	0xd9, 0x36, 0x2c, 0x73, 0xd1, 0x6c, 0x9b, 0x56, 0x63, 0xf6, 0x66, 0x02, 0x3c, 0xea, 0x7e, 0xf7,
	0x86, 0xac, 0xa0, 0xe3, 0xe6, 0x59, 0x44, 0x63, 0xc9, 0x90, 0x9c, 0x69, 0x9b, 0x2a, 0x51, 0x5b,
	0x47, 0x59, 0xe3, 0x7a, 0xeb, 0xe1, 0xf9, 0x4f, 0x34, 0xbd, 0x51, 0x04, 0xbd, 0x3b, 0x14, 0xf7,
	0x0a, 0x9e, 0x67, 0x51, 0x4e, 0x2e, 0x8e, 0x36, 0x7e, 0xe2, 0xec, 0x3c, 0xa6, 0xa7, 0x50, 0xb5,
	0x4f, 0x00, 0x9d, 0x3a, 0xce, 0xda, 0xe3, 0x48, 0x54, 0x07, 0xb7, 0x37, 0x9c, 0x9e, 0x42, 0xad,
	0x3c, 0x6e, 0xaf, 0x57, 0x57, 0x31, 0x15, 0x5b, 0x19, 0x94, 0x13, 0x39, 0xbc, 0xa0, 0x51, 0x84,
	0x4e, 0xf5, 0xfe, 0xd4, 0x53, 0xd3, 0xe7, 0x4a, 0xca, 0xd9, 0x4b, 0x05, 0x2d, 0x05, 0x5f, 0xbe,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x70, 0x6f, 0x4b, 0x29, 0xd8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UpstreamCAClient is the client API for UpstreamCA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpstreamCAClient interface {
	// Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	// Returns the  version and related metadata of the installed plugin. */
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
	// Signs a certificate from the request
	SubmitCSR(ctx context.Context, in *SubmitCSRRequest, opts ...grpc.CallOption) (*SubmitCSRResponse, error)
}

type upstreamCAClient struct {
	cc *grpc.ClientConn
}

func NewUpstreamCAClient(cc *grpc.ClientConn) UpstreamCAClient {
	return &upstreamCAClient{cc}
}

func (c *upstreamCAClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := c.cc.Invoke(ctx, "/spire.server.upstreamca.UpstreamCA/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamCAClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := c.cc.Invoke(ctx, "/spire.server.upstreamca.UpstreamCA/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamCAClient) SubmitCSR(ctx context.Context, in *SubmitCSRRequest, opts ...grpc.CallOption) (*SubmitCSRResponse, error) {
	out := new(SubmitCSRResponse)
	err := c.cc.Invoke(ctx, "/spire.server.upstreamca.UpstreamCA/SubmitCSR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpstreamCAServer is the server API for UpstreamCA service.
type UpstreamCAServer interface {
	// Responsible for configuration of the plugin.
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	// Returns the  version and related metadata of the installed plugin. */
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
	// Signs a certificate from the request
	SubmitCSR(context.Context, *SubmitCSRRequest) (*SubmitCSRResponse, error)
}

func RegisterUpstreamCAServer(s *grpc.Server, srv UpstreamCAServer) {
	s.RegisterService(&_UpstreamCA_serviceDesc, srv)
}

func _UpstreamCA_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamCAServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.upstreamca.UpstreamCA/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamCAServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpstreamCA_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamCAServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.upstreamca.UpstreamCA/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamCAServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpstreamCA_SubmitCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitCSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamCAServer).SubmitCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.upstreamca.UpstreamCA/SubmitCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamCAServer).SubmitCSR(ctx, req.(*SubmitCSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UpstreamCA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.server.upstreamca.UpstreamCA",
	HandlerType: (*UpstreamCAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _UpstreamCA_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _UpstreamCA_GetPluginInfo_Handler,
		},
		{
			MethodName: "SubmitCSR",
			Handler:    _UpstreamCA_SubmitCSR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upstreamca.proto",
}
