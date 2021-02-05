// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/peggy/v1beta1/ethereum_signer.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

// SignType defines messages that have been signed by an orchestrator
type SignType int32

const (
	SIGN_TYPE_UNKNOWN                              SignType = 0
	SIGN_TYPE_ORCHESTRATOR_SIGNED_MULTI_SIG_UPDATE SignType = 1
	SIGN_TYPE_ORCHESTRATOR_SIGNED_WITHDRAW_BATCH   SignType = 2
)

var SignType_name = map[int32]string{
	0: "SIGN_TYPE_UNKNOWN",
	1: "SIGN_TYPE_ORCHESTRATOR_SIGNED_MULTI_SIG_UPDATE",
	2: "SIGN_TYPE_ORCHESTRATOR_SIGNED_WITHDRAW_BATCH",
}

var SignType_value = map[string]int32{
	"SIGN_TYPE_UNKNOWN": 0,
	"SIGN_TYPE_ORCHESTRATOR_SIGNED_MULTI_SIG_UPDATE": 1,
	"SIGN_TYPE_ORCHESTRATOR_SIGNED_WITHDRAW_BATCH":   2,
}

func (SignType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_614115c1c864803f, []int{0}
}

func init() {
	proto.RegisterEnum("injective.peggy.v1beta1.SignType", SignType_name, SignType_value)
}

func init() {
	proto.RegisterFile("injective/peggy/v1beta1/ethereum_signer.proto", fileDescriptor_614115c1c864803f)
}

var fileDescriptor_614115c1c864803f = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcd, 0xcc, 0xcb, 0x4a,
	0x4d, 0x2e, 0xc9, 0x2c, 0x4b, 0xd5, 0x2f, 0x48, 0x4d, 0x4f, 0xaf, 0xd4, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0xd4, 0x4f, 0x2d, 0xc9, 0x48, 0x2d, 0x4a, 0x2d, 0xcd, 0x8d, 0x2f, 0xce, 0x4c,
	0xcf, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0x2b, 0xd7, 0x03, 0x2b,
	0xd7, 0x83, 0x2a, 0x97, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd1, 0x07, 0xb1, 0x20, 0xca,
	0xb5, 0x26, 0x32, 0x72, 0x71, 0x04, 0x67, 0xa6, 0xe7, 0x85, 0x54, 0x16, 0xa4, 0x0a, 0x89, 0x72,
	0x09, 0x06, 0x7b, 0xba, 0xfb, 0xc5, 0x87, 0x44, 0x06, 0xb8, 0xc6, 0x87, 0xfa, 0x79, 0xfb, 0xf9,
	0x87, 0xfb, 0x09, 0x30, 0x08, 0x19, 0x71, 0xe9, 0x21, 0x84, 0xfd, 0x83, 0x9c, 0x3d, 0x5c, 0x83,
	0x43, 0x82, 0x1c, 0x43, 0xfc, 0x83, 0xe2, 0x41, 0xc2, 0xae, 0x2e, 0xf1, 0xbe, 0xa1, 0x3e, 0x21,
	0x9e, 0x20, 0x4e, 0x7c, 0x68, 0x80, 0x8b, 0x63, 0x88, 0xab, 0x00, 0xa3, 0x90, 0x01, 0x97, 0x0e,
	0x7e, 0x3d, 0xe1, 0x9e, 0x21, 0x1e, 0x2e, 0x41, 0x8e, 0xe1, 0xf1, 0x4e, 0x8e, 0x21, 0xce, 0x1e,
	0x02, 0x4c, 0x52, 0x1c, 0x1d, 0x8b, 0xe5, 0x18, 0x56, 0x2c, 0x91, 0x63, 0x70, 0x4a, 0x3d, 0xf1,
	0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8,
	0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xef, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24,
	0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x4f, 0x98, 0x3f, 0x7d, 0x12, 0x93, 0x8a, 0xf5, 0xe1, 0xbe, 0xd6,
	0x4d, 0xce, 0x2f, 0x4a, 0x45, 0xe6, 0x66, 0x24, 0x66, 0xe6, 0xe9, 0xe7, 0xe6, 0xa7, 0x94, 0xe6,
	0xa4, 0x16, 0x43, 0x43, 0xb0, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0x02, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0xfb, 0x44, 0x90, 0x61, 0x01, 0x00, 0x00,
}