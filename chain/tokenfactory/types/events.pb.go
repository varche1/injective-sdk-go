// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/tokenfactory/v1beta1/events.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/x/bank/types"
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

type EventCreateTFDenom struct {
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Denom   string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *EventCreateTFDenom) Reset()         { *m = EventCreateTFDenom{} }
func (m *EventCreateTFDenom) String() string { return proto.CompactTextString(m) }
func (*EventCreateTFDenom) ProtoMessage()    {}
func (*EventCreateTFDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9fd0c5434c2a5b7, []int{0}
}
func (m *EventCreateTFDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateTFDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateTFDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateTFDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateTFDenom.Merge(m, src)
}
func (m *EventCreateTFDenom) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateTFDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateTFDenom.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateTFDenom proto.InternalMessageInfo

func (m *EventCreateTFDenom) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *EventCreateTFDenom) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type EventMintTFDenom struct {
	RecipientAddress string     `protobuf:"bytes,1,opt,name=recipient_address,json=recipientAddress,proto3" json:"recipient_address,omitempty"`
	Amount           types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *EventMintTFDenom) Reset()         { *m = EventMintTFDenom{} }
func (m *EventMintTFDenom) String() string { return proto.CompactTextString(m) }
func (*EventMintTFDenom) ProtoMessage()    {}
func (*EventMintTFDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9fd0c5434c2a5b7, []int{1}
}
func (m *EventMintTFDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMintTFDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMintTFDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMintTFDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMintTFDenom.Merge(m, src)
}
func (m *EventMintTFDenom) XXX_Size() int {
	return m.Size()
}
func (m *EventMintTFDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMintTFDenom.DiscardUnknown(m)
}

var xxx_messageInfo_EventMintTFDenom proto.InternalMessageInfo

func (m *EventMintTFDenom) GetRecipientAddress() string {
	if m != nil {
		return m.RecipientAddress
	}
	return ""
}

func (m *EventMintTFDenom) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type EventBurnTFDenom struct {
	BurnerAddress string     `protobuf:"bytes,1,opt,name=burner_address,json=burnerAddress,proto3" json:"burner_address,omitempty"`
	Amount        types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *EventBurnTFDenom) Reset()         { *m = EventBurnTFDenom{} }
func (m *EventBurnTFDenom) String() string { return proto.CompactTextString(m) }
func (*EventBurnTFDenom) ProtoMessage()    {}
func (*EventBurnTFDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9fd0c5434c2a5b7, []int{2}
}
func (m *EventBurnTFDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBurnTFDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBurnTFDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBurnTFDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBurnTFDenom.Merge(m, src)
}
func (m *EventBurnTFDenom) XXX_Size() int {
	return m.Size()
}
func (m *EventBurnTFDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBurnTFDenom.DiscardUnknown(m)
}

var xxx_messageInfo_EventBurnTFDenom proto.InternalMessageInfo

func (m *EventBurnTFDenom) GetBurnerAddress() string {
	if m != nil {
		return m.BurnerAddress
	}
	return ""
}

func (m *EventBurnTFDenom) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type EventChangeTFAdmin struct {
	Denom           string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	NewAdminAddress string `protobuf:"bytes,2,opt,name=new_admin_address,json=newAdminAddress,proto3" json:"new_admin_address,omitempty"`
}

func (m *EventChangeTFAdmin) Reset()         { *m = EventChangeTFAdmin{} }
func (m *EventChangeTFAdmin) String() string { return proto.CompactTextString(m) }
func (*EventChangeTFAdmin) ProtoMessage()    {}
func (*EventChangeTFAdmin) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9fd0c5434c2a5b7, []int{3}
}
func (m *EventChangeTFAdmin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventChangeTFAdmin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventChangeTFAdmin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventChangeTFAdmin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventChangeTFAdmin.Merge(m, src)
}
func (m *EventChangeTFAdmin) XXX_Size() int {
	return m.Size()
}
func (m *EventChangeTFAdmin) XXX_DiscardUnknown() {
	xxx_messageInfo_EventChangeTFAdmin.DiscardUnknown(m)
}

var xxx_messageInfo_EventChangeTFAdmin proto.InternalMessageInfo

func (m *EventChangeTFAdmin) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *EventChangeTFAdmin) GetNewAdminAddress() string {
	if m != nil {
		return m.NewAdminAddress
	}
	return ""
}

type EventSetTFDenomMetadata struct {
	Denom    string          `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Metadata types1.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata"`
}

