// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package archivist

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ArchivistClient is the client API for Archivist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArchivistClient interface {
	GetBySubject(ctx context.Context, in *GetBySubjectRequest, opts ...grpc.CallOption) (*GetBySubjectResponse, error)
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type archivistClient struct {
	cc grpc.ClientConnInterface
}

func NewArchivistClient(cc grpc.ClientConnInterface) ArchivistClient {
	return &archivistClient{cc}
}

func (c *archivistClient) GetBySubject(ctx context.Context, in *GetBySubjectRequest, opts ...grpc.CallOption) (*GetBySubjectResponse, error) {
	out := new(GetBySubjectResponse)
	err := c.cc.Invoke(ctx, "/archivist.Archivist/GetBySubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archivistClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/archivist.Archivist/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArchivistServer is the server API for Archivist service.
// All implementations must embed UnimplementedArchivistServer
// for forward compatibility
type ArchivistServer interface {
	GetBySubject(context.Context, *GetBySubjectRequest) (*GetBySubjectResponse, error)
	Store(context.Context, *StoreRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedArchivistServer()
}

// UnimplementedArchivistServer must be embedded to have forward compatible implementations.
type UnimplementedArchivistServer struct {
}

func (UnimplementedArchivistServer) GetBySubject(context.Context, *GetBySubjectRequest) (*GetBySubjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBySubject not implemented")
}
func (UnimplementedArchivistServer) Store(context.Context, *StoreRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedArchivistServer) mustEmbedUnimplementedArchivistServer() {}

// UnsafeArchivistServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArchivistServer will
// result in compilation errors.
type UnsafeArchivistServer interface {
	mustEmbedUnimplementedArchivistServer()
}

func RegisterArchivistServer(s grpc.ServiceRegistrar, srv ArchivistServer) {
	s.RegisterService(&Archivist_ServiceDesc, srv)
}

func _Archivist_GetBySubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBySubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchivistServer).GetBySubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archivist.Archivist/GetBySubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchivistServer).GetBySubject(ctx, req.(*GetBySubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Archivist_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchivistServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archivist.Archivist/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchivistServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Archivist_ServiceDesc is the grpc.ServiceDesc for Archivist service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Archivist_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "archivist.Archivist",
	HandlerType: (*ArchivistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBySubject",
			Handler:    _Archivist_GetBySubject_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _Archivist_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "archivist.proto",
}
