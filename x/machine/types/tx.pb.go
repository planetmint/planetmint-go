// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planetmintgo/machine/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgAttestMachine struct {
	Creator string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Machine *Machine `protobuf:"bytes,2,opt,name=machine,proto3" json:"machine,omitempty"`
}

func (m *MsgAttestMachine) Reset()         { *m = MsgAttestMachine{} }
func (m *MsgAttestMachine) String() string { return proto.CompactTextString(m) }
func (*MsgAttestMachine) ProtoMessage()    {}
func (*MsgAttestMachine) Descriptor() ([]byte, []int) {
	return fileDescriptor_85ac37e5c8e5251d, []int{0}
}
func (m *MsgAttestMachine) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAttestMachine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAttestMachine.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAttestMachine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAttestMachine.Merge(m, src)
}
func (m *MsgAttestMachine) XXX_Size() int {
	return m.Size()
}
func (m *MsgAttestMachine) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAttestMachine.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAttestMachine proto.InternalMessageInfo

func (m *MsgAttestMachine) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAttestMachine) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

type MsgAttestMachineResponse struct {
}

func (m *MsgAttestMachineResponse) Reset()         { *m = MsgAttestMachineResponse{} }
func (m *MsgAttestMachineResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAttestMachineResponse) ProtoMessage()    {}
func (*MsgAttestMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85ac37e5c8e5251d, []int{1}
}
func (m *MsgAttestMachineResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAttestMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAttestMachineResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAttestMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAttestMachineResponse.Merge(m, src)
}
func (m *MsgAttestMachineResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAttestMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAttestMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAttestMachineResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAttestMachine)(nil), "planetmintgo.machine.MsgAttestMachine")
	proto.RegisterType((*MsgAttestMachineResponse)(nil), "planetmintgo.machine.MsgAttestMachineResponse")
}

func init() { proto.RegisterFile("planetmintgo/machine/tx.proto", fileDescriptor_85ac37e5c8e5251d) }

var fileDescriptor_85ac37e5c8e5251d = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xc8, 0x49, 0xcc,
	0x4b, 0x2d, 0xc9, 0xcd, 0xcc, 0x2b, 0x49, 0xcf, 0xd7, 0xcf, 0x4d, 0x4c, 0xce, 0xc8, 0xcc, 0x4b,
	0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x41, 0x96, 0xd6, 0x83, 0x4a,
	0x4b, 0x29, 0x61, 0xd5, 0x04, 0xa5, 0x21, 0x3a, 0x95, 0x52, 0xb9, 0x04, 0x7c, 0x8b, 0xd3, 0x1d,
	0x4b, 0x4a, 0x52, 0x8b, 0x4b, 0x7c, 0x21, 0x32, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0x45, 0xa9, 0x89,
	0x25, 0xf9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x39, 0x17, 0x3b,
	0x54, 0xbb, 0x04, 0x93, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xac, 0x1e, 0x36, 0x9b, 0xf5, 0xa0, 0x26,
	0x05, 0xc1, 0x54, 0x2b, 0x49, 0x71, 0x49, 0xa0, 0x5b, 0x13, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57,
	0x9c, 0x6a, 0x94, 0xc7, 0xc5, 0xec, 0x5b, 0x9c, 0x2e, 0x94, 0xce, 0xc5, 0x8b, 0xea, 0x0c, 0x35,
	0x1c, 0x66, 0xa3, 0x99, 0x23, 0xa5, 0x47, 0x9c, 0x3a, 0x98, 0x7d, 0x4e, 0xe6, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x85, 0x14, 0xca, 0xba, 0xe9, 0xf9, 0xfa, 0x15, 0x88,
	0x80, 0xae, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x07, 0x99, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x78, 0x82, 0x2b, 0xac, 0x8d, 0x01, 0x00, 0x00,
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
	AttestMachine(ctx context.Context, in *MsgAttestMachine, opts ...grpc.CallOption) (*MsgAttestMachineResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AttestMachine(ctx context.Context, in *MsgAttestMachine, opts ...grpc.CallOption) (*MsgAttestMachineResponse, error) {
	out := new(MsgAttestMachineResponse)
	err := c.cc.Invoke(ctx, "/planetmintgo.machine.Msg/AttestMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AttestMachine(context.Context, *MsgAttestMachine) (*MsgAttestMachineResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AttestMachine(ctx context.Context, req *MsgAttestMachine) (*MsgAttestMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttestMachine not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AttestMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAttestMachine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AttestMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/planetmintgo.machine.Msg/AttestMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AttestMachine(ctx, req.(*MsgAttestMachine))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "planetmintgo.machine.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AttestMachine",
			Handler:    _Msg_AttestMachine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "planetmintgo/machine/tx.proto",
}

func (m *MsgAttestMachine) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAttestMachine) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAttestMachine) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *MsgAttestMachineResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAttestMachineResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAttestMachineResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgAttestMachine) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Machine != nil {
		l = m.Machine.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAttestMachineResponse) Size() (n int) {
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
func (m *MsgAttestMachine) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAttestMachine: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAttestMachine: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Machine", wireType)
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
			if m.Machine == nil {
				m.Machine = &Machine{}
			}
			if err := m.Machine.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgAttestMachineResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAttestMachineResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAttestMachineResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
