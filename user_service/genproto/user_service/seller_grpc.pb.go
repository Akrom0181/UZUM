// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: seller.proto

package user_service

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
	SellerService_Create_FullMethodName  = "/user_service.SellerService/Create"
	SellerService_GetByID_FullMethodName = "/user_service.SellerService/GetByID"
	SellerService_GetList_FullMethodName = "/user_service.SellerService/GetList"
	SellerService_Update_FullMethodName  = "/user_service.SellerService/Update"
	SellerService_Delete_FullMethodName  = "/user_service.SellerService/Delete"
)

// SellerServiceClient is the client API for SellerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SellerServiceClient interface {
	Create(ctx context.Context, in *CreateSeller, opts ...grpc.CallOption) (*Seller, error)
	GetByID(ctx context.Context, in *SellerPrimaryKey, opts ...grpc.CallOption) (*Seller, error)
	GetList(ctx context.Context, in *GetListSellerRequest, opts ...grpc.CallOption) (*GetListSellerResponse, error)
	Update(ctx context.Context, in *UpdateSeller, opts ...grpc.CallOption) (*Seller, error)
	Delete(ctx context.Context, in *SellerPrimaryKey, opts ...grpc.CallOption) (*Empty2, error)
}

type sellerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSellerServiceClient(cc grpc.ClientConnInterface) SellerServiceClient {
	return &sellerServiceClient{cc}
}

func (c *sellerServiceClient) Create(ctx context.Context, in *CreateSeller, opts ...grpc.CallOption) (*Seller, error) {
	out := new(Seller)
	err := c.cc.Invoke(ctx, SellerService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerServiceClient) GetByID(ctx context.Context, in *SellerPrimaryKey, opts ...grpc.CallOption) (*Seller, error) {
	out := new(Seller)
	err := c.cc.Invoke(ctx, SellerService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerServiceClient) GetList(ctx context.Context, in *GetListSellerRequest, opts ...grpc.CallOption) (*GetListSellerResponse, error) {
	out := new(GetListSellerResponse)
	err := c.cc.Invoke(ctx, SellerService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerServiceClient) Update(ctx context.Context, in *UpdateSeller, opts ...grpc.CallOption) (*Seller, error) {
	out := new(Seller)
	err := c.cc.Invoke(ctx, SellerService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerServiceClient) Delete(ctx context.Context, in *SellerPrimaryKey, opts ...grpc.CallOption) (*Empty2, error) {
	out := new(Empty2)
	err := c.cc.Invoke(ctx, SellerService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SellerServiceServer is the server API for SellerService service.
// All implementations should embed UnimplementedSellerServiceServer
// for forward compatibility
type SellerServiceServer interface {
	Create(context.Context, *CreateSeller) (*Seller, error)
	GetByID(context.Context, *SellerPrimaryKey) (*Seller, error)
	GetList(context.Context, *GetListSellerRequest) (*GetListSellerResponse, error)
	Update(context.Context, *UpdateSeller) (*Seller, error)
	Delete(context.Context, *SellerPrimaryKey) (*Empty2, error)
}

// UnimplementedSellerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSellerServiceServer struct {
}

func (UnimplementedSellerServiceServer) Create(context.Context, *CreateSeller) (*Seller, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSellerServiceServer) GetByID(context.Context, *SellerPrimaryKey) (*Seller, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedSellerServiceServer) GetList(context.Context, *GetListSellerRequest) (*GetListSellerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedSellerServiceServer) Update(context.Context, *UpdateSeller) (*Seller, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSellerServiceServer) Delete(context.Context, *SellerPrimaryKey) (*Empty2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeSellerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SellerServiceServer will
// result in compilation errors.
type UnsafeSellerServiceServer interface {
	mustEmbedUnimplementedSellerServiceServer()
}

func RegisterSellerServiceServer(s grpc.ServiceRegistrar, srv SellerServiceServer) {
	s.RegisterService(&SellerService_ServiceDesc, srv)
}

func _SellerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSeller)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).Create(ctx, req.(*CreateSeller))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).GetByID(ctx, req.(*SellerPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListSellerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).GetList(ctx, req.(*GetListSellerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSeller)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).Update(ctx, req.(*UpdateSeller))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).Delete(ctx, req.(*SellerPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// SellerService_ServiceDesc is the grpc.ServiceDesc for SellerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SellerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.SellerService",
	HandlerType: (*SellerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SellerService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _SellerService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _SellerService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SellerService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SellerService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seller.proto",
}
