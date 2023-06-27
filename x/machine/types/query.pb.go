// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planetmintgo/machine/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7841d43d757203, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7841d43d757203, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetMachineByPublicKeyRequest struct {
	PublicKey string `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (m *QueryGetMachineByPublicKeyRequest) Reset()         { *m = QueryGetMachineByPublicKeyRequest{} }
func (m *QueryGetMachineByPublicKeyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetMachineByPublicKeyRequest) ProtoMessage()    {}
func (*QueryGetMachineByPublicKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7841d43d757203, []int{2}
}
func (m *QueryGetMachineByPublicKeyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetMachineByPublicKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetMachineByPublicKeyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetMachineByPublicKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetMachineByPublicKeyRequest.Merge(m, src)
}
func (m *QueryGetMachineByPublicKeyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetMachineByPublicKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetMachineByPublicKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetMachineByPublicKeyRequest proto.InternalMessageInfo

func (m *QueryGetMachineByPublicKeyRequest) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type QueryGetMachineByPublicKeyResponse struct {
	Machine *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
}

func (m *QueryGetMachineByPublicKeyResponse) Reset()         { *m = QueryGetMachineByPublicKeyResponse{} }
func (m *QueryGetMachineByPublicKeyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetMachineByPublicKeyResponse) ProtoMessage()    {}
func (*QueryGetMachineByPublicKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7841d43d757203, []int{3}
}
func (m *QueryGetMachineByPublicKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetMachineByPublicKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetMachineByPublicKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetMachineByPublicKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetMachineByPublicKeyResponse.Merge(m, src)
}
func (m *QueryGetMachineByPublicKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetMachineByPublicKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetMachineByPublicKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetMachineByPublicKeyResponse proto.InternalMessageInfo

func (m *QueryGetMachineByPublicKeyResponse) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "planetmintgo.machine.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "planetmintgo.machine.QueryParamsResponse")
	proto.RegisterType((*QueryGetMachineByPublicKeyRequest)(nil), "planetmintgo.machine.QueryGetMachineByPublicKeyRequest")
	proto.RegisterType((*QueryGetMachineByPublicKeyResponse)(nil), "planetmintgo.machine.QueryGetMachineByPublicKeyResponse")
}

func init() { proto.RegisterFile("planetmintgo/machine/query.proto", fileDescriptor_bf7841d43d757203) }

var fileDescriptor_bf7841d43d757203 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6a, 0x1a, 0x41,
	0x1c, 0xc7, 0x77, 0xa5, 0xb5, 0x38, 0xbd, 0x4d, 0x2d, 0x94, 0x45, 0x57, 0x5d, 0x28, 0xd8, 0x42,
	0x77, 0xd0, 0x1e, 0x2c, 0xa5, 0x14, 0x2a, 0x85, 0x1e, 0x4a, 0x41, 0xf7, 0x58, 0x28, 0xcb, 0xac,
	0x0c, 0xdb, 0x25, 0xee, 0xcc, 0xe8, 0x8c, 0x21, 0x4b, 0xc8, 0x25, 0xe4, 0x01, 0x02, 0x79, 0x29,
	0x8f, 0x42, 0x2e, 0x39, 0x85, 0x44, 0xf3, 0x20, 0xc1, 0x99, 0x51, 0x23, 0x6e, 0x24, 0x39, 0x39,
	0x8c, 0x9f, 0xef, 0x9f, 0xdf, 0xfc, 0x16, 0xd4, 0xf9, 0x10, 0x53, 0x22, 0xd3, 0x84, 0xca, 0x98,
	0xa1, 0x14, 0x0f, 0xfe, 0x27, 0x94, 0xa0, 0xd1, 0x84, 0x8c, 0x33, 0x9f, 0x8f, 0x99, 0x64, 0xb0,
	0xfc, 0x90, 0xf0, 0x0d, 0xe1, 0x94, 0x63, 0x16, 0x33, 0x05, 0xa0, 0xe5, 0x49, 0xb3, 0x4e, 0x25,
	0x66, 0x2c, 0x1e, 0x12, 0x84, 0x79, 0x82, 0x30, 0xa5, 0x4c, 0x62, 0x99, 0x30, 0x2a, 0xcc, 0xbf,
	0x1f, 0x07, 0x4c, 0xa4, 0x4c, 0xa0, 0x08, 0x0b, 0x13, 0x81, 0x0e, 0x5b, 0x11, 0x91, 0xb8, 0x85,
	0x38, 0x8e, 0x13, 0xaa, 0x60, 0xc3, 0x36, 0x72, 0x7b, 0x71, 0x3c, 0xc6, 0xe9, 0xca, 0xce, 0xcb,
	0x45, 0xcc, 0xaf, 0x66, 0xbc, 0x32, 0x80, 0xfd, 0x65, 0x50, 0x4f, 0x09, 0x03, 0x32, 0x9a, 0x10,
	0x21, 0xbd, 0x3e, 0x78, 0xb3, 0x75, 0x2b, 0x38, 0xa3, 0x82, 0xc0, 0xaf, 0xa0, 0xa8, 0x03, 0xde,
	0xd9, 0x75, 0xbb, 0xf9, 0xba, 0x5d, 0xf1, 0xf3, 0x46, 0xf7, 0xb5, 0xaa, 0xfb, 0x62, 0x7a, 0x5d,
	0xb3, 0x02, 0xa3, 0xf0, 0x7e, 0x80, 0x86, 0xb2, 0xfc, 0x45, 0xe4, 0x1f, 0xcd, 0x75, 0xb3, 0xde,
	0x24, 0x1a, 0x26, 0x83, 0xdf, 0x24, 0x33, 0xb9, 0xb0, 0x02, 0x4a, 0x7c, 0x75, 0xa7, 0x32, 0x4a,
	0xc1, 0xe6, 0xc2, 0xfb, 0x07, 0xbc, 0x7d, 0x16, 0xa6, 0x64, 0x07, 0xbc, 0x32, 0x45, 0x4c, 0xcb,
	0x6a, 0x7e, 0x4b, 0x63, 0x11, 0xac, 0xe8, 0xf6, 0x6d, 0x01, 0xbc, 0x54, 0xfe, 0xf0, 0xcc, 0x06,
	0x45, 0x3d, 0x04, 0x6c, 0xe6, 0x8b, 0x77, 0xdf, 0xcc, 0xf9, 0xf0, 0x04, 0x52, 0x57, 0xf4, 0xde,
	0x9f, 0x5e, 0xde, 0x5d, 0x14, 0x6a, 0xb0, 0x8a, 0x36, 0x92, 0x4f, 0x3b, 0x5b, 0x84, 0x33, 0x1b,
	0xbc, 0xcd, 0x9d, 0x15, 0x76, 0xf6, 0x64, 0xed, 0x7b, 0x60, 0xe7, 0xcb, 0xf3, 0x85, 0xa6, 0xf3,
	0x4f, 0xd5, 0xf9, 0x3b, 0xfc, 0xf6, 0x48, 0xe7, 0x98, 0xc8, 0xd0, 0x9c, 0xc3, 0x28, 0x0b, 0xf5,
	0xd6, 0xc2, 0x03, 0x92, 0xa1, 0xe3, 0xf5, 0x06, 0x4f, 0xba, 0x9d, 0xe9, 0xdc, 0xb5, 0x67, 0x73,
	0xd7, 0xbe, 0x99, 0xbb, 0xf6, 0xf9, 0xc2, 0xb5, 0x66, 0x0b, 0xd7, 0xba, 0x5a, 0xb8, 0xd6, 0xdf,
	0xea, 0xb6, 0xed, 0xd1, 0xda, 0x58, 0x66, 0x9c, 0x88, 0xa8, 0xa8, 0x3e, 0xd7, 0xcf, 0xf7, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xd2, 0x46, 0x91, 0x7e, 0x8f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of GetMachineByPublicKey items.
	GetMachineByPublicKey(ctx context.Context, in *QueryGetMachineByPublicKeyRequest, opts ...grpc.CallOption) (*QueryGetMachineByPublicKeyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/planetmintgo.machine.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetMachineByPublicKey(ctx context.Context, in *QueryGetMachineByPublicKeyRequest, opts ...grpc.CallOption) (*QueryGetMachineByPublicKeyResponse, error) {
	out := new(QueryGetMachineByPublicKeyResponse)
	err := c.cc.Invoke(ctx, "/planetmintgo.machine.Query/GetMachineByPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of GetMachineByPublicKey items.
	GetMachineByPublicKey(context.Context, *QueryGetMachineByPublicKeyRequest) (*QueryGetMachineByPublicKeyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) GetMachineByPublicKey(ctx context.Context, req *QueryGetMachineByPublicKeyRequest) (*QueryGetMachineByPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachineByPublicKey not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/planetmintgo.machine.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetMachineByPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetMachineByPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetMachineByPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/planetmintgo.machine.Query/GetMachineByPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetMachineByPublicKey(ctx, req.(*QueryGetMachineByPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "planetmintgo.machine.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetMachineByPublicKey",
			Handler:    _Query_GetMachineByPublicKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "planetmintgo/machine/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetMachineByPublicKeyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetMachineByPublicKeyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetMachineByPublicKeyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetMachineByPublicKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetMachineByPublicKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetMachineByPublicKeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Machine != nil {
		{
			size, err := m.Machine.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetMachineByPublicKeyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetMachineByPublicKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Machine != nil {
		l = m.Machine.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetMachineByPublicKeyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetMachineByPublicKeyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetMachineByPublicKeyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetMachineByPublicKeyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetMachineByPublicKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetMachineByPublicKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Machine", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Machine == nil {
				m.Machine = &Machine{}
			}
			if err := m.Machine.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
