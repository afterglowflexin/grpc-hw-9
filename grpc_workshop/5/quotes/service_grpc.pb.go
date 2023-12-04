// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: service.proto

package quotes

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	QuotesService_GetQuotes_FullMethodName = "/quotes.QuotesService/GetQuotes"
)

// QuotesServiceClient is the client API for QuotesService service.
//
// For semantics around ctx use and closing/ending streaming.proto RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuotesServiceClient interface {
	// GetQuotes returns a list of quotes or an error
	GetQuotes(ctx context.Context, in *QuotesRequest, opts ...grpc.CallOption) (*QuotesResponse, error)
}

type quotesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuotesServiceClient(cc grpc.ClientConnInterface) QuotesServiceClient {
	return &quotesServiceClient{cc}
}

func (c *quotesServiceClient) GetQuotes(ctx context.Context, in *QuotesRequest, opts ...grpc.CallOption) (*QuotesResponse, error) {
	out := new(QuotesResponse)
	err := c.cc.Invoke(ctx, QuotesService_GetQuotes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuotesServiceServer is the server API for QuotesService service.
// All implementations must embed UnimplementedQuotesServiceServer
// for forward compatibility
type QuotesServiceServer interface {
	// GetQuotes returns a list of quotes or an error
	GetQuotes(context.Context, *QuotesRequest) (*QuotesResponse, error)
	mustEmbedUnimplementedQuotesServiceServer()
}

// UnimplementedQuotesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuotesServiceServer struct {
}

func (UnimplementedQuotesServiceServer) GetQuotes(context.Context, *QuotesRequest) (*QuotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuotes not implemented")
}
func (UnimplementedQuotesServiceServer) mustEmbedUnimplementedQuotesServiceServer() {}

// UnsafeQuotesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuotesServiceServer will
// result in compilation errors.
type UnsafeQuotesServiceServer interface {
	mustEmbedUnimplementedQuotesServiceServer()
}

func RegisterQuotesServiceServer(s grpc.ServiceRegistrar, srv QuotesServiceServer) {
	s.RegisterService(&QuotesService_ServiceDesc, srv)
}

func _QuotesService_GetQuotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotesServiceServer).GetQuotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuotesService_GetQuotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotesServiceServer).GetQuotes(ctx, req.(*QuotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuotesService_ServiceDesc is the grpc.ServiceDesc for QuotesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuotesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quotes.QuotesService",
	HandlerType: (*QuotesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuotes",
			Handler:    _QuotesService_GetQuotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
