// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planetmintgo/dao/reissuance.proto

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

type Reissuance struct {
	Proposer         string `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	RawTx            string `protobuf:"bytes,2,opt,name=rawTx,proto3" json:"rawTx,omitempty"`
	TxID             string `protobuf:"bytes,3,opt,name=txID,proto3" json:"txID,omitempty"`
	BlockHeight      int64  `protobuf:"varint,4,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	FirstIncludedPop int64  `protobuf:"varint,5,opt,name=firstIncludedPop,proto3" json:"firstIncludedPop,omitempty"`
	LastIncludedPop  int64  `protobuf:"varint,6,opt,name=lastIncludedPop,proto3" json:"lastIncludedPop,omitempty"`
}

func (m *Reissuance) Reset()         { *m = Reissuance{} }
func (m *Reissuance) String() string { return proto.CompactTextString(m) }
func (*Reissuance) ProtoMessage()    {}
func (*Reissuance) Descriptor() ([]byte, []int) {
	return fileDescriptor_35cf062bd4436e27, []int{0}
}
func (m *Reissuance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Reissuance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Reissuance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Reissuance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reissuance.Merge(m, src)
}
func (m *Reissuance) XXX_Size() int {
	return m.Size()
}
func (m *Reissuance) XXX_DiscardUnknown() {
	xxx_messageInfo_Reissuance.DiscardUnknown(m)
}

var xxx_messageInfo_Reissuance proto.InternalMessageInfo

func (m *Reissuance) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *Reissuance) GetRawTx() string {
	if m != nil {
		return m.RawTx
	}
	return ""
}

func (m *Reissuance) GetTxID() string {
	if m != nil {
		return m.TxID
	}
	return ""
}

func (m *Reissuance) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Reissuance) GetFirstIncludedPop() int64 {
	if m != nil {
		return m.FirstIncludedPop
	}
	return 0
}

func (m *Reissuance) GetLastIncludedPop() int64 {
	if m != nil {
		return m.LastIncludedPop
	}
	return 0
}

func init() {
	proto.RegisterType((*Reissuance)(nil), "planetmintgo.dao.Reissuance")
}

func init() { proto.RegisterFile("planetmintgo/dao/reissuance.proto", fileDescriptor_35cf062bd4436e27) }

var fileDescriptor_35cf062bd4436e27 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbf, 0x4a, 0xc5, 0x30,
	0x14, 0x87, 0x1b, 0xef, 0x1f, 0xf4, 0x38, 0x78, 0x09, 0x0e, 0xc1, 0x21, 0x54, 0xa7, 0x22, 0xd8,
	0x0c, 0xbe, 0x81, 0x38, 0xd8, 0x4d, 0x8a, 0x93, 0x5b, 0xda, 0xc6, 0xde, 0x60, 0x6f, 0x4f, 0x48,
	0x52, 0xac, 0x6f, 0xe1, 0x63, 0x89, 0xd3, 0x1d, 0x1d, 0xa5, 0x7d, 0x11, 0x21, 0xc2, 0xb5, 0x7a,
	0xb7, 0xf3, 0xfb, 0xce, 0x37, 0x7d, 0x70, 0x6e, 0x1a, 0xd9, 0x2a, 0xbf, 0xd1, 0xad, 0xaf, 0x51,
	0x54, 0x12, 0x85, 0x55, 0xda, 0xb9, 0x4e, 0xb6, 0xa5, 0x4a, 0x8d, 0x45, 0x8f, 0x74, 0x35, 0x55,
	0xd2, 0x4a, 0xe2, 0xc5, 0x07, 0x01, 0xc8, 0x77, 0x1a, 0x3d, 0x83, 0x43, 0x63, 0xd1, 0xa0, 0x53,
	0x96, 0x91, 0x98, 0x24, 0x47, 0xf9, 0x6e, 0xd3, 0x53, 0x58, 0x58, 0xf9, 0xf2, 0xd0, 0xb3, 0x83,
	0xf0, 0xf8, 0x19, 0x94, 0xc2, 0xdc, 0xf7, 0xd9, 0x2d, 0x9b, 0x05, 0x18, 0x6e, 0x1a, 0xc3, 0x71,
	0xd1, 0x60, 0xf9, 0x7c, 0xa7, 0x74, 0xbd, 0xf6, 0x6c, 0x1e, 0x93, 0x64, 0x96, 0x4f, 0x11, 0xbd,
	0x84, 0xd5, 0x93, 0xb6, 0xce, 0x67, 0x6d, 0xd9, 0x74, 0x95, 0xaa, 0xee, 0xd1, 0xb0, 0x45, 0xd0,
	0xf6, 0x38, 0x4d, 0xe0, 0xa4, 0x91, 0x7f, 0xd5, 0x65, 0x50, 0xff, 0xe3, 0x9b, 0xec, 0x7d, 0xe0,
	0x64, 0x3b, 0x70, 0xf2, 0x35, 0x70, 0xf2, 0x36, 0xf2, 0x68, 0x3b, 0xf2, 0xe8, 0x73, 0xe4, 0xd1,
	0xa3, 0xa8, 0xb5, 0x5f, 0x77, 0x45, 0x5a, 0xe2, 0x46, 0xfc, 0x36, 0x98, 0x9c, 0x57, 0x35, 0x8a,
	0x3e, 0x44, 0xf3, 0xaf, 0x46, 0xb9, 0x62, 0x19, 0x82, 0x5d, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xc5, 0x5a, 0x83, 0x1d, 0x55, 0x01, 0x00, 0x00,
}

