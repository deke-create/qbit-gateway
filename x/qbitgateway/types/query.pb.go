// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qbitgateway/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("qbitgateway/query.proto", fileDescriptor_1d1ca50a2d35b5a4) }

var fileDescriptor_1d1ca50a2d35b5a4 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x4e, 0x84, 0x40,
	0x10, 0x86, 0xa1, 0x50, 0x13, 0x4a, 0x1b, 0x13, 0x62, 0xb6, 0xb0, 0x34, 0x81, 0x09, 0x5a, 0xd8,
	0xfb, 0x06, 0xb4, 0x76, 0xb3, 0x38, 0x59, 0x37, 0xca, 0xce, 0xc2, 0x0e, 0x2a, 0x6f, 0x71, 0x8f,
	0x75, 0x25, 0xe5, 0x95, 0x17, 0x78, 0x91, 0x0b, 0x6c, 0x2e, 0xa1, 0x9b, 0xe2, 0xcb, 0x3f, 0xdf,
	0x97, 0x3d, 0x74, 0xda, 0x8a, 0x41, 0xa1, 0x3f, 0x1c, 0xa1, 0x1b, 0xa8, 0x1f, 0x4b, 0xdf, 0xb3,
	0xf0, 0xfd, 0xd3, 0x27, 0x7d, 0x53, 0xd3, 0x13, 0x0a, 0x95, 0x3b, 0x66, 0x7f, 0xe7, 0x8f, 0x86,
	0xd9, 0xfc, 0x10, 0xa0, 0xb7, 0x80, 0xce, 0xb1, 0xa0, 0x58, 0x76, 0x21, 0x2e, 0xe4, 0xcf, 0x0d,
	0x87, 0x96, 0x03, 0x68, 0x0c, 0x14, 0xa7, 0xe1, 0xb7, 0xd2, 0x24, 0x58, 0x81, 0x47, 0x63, 0xdd,
	0x06, 0x47, 0xf6, 0xe5, 0x2e, 0xbb, 0xa9, 0x57, 0xe2, 0xbd, 0x3e, 0xce, 0x2a, 0x9d, 0x66, 0x95,
	0x9e, 0x67, 0x95, 0x1e, 0x16, 0x95, 0x4c, 0x8b, 0x4a, 0x4e, 0x8b, 0x4a, 0x3e, 0xde, 0x8c, 0x95,
	0xaf, 0x41, 0x97, 0x0d, 0xb7, 0xb0, 0xba, 0x15, 0x51, 0x0e, 0x56, 0xa1, 0xe2, 0x5a, 0xf0, 0x0f,
	0xfb, 0x1e, 0x19, 0x3d, 0x05, 0x7d, 0xbb, 0xbd, 0x78, 0xbd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0x5b, 0xad, 0xbd, 0xeb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dekecreate.qbitgateway.qbitgateway.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "qbitgateway/query.proto",
}