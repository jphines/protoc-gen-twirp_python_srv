// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	echo.proto

It has these top-level messages:
	EchoRequest
	EchoMultiRequest
	EchoResponse
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EchoRequest struct {
	Input string `protobuf:"bytes,1,opt,name=input" json:"input,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EchoRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

type EchoMultiRequest struct {
	Input string `protobuf:"bytes,1,opt,name=input" json:"input,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *EchoMultiRequest) Reset()                    { *m = EchoMultiRequest{} }
func (m *EchoMultiRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoMultiRequest) ProtoMessage()               {}
func (*EchoMultiRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EchoMultiRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *EchoMultiRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type EchoResponse struct {
	Output string `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EchoResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "example.echo.EchoRequest")
	proto.RegisterType((*EchoMultiRequest)(nil), "example.echo.EchoMultiRequest")
	proto.RegisterType((*EchoResponse)(nil), "example.echo.EchoResponse")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5,
	0x03, 0x89, 0x29, 0x29, 0x73, 0x71, 0xbb, 0x26, 0x67, 0xe4, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16,
	0x97, 0x08, 0x89, 0x70, 0xb1, 0x66, 0xe6, 0x15, 0x94, 0x96, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x41, 0x38, 0x4a, 0x76, 0x5c, 0x02, 0x20, 0x45, 0xbe, 0xa5, 0x39, 0x25, 0x99, 0x78, 0x55,
	0x82, 0x44, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0x20, 0x1c,
	0x25, 0x35, 0x2e, 0x1e, 0x88, 0x25, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c,
	0xf9, 0xa5, 0x25, 0x08, 0xcd, 0x50, 0x9e, 0xd1, 0x54, 0x46, 0x2e, 0x16, 0x90, 0x42, 0x21, 0x7b,
	0x2e, 0xb6, 0xa0, 0xd4, 0x82, 0xd4, 0xc4, 0x12, 0x21, 0x49, 0x3d, 0x64, 0xe7, 0xea, 0x21, 0xb9,
	0x55, 0x4a, 0x0a, 0x9b, 0x14, 0xd4, 0x06, 0x1f, 0x2e, 0x3e, 0x88, 0x01, 0x60, 0x37, 0x17, 0xe4,
	0xa4, 0x0a, 0xc9, 0x61, 0xaa, 0x46, 0xf6, 0x0f, 0x3e, 0xd3, 0x9c, 0x38, 0xa3, 0xd8, 0xa1, 0x92,
	0x49, 0x6c, 0xe0, 0x40, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x1d, 0x36, 0xba, 0x52,
	0x01, 0x00, 0x00,
}
