// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/insurance/v1beta1/insurance.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Params struct {
	// default_redemption_notice_period_duration defines the default minimum notice period duration that must pass after an underwriter sends
	// a redemption request before the underwriter can claim his tokens
	DefaultRedemptionNoticePeriodDuration int64 `protobuf:"varint,1,opt,name=default_redemption_notice_period_duration,json=defaultRedemptionNoticePeriodDuration,proto3" json:"default_redemption_notice_period_duration,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc47a7b76393948, []int{0}
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

func (m *Params) GetDefaultRedemptionNoticePeriodDuration() int64 {
	if m != nil {
		return m.DefaultRedemptionNoticePeriodDuration
	}
	return 0
}

type InsuranceFund struct {
	// deposit denomination for the given insurance fund
	DepositDenom string `protobuf:"bytes,1,opt,name=deposit_denom,json=depositDenom,proto3" json:"deposit_denom,omitempty"`
	// insurance fund pool token denomination for the given insurance fund
	InsurancePoolTokenDenom string `protobuf:"bytes,2,opt,name=insurance_pool_token_denom,json=insurancePoolTokenDenom,proto3" json:"insurance_pool_token_denom,omitempty"`
	// redemption_notice_period_duration defines the minimum notice period duration that must pass after an underwriter sends
	// a redemption request before the underwriter can claim his tokens
	RedemptionNoticePeriodDuration int64 `protobuf:"varint,3,opt,name=redemption_notice_period_duration,json=redemptionNoticePeriodDuration,proto3" json:"redemption_notice_period_duration,omitempty"`
	// market ID of the derivative market
	MarketId string `protobuf:"bytes,4,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
}

func (m *InsuranceFund) Reset()         { *m = InsuranceFund{} }
func (m *InsuranceFund) String() string { return proto.CompactTextString(m) }
func (*InsuranceFund) ProtoMessage()    {}
func (*InsuranceFund) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc47a7b76393948, []int{1}
}
func (m *InsuranceFund) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InsuranceFund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InsuranceFund.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InsuranceFund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsuranceFund.Merge(m, src)
}
func (m *InsuranceFund) XXX_Size() int {
	return m.Size()
}
func (m *InsuranceFund) XXX_DiscardUnknown() {
	xxx_messageInfo_InsuranceFund.DiscardUnknown(m)
}

var xxx_messageInfo_InsuranceFund proto.InternalMessageInfo

func (m *InsuranceFund) GetDepositDenom() string {
	if m != nil {
		return m.DepositDenom
	}
	return ""
}

func (m *InsuranceFund) GetInsurancePoolTokenDenom() string {
	if m != nil {
		return m.InsurancePoolTokenDenom
	}
	return ""
}

func (m *InsuranceFund) GetRedemptionNoticePeriodDuration() int64 {
	if m != nil {
		return m.RedemptionNoticePeriodDuration
	}
	return 0
}

func (m *InsuranceFund) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

type RedemptionSchedule struct {
	// address of the redeemer
	Redeemer string `protobuf:"bytes,1,opt,name=redeemer,proto3" json:"redeemer,omitempty"`
	// the time after which the redemption can be claimed
	ClaimableRedemptionTime int64 `protobuf:"varint,2,opt,name=claimable_redemption_time,json=claimableRedemptionTime,proto3" json:"claimable_redemption_time,omitempty"`
	// the insurance_pool_token amount to redeem
	RedemptionAmount types.Coin `protobuf:"bytes,3,opt,name=redemption_amount,json=redemptionAmount,proto3" json:"redemption_amount"`
}

func (m *RedemptionSchedule) Reset()         { *m = RedemptionSchedule{} }
func (m *RedemptionSchedule) String() string { return proto.CompactTextString(m) }
func (*RedemptionSchedule) ProtoMessage()    {}
func (*RedemptionSchedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc47a7b76393948, []int{2}
}
func (m *RedemptionSchedule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RedemptionSchedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RedemptionSchedule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RedemptionSchedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedemptionSchedule.Merge(m, src)
}
func (m *RedemptionSchedule) XXX_Size() int {
	return m.Size()
}
func (m *RedemptionSchedule) XXX_DiscardUnknown() {
	xxx_messageInfo_RedemptionSchedule.DiscardUnknown(m)
}

