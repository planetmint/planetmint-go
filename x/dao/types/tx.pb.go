// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planetmintgo/dao/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgReissueRDDLProposal struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Proposer    string `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Tx          string `protobuf:"bytes,3,opt,name=tx,proto3" json:"tx,omitempty"`
	Blockheight uint64 `protobuf:"varint,4,opt,name=blockheight,proto3" json:"blockheight,omitempty"`
}

func (m *MsgReissueRDDLProposal) Reset()         { *m = MsgReissueRDDLProposal{} }
func (m *MsgReissueRDDLProposal) String() string { return proto.CompactTextString(m) }
func (*MsgReissueRDDLProposal) ProtoMessage()    {}
func (*MsgReissueRDDLProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_7117c47dbc1828c7, []int{0}
}
func (m *MsgReissueRDDLProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReissueRDDLProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReissueRDDLProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReissueRDDLProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReissueRDDLProposal.Merge(m, src)
}
func (m *MsgReissueRDDLProposal) XXX_Size() int {
	return m.Size()
}
func (m *MsgReissueRDDLProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReissueRDDLProposal.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReissueRDDLProposal proto.InternalMessageInfo

func (m *MsgReissueRDDLProposal) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgReissueRDDLProposal) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *MsgReissueRDDLProposal) GetTx() string {
	if m != nil {
		return m.Tx
	}
	return ""
}

func (m *MsgReissueRDDLProposal) GetBlockheight() uint64 {
	if m != nil {
		return m.Blockheight
	}
	return 0
}

type MsgReissueRDDLProposalResponse struct {
}