func (m *Reissuance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reissuance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Reissuance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastIncludedPop != 0 {
		i = encodeVarintReissuance(dAtA, i, uint64(m.LastIncludedPop))
		i--
		dAtA[i] = 0x30
	}
	if m.FirstIncludedPop != 0 {
		i = encodeVarintReissuance(dAtA, i, uint64(m.FirstIncludedPop))
		i--
		dAtA[i] = 0x28
	}
	if m.BlockHeight != 0 {
		i = encodeVarintReissuance(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.TxID) > 0 {
		i -= len(m.TxID)
		copy(dAtA[i:], m.TxID)
		i = encodeVarintReissuance(dAtA, i, uint64(len(m.TxID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RawTx) > 0 {
		i -= len(m.RawTx)
		copy(dAtA[i:], m.RawTx)
		i = encodeVarintReissuance(dAtA, i, uint64(len(m.RawTx)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintReissuance(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintReissuance(dAtA []byte, offset int, v uint64) int {
	offset -= sovReissuance(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Reissuance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovReissuance(uint64(l))
	}
	l = len(m.RawTx)
	if l > 0 {
		n += 1 + l + sovReissuance(uint64(l))
	}
	l = len(m.TxID)
	if l > 0 {
		n += 1 + l + sovReissuance(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovReissuance(uint64(m.BlockHeight))
	}
	if m.FirstIncludedPop != 0 {
		n += 1 + sovReissuance(uint64(m.FirstIncludedPop))
	}
	if m.LastIncludedPop != 0 {
		n += 1 + sovReissuance(uint64(m.LastIncludedPop))
	}
	return n
}

func sovReissuance(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReissuance(x uint64) (n int) {
	return sovReissuance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Reissuance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReissuance
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
			return fmt.Errorf("proto: Reissuance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reissuance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReissuance
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
				return ErrInvalidLengthReissuance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReissuance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawTx", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReissuance
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
				return ErrInvalidLengthReissuance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReissuance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawTx = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReissuance
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
				return ErrInvalidLengthReissuance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReissuance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReissuance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstIncludedPop", wireType)
			}
			m.FirstIncludedPop = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReissuance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FirstIncludedPop |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIncludedPop", wireType)
			}
			m.LastIncludedPop = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReissuance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastIncludedPop |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReissuance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReissuance
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
func skipReissuance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReissuance
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
					return 0, ErrIntOverflowReissuance
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
					return 0, ErrIntOverflowReissuance
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
				return 0, ErrInvalidLengthReissuance
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReissuance
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReissuance
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReissuance        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReissuance          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReissuance = fmt.Errorf("proto: unexpected end of group")
)
