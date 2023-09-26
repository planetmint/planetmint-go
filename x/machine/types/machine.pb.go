// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planetmintgo/machine/machine.proto

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

type Machine struct {
	Name               string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ticker             string    `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Domain             string    `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Reissue            bool      `protobuf:"varint,4,opt,name=reissue,proto3" json:"reissue,omitempty"`
	Amount             uint64    `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Precision          uint64    `protobuf:"varint,6,opt,name=precision,proto3" json:"precision,omitempty"`
	IssuerPlanetmint   string    `protobuf:"bytes,7,opt,name=issuerPlanetmint,proto3" json:"issuerPlanetmint,omitempty"`
	IssuerLiquid       string    `protobuf:"bytes,8,opt,name=issuerLiquid,proto3" json:"issuerLiquid,omitempty"`
	MachineId          string    `protobuf:"bytes,9,opt,name=machineId,proto3" json:"machineId,omitempty"`
	Metadata           *Metadata `protobuf:"bytes,10,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Type               uint32    `protobuf:"varint,11,opt,name=type,proto3" json:"type,omitempty"`
	MachineIdSignature string    `protobuf:"bytes,12,opt,name=machineIdSignature,proto3" json:"machineIdSignature,omitempty"`
	Address            string    `protobuf:"bytes,13,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Machine) Reset()         { *m = Machine{} }
func (m *Machine) String() string { return proto.CompactTextString(m) }
func (*Machine) ProtoMessage()    {}
func (*Machine) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bb279745bef7c4b, []int{0}
}
func (m *Machine) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Machine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Machine.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Machine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Machine.Merge(m, src)
}
func (m *Machine) XXX_Size() int {
	return m.Size()
}
func (m *Machine) XXX_DiscardUnknown() {
	xxx_messageInfo_Machine.DiscardUnknown(m)
}

var xxx_messageInfo_Machine proto.InternalMessageInfo

func (m *Machine) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Machine) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *Machine) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Machine) GetReissue() bool {
	if m != nil {
		return m.Reissue
	}
	return false
}

func (m *Machine) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Machine) GetPrecision() uint64 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *Machine) GetIssuerPlanetmint() string {
	if m != nil {
		return m.IssuerPlanetmint
	}
	return ""
}

func (m *Machine) GetIssuerLiquid() string {
	if m != nil {
		return m.IssuerLiquid
	}
	return ""
}

func (m *Machine) GetMachineId() string {
	if m != nil {
		return m.MachineId
	}
	return ""
}

func (m *Machine) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Machine) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Machine) GetMachineIdSignature() string {
	if m != nil {
		return m.MachineIdSignature
	}
	return ""
}

func (m *Machine) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Metadata struct {
	Gps               string `protobuf:"bytes,1,opt,name=gps,proto3" json:"gps,omitempty"`
	Device            string `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	AssetDefinition   string `protobuf:"bytes,3,opt,name=assetDefinition,proto3" json:"assetDefinition,omitempty"`
	AdditionalDataCID string `protobuf:"bytes,4,opt,name=additionalDataCID,proto3" json:"additionalDataCID,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bb279745bef7c4b, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetGps() string {
	if m != nil {
		return m.Gps
	}
	return ""
}

func (m *Metadata) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *Metadata) GetAssetDefinition() string {
	if m != nil {
		return m.AssetDefinition
	}
	return ""
}

func (m *Metadata) GetAdditionalDataCID() string {
	if m != nil {
		return m.AdditionalDataCID
	}
	return ""
}

type MachineIndex struct {
	MachineId        string `protobuf:"bytes,1,opt,name=machineId,proto3" json:"machineId,omitempty"`
	IssuerPlanetmint string `protobuf:"bytes,2,opt,name=issuerPlanetmint,proto3" json:"issuerPlanetmint,omitempty"`
	IssuerLiquid     string `protobuf:"bytes,3,opt,name=issuerLiquid,proto3" json:"issuerLiquid,omitempty"`
	Address          string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MachineIndex) Reset()         { *m = MachineIndex{} }
func (m *MachineIndex) String() string { return proto.CompactTextString(m) }
func (*MachineIndex) ProtoMessage()    {}
func (*MachineIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bb279745bef7c4b, []int{2}
}
func (m *MachineIndex) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MachineIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MachineIndex.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MachineIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineIndex.Merge(m, src)
}
func (m *MachineIndex) XXX_Size() int {
	return m.Size()
}
func (m *MachineIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineIndex.DiscardUnknown(m)
}

var xxx_messageInfo_MachineIndex proto.InternalMessageInfo

func (m *MachineIndex) GetMachineId() string {
	if m != nil {
		return m.MachineId
	}
	return ""
}

func (m *MachineIndex) GetIssuerPlanetmint() string {
	if m != nil {
		return m.IssuerPlanetmint
	}
	return ""
}

func (m *MachineIndex) GetIssuerLiquid() string {
	if m != nil {
		return m.IssuerLiquid
	}
	return ""
}

func (m *MachineIndex) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Machine)(nil), "planetmintgo.machine.Machine")
	proto.RegisterType((*Metadata)(nil), "planetmintgo.machine.Metadata")
	proto.RegisterType((*MachineIndex)(nil), "planetmintgo.machine.MachineIndex")
}

func init() {
	proto.RegisterFile("planetmintgo/machine/machine.proto", fileDescriptor_1bb279745bef7c4b)
}

var fileDescriptor_1bb279745bef7c4b = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xbd, 0x8e, 0xd3, 0x40,
	0x18, 0xcc, 0x5e, 0x42, 0x7e, 0xbe, 0xcb, 0x89, 0x63, 0x85, 0xd0, 0x16, 0xc8, 0xb2, 0x5c, 0x59,
	0x08, 0x1c, 0x89, 0xeb, 0x28, 0x21, 0x4d, 0x24, 0x22, 0x21, 0xd3, 0xd1, 0xed, 0x79, 0x3f, 0x7c,
	0x2b, 0xce, 0xbb, 0xc6, 0xbb, 0x46, 0xc7, 0x3b, 0x50, 0x50, 0xf1, 0x06, 0xbc, 0x0b, 0xe5, 0x95,
	0x94, 0x28, 0x79, 0x11, 0xb4, 0x6b, 0x3b, 0x3f, 0x77, 0x2e, 0xae, 0xca, 0xcc, 0x7c, 0xb3, 0xf9,
	0xec, 0x19, 0x2f, 0x44, 0xe5, 0x35, 0x57, 0x68, 0x0b, 0xa9, 0x6c, 0xae, 0x17, 0x05, 0xcf, 0xae,
	0xa4, 0xc2, 0xee, 0x37, 0x29, 0x2b, 0x6d, 0x35, 0x7d, 0x7a, 0xe8, 0x49, 0xda, 0x59, 0xf4, 0x7b,
	0x08, 0x93, 0x75, 0x83, 0x29, 0x85, 0x91, 0xe2, 0x05, 0x32, 0x12, 0x92, 0x78, 0x96, 0x7a, 0x4c,
	0x9f, 0xc1, 0xd8, 0xca, 0xec, 0x0b, 0x56, 0xec, 0xc4, 0xab, 0x2d, 0x73, 0xba, 0xd0, 0x05, 0x97,
	0x8a, 0x0d, 0x1b, 0xbd, 0x61, 0x94, 0xc1, 0xa4, 0x42, 0x69, 0x4c, 0x8d, 0x6c, 0x14, 0x92, 0x78,
	0x9a, 0x76, 0xd4, 0x9d, 0xe0, 0x85, 0xae, 0x95, 0x65, 0x8f, 0x42, 0x12, 0x8f, 0xd2, 0x96, 0xd1,
	0xe7, 0x30, 0x2b, 0x2b, 0xcc, 0xa4, 0x91, 0x5a, 0xb1, 0xb1, 0x1f, 0xed, 0x05, 0xfa, 0x02, 0xce,
	0xfd, 0xf1, 0xea, 0xc3, 0xee, 0xe9, 0xd9, 0xc4, 0x6f, 0xbc, 0xa7, 0xd3, 0x08, 0xe6, 0x8d, 0xf6,
	0x5e, 0x7e, 0xad, 0xa5, 0x60, 0x53, 0xef, 0x3b, 0xd2, 0xdc, 0xb6, 0xf6, 0xd5, 0x57, 0x82, 0xcd,
	0xbc, 0x61, 0x2f, 0xd0, 0x37, 0x30, 0x2d, 0xd0, 0x72, 0xc1, 0x2d, 0x67, 0x10, 0x92, 0xf8, 0xf4,
	0x75, 0x90, 0xf4, 0xc5, 0x96, 0xac, 0x5b, 0x57, 0xba, 0xf3, 0xbb, 0xf4, 0xec, 0xf7, 0x12, 0xd9,
	0x69, 0x48, 0xe2, 0xb3, 0xd4, 0x63, 0x9a, 0x00, 0xdd, 0xfd, 0xf9, 0x47, 0x99, 0x2b, 0x6e, 0xeb,
	0x0a, 0xd9, 0xdc, 0xaf, 0xed, 0x99, 0xb8, 0xf4, 0xb8, 0x10, 0x15, 0x1a, 0xc3, 0xce, 0xbc, 0xa9,
	0xa3, 0xd1, 0x0f, 0x02, 0xd3, 0x6e, 0x29, 0x3d, 0x87, 0x61, 0x5e, 0x9a, 0xb6, 0x27, 0x07, 0x7d,
	0x1d, 0xf8, 0x4d, 0x66, 0xd8, 0xd5, 0xd4, 0x30, 0x1a, 0xc3, 0x63, 0x6e, 0x0c, 0xda, 0x25, 0x7e,
	0x96, 0x4a, 0x5a, 0x17, 0x71, 0xd3, 0xd7, 0x5d, 0x99, 0xbe, 0x84, 0x27, 0x5c, 0x08, 0x8f, 0xf9,
	0xf5, 0x92, 0x5b, 0xfe, 0x6e, 0xb5, 0xf4, 0x15, 0xce, 0xd2, 0xfb, 0x83, 0xe8, 0x17, 0x81, 0x79,
	0xfb, 0xd9, 0xac, 0x94, 0xc0, 0x9b, 0xe3, 0x5c, 0xc9, 0xdd, 0x5c, 0xfb, 0x5a, 0x3c, 0x79, 0x60,
	0x8b, 0xc3, 0x9e, 0x16, 0x0f, 0x72, 0x1a, 0x1d, 0xe5, 0xf4, 0x76, 0xfd, 0x67, 0x13, 0x90, 0xdb,
	0x4d, 0x40, 0xfe, 0x6d, 0x02, 0xf2, 0x73, 0x1b, 0x0c, 0x6e, 0xb7, 0xc1, 0xe0, 0xef, 0x36, 0x18,
	0x7c, 0xba, 0xc8, 0xa5, 0xbd, 0xaa, 0x2f, 0x93, 0x4c, 0x17, 0x8b, 0x7d, 0xa7, 0x07, 0xf0, 0x55,
	0xae, 0x17, 0x37, 0xbb, 0xcb, 0xe3, 0xfa, 0x33, 0x97, 0x63, 0x7f, 0x77, 0x2e, 0xfe, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x44, 0x7f, 0xea, 0x9d, 0x61, 0x03, 0x00, 0x00,
}

func (m *Machine) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Machine) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Machine) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.MachineIdSignature) > 0 {
		i -= len(m.MachineIdSignature)
		copy(dAtA[i:], m.MachineIdSignature)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.MachineIdSignature)))
		i--
		dAtA[i] = 0x62
	}
	if m.Type != 0 {
		i = encodeVarintMachine(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x58
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMachine(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if len(m.MachineId) > 0 {
		i -= len(m.MachineId)
		copy(dAtA[i:], m.MachineId)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.MachineId)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.IssuerLiquid) > 0 {
		i -= len(m.IssuerLiquid)
		copy(dAtA[i:], m.IssuerLiquid)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.IssuerLiquid)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.IssuerPlanetmint) > 0 {
		i -= len(m.IssuerPlanetmint)
		copy(dAtA[i:], m.IssuerPlanetmint)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.IssuerPlanetmint)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Precision != 0 {
		i = encodeVarintMachine(dAtA, i, uint64(m.Precision))
		i--
		dAtA[i] = 0x30
	}
	if m.Amount != 0 {
		i = encodeVarintMachine(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x28
	}
	if m.Reissue {
		i--
		if m.Reissue {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Ticker) > 0 {
		i -= len(m.Ticker)
		copy(dAtA[i:], m.Ticker)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.Ticker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AdditionalDataCID) > 0 {
		i -= len(m.AdditionalDataCID)
		copy(dAtA[i:], m.AdditionalDataCID)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.AdditionalDataCID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AssetDefinition) > 0 {
		i -= len(m.AssetDefinition)
		copy(dAtA[i:], m.AssetDefinition)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.AssetDefinition)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Device) > 0 {
		i -= len(m.Device)
		copy(dAtA[i:], m.Device)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.Device)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Gps) > 0 {
		i -= len(m.Gps)
		copy(dAtA[i:], m.Gps)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.Gps)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MachineIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MachineIndex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MachineIndex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.IssuerLiquid) > 0 {
		i -= len(m.IssuerLiquid)
		copy(dAtA[i:], m.IssuerLiquid)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.IssuerLiquid)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.IssuerPlanetmint) > 0 {
		i -= len(m.IssuerPlanetmint)
		copy(dAtA[i:], m.IssuerPlanetmint)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.IssuerPlanetmint)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MachineId) > 0 {
		i -= len(m.MachineId)
		copy(dAtA[i:], m.MachineId)
		i = encodeVarintMachine(dAtA, i, uint64(len(m.MachineId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMachine(dAtA []byte, offset int, v uint64) int {
	offset -= sovMachine(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Machine) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	l = len(m.Ticker)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	if m.Reissue {
		n += 2
	}
	if m.Amount != 0 {
		n += 1 + sovMachine(uint64(m.Amount))
	}
	if m.Precision != 0 {
		n += 1 + sovMachine(uint64(m.Precision))
	}
	l = len(m.IssuerPlanetmint)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	l = len(m.IssuerLiquid)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	l = len(m.MachineId)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovMachine(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovMachine(uint64(m.Type))
	}
	l = len(m.MachineIdSignature)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	return n
}

func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Gps)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	l = len(m.AssetDefinition)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	l = len(m.AdditionalDataCID)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	return n
}

func (m *MachineIndex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MachineId)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	l = len(m.IssuerPlanetmint)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	l = len(m.IssuerLiquid)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMachine(uint64(l))
	}
	return n
}

func sovMachine(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMachine(x uint64) (n int) {
	return sovMachine(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Machine) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMachine
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
			return fmt.Errorf("proto: Machine: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Machine: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ticker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reissue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
			m.Reissue = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
			}
			m.Precision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Precision |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerPlanetmint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuerPlanetmint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerLiquid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuerLiquid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MachineId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MachineId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MachineIdSignature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MachineIdSignature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMachine(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMachine
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
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMachine
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gps", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gps = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetDefinition", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetDefinition = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalDataCID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalDataCID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMachine(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMachine
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
func (m *MachineIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMachine
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
			return fmt.Errorf("proto: MachineIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MachineIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MachineId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MachineId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerPlanetmint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuerPlanetmint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerLiquid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuerLiquid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMachine
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
				return ErrInvalidLengthMachine
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMachine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMachine(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMachine
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
func skipMachine(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMachine
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
					return 0, ErrIntOverflowMachine
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
					return 0, ErrIntOverflowMachine
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
				return 0, ErrInvalidLengthMachine
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMachine
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMachine
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMachine        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMachine          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMachine = fmt.Errorf("proto: unexpected end of group")
)
