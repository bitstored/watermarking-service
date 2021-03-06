// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	EncodeMessageRequest
	EncodeMessageResponse
	DecodeMessageRequest
	DecodeMessageResponse
	WatermarkImageWithTextRequest
	WatermarkImageWithTextResponse
	WatermarkImageWithImageRequest
	WatermarkImageWithImageResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EncodeMessageRequest struct {
	Image       []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Text        []byte `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	FullMessage bool   `protobuf:"varint,3,opt,name=full_message,json=fullMessage" json:"full_message,omitempty"`
}

func (m *EncodeMessageRequest) Reset()                    { *m = EncodeMessageRequest{} }
func (m *EncodeMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*EncodeMessageRequest) ProtoMessage()               {}
func (*EncodeMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EncodeMessageRequest) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *EncodeMessageRequest) GetText() []byte {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *EncodeMessageRequest) GetFullMessage() bool {
	if m != nil {
		return m.FullMessage
	}
	return false
}

type EncodeMessageResponse struct {
	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (m *EncodeMessageResponse) Reset()                    { *m = EncodeMessageResponse{} }
func (m *EncodeMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*EncodeMessageResponse) ProtoMessage()               {}
func (*EncodeMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EncodeMessageResponse) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

type DecodeMessageRequest struct {
	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (m *DecodeMessageRequest) Reset()                    { *m = DecodeMessageRequest{} }
func (m *DecodeMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*DecodeMessageRequest) ProtoMessage()               {}
func (*DecodeMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DecodeMessageRequest) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

type DecodeMessageResponse struct {
	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Text  []byte `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *DecodeMessageResponse) Reset()                    { *m = DecodeMessageResponse{} }
func (m *DecodeMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*DecodeMessageResponse) ProtoMessage()               {}
func (*DecodeMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DecodeMessageResponse) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *DecodeMessageResponse) GetText() []byte {
	if m != nil {
		return m.Text
	}
	return nil
}

type WatermarkImageWithTextRequest struct {
	Image     []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Watermark string `protobuf:"bytes,2,opt,name=watermark" json:"watermark,omitempty"`
}

func (m *WatermarkImageWithTextRequest) Reset()                    { *m = WatermarkImageWithTextRequest{} }
func (m *WatermarkImageWithTextRequest) String() string            { return proto.CompactTextString(m) }
func (*WatermarkImageWithTextRequest) ProtoMessage()               {}
func (*WatermarkImageWithTextRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WatermarkImageWithTextRequest) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *WatermarkImageWithTextRequest) GetWatermark() string {
	if m != nil {
		return m.Watermark
	}
	return ""
}

type WatermarkImageWithTextResponse struct {
	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (m *WatermarkImageWithTextResponse) Reset()                    { *m = WatermarkImageWithTextResponse{} }
func (m *WatermarkImageWithTextResponse) String() string            { return proto.CompactTextString(m) }
func (*WatermarkImageWithTextResponse) ProtoMessage()               {}
func (*WatermarkImageWithTextResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WatermarkImageWithTextResponse) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

type WatermarkImageWithImageRequest struct {
	Image     []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Watermark []byte `protobuf:"bytes,2,opt,name=watermark,proto3" json:"watermark,omitempty"`
}

func (m *WatermarkImageWithImageRequest) Reset()                    { *m = WatermarkImageWithImageRequest{} }
func (m *WatermarkImageWithImageRequest) String() string            { return proto.CompactTextString(m) }
func (*WatermarkImageWithImageRequest) ProtoMessage()               {}
func (*WatermarkImageWithImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *WatermarkImageWithImageRequest) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *WatermarkImageWithImageRequest) GetWatermark() []byte {
	if m != nil {
		return m.Watermark
	}
	return nil
}

