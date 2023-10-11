// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planetmintgo/dao/po_p_distribution.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type PoPDistribution struct {
	RddlAmount string `protobuf:"bytes,1,opt,name=rddlAmount,proto3" json:"rddlAmount,omitempty"`
	FirstPop   uint64 `protobuf:"varint,2,opt,name=firstPop,proto3" json:"firstPop,omitempty"`
	LastPop    uint64 `protobuf:"varint,3,opt,name=lastPop,proto3" json:"lastPop,omitempty"`
}

func (m *PoPDistribution) Reset()         { *m = PoPDistribution{} }
func (m *PoPDistribution) String() string { return proto.CompactTextString(m) }
func (*PoPDistribution) ProtoMessage()    {}
func (*PoPDistribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_456968989719b0d9, []int{0}
}
func (m *PoPDistribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoPDistribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoPDistribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoPDistribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoPDistribution.Merge(m, src)
}
func (m *PoPDistribution) XXX_Size() int {
	return m.Size()
}
func (m *PoPDistribution) XXX_DiscardUnknown() {
	xxx_messageInfo_PoPDistribution.DiscardUnknown(m)
}

var xxx_messageInfo_PoPDistribution proto.InternalMessageInfo

func (m *PoPDistribution) GetRddlAmount() string {
	if m != nil {
		return m.RddlAmount
	}
	return ""
}

func (m *PoPDistribution) GetFirstPop() uint64 {
	if m != nil {
		return m.FirstPop
	}
	return 0
}

func (m *PoPDistribution) GetLastPop() uint64 {
	if m != nil {
		return m.LastPop
	}
	return 0
}

func init() {
	proto.RegisterType((*PoPDistribution)(nil), "planetmintgo.dao.PoPDistribution")
}

func init() {
	proto.RegisterFile("planetmintgo/dao/po_p_distribution.proto", fileDescriptor_456968989719b0d9)
}

var fileDescriptor_456968989719b0d9 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0xc8, 0x49, 0xcc,
	0x4b, 0x2d, 0xc9, 0xcd, 0xcc, 0x2b, 0x49, 0xcf, 0xd7, 0x4f, 0x49, 0xcc, 0xd7, 0x2f, 0xc8, 0x8f,
	0x2f, 0x88, 0x4f, 0xc9, 0x2c, 0x2e, 0x29, 0xca, 0x4c, 0x2a, 0x2d, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x40, 0x56, 0xa9, 0x97, 0x92, 0x98, 0xaf, 0x94, 0xce, 0xc5,
	0x1f, 0x90, 0x1f, 0xe0, 0x82, 0xa4, 0x54, 0x48, 0x8e, 0x8b, 0xab, 0x28, 0x25, 0x25, 0xc7, 0x31,
	0x37, 0xbf, 0x34, 0xaf, 0x44, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0x49, 0x44, 0x48, 0x8a,
	0x8b, 0x23, 0x2d, 0xb3, 0xa8, 0xb8, 0x24, 0x20, 0xbf, 0x40, 0x82, 0x49, 0x81, 0x51, 0x83, 0x25,
	0x08, 0xce, 0x17, 0x92, 0xe0, 0x62, 0xcf, 0x49, 0x84, 0x48, 0x31, 0x83, 0xa5, 0x60, 0x5c, 0x27,
	0xcf, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0xb8, 0x0f, 0x89, 0xa9, 0x9b, 0x9e, 0xaf, 0x5f,
	0x01, 0xf6, 0x57, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x33, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf1, 0x98, 0x52, 0xec, 0xf8, 0x00, 0x00, 0x00,
}

func (m *PoPDistribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoPDistribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoPDistribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastPop != 0 {
		i = encodeVarintPoPDistribution(dAtA, i, uint64(m.LastPop))
		i--
		dAtA[i] = 0x18
	}
	if m.FirstPop != 0 {
		i = encodeVarintPoPDistribution(dAtA, i, uint64(m.FirstPop))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RddlAmount) > 0 {
		i -= len(m.RddlAmount)
		copy(dAtA[i:], m.RddlAmount)
		i = encodeVarintPoPDistribution(dAtA, i, uint64(len(m.RddlAmount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoPDistribution(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoPDistribution(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoPDistribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RddlAmount)
	if l > 0 {
		n += 1 + l + sovPoPDistribution(uint64(l))
	}
	if m.FirstPop != 0 {
		n += 1 + sovPoPDistribution(uint64(m.FirstPop))
	}
	if m.LastPop != 0 {
		n += 1 + sovPoPDistribution(uint64(m.LastPop))
	}
	return n
}

func sovPoPDistribution(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoPDistribution(x uint64) (n int) {
	return sovPoPDistribution(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoPDistribution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoPDistribution
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
			return fmt.Errorf("proto: PoPDistribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoPDistribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RddlAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoPDistribution
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
				return ErrInvalidLengthPoPDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoPDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RddlAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstPop", wireType)
			}
			m.FirstPop = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoPDistribution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FirstPop |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastPop", wireType)
			}
			m.LastPop = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoPDistribution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastPop |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPoPDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoPDistribution
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
func skipPoPDistribution(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoPDistribution
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
					return 0, ErrIntOverflowPoPDistribution
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
					return 0, ErrIntOverflowPoPDistribution
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
				return 0, ErrInvalidLengthPoPDistribution
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoPDistribution
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoPDistribution
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoPDistribution        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoPDistribution          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoPDistribution = fmt.Errorf("proto: unexpected end of group")
)
