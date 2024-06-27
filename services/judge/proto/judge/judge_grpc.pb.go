// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: judge.proto

package judge

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

// JudgeServiceClient is the client API for JudgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JudgeServiceClient interface {
	GetTestCase(ctx context.Context, in *GetTestCasesRequest, opts ...grpc.CallOption) (*GetTestCasesResponse, error)
	RunTests(ctx context.Context, in *RunTestsRequest, opts ...grpc.CallOption) (*RunTestsResponse, error)
}

type judgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJudgeServiceClient(cc grpc.ClientConnInterface) JudgeServiceClient {
	return &judgeServiceClient{cc}
}

func (c *judgeServiceClient) GetTestCase(ctx context.Context, in *GetTestCasesRequest, opts ...grpc.CallOption) (*GetTestCasesResponse, error) {
	out := new(GetTestCasesResponse)
	err := c.cc.Invoke(ctx, "/judge.JudgeService/GetTestCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeServiceClient) RunTests(ctx context.Context, in *RunTestsRequest, opts ...grpc.CallOption) (*RunTestsResponse, error) {
	out := new(RunTestsResponse)
	err := c.cc.Invoke(ctx, "/judge.JudgeService/RunTests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JudgeServiceServer is the server API for JudgeService service.
// All implementations must embed UnimplementedJudgeServiceServer
// for forward compatibility
type JudgeServiceServer interface {
	GetTestCase(context.Context, *GetTestCasesRequest) (*GetTestCasesResponse, error)
	RunTests(context.Context, *RunTestsRequest) (*RunTestsResponse, error)
	mustEmbedUnimplementedJudgeServiceServer()
}

// UnimplementedJudgeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJudgeServiceServer struct {
}

func (UnimplementedJudgeServiceServer) GetTestCase(context.Context, *GetTestCasesRequest) (*GetTestCasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestCase not implemented")
}
func (UnimplementedJudgeServiceServer) RunTests(context.Context, *RunTestsRequest) (*RunTestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunTests not implemented")
}
func (UnimplementedJudgeServiceServer) mustEmbedUnimplementedJudgeServiceServer() {}

// UnsafeJudgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JudgeServiceServer will
// result in compilation errors.
type UnsafeJudgeServiceServer interface {
	mustEmbedUnimplementedJudgeServiceServer()
}

func RegisterJudgeServiceServer(s grpc.ServiceRegistrar, srv JudgeServiceServer) {
	s.RegisterService(&JudgeService_ServiceDesc, srv)
}

func _JudgeService_GetTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestCasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServiceServer).GetTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judge.JudgeService/GetTestCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServiceServer).GetTestCase(ctx, req.(*GetTestCasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JudgeService_RunTests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunTestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServiceServer).RunTests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judge.JudgeService/RunTests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServiceServer).RunTests(ctx, req.(*RunTestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JudgeService_ServiceDesc is the grpc.ServiceDesc for JudgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JudgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "judge.JudgeService",
	HandlerType: (*JudgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTestCase",
			Handler:    _JudgeService_GetTestCase_Handler,
		},
		{
			MethodName: "RunTests",
			Handler:    _JudgeService_RunTests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "judge.proto",
}