func (m *MsgReissueRDDLProposalResponse) Reset()         { *m = MsgReissueRDDLProposalResponse{} }
func (m *MsgReissueRDDLProposalResponse) String() string { return proto.CompactTextString(m) }
func (*MsgReissueRDDLProposalResponse) ProtoMessage()    {}
func (*MsgReissueRDDLProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7117c47dbc1828c7, []int{1}
}
func (m *MsgReissueRDDLProposalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReissueRDDLProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReissueRDDLProposalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReissueRDDLProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReissueRDDLProposalResponse.Merge(m, src)
}
func (m *MsgReissueRDDLProposalResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgReissueRDDLProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReissueRDDLProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReissueRDDLProposalResponse proto.InternalMessageInfo

type MsgMintToken struct {
	Creator     string       `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	MintRequest *MintRequest `protobuf:"bytes,2,opt,name=mintRequest,proto3" json:"mintRequest,omitempty"`
}

func (m *MsgMintToken) Reset()         { *m = MsgMintToken{} }
func (m *MsgMintToken) String() string { return proto.CompactTextString(m) }
func (*MsgMintToken) ProtoMessage()    {}
func (*MsgMintToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_7117c47dbc1828c7, []int{2}
}
func (m *MsgMintToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintToken.Merge(m, src)
}
func (m *MsgMintToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintToken proto.InternalMessageInfo

func (m *MsgMintToken) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgMintToken) GetMintRequest() *MintRequest {
	if m != nil {
		return m.MintRequest
	}
	return nil
}

type MsgMintTokenResponse struct {
}

func (m *MsgMintTokenResponse) Reset()         { *m = MsgMintTokenResponse{} }
func (m *MsgMintTokenResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMintTokenResponse) ProtoMessage()    {}
func (*MsgMintTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7117c47dbc1828c7, []int{3}
}
func (m *MsgMintTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintTokenResponse.Merge(m, src)
}
func (m *MsgMintTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintTokenResponse proto.InternalMessageInfo

type MsgReissueRDDLResult struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Proposer    string `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	TxId        string `protobuf:"bytes,3,opt,name=txId,proto3" json:"txId,omitempty"`
	BlockHeight uint64 `protobuf:"varint,4,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
}

func (m *MsgReissueRDDLResult) Reset()         { *m = MsgReissueRDDLResult{} }
func (m *MsgReissueRDDLResult) String() string { return proto.CompactTextString(m) }
func (*MsgReissueRDDLResult) ProtoMessage()    {}
func (*MsgReissueRDDLResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_7117c47dbc1828c7, []int{4}
}
func (m *MsgReissueRDDLResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReissueRDDLResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReissueRDDLResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReissueRDDLResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReissueRDDLResult.Merge(m, src)
}
func (m *MsgReissueRDDLResult) XXX_Size() int {
	return m.Size()
}
func (m *MsgReissueRDDLResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReissueRDDLResult.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReissueRDDLResult proto.InternalMessageInfo

func (m *MsgReissueRDDLResult) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgReissueRDDLResult) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *MsgReissueRDDLResult) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *MsgReissueRDDLResult) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type MsgReissueRDDLResultResponse struct {
}

func (m *MsgReissueRDDLResultResponse) Reset()         { *m = MsgReissueRDDLResultResponse{} }
func (m *MsgReissueRDDLResultResponse) String() string { return proto.CompactTextString(m) }
func (*MsgReissueRDDLResultResponse) ProtoMessage()    {}
func (*MsgReissueRDDLResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7117c47dbc1828c7, []int{5}
}
func (m *MsgReissueRDDLResultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReissueRDDLResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReissueRDDLResultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReissueRDDLResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReissueRDDLResultResponse.Merge(m, src)
}
func (m *MsgReissueRDDLResultResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgReissueRDDLResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReissueRDDLResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReissueRDDLResultResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgReissueRDDLProposal)(nil), "planetmintgo.dao.MsgReissueRDDLProposal")
	proto.RegisterType((*MsgReissueRDDLProposalResponse)(nil), "planetmintgo.dao.MsgReissueRDDLProposalResponse")
	proto.RegisterType((*MsgMintToken)(nil), "planetmintgo.dao.MsgMintToken")
	proto.RegisterType((*MsgMintTokenResponse)(nil), "planetmintgo.dao.MsgMintTokenResponse")
	proto.RegisterType((*MsgReissueRDDLResult)(nil), "planetmintgo.dao.MsgReissueRDDLResult")
	proto.RegisterType((*MsgReissueRDDLResultResponse)(nil), "planetmintgo.dao.MsgReissueRDDLResultResponse")
}

func init() { proto.RegisterFile("planetmintgo/dao/tx.proto", fileDescriptor_7117c47dbc1828c7) }

var fileDescriptor_7117c47dbc1828c7 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xf5, 0xca, 0xa6, 0xad, 0xc7, 0xa5, 0xb4, 0xdb, 0x62, 0x54, 0xd1, 0x2e, 0x42, 0x01, 0xe3,
	0x4b, 0xa4, 0xe0, 0x7c, 0x40, 0x20, 0xf8, 0x10, 0x43, 0x04, 0x61, 0x93, 0x53, 0x2e, 0x41, 0xb6,
	0x17, 0x59, 0x58, 0xd6, 0xca, 0xda, 0x15, 0x28, 0xb7, 0x90, 0x2f, 0xc8, 0xf7, 0xe4, 0x0b, 0x72,
	0xf4, 0x31, 0xc7, 0x60, 0xff, 0x48, 0xb0, 0x14, 0xc9, 0x22, 0x16, 0x8e, 0xc9, 0x6d, 0x76, 0xde,
	0x9b, 0x9d, 0x37, 0x6f, 0x18, 0xf8, 0x1b, 0xfa, 0x4e, 0xc0, 0xe4, 0xcc, 0x0b, 0xa4, 0xcb, 0xad,
	0xb1, 0xc3, 0x2d, 0x99, 0x98, 0x61, 0xc4, 0x25, 0xc7, 0x3f, 0xcb, 0x90, 0x39, 0x76, 0xb8, 0x76,
	0xb0, 0x45, 0x5e, 0x87, 0x37, 0x11, 0x9b, 0xc7, 0x4c, 0xc8, 0xac, 0xcc, 0xb8, 0x43, 0xd0, 0xb6,
	0x85, 0x4b, 0x99, 0x27, 0x44, 0xcc, 0x68, 0xbf, 0x7f, 0x7e, 0x11, 0xf1, 0x90, 0x0b, 0xc7, 0xc7,
	0x2a, 0x7c, 0x1d, 0x45, 0xcc, 0x91, 0x3c, 0x52, 0x91, 0x8e, 0xba, 0x4d, 0x9a, 0x3f, 0xb1, 0x06,
	0xdf, 0xc2, 0x94, 0xc5, 0x22, 0x55, 0x49, 0xa1, 0xe2, 0x8d, 0x7f, 0x80, 0x22, 0x13, 0xb5, 0x9e,
	0x66, 0x15, 0x99, 0x60, 0x1d, 0x5a, 0x43, 0x9f, 0x8f, 0xa6, 0x13, 0xe6, 0xb9, 0x13, 0xa9, 0x36,
	0x74, 0xd4, 0x6d, 0xd0, 0x72, 0xca, 0xd0, 0x81, 0x54, 0x2b, 0xa0, 0x4c, 0x84, 0x3c, 0x10, 0xcc,
	0xf0, 0xe0, 0xbb, 0x2d, 0x5c, 0xdb, 0x0b, 0xe4, 0x15, 0x9f, 0xb2, 0x60, 0x87, 0xb2, 0x13, 0x68,
	0xad, 0x87, 0xa4, 0xd9, 0x8c, 0xa9, 0xb8, 0x56, 0xef, 0xbf, 0xf9, 0xde, 0x1b, 0xd3, 0xde, 0x90,
	0x68, 0xb9, 0xc2, 0x68, 0xc3, 0x9f, 0x72, 0xab, 0x42, 0xc2, 0x3d, 0x4a, 0x81, 0x92, 0x4a, 0xca,
	0x44, 0xec, 0xcb, 0x4f, 0xba, 0x84, 0xa1, 0x21, 0x93, 0xc1, 0xf8, 0xcd, 0xa7, 0x34, 0x2e, 0x9c,
	0x3a, 0xdb, 0x76, 0x2a, 0x4b, 0x19, 0x04, 0xfe, 0x55, 0x69, 0xc8, 0x45, 0xf6, 0x1e, 0x15, 0xa8,
	0xdb, 0xc2, 0xc5, 0x73, 0xf8, 0x5d, 0xb5, 0xd0, 0x6e, 0x85, 0x0f, 0x95, 0xc6, 0x6b, 0x47, 0xfb,
	0x32, 0xf3, 0xd6, 0xf8, 0x12, 0x9a, 0x9b, 0xfd, 0x90, 0xca, 0xf2, 0x02, 0xd7, 0x3a, 0xbb, 0xf1,
	0xe2, 0xd3, 0x29, 0xfc, 0xda, 0x36, 0xbc, 0xf3, 0x91, 0xb6, 0x8c, 0xa7, 0x99, 0xfb, 0xf1, 0xf2,
	0x66, 0xa7, 0x83, 0xa7, 0x25, 0x41, 0x8b, 0x25, 0x41, 0x2f, 0x4b, 0x82, 0x1e, 0x56, 0xa4, 0xb6,
	0x58, 0x91, 0xda, 0xf3, 0x8a, 0xd4, 0xae, 0x2d, 0xd7, 0x93, 0x93, 0x78, 0x68, 0x8e, 0xf8, 0xcc,
	0xda, 0xfc, 0x59, 0x0a, 0x0f, 0x5d, 0x6e, 0x25, 0xd9, 0x39, 0xde, 0x86, 0x4c, 0x0c, 0xbf, 0xa4,
	0xb7, 0x75, 0xfc, 0x1a, 0x00, 0x00, 0xff, 0xff, 0x42, 0x3f, 0x72, 0xd4, 0xaf, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	ReissueRDDLProposal(ctx context.Context, in *MsgReissueRDDLProposal, opts ...grpc.CallOption) (*MsgReissueRDDLProposalResponse, error)
	MintToken(ctx context.Context, in *MsgMintToken, opts ...grpc.CallOption) (*MsgMintTokenResponse, error)
	ReissueRDDLResult(ctx context.Context, in *MsgReissueRDDLResult, opts ...grpc.CallOption) (*MsgReissueRDDLResultResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ReissueRDDLProposal(ctx context.Context, in *MsgReissueRDDLProposal, opts ...grpc.CallOption) (*MsgReissueRDDLProposalResponse, error) {
	out := new(MsgReissueRDDLProposalResponse)
	err := c.cc.Invoke(ctx, "/planetmintgo.dao.Msg/ReissueRDDLProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MintToken(ctx context.Context, in *MsgMintToken, opts ...grpc.CallOption) (*MsgMintTokenResponse, error) {
	out := new(MsgMintTokenResponse)
	err := c.cc.Invoke(ctx, "/planetmintgo.dao.Msg/MintToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ReissueRDDLResult(ctx context.Context, in *MsgReissueRDDLResult, opts ...grpc.CallOption) (*MsgReissueRDDLResultResponse, error) {
	out := new(MsgReissueRDDLResultResponse)
	err := c.cc.Invoke(ctx, "/planetmintgo.dao.Msg/ReissueRDDLResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	ReissueRDDLProposal(context.Context, *MsgReissueRDDLProposal) (*MsgReissueRDDLProposalResponse, error)
	MintToken(context.Context, *MsgMintToken) (*MsgMintTokenResponse, error)
	ReissueRDDLResult(context.Context, *MsgReissueRDDLResult) (*MsgReissueRDDLResultResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ReissueRDDLProposal(ctx context.Context, req *MsgReissueRDDLProposal) (*MsgReissueRDDLProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReissueRDDLProposal not implemented")
}
func (*UnimplementedMsgServer) MintToken(ctx context.Context, req *MsgMintToken) (*MsgMintTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintToken not implemented")
}
func (*UnimplementedMsgServer) ReissueRDDLResult(ctx context.Context, req *MsgReissueRDDLResult) (*MsgReissueRDDLResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReissueRDDLResult not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ReissueRDDLProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgReissueRDDLProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ReissueRDDLProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/planetmintgo.dao.Msg/ReissueRDDLProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ReissueRDDLProposal(ctx, req.(*MsgReissueRDDLProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MintToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/planetmintgo.dao.Msg/MintToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintToken(ctx, req.(*MsgMintToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ReissueRDDLResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgReissueRDDLResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ReissueRDDLResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/planetmintgo.dao.Msg/ReissueRDDLResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ReissueRDDLResult(ctx, req.(*MsgReissueRDDLResult))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "planetmintgo.dao.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReissueRDDLProposal",
			Handler:    _Msg_ReissueRDDLProposal_Handler,
		},
		{
			MethodName: "MintToken",
			Handler:    _Msg_MintToken_Handler,
		},
		{
			MethodName: "ReissueRDDLResult",
			Handler:    _Msg_ReissueRDDLResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "planetmintgo/dao/tx.proto",
}

func (m *MsgReissueRDDLProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReissueRDDLProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReissueRDDLProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Blockheight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Blockheight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgReissueRDDLProposalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReissueRDDLProposalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReissueRDDLProposalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgMintToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MintRequest != nil {
		{
			size, err := m.MintRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMintTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgReissueRDDLResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReissueRDDLResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReissueRDDLResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgReissueRDDLResultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReissueRDDLResultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReissueRDDLResultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgReissueRDDLProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Blockheight != 0 {
		n += 1 + sovTx(uint64(m.Blockheight))
	}
	return n
}

func (m *MsgReissueRDDLProposalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgMintToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.MintRequest != nil {
		l = m.MintRequest.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgMintTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgReissueRDDLResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTx(uint64(m.BlockHeight))
	}
	return n
}

func (m *MsgReissueRDDLResultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgReissueRDDLProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgReissueRDDLProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReissueRDDLProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tx = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blockheight", wireType)
			}
			m.Blockheight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Blockheight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgReissueRDDLProposalResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgReissueRDDLProposalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReissueRDDLProposalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgMintToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMintToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MintRequest == nil {
				m.MintRequest = &MintRequest{}
			}
			if err := m.MintRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgMintTokenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMintTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgReissueRDDLResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgReissueRDDLResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReissueRDDLResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgReissueRDDLResultResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgReissueRDDLResultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReissueRDDLResultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
