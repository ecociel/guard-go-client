// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: proto/iam.proto

package __

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

// CheckServiceClient is the client API for CheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckServiceClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Add(ctx context.Context, in *AddTuplesRequest, opts ...grpc.CallOption) (*AddTuplesResponse, error)
	Delete(ctx context.Context, in *DeleteTuplesRequest, opts ...grpc.CallOption) (*DeleteTuplesResponse, error)
	ContentChangeCheck(ctx context.Context, in *ContentChangeRequest, opts ...grpc.CallOption) (*ContentChangeResponse, error)
}

type checkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckServiceClient(cc grpc.ClientConnInterface) CheckServiceClient {
	return &checkServiceClient{cc}
}

func (c *checkServiceClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/am.CheckService/check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/am.CheckService/list", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkServiceClient) Add(ctx context.Context, in *AddTuplesRequest, opts ...grpc.CallOption) (*AddTuplesResponse, error) {
	out := new(AddTuplesResponse)
	err := c.cc.Invoke(ctx, "/am.CheckService/add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkServiceClient) Delete(ctx context.Context, in *DeleteTuplesRequest, opts ...grpc.CallOption) (*DeleteTuplesResponse, error) {
	out := new(DeleteTuplesResponse)
	err := c.cc.Invoke(ctx, "/am.CheckService/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkServiceClient) ContentChangeCheck(ctx context.Context, in *ContentChangeRequest, opts ...grpc.CallOption) (*ContentChangeResponse, error) {
	out := new(ContentChangeResponse)
	err := c.cc.Invoke(ctx, "/am.CheckService/content_change_check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckServiceServer is the server API for CheckService service.
// All implementations must embed UnimplementedCheckServiceServer
// for forward compatibility
type CheckServiceServer interface {
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Add(context.Context, *AddTuplesRequest) (*AddTuplesResponse, error)
	Delete(context.Context, *DeleteTuplesRequest) (*DeleteTuplesResponse, error)
	ContentChangeCheck(context.Context, *ContentChangeRequest) (*ContentChangeResponse, error)
	mustEmbedUnimplementedCheckServiceServer()
}

// UnimplementedCheckServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCheckServiceServer struct {
}

func (UnimplementedCheckServiceServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedCheckServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCheckServiceServer) Add(context.Context, *AddTuplesRequest) (*AddTuplesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCheckServiceServer) Delete(context.Context, *DeleteTuplesRequest) (*DeleteTuplesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCheckServiceServer) ContentChangeCheck(context.Context, *ContentChangeRequest) (*ContentChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContentChangeCheck not implemented")
}
func (UnimplementedCheckServiceServer) mustEmbedUnimplementedCheckServiceServer() {}

// UnsafeCheckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckServiceServer will
// result in compilation errors.
type UnsafeCheckServiceServer interface {
	mustEmbedUnimplementedCheckServiceServer()
}

func RegisterCheckServiceServer(s grpc.ServiceRegistrar, srv CheckServiceServer) {
	s.RegisterService(&CheckService_ServiceDesc, srv)
}

func _CheckService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/am.CheckService/check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServiceServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/am.CheckService/list",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTuplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/am.CheckService/add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServiceServer).Add(ctx, req.(*AddTuplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTuplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/am.CheckService/delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServiceServer).Delete(ctx, req.(*DeleteTuplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckService_ContentChangeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServiceServer).ContentChangeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/am.CheckService/content_change_check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServiceServer).ContentChangeCheck(ctx, req.(*ContentChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckService_ServiceDesc is the grpc.ServiceDesc for CheckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "am.CheckService",
	HandlerType: (*CheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "check",
			Handler:    _CheckService_Check_Handler,
		},
		{
			MethodName: "list",
			Handler:    _CheckService_List_Handler,
		},
		{
			MethodName: "add",
			Handler:    _CheckService_Add_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _CheckService_Delete_Handler,
		},
		{
			MethodName: "content_change_check",
			Handler:    _CheckService_ContentChangeCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/iam.proto",
}