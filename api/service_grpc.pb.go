// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: service.proto

package api

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
	Expenses_CreateExpense_FullMethodName = "/expense.Expenses/CreateExpense"
	Expenses_GetBalances_FullMethodName   = "/expense.Expenses/GetBalances"
)

// ExpensesClient is the client API for Expenses service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Expense is a service for registering and tracking expenses.
type ExpensesClient interface {
	// CreateExpense registers an expense.
	CreateExpense(ctx context.Context, in *CreateExpenseRequest, opts ...grpc.CallOption) (*CreateExpenseResponse, error)
	// GetBalances periodically streams each user's balance in a round-robin manner.
	GetBalances(ctx context.Context, in *GetBalancesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetBalancesResponse], error)
}

type expensesClient struct {
	cc grpc.ClientConnInterface
}

func NewExpensesClient(cc grpc.ClientConnInterface) ExpensesClient {
	return &expensesClient{cc}
}

func (c *expensesClient) CreateExpense(ctx context.Context, in *CreateExpenseRequest, opts ...grpc.CallOption) (*CreateExpenseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateExpenseResponse)
	err := c.cc.Invoke(ctx, Expenses_CreateExpense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expensesClient) GetBalances(ctx context.Context, in *GetBalancesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetBalancesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Expenses_ServiceDesc.Streams[0], Expenses_GetBalances_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetBalancesRequest, GetBalancesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Expenses_GetBalancesClient = grpc.ServerStreamingClient[GetBalancesResponse]

// ExpensesServer is the server API for Expenses service.
// All implementations must embed UnimplementedExpensesServer
// for forward compatibility.
//
// Expense is a service for registering and tracking expenses.
type ExpensesServer interface {
	// CreateExpense registers an expense.
	CreateExpense(context.Context, *CreateExpenseRequest) (*CreateExpenseResponse, error)
	// GetBalances periodically streams each user's balance in a round-robin manner.
	GetBalances(*GetBalancesRequest, grpc.ServerStreamingServer[GetBalancesResponse]) error
	mustEmbedUnimplementedExpensesServer()
}

// UnimplementedExpensesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExpensesServer struct{}

func (UnimplementedExpensesServer) CreateExpense(context.Context, *CreateExpenseRequest) (*CreateExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExpense not implemented")
}
func (UnimplementedExpensesServer) GetBalances(*GetBalancesRequest, grpc.ServerStreamingServer[GetBalancesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetBalances not implemented")
}
func (UnimplementedExpensesServer) mustEmbedUnimplementedExpensesServer() {}
func (UnimplementedExpensesServer) testEmbeddedByValue()                  {}

// UnsafeExpensesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpensesServer will
// result in compilation errors.
type UnsafeExpensesServer interface {
	mustEmbedUnimplementedExpensesServer()
}

func RegisterExpensesServer(s grpc.ServiceRegistrar, srv ExpensesServer) {
	// If the following call pancis, it indicates UnimplementedExpensesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Expenses_ServiceDesc, srv)
}

func _Expenses_CreateExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpensesServer).CreateExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Expenses_CreateExpense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpensesServer).CreateExpense(ctx, req.(*CreateExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Expenses_GetBalances_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBalancesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExpensesServer).GetBalances(m, &grpc.GenericServerStream[GetBalancesRequest, GetBalancesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Expenses_GetBalancesServer = grpc.ServerStreamingServer[GetBalancesResponse]

// Expenses_ServiceDesc is the grpc.ServiceDesc for Expenses service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Expenses_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "expense.Expenses",
	HandlerType: (*ExpensesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExpense",
			Handler:    _Expenses_CreateExpense_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBalances",
			Handler:       _Expenses_GetBalances_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
