// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planetmintgo/dao/params.proto

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
	MintAddress                  string `protobuf:"bytes,1,opt,name=mint_address,json=mintAddress,proto3" json:"mint_address,omitempty"`
	TokenDenom                   string `protobuf:"bytes,2,opt,name=token_denom,json=tokenDenom,proto3" json:"token_denom,omitempty"`
	FeeDenom                     string `protobuf:"bytes,3,opt,name=fee_denom,json=feeDenom,proto3" json:"fee_denom,omitempty"`
	StagedDenom                  string `protobuf:"bytes,4,opt,name=staged_denom,json=stagedDenom,proto3" json:"staged_denom,omitempty"`
	ClaimDenom                   string `protobuf:"bytes,5,opt,name=claim_denom,json=claimDenom,proto3" json:"claim_denom,omitempty"`
	ReissuanceAsset              string `protobuf:"bytes,6,opt,name=reissuance_asset,json=reissuanceAsset,proto3" json:"reissuance_asset,omitempty"`
	ReissuanceEpochs             int64  `protobuf:"varint,7,opt,name=reissuance_epochs,json=reissuanceEpochs,proto3" json:"reissuance_epochs,omitempty"`
	PopEpochs                    int64  `protobuf:"varint,8,opt,name=pop_epochs,json=popEpochs,proto3" json:"pop_epochs,omitempty"`
	DistributionOffset           int64  `protobuf:"varint,9,opt,name=distribution_offset,json=distributionOffset,proto3" json:"distribution_offset,omitempty"`
	DistributionAddressEarlyInv  string `protobuf:"bytes,10,opt,name=distribution_address_early_inv,json=distributionAddressEarlyInv,proto3" json:"distribution_address_early_inv,omitempty"`
	DistributionAddressInvestor  string `protobuf:"bytes,11,opt,name=distribution_address_investor,json=distributionAddressInvestor,proto3" json:"distribution_address_investor,omitempty"`
	DistributionAddressStrategic string `protobuf:"bytes,12,opt,name=distribution_address_strategic,json=distributionAddressStrategic,proto3" json:"distribution_address_strategic,omitempty"`
	DistributionAddressDao       string `protobuf:"bytes,13,opt,name=distribution_address_dao,json=distributionAddressDao,proto3" json:"distribution_address_dao,omitempty"`
	DistributionAddressPop       string `protobuf:"bytes,14,opt,name=distribution_address_pop,json=distributionAddressPop,proto3" json:"distribution_address_pop,omitempty"`
	MqttResponseTimeout          int64  `protobuf:"varint,15,opt,name=mqtt_response_timeout,json=mqttResponseTimeout,proto3" json:"mqtt_response_timeout,omitempty"`
	ClaimAddress                 string `protobuf:"bytes,16,opt,name=claim_address,json=claimAddress,proto3" json:"claim_address,omitempty"`
	TxGasLimit                   uint64 `protobuf:"varint,17,opt,name=tx_gas_limit,json=txGasLimit,proto3" json:"tx_gas_limit,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a58575036b3ad531, []int{0}
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

func (m *Params) GetMintAddress() string {
	if m != nil {
		return m.MintAddress
	}
	return ""
}

func (m *Params) GetTokenDenom() string {
	if m != nil {
		return m.TokenDenom
	}
	return ""
}

func (m *Params) GetFeeDenom() string {
	if m != nil {
		return m.FeeDenom
	}
	return ""
}

func (m *Params) GetStagedDenom() string {
	if m != nil {
		return m.StagedDenom
	}
	return ""
}

func (m *Params) GetClaimDenom() string {
	if m != nil {
		return m.ClaimDenom
	}
	return ""
}

func (m *Params) GetReissuanceAsset() string {
	if m != nil {
		return m.ReissuanceAsset
	}
	return ""
}

func (m *Params) GetReissuanceEpochs() int64 {
	if m != nil {
		return m.ReissuanceEpochs
	}
	return 0
}

func (m *Params) GetPopEpochs() int64 {
	if m != nil {
		return m.PopEpochs
	}
	return 0
}

func (m *Params) GetDistributionOffset() int64 {
	if m != nil {
		return m.DistributionOffset
	}
	return 0
}

func (m *Params) GetDistributionAddressEarlyInv() string {
	if m != nil {
		return m.DistributionAddressEarlyInv
	}
	return ""
}

func (m *Params) GetDistributionAddressInvestor() string {
	if m != nil {
		return m.DistributionAddressInvestor
	}
	return ""
}

func (m *Params) GetDistributionAddressStrategic() string {
	if m != nil {
		return m.DistributionAddressStrategic
	}
	return ""
}

func (m *Params) GetDistributionAddressDao() string {
	if m != nil {
		return m.DistributionAddressDao
	}
	return ""
}

func (m *Params) GetDistributionAddressPop() string {
	if m != nil {
		return m.DistributionAddressPop
	}
	return ""
}

func (m *Params) GetMqttResponseTimeout() int64 {
	if m != nil {
		return m.MqttResponseTimeout
	}
	return 0
}

func (m *Params) GetClaimAddress() string {
	if m != nil {
		return m.ClaimAddress
	}
	return ""
}

func (m *Params) GetTxGasLimit() uint64 {
	if m != nil {
		return m.TxGasLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "planetmintgo.dao.Params")
}

func init() { proto.RegisterFile("planetmintgo/dao/params.proto", fileDescriptor_a58575036b3ad531) }

var fileDescriptor_a58575036b3ad531 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x41, 0x6f, 0x12, 0x41,
	0x14, 0xc7, 0x59, 0x8b, 0x58, 0x1e, 0xd4, 0xd2, 0xa9, 0x9a, 0x89, 0x95, 0x85, 0xea, 0x05, 0x63,
	0x64, 0x13, 0xbd, 0x18, 0x6f, 0xad, 0x34, 0x86, 0xc4, 0xc4, 0x06, 0x3d, 0x79, 0xd9, 0x0c, 0xec,
	0xb0, 0x9d, 0xc8, 0xee, 0x1b, 0x77, 0x1e, 0x84, 0x7e, 0x0b, 0x8f, 0x7a, 0xf3, 0xe3, 0x78, 0xec,
	0xd1, 0xa3, 0x81, 0x2f, 0x62, 0x76, 0x66, 0x49, 0xd7, 0x04, 0x6e, 0x93, 0xff, 0xef, 0xf7, 0xde,
	0x9b, 0x4c, 0xe6, 0x41, 0x5b, 0xcf, 0x44, 0x2a, 0x29, 0x51, 0x29, 0xc5, 0x18, 0x44, 0x02, 0x03,
	0x2d, 0x32, 0x91, 0x98, 0xbe, 0xce, 0x90, 0x90, 0xb5, 0xca, 0xb8, 0x1f, 0x09, 0x7c, 0xfc, 0x20,
	0xc6, 0x18, 0x2d, 0x0c, 0xf2, 0x93, 0xf3, 0x9e, 0xfe, 0xac, 0x41, 0xed, 0xd2, 0x16, 0xb2, 0x53,
	0x68, 0xe6, 0x7a, 0x28, 0xa2, 0x28, 0x93, 0xc6, 0x70, 0xaf, 0xeb, 0xf5, 0xea, 0xa3, 0x46, 0x9e,
	0x9d, 0xb9, 0x88, 0x75, 0xa0, 0x41, 0xf8, 0x55, 0xa6, 0x61, 0x24, 0x53, 0x4c, 0xf8, 0x1d, 0x6b,
	0x80, 0x8d, 0x06, 0x79, 0xc2, 0x4e, 0xa0, 0x3e, 0x95, 0xb2, 0xc0, 0x7b, 0x16, 0xef, 0x4f, 0xa5,
	0x74, 0xf0, 0x14, 0x9a, 0x86, 0x44, 0x2c, 0xa3, 0x82, 0x57, 0xdd, 0x00, 0x97, 0x39, 0xa5, 0x03,
	0x8d, 0xc9, 0x4c, 0xa8, 0xa4, 0x30, 0xee, 0xba, 0x01, 0x36, 0x72, 0xc2, 0x73, 0x68, 0x65, 0x52,
	0x19, 0x33, 0x17, 0xe9, 0x44, 0x86, 0xc2, 0x18, 0x49, 0xbc, 0x66, 0xad, 0xc3, 0xdb, 0xfc, 0x2c,
	0x8f, 0xd9, 0x0b, 0x38, 0x2a, 0xa9, 0x52, 0xe3, 0xe4, 0xca, 0xf0, 0x7b, 0x5d, 0xaf, 0xb7, 0x37,
	0x2a, 0xf5, 0xb8, 0xb0, 0x39, 0x6b, 0x03, 0x68, 0xd4, 0x1b, 0x6b, 0xdf, 0x5a, 0x75, 0x8d, 0xba,
	0xc0, 0x01, 0x1c, 0x47, 0xca, 0x50, 0xa6, 0xc6, 0x73, 0x52, 0x98, 0x86, 0x38, 0x9d, 0xe6, 0x93,
	0xeb, 0xd6, 0x63, 0x65, 0xf4, 0xd1, 0x12, 0xf6, 0x0e, 0xfc, 0xff, 0x0a, 0x8a, 0x47, 0x0d, 0xa5,
	0xc8, 0x66, 0xd7, 0xa1, 0x4a, 0x17, 0x1c, 0xec, 0xad, 0x4f, 0xca, 0x56, 0xf1, 0xcc, 0x17, 0xb9,
	0x33, 0x4c, 0x17, 0xec, 0x1c, 0xda, 0x5b, 0x9b, 0xa8, 0x74, 0x21, 0x0d, 0x61, 0xc6, 0x1b, 0x3b,
	0x7b, 0x0c, 0x0b, 0x85, 0x0d, 0x76, 0x5c, 0xc4, 0x50, 0x26, 0x48, 0xc6, 0x6a, 0xc2, 0x9b, 0xb6,
	0xc9, 0x93, 0x2d, 0x4d, 0x3e, 0x6d, 0x1c, 0xf6, 0x06, 0xf8, 0xd6, 0x2e, 0x91, 0x40, 0x7e, 0x60,
	0xeb, 0x1f, 0x6d, 0xa9, 0x1f, 0x08, 0xdc, 0x59, 0xa9, 0x51, 0xf3, 0xfb, 0x3b, 0x2b, 0x2f, 0x51,
	0xb3, 0x57, 0xf0, 0x30, 0xf9, 0x46, 0x14, 0x66, 0xd2, 0x68, 0x4c, 0x8d, 0x0c, 0x49, 0x25, 0x12,
	0xe7, 0xc4, 0x0f, 0xed, 0xab, 0x1f, 0xe7, 0x70, 0x54, 0xb0, 0xcf, 0x0e, 0xb1, 0x67, 0x70, 0xe0,
	0xfe, 0xcf, 0xe6, 0x13, 0xb7, 0xec, 0x88, 0xa6, 0x0d, 0x37, 0xbf, 0xb8, 0x0b, 0x4d, 0x5a, 0x86,
	0xb1, 0x30, 0xe1, 0x4c, 0x25, 0x8a, 0xf8, 0x51, 0xd7, 0xeb, 0x55, 0x47, 0x40, 0xcb, 0xf7, 0xc2,
	0x7c, 0xc8, 0x93, 0xb7, 0xd5, 0x1f, 0xbf, 0x3a, 0x95, 0xf3, 0xe1, 0xef, 0x95, 0xef, 0xdd, 0xac,
	0x7c, 0xef, 0xef, 0xca, 0xf7, 0xbe, 0xaf, 0xfd, 0xca, 0xcd, 0xda, 0xaf, 0xfc, 0x59, 0xfb, 0x95,
	0x2f, 0x41, 0xac, 0xe8, 0x6a, 0x3e, 0xee, 0x4f, 0x30, 0x09, 0x6e, 0x17, 0xad, 0x74, 0x7c, 0x19,
	0x63, 0xb0, 0xb4, 0x5b, 0x49, 0xd7, 0x5a, 0x9a, 0x71, 0xcd, 0x6e, 0xdb, 0xeb, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x45, 0xb8, 0xbd, 0x25, 0xb6, 0x03, 0x00, 0x00,
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
	if m.TxGasLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TxGasLimit))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if len(m.ClaimAddress) > 0 {
		i -= len(m.ClaimAddress)
		copy(dAtA[i:], m.ClaimAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ClaimAddress)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if m.MqttResponseTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MqttResponseTimeout))
		i--
		dAtA[i] = 0x78
	}
	if len(m.DistributionAddressPop) > 0 {
		i -= len(m.DistributionAddressPop)
		copy(dAtA[i:], m.DistributionAddressPop)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DistributionAddressPop)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.DistributionAddressDao) > 0 {
		i -= len(m.DistributionAddressDao)
		copy(dAtA[i:], m.DistributionAddressDao)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DistributionAddressDao)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.DistributionAddressStrategic) > 0 {
		i -= len(m.DistributionAddressStrategic)
		copy(dAtA[i:], m.DistributionAddressStrategic)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DistributionAddressStrategic)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.DistributionAddressInvestor) > 0 {
		i -= len(m.DistributionAddressInvestor)
		copy(dAtA[i:], m.DistributionAddressInvestor)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DistributionAddressInvestor)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.DistributionAddressEarlyInv) > 0 {
		i -= len(m.DistributionAddressEarlyInv)
		copy(dAtA[i:], m.DistributionAddressEarlyInv)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DistributionAddressEarlyInv)))
		i--
		dAtA[i] = 0x52
	}
	if m.DistributionOffset != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DistributionOffset))
		i--
		dAtA[i] = 0x48
	}
	if m.PopEpochs != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PopEpochs))
		i--
		dAtA[i] = 0x40
	}
	if m.ReissuanceEpochs != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ReissuanceEpochs))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ReissuanceAsset) > 0 {
		i -= len(m.ReissuanceAsset)
		copy(dAtA[i:], m.ReissuanceAsset)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ReissuanceAsset)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ClaimDenom) > 0 {
		i -= len(m.ClaimDenom)
		copy(dAtA[i:], m.ClaimDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ClaimDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.StagedDenom) > 0 {
		i -= len(m.StagedDenom)
		copy(dAtA[i:], m.StagedDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.StagedDenom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FeeDenom) > 0 {
		i -= len(m.FeeDenom)
		copy(dAtA[i:], m.FeeDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.FeeDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenDenom) > 0 {
		i -= len(m.TokenDenom)
		copy(dAtA[i:], m.TokenDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TokenDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MintAddress) > 0 {
		i -= len(m.MintAddress)
		copy(dAtA[i:], m.MintAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MintAddress)))
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
	l = len(m.MintAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TokenDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.FeeDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.StagedDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ClaimDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ReissuanceAsset)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.ReissuanceEpochs != 0 {
		n += 1 + sovParams(uint64(m.ReissuanceEpochs))
	}
	if m.PopEpochs != 0 {
		n += 1 + sovParams(uint64(m.PopEpochs))
	}
	if m.DistributionOffset != 0 {
		n += 1 + sovParams(uint64(m.DistributionOffset))
	}
	l = len(m.DistributionAddressEarlyInv)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DistributionAddressInvestor)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DistributionAddressStrategic)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DistributionAddressDao)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DistributionAddressPop)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MqttResponseTimeout != 0 {
		n += 1 + sovParams(uint64(m.MqttResponseTimeout))
	}
	l = len(m.ClaimAddress)
	if l > 0 {
		n += 2 + l + sovParams(uint64(l))
	}
	if m.TxGasLimit != 0 {
		n += 2 + sovParams(uint64(m.TxGasLimit))
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
				return fmt.Errorf("proto: wrong wireType = %d for field MintAddress", wireType)
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
			m.MintAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenDenom", wireType)
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
			m.TokenDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeDenom", wireType)
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
			m.FeeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StagedDenom", wireType)
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
			m.StagedDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimDenom", wireType)
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
			m.ClaimDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReissuanceAsset", wireType)
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
			m.ReissuanceAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReissuanceEpochs", wireType)
			}
			m.ReissuanceEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReissuanceEpochs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PopEpochs", wireType)
			}
			m.PopEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PopEpochs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionOffset", wireType)
			}
			m.DistributionOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistributionOffset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionAddressEarlyInv", wireType)
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
			m.DistributionAddressEarlyInv = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionAddressInvestor", wireType)
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
			m.DistributionAddressInvestor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionAddressStrategic", wireType)
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
			m.DistributionAddressStrategic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionAddressDao", wireType)
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
			m.DistributionAddressDao = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionAddressPop", wireType)
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
			m.DistributionAddressPop = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MqttResponseTimeout", wireType)
			}
			m.MqttResponseTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MqttResponseTimeout |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimAddress", wireType)
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
			m.ClaimAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxGasLimit", wireType)
			}
			m.TxGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxGasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
