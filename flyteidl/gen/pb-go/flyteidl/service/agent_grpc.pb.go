// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: flyteidl/service/agent.proto

package service

import (
	context "context"

	admin "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AsyncAgentService_CreateTask_FullMethodName     = "/flyteidl.service.AsyncAgentService/CreateTask"
	AsyncAgentService_GetTask_FullMethodName        = "/flyteidl.service.AsyncAgentService/GetTask"
	AsyncAgentService_DeleteTask_FullMethodName     = "/flyteidl.service.AsyncAgentService/DeleteTask"
	AsyncAgentService_GetTaskMetrics_FullMethodName = "/flyteidl.service.AsyncAgentService/GetTaskMetrics"
	AsyncAgentService_GetTaskLogs_FullMethodName    = "/flyteidl.service.AsyncAgentService/GetTaskLogs"
)

// AsyncAgentServiceClient is the client API for AsyncAgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AsyncAgentServiceClient interface {
	// Send a task create request to the agent server.
	CreateTask(ctx context.Context, in *admin.CreateTaskRequest, opts ...grpc.CallOption) (*admin.CreateTaskResponse, error)
	// Get job status.
	GetTask(ctx context.Context, in *admin.GetTaskRequest, opts ...grpc.CallOption) (*admin.GetTaskResponse, error)
	// Delete the task resource.
	DeleteTask(ctx context.Context, in *admin.DeleteTaskRequest, opts ...grpc.CallOption) (*admin.DeleteTaskResponse, error)
	// GetTaskMetrics returns one or more task execution metrics, if available.
	//
	// Errors include
	//   - OutOfRange if metrics are not available for the specified task time range
	//   - various other errors
	GetTaskMetrics(ctx context.Context, in *admin.GetTaskMetricsRequest, opts ...grpc.CallOption) (*admin.GetTaskMetricsResponse, error)
	// GetTaskLogs returns task execution logs, if available.
	GetTaskLogs(ctx context.Context, in *admin.GetTaskLogsRequest, opts ...grpc.CallOption) (*admin.GetTaskLogsResponse, error)
}

type asyncAgentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAsyncAgentServiceClient(cc grpc.ClientConnInterface) AsyncAgentServiceClient {
	return &asyncAgentServiceClient{cc}
}

