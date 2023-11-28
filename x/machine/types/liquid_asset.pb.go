// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planetmintgo/machine/liquid_asset.proto

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

type LiquidAsset struct {
	MachineID      string `protobuf:"bytes,1,opt,name=machineID,proto3" json:"machineID,omitempty"`
	MachineAddress string `protobuf:"bytes,2,opt,name=machineAddress,proto3" json:"machineAddress,omitempty"`
	AssetID        string `protobuf:"bytes,3,opt,name=assetID,proto3" json:"assetID,omitempty"`
	Registered     bool   `protobuf:"varint,4,opt,name=registered,proto3" json:"registered,omitempty"`
}

func (m *LiquidAsset) Reset()         { *m = LiquidAsset{} }
func (m *LiquidAsset) String() string { return proto.CompactTextString(m) }
func (*LiquidAsset) ProtoMessage()    {}
func (*LiquidAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_fae3c910c6dc0f57, []int{0}
}
func (m *LiquidAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidAsset.Merge(m, src)
}
func (m *LiquidAsset) XXX_Size() int {
	return m.Size()
}
func (m *LiquidAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidAsset.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidAsset proto.InternalMessageInfo

func (m *LiquidAsset) GetMachineID() string {
	if m != nil {
		return m.MachineID
	}
	return ""
}

func (m *LiquidAsset) GetMachineAddress() string {
	if m != nil {
		return m.MachineAddress
	}
	return ""
}

func (m *LiquidAsset) GetAssetID() string {
	if m != nil {
		return m.AssetID
	}
	return ""
}

func (m *LiquidAsset) GetRegistered() bool {
	if m != nil {
		return m.Registered
	}
	return false
}

func init() {
	proto.RegisterType((*LiquidAsset)(nil), "planetmintgo.machine.LiquidAsset")
}

func init() {
	proto.RegisterFile("planetmintgo/machine/liquid_asset.proto", fileDescriptor_fae3c910c6dc0f57)
}

var fileDescriptor_fae3c910c6dc0f57 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0xc8, 0x49, 0xcc,
	0x4b, 0x2d, 0xc9, 0xcd, 0xcc, 0x2b, 0x49, 0xcf, 0xd7, 0xcf, 0x4d, 0x4c, 0xce, 0xc8, 0xcc, 0x4b,
	0xd5, 0xcf, 0xc9, 0x2c, 0x2c, 0xcd, 0x4c, 0x89, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x41, 0x56, 0xa8, 0x07, 0x55, 0xa8, 0xd4, 0xcb, 0xc8, 0xc5, 0xed,
	0x03, 0x56, 0xec, 0x08, 0x52, 0x2b, 0x24, 0xc3, 0xc5, 0x09, 0x95, 0xf2, 0x74, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x08, 0xa9, 0x71, 0xf1, 0x41, 0x39, 0x8e, 0x29, 0x29, 0x45,
	0xa9, 0xc5, 0xc5, 0x12, 0x4c, 0x60, 0x25, 0x68, 0xa2, 0x42, 0x12, 0x5c, 0xec, 0x60, 0xab, 0x3d,
	0x5d, 0x24, 0x98, 0xc1, 0x0a, 0x60, 0x5c, 0x21, 0x39, 0x2e, 0xae, 0xa2, 0xd4, 0xf4, 0xcc, 0xe2,
	0x92, 0xd4, 0xa2, 0xd4, 0x14, 0x09, 0x16, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x24, 0x11, 0x27, 0xdf,
	0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39,
	0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4e, 0xcf, 0x2c, 0xc9, 0x28,
	0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0x78, 0x05, 0x89, 0xa9, 0x9b, 0x9e, 0xaf, 0x5f, 0x01,
	0x0f, 0x81, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xdf, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xed, 0xa6, 0xca, 0xe0, 0x26, 0x01, 0x00, 0x00,
}

func (m *LiquidAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Registered {
		i--
		if m.Registered {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.AssetID) > 0 {
		i -= len(m.AssetID)
		copy(dAtA[i:], m.AssetID)
		i = encodeVarintLiquidAsset(dAtA, i, uint64(len(m.AssetID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MachineAddress) > 0 {
		i -= len(m.MachineAddress)
		copy(dAtA[i:], m.MachineAddress)
		i = encodeVarintLiquidAsset(dAtA, i, uint64(len(m.MachineAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MachineID) > 0 {
		i -= len(m.MachineID)
		copy(dAtA[i:], m.MachineID)
		i = encodeVarintLiquidAsset(dAtA, i, uint64(len(m.MachineID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLiquidAsset(dAtA []byte, offset int, v uint64) int {
	offset -= sovLiquidAsset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LiquidAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MachineID)
	if l > 0 {
		n += 1 + l + sovLiquidAsset(uint64(l))
	}
	l = len(m.MachineAddress)
	if l > 0 {
		n += 1 + l + sovLiquidAsset(uint64(l))
	}
	l = len(m.AssetID)
	if l > 0 {
		n += 1 + l + sovLiquidAsset(uint64(l))
	}
	if m.Registered {
		n += 2
	}
	return n
}

func sovLiquidAsset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLiquidAsset(x uint64) (n int) {
	return sovLiquidAsset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LiquidAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidAsset
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
			return fmt.Errorf("proto: LiquidAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MachineID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidAsset
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
				return ErrInvalidLengthLiquidAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MachineID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MachineAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidAsset
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
				return ErrInvalidLengthLiquidAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MachineAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidAsset
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
				return ErrInvalidLengthLiquidAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registered", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Registered = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidAsset
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
func skipLiquidAsset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiquidAsset
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
					return 0, ErrIntOverflowLiquidAsset
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
					return 0, ErrIntOverflowLiquidAsset
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
				return 0, ErrInvalidLengthLiquidAsset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLiquidAsset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLiquidAsset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLiquidAsset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiquidAsset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLiquidAsset = fmt.Errorf("proto: unexpected end of group")
)