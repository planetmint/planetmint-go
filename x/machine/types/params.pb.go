// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planetmintgo/machine/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters for the module.
type Params struct {
	AssetRegistryScheme     string `protobuf:"bytes,1,opt,name=asset_registry_scheme,json=assetRegistryScheme,proto3" json:"asset_registry_scheme,omitempty"`
	AssetRegistryDomain     string `protobuf:"bytes,2,opt,name=asset_registry_domain,json=assetRegistryDomain,proto3" json:"asset_registry_domain,omitempty"`
	AssetRegistryPath       string `protobuf:"bytes,3,opt,name=asset_registry_path,json=assetRegistryPath,proto3" json:"asset_registry_path,omitempty"`
	DaoMachineFundingAmount uint64 `protobuf:"varint,4,opt,name=dao_machine_funding_amount,json=daoMachineFundingAmount,proto3" json:"dao_machine_funding_amount,omitempty"`
	DaoMachineFundingDenom  string `protobuf:"bytes,5,opt,name=dao_machine_funding_denom,json=daoMachineFundingDenom,proto3" json:"dao_machine_funding_denom,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_84cd778d65e6639c, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAssetRegistryScheme() string {
	if m != nil {
		return m.AssetRegistryScheme
	}
	return ""
}

func (m *Params) GetAssetRegistryDomain() string {
	if m != nil {
		return m.AssetRegistryDomain
	}
	return ""
}

func (m *Params) GetAssetRegistryPath() string {
	if m != nil {
		return m.AssetRegistryPath
	}
	return ""
}

func (m *Params) GetDaoMachineFundingAmount() uint64 {
	if m != nil {
		return m.DaoMachineFundingAmount
	}
	return 0
}

func (m *Params) GetDaoMachineFundingDenom() string {
	if m != nil {
		return m.DaoMachineFundingDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "planetmintgo.machine.Params")
}

func init() { proto.RegisterFile("planetmintgo/machine/params.proto", fileDescriptor_84cd778d65e6639c) }

var fileDescriptor_84cd778d65e6639c = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x40, 0xe3, 0xbe, 0xbc, 0x4a, 0x78, 0x23, 0x2d, 0x10, 0x3a, 0x98, 0xc2, 0xd4, 0x85, 0x44,
	0xa2, 0x13, 0x30, 0x81, 0x2a, 0xb6, 0x4a, 0x55, 0xd8, 0x58, 0x22, 0x37, 0x31, 0x8e, 0x25, 0xec,
	0x1b, 0xc5, 0x8e, 0x44, 0xff, 0x82, 0x91, 0x09, 0xf1, 0x39, 0x8c, 0x1d, 0x19, 0x51, 0xf2, 0x23,
	0xa8, 0x4e, 0x04, 0x05, 0xb2, 0x5d, 0xe9, 0x9c, 0x23, 0xcb, 0xf7, 0xe2, 0xe3, 0xfc, 0x81, 0x2a,
	0x66, 0xa4, 0x50, 0x86, 0x43, 0x28, 0x69, 0x92, 0x09, 0xc5, 0xc2, 0x9c, 0x16, 0x54, 0xea, 0x20,
	0x2f, 0xc0, 0x80, 0x37, 0xdc, 0x56, 0x82, 0x56, 0x19, 0x0d, 0x39, 0x70, 0xb0, 0x42, 0xb8, 0x99,
	0x1a, 0xf7, 0xe4, 0xa5, 0x87, 0xfb, 0x0b, 0x1b, 0x7b, 0x67, 0x78, 0x8f, 0x6a, 0xcd, 0x4c, 0x5c,
	0x30, 0x2e, 0xb4, 0x29, 0x56, 0xb1, 0x4e, 0x32, 0x26, 0x99, 0x8f, 0xc6, 0x68, 0xb2, 0x13, 0x0d,
	0x2c, 0x8c, 0x5a, 0x76, 0x6b, 0x51, 0x47, 0x93, 0x82, 0xa4, 0x42, 0xf9, 0xbd, 0x8e, 0x66, 0x66,
	0x91, 0x17, 0xe0, 0xc1, 0xaf, 0x26, 0xa7, 0x26, 0xf3, 0xff, 0xd9, 0x62, 0xf7, 0x47, 0xb1, 0xa0,
	0x26, 0xf3, 0x2e, 0xf1, 0x28, 0xa5, 0x10, 0xb7, 0xff, 0x88, 0xef, 0x4b, 0x95, 0x0a, 0xc5, 0x63,
	0x2a, 0xa1, 0x54, 0xc6, 0x77, 0xc7, 0x68, 0xe2, 0x46, 0x07, 0x29, 0x85, 0x79, 0x23, 0xdc, 0x34,
	0xfc, 0xca, 0x62, 0xef, 0x1c, 0x1f, 0x76, 0xc5, 0x29, 0x53, 0x20, 0xfd, 0xff, 0xf6, 0xc9, 0xfd,
	0x3f, 0xed, 0x6c, 0x43, 0x2f, 0xdc, 0xe7, 0xd7, 0x23, 0xe7, 0x7a, 0xfe, 0x56, 0x11, 0xb4, 0xae,
	0x08, 0xfa, 0xa8, 0x08, 0x7a, 0xaa, 0x89, 0xb3, 0xae, 0x89, 0xf3, 0x5e, 0x13, 0xe7, 0x6e, 0xca,
	0x85, 0xc9, 0xca, 0x65, 0x90, 0x80, 0x0c, 0xbf, 0x37, 0xbe, 0x35, 0x9e, 0x72, 0x08, 0x1f, 0xbf,
	0x4e, 0x64, 0x56, 0x39, 0xd3, 0xcb, 0xbe, 0x5d, 0xfb, 0xf4, 0x33, 0x00, 0x00, 0xff, 0xff, 0xb0,
	0xd8, 0x7a, 0xbe, 0xc7, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DaoMachineFundingDenom) > 0 {
		i -= len(m.DaoMachineFundingDenom)
		copy(dAtA[i:], m.DaoMachineFundingDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DaoMachineFundingDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if m.DaoMachineFundingAmount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DaoMachineFundingAmount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.AssetRegistryPath) > 0 {
		i -= len(m.AssetRegistryPath)
		copy(dAtA[i:], m.AssetRegistryPath)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AssetRegistryPath)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AssetRegistryDomain) > 0 {
		i -= len(m.AssetRegistryDomain)
		copy(dAtA[i:], m.AssetRegistryDomain)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AssetRegistryDomain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AssetRegistryScheme) > 0 {
		i -= len(m.AssetRegistryScheme)
		copy(dAtA[i:], m.AssetRegistryScheme)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AssetRegistryScheme)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AssetRegistryScheme)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.AssetRegistryDomain)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.AssetRegistryPath)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.DaoMachineFundingAmount != 0 {
		n += 1 + sovParams(uint64(m.DaoMachineFundingAmount))
	}
	l = len(m.DaoMachineFundingDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetRegistryScheme", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetRegistryScheme = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetRegistryDomain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetRegistryDomain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetRegistryPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetRegistryPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaoMachineFundingAmount", wireType)
			}
			m.DaoMachineFundingAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DaoMachineFundingAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaoMachineFundingDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DaoMachineFundingDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