func (c *asyncAgentServiceClient) CreateTask(ctx context.Context, in *admin.CreateTaskRequest, opts ...grpc.CallOption) (*admin.CreateTaskResponse, error) {
	out := new(admin.CreateTaskResponse)
	err := c.cc.Invoke(ctx, AsyncAgentService_CreateTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asyncAgentServiceClient) GetTask(ctx context.Context, in *admin.GetTaskRequest, opts ...grpc.CallOption) (*admin.GetTaskResponse, error) {
	out := new(admin.GetTaskResponse)
	err := c.cc.Invoke(ctx, AsyncAgentService_GetTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asyncAgentServiceClient) DeleteTask(ctx context.Context, in *admin.DeleteTaskRequest, opts ...grpc.CallOption) (*admin.DeleteTaskResponse, error) {
	out := new(admin.DeleteTaskResponse)
	err := c.cc.Invoke(ctx, AsyncAgentService_DeleteTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asyncAgentServiceClient) GetTaskMetrics(ctx context.Context, in *admin.GetTaskMetricsRequest, opts ...grpc.CallOption) (*admin.GetTaskMetricsResponse, error) {
	out := new(admin.GetTaskMetricsResponse)
	err := c.cc.Invoke(ctx, AsyncAgentService_GetTaskMetrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asyncAgentServiceClient) GetTaskLogs(ctx context.Context, in *admin.GetTaskLogsRequest, opts ...grpc.CallOption) (*admin.GetTaskLogsResponse, error) {
	out := new(admin.GetTaskLogsResponse)
	err := c.cc.Invoke(ctx, AsyncAgentService_GetTaskLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AsyncAgentServiceServer is the server API for AsyncAgentService service.
// All implementations should embed UnimplementedAsyncAgentServiceServer
// for forward compatibility
type AsyncAgentServiceServer interface {
	// Send a task create request to the agent server.
	CreateTask(context.Context, *admin.CreateTaskRequest) (*admin.CreateTaskResponse, error)
	// Get job status.
	GetTask(context.Context, *admin.GetTaskRequest) (*admin.GetTaskResponse, error)
	// Delete the task resource.
	DeleteTask(context.Context, *admin.DeleteTaskRequest) (*admin.DeleteTaskResponse, error)
	// GetTaskMetrics returns one or more task execution metrics, if available.
	//
	// Errors include
	//   - OutOfRange if metrics are not available for the specified task time range
	//   - various other errors
	GetTaskMetrics(context.Context, *admin.GetTaskMetricsRequest) (*admin.GetTaskMetricsResponse, error)
	// GetTaskLogs returns task execution logs, if available.
	GetTaskLogs(context.Context, *admin.GetTaskLogsRequest) (*admin.GetTaskLogsResponse, error)
}

// UnimplementedAsyncAgentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAsyncAgentServiceServer struct {
}

func (UnimplementedAsyncAgentServiceServer) CreateTask(context.Context, *admin.CreateTaskRequest) (*admin.CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedAsyncAgentServiceServer) GetTask(context.Context, *admin.GetTaskRequest) (*admin.GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedAsyncAgentServiceServer) DeleteTask(context.Context, *admin.DeleteTaskRequest) (*admin.DeleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedAsyncAgentServiceServer) GetTaskMetrics(context.Context, *admin.GetTaskMetricsRequest) (*admin.GetTaskMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskMetrics not implemented")
}
func (UnimplementedAsyncAgentServiceServer) GetTaskLogs(context.Context, *admin.GetTaskLogsRequest) (*admin.GetTaskLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskLogs not implemented")
}

// UnsafeAsyncAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AsyncAgentServiceServer will
// result in compilation errors.
type UnsafeAsyncAgentServiceServer interface {
	mustEmbedUnimplementedAsyncAgentServiceServer()
}

func RegisterAsyncAgentServiceServer(s grpc.ServiceRegistrar, srv AsyncAgentServiceServer) {
	s.RegisterService(&AsyncAgentService_ServiceDesc, srv)
}

func _AsyncAgentService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsyncAgentServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsyncAgentService_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsyncAgentServiceServer).CreateTask(ctx, req.(*admin.CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsyncAgentService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsyncAgentServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsyncAgentService_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsyncAgentServiceServer).GetTask(ctx, req.(*admin.GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsyncAgentService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsyncAgentServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsyncAgentService_DeleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsyncAgentServiceServer).DeleteTask(ctx, req.(*admin.DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsyncAgentService_GetTaskMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.GetTaskMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsyncAgentServiceServer).GetTaskMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsyncAgentService_GetTaskMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsyncAgentServiceServer).GetTaskMetrics(ctx, req.(*admin.GetTaskMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsyncAgentService_GetTaskLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.GetTaskLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsyncAgentServiceServer).GetTaskLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsyncAgentService_GetTaskLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsyncAgentServiceServer).GetTaskLogs(ctx, req.(*admin.GetTaskLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AsyncAgentService_ServiceDesc is the grpc.ServiceDesc for AsyncAgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AsyncAgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.AsyncAgentService",
	HandlerType: (*AsyncAgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _AsyncAgentService_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _AsyncAgentService_GetTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _AsyncAgentService_DeleteTask_Handler,
		},
		{
			MethodName: "GetTaskMetrics",
			Handler:    _AsyncAgentService_GetTaskMetrics_Handler,
		},
		{
			MethodName: "GetTaskLogs",
			Handler:    _AsyncAgentService_GetTaskLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/agent.proto",
}

const (
	AgentMetadataService_GetAgent_FullMethodName   = "/flyteidl.service.AgentMetadataService/GetAgent"
	AgentMetadataService_ListAgents_FullMethodName = "/flyteidl.service.AgentMetadataService/ListAgents"
)

// AgentMetadataServiceClient is the client API for AgentMetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentMetadataServiceClient interface {
	// Fetch a :ref:`ref_flyteidl.admin.Agent` definition.
	GetAgent(ctx context.Context, in *admin.GetAgentRequest, opts ...grpc.CallOption) (*admin.GetAgentResponse, error)
	// Fetch a list of :ref:`ref_flyteidl.admin.Agent` definitions.
	ListAgents(ctx context.Context, in *admin.ListAgentsRequest, opts ...grpc.CallOption) (*admin.ListAgentsResponse, error)
}

type agentMetadataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentMetadataServiceClient(cc grpc.ClientConnInterface) AgentMetadataServiceClient {
	return &agentMetadataServiceClient{cc}
}

func (c *agentMetadataServiceClient) GetAgent(ctx context.Context, in *admin.GetAgentRequest, opts ...grpc.CallOption) (*admin.GetAgentResponse, error) {
	out := new(admin.GetAgentResponse)
	err := c.cc.Invoke(ctx, AgentMetadataService_GetAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentMetadataServiceClient) ListAgents(ctx context.Context, in *admin.ListAgentsRequest, opts ...grpc.CallOption) (*admin.ListAgentsResponse, error) {
	out := new(admin.ListAgentsResponse)
	err := c.cc.Invoke(ctx, AgentMetadataService_ListAgents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentMetadataServiceServer is the server API for AgentMetadataService service.
// All implementations should embed UnimplementedAgentMetadataServiceServer
// for forward compatibility
type AgentMetadataServiceServer interface {
	// Fetch a :ref:`ref_flyteidl.admin.Agent` definition.
	GetAgent(context.Context, *admin.GetAgentRequest) (*admin.GetAgentResponse, error)
	// Fetch a list of :ref:`ref_flyteidl.admin.Agent` definitions.
	ListAgents(context.Context, *admin.ListAgentsRequest) (*admin.ListAgentsResponse, error)
}

// UnimplementedAgentMetadataServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAgentMetadataServiceServer struct {
}

func (UnimplementedAgentMetadataServiceServer) GetAgent(context.Context, *admin.GetAgentRequest) (*admin.GetAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgent not implemented")
}
func (UnimplementedAgentMetadataServiceServer) ListAgents(context.Context, *admin.ListAgentsRequest) (*admin.ListAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgents not implemented")
}

// UnsafeAgentMetadataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentMetadataServiceServer will
// result in compilation errors.
type UnsafeAgentMetadataServiceServer interface {
	mustEmbedUnimplementedAgentMetadataServiceServer()
}

func RegisterAgentMetadataServiceServer(s grpc.ServiceRegistrar, srv AgentMetadataServiceServer) {
	s.RegisterService(&AgentMetadataService_ServiceDesc, srv)
}

func _AgentMetadataService_GetAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.GetAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentMetadataServiceServer).GetAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentMetadataService_GetAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentMetadataServiceServer).GetAgent(ctx, req.(*admin.GetAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentMetadataService_ListAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.ListAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentMetadataServiceServer).ListAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentMetadataService_ListAgents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentMetadataServiceServer).ListAgents(ctx, req.(*admin.ListAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentMetadataService_ServiceDesc is the grpc.ServiceDesc for AgentMetadataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentMetadataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.AgentMetadataService",
	HandlerType: (*AgentMetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgent",
			Handler:    _AgentMetadataService_GetAgent_Handler,
		},
		{
			MethodName: "ListAgents",
			Handler:    _AgentMetadataService_ListAgents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/agent.proto",
}