func (m *EventSetTFDenomMetadata) Reset()         { *m = EventSetTFDenomMetadata{} }
func (m *EventSetTFDenomMetadata) String() string { return proto.CompactTextString(m) }
func (*EventSetTFDenomMetadata) ProtoMessage()    {}
func (*EventSetTFDenomMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9fd0c5434c2a5b7, []int{4}
}
func (m *EventSetTFDenomMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetTFDenomMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetTFDenomMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetTFDenomMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetTFDenomMetadata.Merge(m, src)
}
func (m *EventSetTFDenomMetadata) XXX_Size() int {
	return m.Size()
}
func (m *EventSetTFDenomMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetTFDenomMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetTFDenomMetadata proto.InternalMessageInfo

func (m *EventSetTFDenomMetadata) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *EventSetTFDenomMetadata) GetMetadata() types1.Metadata {
	if m != nil {
		return m.Metadata
	}
	return types1.Metadata{}
}

func init() {
	proto.RegisterType((*EventCreateTFDenom)(nil), "injective.tokenfactory.v1beta1.EventCreateTFDenom")
	proto.RegisterType((*EventMintTFDenom)(nil), "injective.tokenfactory.v1beta1.EventMintTFDenom")
	proto.RegisterType((*EventBurnTFDenom)(nil), "injective.tokenfactory.v1beta1.EventBurnTFDenom")
	proto.RegisterType((*EventChangeTFAdmin)(nil), "injective.tokenfactory.v1beta1.EventChangeTFAdmin")
	proto.RegisterType((*EventSetTFDenomMetadata)(nil), "injective.tokenfactory.v1beta1.EventSetTFDenomMetadata")
}

func init() {
	proto.RegisterFile("injective/tokenfactory/v1beta1/events.proto", fileDescriptor_b9fd0c5434c2a5b7)
}

