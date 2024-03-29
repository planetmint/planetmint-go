// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planetmintgo/dao/mint_request.proto

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

type MintRequest struct {
	Beneficiary  string `protobuf:"bytes,1,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	Amount       uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	LiquidTxHash string `protobuf:"bytes,3,opt,name=liquidTxHash,proto3" json:"liquidTxHash,omitempty"`
}

func (m *MintRequest) Reset()         { *m = MintRequest{} }
func (m *MintRequest) String() string { return proto.CompactTextString(m) }
func (*MintRequest) ProtoMessage()    {}
func (*MintRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc5c152bf52708dc, []int{0}
}
func (m *MintRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintRequest.Merge(m, src)
}
func (m *MintRequest) XXX_Size() int {
	return m.Size()
}
func (m *MintRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MintRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MintRequest proto.InternalMessageInfo

func (m *MintRequest) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *MintRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *MintRequest) GetLiquidTxHash() string {
	if m != nil {
		return m.LiquidTxHash
	}
	return ""
}

func init() {
	proto.RegisterType((*MintRequest)(nil), "planetmintgo.dao.MintRequest")
}

func init() {
	proto.RegisterFile("planetmintgo/dao/mint_request.proto", fileDescriptor_dc5c152bf52708dc)
}

var fileDescriptor_dc5c152bf52708dc = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0xc8, 0x49, 0xcc,
	0x4b, 0x2d, 0xc9, 0xcd, 0xcc, 0x2b, 0x49, 0xcf, 0xd7, 0x4f, 0x49, 0xcc, 0xd7, 0x07, 0x31, 0xe3,
	0x8b, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x90,
	0x15, 0xe9, 0xa5, 0x24, 0xe6, 0x2b, 0x65, 0x73, 0x71, 0xfb, 0x66, 0xe6, 0x95, 0x04, 0x41, 0x94,
	0x09, 0x29, 0x70, 0x71, 0x27, 0xa5, 0xe6, 0xa5, 0xa6, 0x65, 0x26, 0x67, 0x26, 0x16, 0x55, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x0b, 0x09, 0x89, 0x71, 0xb1, 0x25, 0xe6, 0xe6, 0x97,
	0xe6, 0x95, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x04, 0x41, 0x79, 0x42, 0x4a, 0x5c, 0x3c, 0x39,
	0x99, 0x85, 0xa5, 0x99, 0x29, 0x21, 0x15, 0x1e, 0x89, 0xc5, 0x19, 0x12, 0xcc, 0x60, 0xad, 0x28,
	0x62, 0x4e, 0x9e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9f, 0x9e,
	0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f, 0x70, 0x23, 0x12, 0x53, 0x37, 0x3d,
	0x5f, 0xbf, 0x02, 0xec, 0xad, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x87, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xca, 0x8a, 0xc6, 0x1d, 0xf7, 0x00, 0x00, 0x00,
}

func (m *MintRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LiquidTxHash) > 0 {
		i -= len(m.LiquidTxHash)
		copy(dAtA[i:], m.LiquidTxHash)
		i = encodeVarintMintRequest(dAtA, i, uint64(len(m.LiquidTxHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Amount != 0 {
		i = encodeVarintMintRequest(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Beneficiary) > 0 {
		i -= len(m.Beneficiary)
		copy(dAtA[i:], m.Beneficiary)
		i = encodeVarintMintRequest(dAtA, i, uint64(len(m.Beneficiary)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMintRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovMintRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MintRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Beneficiary)
	if l > 0 {
		n += 1 + l + sovMintRequest(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovMintRequest(uint64(m.Amount))
	}
	l = len(m.LiquidTxHash)
	if l > 0 {
		n += 1 + l + sovMintRequest(uint64(l))
	}
	return n
}

func sovMintRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMintRequest(x uint64) (n int) {
	return sovMintRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MintRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMintRequest
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
			return fmt.Errorf("proto: MintRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beneficiary", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
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
				return ErrInvalidLengthMintRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMintRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Beneficiary = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
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
				return ErrInvalidLengthMintRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMintRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMintRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMintRequest
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
func skipMintRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMintRequest
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
					return 0, ErrIntOverflowMintRequest
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
					return 0, ErrIntOverflowMintRequest
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
				return 0, ErrInvalidLengthMintRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMintRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMintRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMintRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMintRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMintRequest = fmt.Errorf("proto: unexpected end of group")
)