type WatermarkImageWithImageResponse struct {
	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (m *WatermarkImageWithImageResponse) Reset()                    { *m = WatermarkImageWithImageResponse{} }
func (m *WatermarkImageWithImageResponse) String() string            { return proto.CompactTextString(m) }
func (*WatermarkImageWithImageResponse) ProtoMessage()               {}
func (*WatermarkImageWithImageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *WatermarkImageWithImageResponse) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func init() {
	proto.RegisterType((*EncodeMessageRequest)(nil), "watermarking_service.EncodeMessageRequest")
	proto.RegisterType((*EncodeMessageResponse)(nil), "watermarking_service.EncodeMessageResponse")
	proto.RegisterType((*DecodeMessageRequest)(nil), "watermarking_service.DecodeMessageRequest")
	proto.RegisterType((*DecodeMessageResponse)(nil), "watermarking_service.DecodeMessageResponse")
	proto.RegisterType((*WatermarkImageWithTextRequest)(nil), "watermarking_service.WatermarkImageWithTextRequest")
	proto.RegisterType((*WatermarkImageWithTextResponse)(nil), "watermarking_service.WatermarkImageWithTextResponse")
	proto.RegisterType((*WatermarkImageWithImageRequest)(nil), "watermarking_service.WatermarkImageWithImageRequest")
	proto.RegisterType((*WatermarkImageWithImageResponse)(nil), "watermarking_service.WatermarkImageWithImageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Watermarking service

type WatermarkingClient interface {
	EncodeMessage(ctx context.Context, in *EncodeMessageRequest, opts ...grpc.CallOption) (*EncodeMessageResponse, error)
	DecodeMessage(ctx context.Context, in *DecodeMessageRequest, opts ...grpc.CallOption) (*DecodeMessageResponse, error)
	WatermarkImageWithText(ctx context.Context, in *WatermarkImageWithTextRequest, opts ...grpc.CallOption) (*WatermarkImageWithTextResponse, error)
	WatermarkImageWithImage(ctx context.Context, in *WatermarkImageWithImageRequest, opts ...grpc.CallOption) (*WatermarkImageWithImageResponse, error)
}

type watermarkingClient struct {
	cc *grpc.ClientConn
}

func NewWatermarkingClient(cc *grpc.ClientConn) WatermarkingClient {
	return &watermarkingClient{cc}
}

func (c *watermarkingClient) EncodeMessage(ctx context.Context, in *EncodeMessageRequest, opts ...grpc.CallOption) (*EncodeMessageResponse, error) {
	out := new(EncodeMessageResponse)
	err := grpc.Invoke(ctx, "/watermarking_service.Watermarking/EncodeMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkingClient) DecodeMessage(ctx context.Context, in *DecodeMessageRequest, opts ...grpc.CallOption) (*DecodeMessageResponse, error) {
	out := new(DecodeMessageResponse)
	err := grpc.Invoke(ctx, "/watermarking_service.Watermarking/DecodeMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkingClient) WatermarkImageWithText(ctx context.Context, in *WatermarkImageWithTextRequest, opts ...grpc.CallOption) (*WatermarkImageWithTextResponse, error) {
	out := new(WatermarkImageWithTextResponse)
	err := grpc.Invoke(ctx, "/watermarking_service.Watermarking/WatermarkImageWithText", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkingClient) WatermarkImageWithImage(ctx context.Context, in *WatermarkImageWithImageRequest, opts ...grpc.CallOption) (*WatermarkImageWithImageResponse, error) {
	out := new(WatermarkImageWithImageResponse)
	err := grpc.Invoke(ctx, "/watermarking_service.Watermarking/WatermarkImageWithImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Watermarking service

type WatermarkingServer interface {
	EncodeMessage(context.Context, *EncodeMessageRequest) (*EncodeMessageResponse, error)
	DecodeMessage(context.Context, *DecodeMessageRequest) (*DecodeMessageResponse, error)
	WatermarkImageWithText(context.Context, *WatermarkImageWithTextRequest) (*WatermarkImageWithTextResponse, error)
	WatermarkImageWithImage(context.Context, *WatermarkImageWithImageRequest) (*WatermarkImageWithImageResponse, error)
}

func RegisterWatermarkingServer(s *grpc.Server, srv WatermarkingServer) {
	s.RegisterService(&_Watermarking_serviceDesc, srv)
}

func _Watermarking_EncodeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodeMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkingServer).EncodeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/watermarking_service.Watermarking/EncodeMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkingServer).EncodeMessage(ctx, req.(*EncodeMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermarking_DecodeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkingServer).DecodeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/watermarking_service.Watermarking/DecodeMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkingServer).DecodeMessage(ctx, req.(*DecodeMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermarking_WatermarkImageWithText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatermarkImageWithTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkingServer).WatermarkImageWithText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/watermarking_service.Watermarking/WatermarkImageWithText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkingServer).WatermarkImageWithText(ctx, req.(*WatermarkImageWithTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermarking_WatermarkImageWithImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatermarkImageWithImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkingServer).WatermarkImageWithImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/watermarking_service.Watermarking/WatermarkImageWithImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkingServer).WatermarkImageWithImage(ctx, req.(*WatermarkImageWithImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Watermarking_serviceDesc = grpc.ServiceDesc{
	ServiceName: "watermarking_service.Watermarking",
	HandlerType: (*WatermarkingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EncodeMessage",
			Handler:    _Watermarking_EncodeMessage_Handler,
		},
		{
			MethodName: "DecodeMessage",
			Handler:    _Watermarking_DecodeMessage_Handler,
		},
		{
			MethodName: "WatermarkImageWithText",
			Handler:    _Watermarking_WatermarkImageWithText_Handler,
		},
		{
			MethodName: "WatermarkImageWithImage",
			Handler:    _Watermarking_WatermarkImageWithImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0x33, 0x78, 0xaf, 0xf1, 0x1e, 0x7b, 0x37, 0x93, 0xa2, 0xa4, 0x41, 0xc1, 0x51, 0x13,
	0x52, 0xb4, 0x8d, 0xe2, 0x9f, 0x84, 0x9d, 0x46, 0x17, 0x2e, 0xdc, 0x54, 0x12, 0x12, 0x37, 0x64,
	0x80, 0x63, 0x6d, 0x84, 0x4e, 0xed, 0x0c, 0xc8, 0xda, 0x57, 0x70, 0xe3, 0x43, 0x98, 0xb8, 0xf0,
	0x51, 0x7c, 0x05, 0x1f, 0xc4, 0x74, 0x5a, 0xb4, 0x90, 0x69, 0x53, 0x36, 0x64, 0x98, 0x9e, 0xef,
	0x3b, 0xbf, 0x39, 0xf3, 0xb5, 0x70, 0x29, 0x31, 0xdd, 0x46, 0x0b, 0xf4, 0x92, 0x54, 0x28, 0x41,
	0xed, 0x2f, 0x5c, 0x61, 0xba, 0xe6, 0xe9, 0xa7, 0x28, 0x0e, 0x67, 0xc5, 0x33, 0xa7, 0x1b, 0x0a,
	0x11, 0xae, 0xd0, 0xe7, 0x49, 0xe4, 0xf3, 0x38, 0x16, 0x8a, 0xab, 0x48, 0xc4, 0x32, 0xd7, 0xb0,
	0x05, 0xd8, 0xaf, 0xe3, 0x85, 0x58, 0xe2, 0x5b, 0x94, 0x92, 0x87, 0x18, 0xe0, 0xe7, 0x0d, 0x4a,
	0x45, 0x6d, 0x38, 0x8f, 0xd6, 0x3c, 0xc4, 0x0e, 0xe9, 0x93, 0x81, 0x15, 0xe4, 0x7f, 0x28, 0x85,
	0x33, 0x85, 0x3b, 0xd5, 0x69, 0xe9, 0x4d, 0xbd, 0xa6, 0x77, 0xc0, 0xfa, 0xb0, 0x59, 0xad, 0x66,
	0xeb, 0xdc, 0xa0, 0x73, 0xa5, 0x4f, 0x06, 0xd7, 0x82, 0xeb, 0xd9, 0x5e, 0xe1, 0xc9, 0x1e, 0x42,
	0xfb, 0xa8, 0x89, 0x4c, 0x44, 0x2c, 0xd1, 0xdc, 0x85, 0x3d, 0x00, 0xfb, 0x15, 0x36, 0x65, 0x62,
	0x2f, 0xa0, 0x7d, 0x54, 0x5d, 0x67, 0x6e, 0x3a, 0x02, 0x7b, 0x07, 0xb7, 0xa6, 0xfb, 0xd1, 0xbd,
	0xc9, 0xaa, 0xa6, 0x91, 0xfa, 0x38, 0xc1, 0x9d, 0xaa, 0x9f, 0x46, 0x17, 0x2e, 0xfe, 0x4d, 0x5c,
	0xfb, 0x5d, 0x04, 0xff, 0x37, 0xd8, 0x33, 0xb8, 0x5d, 0x65, 0x5a, 0x7b, 0xfa, 0x89, 0x49, 0xa7,
	0x17, 0x27, 0xd2, 0x58, 0x65, 0x9a, 0xe7, 0xd0, 0xab, 0x74, 0xad, 0xc3, 0x79, 0xfc, 0xe3, 0x1c,
	0xac, 0x69, 0x29, 0x57, 0xf4, 0x3b, 0x81, 0xcb, 0x83, 0xdb, 0xa4, 0xae, 0x67, 0x0a, 0x9e, 0x67,
	0xca, 0x95, 0x33, 0x6c, 0x54, 0x9b, 0x13, 0x31, 0xef, 0xeb, 0xef, 0x3f, 0xdf, 0x5a, 0x03, 0x76,
	0xd7, 0x2f, 0x8b, 0x74, 0x92, 0xb7, 0x8f, 0x7c, 0xa9, 0x30, 0xe4, 0xb1, 0xf0, 0x51, 0x6b, 0xc7,
	0xc4, 0xd5, 0x68, 0x07, 0x59, 0xa8, 0x42, 0x33, 0xc5, 0xab, 0x0a, 0xcd, 0x18, 0xae, 0x86, 0x68,
	0x4b, 0xdc, 0xa3, 0xfd, 0x24, 0x70, 0xc3, 0x1c, 0x07, 0x3a, 0x32, 0xf7, 0xad, 0x4d, 0xa4, 0xf3,
	0xe4, 0x34, 0x51, 0x41, 0xed, 0x6a, 0xea, 0x7b, 0xac, 0x67, 0xa4, 0xd6, 0x17, 0xee, 0x67, 0x6f,
	0x44, 0x46, 0xfc, 0x8b, 0xc0, 0xcd, 0x8a, 0xc8, 0xd0, 0xc6, 0xdd, 0xcb, 0xb9, 0x75, 0x9e, 0x9e,
	0xa8, 0x2a, 0xa0, 0x87, 0x1a, 0xfa, 0x3e, 0xeb, 0xd7, 0x40, 0xeb, 0xdf, 0x31, 0x71, 0x5f, 0x9e,
	0xbd, 0x6f, 0x25, 0xf3, 0xf9, 0x55, 0xfd, 0x71, 0x1b, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x05,
	0xed, 0xb6, 0xb8, 0x21, 0x05, 0x00, 0x00,
}