var fileDescriptor_b9fd0c5434c2a5b7 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x6a, 0x14, 0x41,
	0x10, 0xdd, 0x11, 0x8d, 0xda, 0xa2, 0x26, 0x43, 0xc0, 0x35, 0xe0, 0x28, 0x0b, 0x82, 0x18, 0x9c,
	0x21, 0x0a, 0x7a, 0x94, 0x6c, 0x62, 0x40, 0x30, 0x97, 0x35, 0x78, 0xf0, 0x12, 0x7a, 0x7a, 0xca,
	0xdd, 0x76, 0x33, 0x55, 0x43, 0x77, 0xcd, 0xc6, 0xfd, 0x0b, 0x3f, 0x2b, 0xc7, 0x1c, 0x3d, 0x89,
	0xec, 0xfe, 0x88, 0xcc, 0x4c, 0x77, 0xef, 0xaa, 0x88, 0x07, 0x6f, 0x53, 0x5d, 0xaf, 0xde, 0xab,
	0xf7, 0xa6, 0xc4, 0xae, 0xc6, 0xcf, 0xa0, 0x58, 0xcf, 0x20, 0x63, 0x9a, 0x02, 0x7e, 0x92, 0x8a,
	0xc9, 0xcc, 0xb3, 0xd9, 0x5e, 0x0e, 0x2c, 0xf7, 0x32, 0x98, 0x01, 0xb2, 0x4d, 0x2b, 0x43, 0x4c,
	0x71, 0x12, 0xc0, 0xe9, 0x3a, 0x38, 0x75, 0xe0, 0x9d, 0xed, 0x31, 0x8d, 0xa9, 0x85, 0x66, 0xcd,
	0x57, 0x37, 0xb5, 0x93, 0x28, 0xb2, 0x25, 0xd9, 0x2c, 0x97, 0x16, 0x02, 0xaf, 0x22, 0x8d, 0x7f,
	0xf4, 0x71, 0x1a, 0xfa, 0x4d, 0xe1, 0xfa, 0x2f, 0xff, 0xb1, 0xa2, 0xac, 0x79, 0x42, 0x46, 0xf3,
	0xfc, 0x18, 0x58, 0x16, 0x92, 0x65, 0x37, 0x37, 0x38, 0x14, 0xf1, 0x9b, 0x66, 0xfb, 0x03, 0x03,
	0x92, 0xe1, 0xe4, 0xe8, 0x10, 0x90, 0xca, 0xb8, 0x2f, 0xae, 0x4b, 0xa5, 0xa8, 0x46, 0xee, 0x47,
	0x8f, 0xa2, 0x27, 0x37, 0x47, 0xbe, 0x8c, 0xb7, 0xc5, 0xb5, 0xa2, 0x81, 0xf4, 0xaf, 0xb4, 0xef,
	0x5d, 0x31, 0xf8, 0x22, 0x36, 0x5b, 0x96, 0x63, 0x8d, 0xec, 0x39, 0x76, 0xc5, 0x96, 0x01, 0xa5,
	0x2b, 0x0d, 0xc8, 0xa7, 0xb2, 0x28, 0x0c, 0x58, 0xeb, 0xd8, 0x36, 0x43, 0x63, 0xbf, 0x7b, 0x8f,
	0x5f, 0x89, 0x0d, 0x59, 0xb6, 0x7a, 0x0d, 0xef, 0xad, 0xe7, 0xf7, 0xd3, 0xce, 0x6f, 0xda, 0xe4,
	0xe1, 0xa3, 0x4b, 0x0f, 0x48, 0xe3, 0xf0, 0xea, 0xc5, 0xf7, 0x87, 0xbd, 0x91, 0x83, 0x0f, 0x8c,
	0x53, 0x1e, 0xd6, 0x06, 0xbd, 0xf2, 0x63, 0x71, 0x27, 0xaf, 0x0d, 0x82, 0xf9, 0x4d, 0xf6, 0x76,
	0xf7, 0xfa, 0xdf, 0x9a, 0x1f, 0x7c, 0x66, 0x13, 0x89, 0x63, 0x38, 0x39, 0xda, 0x2f, 0x4a, 0x8d,
	0xab, 0x64, 0xa2, 0xb5, 0x64, 0xe2, 0xa7, 0x62, 0x0b, 0xe1, 0xfc, 0x54, 0x36, 0x90, 0xb0, 0x4e,
	0x97, 0xdd, 0x5d, 0x84, 0xf3, 0x76, 0xd4, 0x2d, 0x34, 0xa8, 0xc4, 0xbd, 0x96, 0xf7, 0x3d, 0xf8,
	0x10, 0xfd, 0xcf, 0xfa, 0x0b, 0xf9, 0x6b, 0x71, 0xa3, 0x74, 0x08, 0xe7, 0xe1, 0xc1, 0xca, 0x03,
	0x4e, 0x83, 0x07, 0x4f, 0xe3, 0x7c, 0x84, 0xa1, 0xe1, 0xd9, 0xc5, 0x22, 0x89, 0x2e, 0x17, 0x49,
	0xf4, 0x63, 0x91, 0x44, 0x5f, 0x97, 0x49, 0xef, 0x72, 0x99, 0xf4, 0xbe, 0x2d, 0x93, 0xde, 0xc7,
	0xd1, 0x58, 0xf3, 0xa4, 0xce, 0x53, 0x45, 0x65, 0xf6, 0xd6, 0x9f, 0xd6, 0x3b, 0x99, 0xdb, 0x2c,
	0x1c, 0xda, 0x33, 0x45, 0x06, 0xd6, 0xcb, 0x89, 0xd4, 0x98, 0x95, 0x54, 0xd4, 0x67, 0x60, 0x7f,
	0xbd, 0x42, 0x9e, 0x57, 0x60, 0xf3, 0x8d, 0xf6, 0xe4, 0x5e, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0xf2, 0x1b, 0xa1, 0x5f, 0x4f, 0x03, 0x00, 0x00,
}

func (m *EventCreateTFDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateTFDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateTFDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventMintTFDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMintTFDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMintTFDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.RecipientAddress) > 0 {
		i -= len(m.RecipientAddress)
		copy(dAtA[i:], m.RecipientAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RecipientAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBurnTFDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBurnTFDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBurnTFDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.BurnerAddress) > 0 {
		i -= len(m.BurnerAddress)
		copy(dAtA[i:], m.BurnerAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.BurnerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventChangeTFAdmin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventChangeTFAdmin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventChangeTFAdmin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewAdminAddress) > 0 {
		i -= len(m.NewAdminAddress)
		copy(dAtA[i:], m.NewAdminAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NewAdminAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSetTFDenomMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetTFDenomMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetTFDenomMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreateTFDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventMintTFDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RecipientAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *EventBurnTFDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BurnerAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *EventChangeTFAdmin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NewAdminAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventSetTFDenomMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreateTFDenom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateTFDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateTFDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventMintTFDenom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventMintTFDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMintTFDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventBurnTFDenom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventBurnTFDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBurnTFDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventChangeTFAdmin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventChangeTFAdmin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventChangeTFAdmin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewAdminAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewAdminAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSetTFDenomMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSetTFDenomMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetTFDenomMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
