// Code generated by protoc-gen-go-rpc. DO NOT EDIT.

package proto

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

// QuestionnaireServiceClient is the client API for QuestionnaireService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionnaireServiceClient interface {
	CreateQuestionnaire(ctx context.Context, in *CreateQuestionnaireRequest, opts ...grpc.CallOption) (*CreateQuestionnaireResponse, error)
	GetQuestionnaire(ctx context.Context, in *GetQuestionnaireRequest, opts ...grpc.CallOption) (*GetQuestionnaireResponse, error)
	UpdateQuestionnaire(ctx context.Context, in *UpdateQuestionnaireRequest, opts ...grpc.CallOption) (*UpdateQuestionnaireResponse, error)
}

type questionnaireServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionnaireServiceClient(cc grpc.ClientConnInterface) QuestionnaireServiceClient {
	return &questionnaireServiceClient{cc}
}

func (c *questionnaireServiceClient) CreateQuestionnaire(ctx context.Context, in *CreateQuestionnaireRequest, opts ...grpc.CallOption) (*CreateQuestionnaireResponse, error) {
	out := new(CreateQuestionnaireResponse)
	err := c.cc.Invoke(ctx, "/proto.QuestionnaireService/CreateQuestionnaire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireServiceClient) GetQuestionnaire(ctx context.Context, in *GetQuestionnaireRequest, opts ...grpc.CallOption) (*GetQuestionnaireResponse, error) {
	out := new(GetQuestionnaireResponse)
	err := c.cc.Invoke(ctx, "/proto.QuestionnaireService/GetQuestionnaire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireServiceClient) UpdateQuestionnaire(ctx context.Context, in *UpdateQuestionnaireRequest, opts ...grpc.CallOption) (*UpdateQuestionnaireResponse, error) {
	out := new(UpdateQuestionnaireResponse)
	err := c.cc.Invoke(ctx, "/proto.QuestionnaireService/UpdateQuestionnaire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionnaireServiceServer is the server API for QuestionnaireService service.
// All implementations must embed UnimplementedQuestionnaireServiceServer
// for forward compatibility
type QuestionnaireServiceServer interface {
	CreateQuestionnaire(context.Context, *CreateQuestionnaireRequest) (*CreateQuestionnaireResponse, error)
	GetQuestionnaire(context.Context, *GetQuestionnaireRequest) (*GetQuestionnaireResponse, error)
	UpdateQuestionnaire(context.Context, *UpdateQuestionnaireRequest) (*UpdateQuestionnaireResponse, error)
	mustEmbedUnimplementedQuestionnaireServiceServer()
}

// UnimplementedQuestionnaireServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuestionnaireServiceServer struct {
}

func (UnimplementedQuestionnaireServiceServer) CreateQuestionnaire(context.Context, *CreateQuestionnaireRequest) (*CreateQuestionnaireResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestionnaire not implemented")
}
func (UnimplementedQuestionnaireServiceServer) GetQuestionnaire(context.Context, *GetQuestionnaireRequest) (*GetQuestionnaireResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionnaire not implemented")
}
func (UnimplementedQuestionnaireServiceServer) UpdateQuestionnaire(context.Context, *UpdateQuestionnaireRequest) (*UpdateQuestionnaireResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuestionnaire not implemented")
}
func (UnimplementedQuestionnaireServiceServer) mustEmbedUnimplementedQuestionnaireServiceServer() {}

// UnsafeQuestionnaireServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionnaireServiceServer will
// result in compilation errors.
type UnsafeQuestionnaireServiceServer interface {
	mustEmbedUnimplementedQuestionnaireServiceServer()
}

func RegisterQuestionnaireServiceServer(s grpc.ServiceRegistrar, srv QuestionnaireServiceServer) {
	s.RegisterService(&QuestionnaireService_ServiceDesc, srv)
}

func _QuestionnaireService_CreateQuestionnaire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuestionnaireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServiceServer).CreateQuestionnaire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.QuestionnaireService/CreateQuestionnaire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServiceServer).CreateQuestionnaire(ctx, req.(*CreateQuestionnaireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionnaireService_GetQuestionnaire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionnaireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServiceServer).GetQuestionnaire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.QuestionnaireService/GetQuestionnaire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServiceServer).GetQuestionnaire(ctx, req.(*GetQuestionnaireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionnaireService_UpdateQuestionnaire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuestionnaireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServiceServer).UpdateQuestionnaire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.QuestionnaireService/UpdateQuestionnaire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServiceServer).UpdateQuestionnaire(ctx, req.(*UpdateQuestionnaireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuestionnaireService_ServiceDesc is the grpc.ServiceDesc for QuestionnaireService service.
// It's only intended for direct use with rpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuestionnaireService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.QuestionnaireService",
	HandlerType: (*QuestionnaireServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuestionnaire",
			Handler:    _QuestionnaireService_CreateQuestionnaire_Handler,
		},
		{
			MethodName: "GetQuestionnaire",
			Handler:    _QuestionnaireService_GetQuestionnaire_Handler,
		},
		{
			MethodName: "UpdateQuestionnaire",
			Handler:    _QuestionnaireService_UpdateQuestionnaire_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "questionnaire.proto",
}
