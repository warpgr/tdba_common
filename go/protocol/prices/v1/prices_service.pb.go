// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prices/v1/prices_service.proto

package prices

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("prices/v1/prices_service.proto", fileDescriptor_b5d2f47953755fd6)
}

var fileDescriptor_b5d2f47953755fd6 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x28, 0xca, 0x4c,
	0x4e, 0x2d, 0xd6, 0x2f, 0x33, 0xd4, 0x87, 0xb0, 0xe2, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53,
	0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x21, 0xa2, 0x7a, 0x65, 0x86, 0x52, 0x62, 0xe8,
	0x4a, 0x21, 0x4a, 0x8c, 0xe6, 0xb0, 0x70, 0xf1, 0x06, 0x80, 0x05, 0x82, 0x21, 0x5a, 0x85, 0xfc,
	0xb8, 0xb8, 0xdd, 0x53, 0x4b, 0xc0, 0x62, 0x6e, 0xf9, 0x45, 0x42, 0xb2, 0x7a, 0x70, 0x43, 0xf4,
	0x90, 0xc4, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0xe4, 0x70, 0x49, 0x17, 0x17, 0xe4,
	0xe7, 0x15, 0xa7, 0x2a, 0x31, 0x08, 0x05, 0x72, 0xf1, 0x04, 0xe4, 0xe7, 0xe4, 0x50, 0xcd, 0x40,
	0x03, 0x46, 0x90, 0x91, 0x30, 0xa9, 0x62, 0x90, 0x91, 0xd8, 0xf4, 0x14, 0x23, 0x99, 0x29, 0x8f,
	0x53, 0x1e, 0xee, 0xca, 0x10, 0x2e, 0x5e, 0xb8, 0x2b, 0xa9, 0x64, 0xa6, 0x01, 0xa3, 0x50, 0x2c,
	0x97, 0x10, 0xc2, 0xef, 0x45, 0xf9, 0xb9, 0x21, 0xf9, 0x20, 0xa3, 0x15, 0xb0, 0x79, 0x11, 0x2c,
	0x0b, 0x33, 0x5c, 0x11, 0x8f, 0x0a, 0x24, 0xe3, 0x13, 0xb8, 0x84, 0x91, 0x1c, 0x0d, 0x37, 0x5f,
	0x11, 0xab, 0xd3, 0x50, 0x2c, 0x50, 0xc2, 0xa7, 0x04, 0x61, 0x83, 0x93, 0x4d, 0x94, 0x55, 0x7a,
	0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x79, 0x62, 0x51, 0x41, 0x7a, 0x91,
	0x7e, 0x49, 0x4a, 0x52, 0x62, 0x7c, 0x72, 0x7e, 0x6e, 0x6e, 0x7e, 0x9e, 0x7e, 0x7a, 0xbe, 0x3e,
	0x38, 0x21, 0x25, 0xe7, 0xe7, 0xe8, 0xc3, 0x93, 0x98, 0x35, 0x84, 0x95, 0xc4, 0x06, 0x96, 0x32,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x57, 0xdd, 0x90, 0xa8, 0x02, 0x00, 0x00,
}
