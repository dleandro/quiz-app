// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: quiz.proto

package quizz

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	QuizService_GetQuestions_FullMethodName   = "/protofiles.QuizService/GetQuestions"
	QuizService_SaveResults_FullMethodName    = "/protofiles.QuizService/SaveResults"
	QuizService_GetStatistics_FullMethodName  = "/protofiles.QuizService/GetStatistics"
	QuizService_CreateQuestion_FullMethodName = "/protofiles.QuizService/CreateQuestion"
	QuizService_DeleteQuestion_FullMethodName = "/protofiles.QuizService/DeleteQuestion"
)

// QuizServiceClient is the client API for QuizService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuizServiceClient interface {
	GetQuestions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*QuestionsResponse, error)
	SaveResults(ctx context.Context, in *ResultsRequest, opts ...grpc.CallOption) (*ResultsResponse, error)
	GetStatistics(ctx context.Context, in *ResultsRequest, opts ...grpc.CallOption) (*StatisticsResponse, error)
	CreateQuestion(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*CreateQuestionResponse, error)
	DeleteQuestion(ctx context.Context, in *DeleteQuestionRequest, opts ...grpc.CallOption) (*DeleteQuestionResponse, error)
}

type quizServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuizServiceClient(cc grpc.ClientConnInterface) QuizServiceClient {
	return &quizServiceClient{cc}
}

func (c *quizServiceClient) GetQuestions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*QuestionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuestionsResponse)
	err := c.cc.Invoke(ctx, QuizService_GetQuestions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) SaveResults(ctx context.Context, in *ResultsRequest, opts ...grpc.CallOption) (*ResultsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResultsResponse)
	err := c.cc.Invoke(ctx, QuizService_SaveResults_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) GetStatistics(ctx context.Context, in *ResultsRequest, opts ...grpc.CallOption) (*StatisticsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatisticsResponse)
	err := c.cc.Invoke(ctx, QuizService_GetStatistics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) CreateQuestion(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*CreateQuestionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateQuestionResponse)
	err := c.cc.Invoke(ctx, QuizService_CreateQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) DeleteQuestion(ctx context.Context, in *DeleteQuestionRequest, opts ...grpc.CallOption) (*DeleteQuestionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteQuestionResponse)
	err := c.cc.Invoke(ctx, QuizService_DeleteQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuizServiceServer is the server API for QuizService service.
// All implementations must embed UnimplementedQuizServiceServer
// for forward compatibility.
type QuizServiceServer interface {
	GetQuestions(context.Context, *Empty) (*QuestionsResponse, error)
	SaveResults(context.Context, *ResultsRequest) (*ResultsResponse, error)
	GetStatistics(context.Context, *ResultsRequest) (*StatisticsResponse, error)
	CreateQuestion(context.Context, *CreateQuestionRequest) (*CreateQuestionResponse, error)
	DeleteQuestion(context.Context, *DeleteQuestionRequest) (*DeleteQuestionResponse, error)
	mustEmbedUnimplementedQuizServiceServer()
}

// UnimplementedQuizServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQuizServiceServer struct{}

func (UnimplementedQuizServiceServer) GetQuestions(context.Context, *Empty) (*QuestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestions not implemented")
}
func (UnimplementedQuizServiceServer) SaveResults(context.Context, *ResultsRequest) (*ResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveResults not implemented")
}
func (UnimplementedQuizServiceServer) GetStatistics(context.Context, *ResultsRequest) (*StatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatistics not implemented")
}
func (UnimplementedQuizServiceServer) CreateQuestion(context.Context, *CreateQuestionRequest) (*CreateQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestion not implemented")
}
func (UnimplementedQuizServiceServer) DeleteQuestion(context.Context, *DeleteQuestionRequest) (*DeleteQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuestion not implemented")
}
func (UnimplementedQuizServiceServer) mustEmbedUnimplementedQuizServiceServer() {}
func (UnimplementedQuizServiceServer) testEmbeddedByValue()                     {}

// UnsafeQuizServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuizServiceServer will
// result in compilation errors.
type UnsafeQuizServiceServer interface {
	mustEmbedUnimplementedQuizServiceServer()
}

func RegisterQuizServiceServer(s grpc.ServiceRegistrar, srv QuizServiceServer) {
	// If the following call pancis, it indicates UnimplementedQuizServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&QuizService_ServiceDesc, srv)
}

func _QuizService_GetQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).GetQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_GetQuestions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).GetQuestions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_SaveResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).SaveResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_SaveResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).SaveResults(ctx, req.(*ResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_GetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).GetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_GetStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).GetStatistics(ctx, req.(*ResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_CreateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).CreateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_CreateQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).CreateQuestion(ctx, req.(*CreateQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_DeleteQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).DeleteQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_DeleteQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).DeleteQuestion(ctx, req.(*DeleteQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuizService_ServiceDesc is the grpc.ServiceDesc for QuizService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuizService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.QuizService",
	HandlerType: (*QuizServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuestions",
			Handler:    _QuizService_GetQuestions_Handler,
		},
		{
			MethodName: "SaveResults",
			Handler:    _QuizService_SaveResults_Handler,
		},
		{
			MethodName: "GetStatistics",
			Handler:    _QuizService_GetStatistics_Handler,
		},
		{
			MethodName: "CreateQuestion",
			Handler:    _QuizService_CreateQuestion_Handler,
		},
		{
			MethodName: "DeleteQuestion",
			Handler:    _QuizService_DeleteQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quiz.proto",
}
