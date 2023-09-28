//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package grpc

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

// WeaviateClient is the client API for Weaviate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeaviateClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error)
	BatchObjects(ctx context.Context, in *BatchObjectsRequest, opts ...grpc.CallOption) (*BatchObjectsReply, error)
}

type weaviateClient struct {
	cc grpc.ClientConnInterface
}

func NewWeaviateClient(cc grpc.ClientConnInterface) WeaviateClient {
	return &weaviateClient{cc}
}

func (c *weaviateClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error) {
	out := new(SearchReply)
	err := c.cc.Invoke(ctx, "/weaviategrpc.Weaviate/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weaviateClient) BatchObjects(ctx context.Context, in *BatchObjectsRequest, opts ...grpc.CallOption) (*BatchObjectsReply, error) {
	out := new(BatchObjectsReply)
	err := c.cc.Invoke(ctx, "/weaviategrpc.Weaviate/BatchObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeaviateServer is the server API for Weaviate service.
// All implementations must embed UnimplementedWeaviateServer
// for forward compatibility
type WeaviateServer interface {
	Search(context.Context, *SearchRequest) (*SearchReply, error)
	BatchObjects(context.Context, *BatchObjectsRequest) (*BatchObjectsReply, error)
	mustEmbedUnimplementedWeaviateServer()
}

// UnimplementedWeaviateServer must be embedded to have forward compatible implementations.
type UnimplementedWeaviateServer struct{}

func (UnimplementedWeaviateServer) Search(context.Context, *SearchRequest) (*SearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func (UnimplementedWeaviateServer) BatchObjects(context.Context, *BatchObjectsRequest) (*BatchObjectsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchObjects not implemented")
}
func (UnimplementedWeaviateServer) mustEmbedUnimplementedWeaviateServer() {}

// UnsafeWeaviateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeaviateServer will
// result in compilation errors.
type UnsafeWeaviateServer interface {
	mustEmbedUnimplementedWeaviateServer()
}

func RegisterWeaviateServer(s grpc.ServiceRegistrar, srv WeaviateServer) {
	s.RegisterService(&Weaviate_ServiceDesc, srv)
}

func _Weaviate_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeaviateServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weaviategrpc.Weaviate/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeaviateServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weaviate_BatchObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeaviateServer).BatchObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weaviategrpc.Weaviate/BatchObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeaviateServer).BatchObjects(ctx, req.(*BatchObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Weaviate_ServiceDesc is the grpc.ServiceDesc for Weaviate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Weaviate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weaviategrpc.Weaviate",
	HandlerType: (*WeaviateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Weaviate_Search_Handler,
		},
		{
			MethodName: "BatchObjects",
			Handler:    _Weaviate_BatchObjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weaviate.proto",
}