var xxx_messageInfo_RedemptionSchedule proto.InternalMessageInfo

func (m *RedemptionSchedule) GetRedeemer() string {
	if m != nil {
		return m.Redeemer
	}
	return ""
}

func (m *RedemptionSchedule) GetClaimableRedemptionTime() int64 {
	if m != nil {
		return m.ClaimableRedemptionTime
	}
	return 0
}

func (m *RedemptionSchedule) GetRedemptionAmount() types.Coin {
	if m != nil {
		return m.RedemptionAmount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Params)(nil), "injective.insurance.v1beta1.Params")
	proto.RegisterType((*InsuranceFund)(nil), "injective.insurance.v1beta1.InsuranceFund")
	proto.RegisterType((*RedemptionSchedule)(nil), "injective.insurance.v1beta1.RedemptionSchedule")
}

func init() {
	proto.RegisterFile("injective/insurance/v1beta1/insurance.proto", fileDescriptor_dbc47a7b76393948)
}

var fileDescriptor_dbc47a7b76393948 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6a, 0xd4, 0x40,
	0x18, 0xc7, 0x37, 0xee, 0x52, 0xda, 0xd1, 0x82, 0x0e, 0x42, 0xdb, 0x2d, 0xa4, 0xba, 0x22, 0x28,
	0x62, 0x42, 0xf5, 0x56, 0x4f, 0xd6, 0x22, 0x2c, 0x14, 0x5d, 0x62, 0x0f, 0xe2, 0x25, 0x4c, 0x66,
	0x3e, 0x77, 0xc7, 0x66, 0xe6, 0x0b, 0x33, 0x93, 0x82, 0x6f, 0xe1, 0x23, 0xf8, 0x1c, 0x3e, 0x41,
	0x8f, 0x7b, 0xd4, 0x8b, 0xc8, 0xee, 0xc5, 0xc7, 0x90, 0x4c, 0xb2, 0x49, 0x4e, 0xf6, 0xb6, 0xf3,
	0xff, 0x7e, 0xff, 0x6f, 0xf3, 0x9f, 0xf9, 0x93, 0x67, 0x52, 0x7f, 0x01, 0xee, 0xe4, 0x15, 0xc4,
	0x52, 0xdb, 0xd2, 0x30, 0xcd, 0x21, 0xbe, 0x3a, 0xce, 0xc0, 0xb1, 0xe3, 0x4e, 0x89, 0x0a, 0x83,
	0x0e, 0xe9, 0x61, 0x0b, 0x47, 0xdd, 0xa8, 0x81, 0xc7, 0xf7, 0xe7, 0x38, 0x47, 0xcf, 0xc5, 0xd5,
	0xaf, 0xda, 0x32, 0x0e, 0x39, 0x5a, 0x85, 0x36, 0xce, 0x98, 0xed, 0xf6, 0x72, 0x94, 0xba, 0x9e,
	0x4f, 0x16, 0x64, 0x6b, 0xc6, 0x0c, 0x53, 0x96, 0x7e, 0x24, 0x4f, 0x05, 0x7c, 0x66, 0x65, 0xee,
	0x52, 0x03, 0x02, 0x54, 0xe1, 0x24, 0xea, 0x54, 0xa3, 0x93, 0x1c, 0xd2, 0x02, 0x8c, 0x44, 0x91,
	0x8a, 0xd2, 0xb0, 0x4a, 0xde, 0x0f, 0x1e, 0x04, 0x4f, 0x86, 0xc9, 0xe3, 0xc6, 0x90, 0xb4, 0xfc,
	0x3b, 0x8f, 0xcf, 0x3c, 0x7d, 0xd6, 0xc0, 0x27, 0xa3, 0xbf, 0xdf, 0x8f, 0x82, 0xc9, 0xaf, 0x80,
	0xec, 0x4e, 0x37, 0x5f, 0xfd, 0xb6, 0xd4, 0x82, 0x3e, 0x22, 0xbb, 0x02, 0x0a, 0xb4, 0xd2, 0xa5,
	0x02, 0x34, 0x2a, 0xbf, 0x75, 0x27, 0xb9, 0xd3, 0x88, 0x67, 0x95, 0x46, 0x5f, 0x91, 0x71, 0x9b,
	0x35, 0x2d, 0x10, 0xf3, 0xd4, 0xe1, 0x25, 0xe8, 0xc6, 0x71, 0xcb, 0x3b, 0xf6, 0x5a, 0x62, 0x86,
	0x98, 0x5f, 0x54, 0xf3, 0xda, 0x3c, 0x25, 0x0f, 0x6f, 0xce, 0x32, 0xf4, 0x59, 0x42, 0xf3, 0xdf,
	0x10, 0xf4, 0x90, 0xec, 0x28, 0x66, 0x2e, 0xc1, 0xa5, 0x52, 0xec, 0x8f, 0xfc, 0xdf, 0x6e, 0xd7,
	0xc2, 0x54, 0x4c, 0x7e, 0x04, 0x84, 0x76, 0x97, 0xf0, 0x81, 0x2f, 0x40, 0x94, 0x39, 0xd0, 0x31,
	0xd9, 0xae, 0xb6, 0x82, 0x02, 0xd3, 0x64, 0x6b, 0xcf, 0xf4, 0x84, 0x1c, 0xf0, 0x9c, 0x49, 0xc5,
	0xb2, 0x1c, 0xfa, 0x17, 0xee, 0xa4, 0x02, 0x1f, 0x6b, 0x98, 0xec, 0xb5, 0x40, 0xb7, 0xfb, 0x42,
	0x2a, 0xa0, 0xe7, 0xe4, 0x5e, 0xcf, 0xc1, 0x14, 0x96, 0xda, 0xf9, 0x18, 0xb7, 0x5f, 0x1c, 0x44,
	0xf5, 0x83, 0x47, 0xd5, 0x83, 0x6f, 0xba, 0x11, 0xbd, 0x41, 0xa9, 0x4f, 0x47, 0xd7, 0xbf, 0x8f,
	0x06, 0xc9, 0xdd, 0xce, 0xf9, 0xda, 0x1b, 0x4f, 0xe5, 0xf5, 0x2a, 0x0c, 0x96, 0xab, 0x30, 0xf8,
	0xb3, 0x0a, 0x83, 0x6f, 0xeb, 0x70, 0xb0, 0x5c, 0x87, 0x83, 0x9f, 0xeb, 0x70, 0xf0, 0xe9, 0xfd,
	0x5c, 0xba, 0x45, 0x99, 0x45, 0x1c, 0x55, 0x3c, 0xdd, 0x54, 0xef, 0x9c, 0x65, 0x36, 0x6e, 0x8b,
	0xf8, 0x9c, 0xa3, 0x81, 0xfe, 0x71, 0xc1, 0xa4, 0x8e, 0x15, 0x56, 0xf9, 0x6d, 0xaf, 0xd2, 0xee,
	0x6b, 0x01, 0x36, 0xdb, 0xf2, 0xa5, 0x7b, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xeb, 0x11,
	0x5c, 0xf6, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.DefaultRedemptionNoticePeriodDuration != that1.DefaultRedemptionNoticePeriodDuration {
		return false
	}
	return true
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
	if m.DefaultRedemptionNoticePeriodDuration != 0 {
		i = encodeVarintInsurance(dAtA, i, uint64(m.DefaultRedemptionNoticePeriodDuration))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *InsuranceFund) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InsuranceFund) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InsuranceFund) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MarketId) > 0 {
		i -= len(m.MarketId)
		copy(dAtA[i:], m.MarketId)
		i = encodeVarintInsurance(dAtA, i, uint64(len(m.MarketId)))
		i--
		dAtA[i] = 0x22
	}
	if m.RedemptionNoticePeriodDuration != 0 {
		i = encodeVarintInsurance(dAtA, i, uint64(m.RedemptionNoticePeriodDuration))
		i--
		dAtA[i] = 0x18
	}
	if len(m.InsurancePoolTokenDenom) > 0 {
		i -= len(m.InsurancePoolTokenDenom)
		copy(dAtA[i:], m.InsurancePoolTokenDenom)
		i = encodeVarintInsurance(dAtA, i, uint64(len(m.InsurancePoolTokenDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DepositDenom) > 0 {
		i -= len(m.DepositDenom)
		copy(dAtA[i:], m.DepositDenom)
		i = encodeVarintInsurance(dAtA, i, uint64(len(m.DepositDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RedemptionSchedule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RedemptionSchedule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RedemptionSchedule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RedemptionAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintInsurance(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.ClaimableRedemptionTime != 0 {
		i = encodeVarintInsurance(dAtA, i, uint64(m.ClaimableRedemptionTime))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Redeemer) > 0 {
		i -= len(m.Redeemer)
		copy(dAtA[i:], m.Redeemer)
		i = encodeVarintInsurance(dAtA, i, uint64(len(m.Redeemer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInsurance(dAtA []byte, offset int, v uint64) int {
	offset -= sovInsurance(v)
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
	if m.DefaultRedemptionNoticePeriodDuration != 0 {
		n += 1 + sovInsurance(uint64(m.DefaultRedemptionNoticePeriodDuration))
	}
	return n
}

func (m *InsuranceFund) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DepositDenom)
	if l > 0 {
		n += 1 + l + sovInsurance(uint64(l))
	}
	l = len(m.InsurancePoolTokenDenom)
	if l > 0 {
		n += 1 + l + sovInsurance(uint64(l))
	}
	if m.RedemptionNoticePeriodDuration != 0 {
		n += 1 + sovInsurance(uint64(m.RedemptionNoticePeriodDuration))
	}
	l = len(m.MarketId)
	if l > 0 {
		n += 1 + l + sovInsurance(uint64(l))
	}
	return n
}

func (m *RedemptionSchedule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Redeemer)
	if l > 0 {
		n += 1 + l + sovInsurance(uint64(l))
	}
	if m.ClaimableRedemptionTime != 0 {
		n += 1 + sovInsurance(uint64(m.ClaimableRedemptionTime))
	}
	l = m.RedemptionAmount.Size()
	n += 1 + l + sovInsurance(uint64(l))
	return n
}

func sovInsurance(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInsurance(x uint64) (n int) {
	return sovInsurance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInsurance
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultRedemptionNoticePeriodDuration", wireType)
			}
			m.DefaultRedemptionNoticePeriodDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultRedemptionNoticePeriodDuration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInsurance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInsurance
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
func (m *InsuranceFund) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInsurance
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
			return fmt.Errorf("proto: InsuranceFund: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InsuranceFund: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsurancePoolTokenDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InsurancePoolTokenDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionNoticePeriodDuration", wireType)
			}
			m.RedemptionNoticePeriodDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RedemptionNoticePeriodDuration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInsurance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInsurance
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
func (m *RedemptionSchedule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInsurance
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
			return fmt.Errorf("proto: RedemptionSchedule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RedemptionSchedule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Redeemer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Redeemer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimableRedemptionTime", wireType)
			}
			m.ClaimableRedemptionTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimableRedemptionTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RedemptionAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInsurance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInsurance
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
func skipInsurance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInsurance
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
					return 0, ErrIntOverflowInsurance
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
					return 0, ErrIntOverflowInsurance
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
				return 0, ErrInvalidLengthInsurance
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInsurance
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInsurance
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInsurance        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInsurance          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInsurance = fmt.Errorf("proto: unexpected end of group")
)