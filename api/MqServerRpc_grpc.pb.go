// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: MqServerRpc.proto

package api

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

// MqServerCallClient is the client API for MqServerCall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MqServerCallClient interface {
	JoinConsumerGroup(ctx context.Context, in *JoinConsumerGroupRequest, opts ...grpc.CallOption) (*JoinConsumerGroupResponse, error)
	LeaveConsumerGroup(ctx context.Context, in *LeaveConsumerGroupRequest, opts ...grpc.CallOption) (*LeaveConsumerGroupResponse, error)
	CheckSourceTerm(ctx context.Context, in *CheckSourceTermRequest, opts ...grpc.CallOption) (*CheckSourceTermResponse, error)
	// 订阅
	SubscribeTopic(ctx context.Context, in *SubscribeTopicRequest, opts ...grpc.CallOption) (*SubscribeTopicResponse, error)
	UnSubscribeTopic(ctx context.Context, in *UnSubscribeTopicRequest, opts ...grpc.CallOption) (*UnSubscribeTopicResponse, error)
	// 注册消费者组
	RegisterConsumerGroup(ctx context.Context, in *RegisterConsumerGroupRequest, opts ...grpc.CallOption) (*RegisterConsumerGroupResponse, error)
	// 创建话题
	CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicResponse, error)
	QueryTopic(ctx context.Context, in *QueryTopicRequest, opts ...grpc.CallOption) (*QueryTopicResponse, error)
	// Add / Del Part
	AddPart(ctx context.Context, in *AddPartRequest, opts ...grpc.CallOption) (*AddPartResponse, error)
	RemovePart(ctx context.Context, in *RemovePartRequest, opts ...grpc.CallOption) (*RemovePartResponse, error)
	// 注册
	RegisterConsumer(ctx context.Context, in *RegisterConsumerRequest, opts ...grpc.CallOption) (*RegisterConsumerResponse, error)
	RegisterProducer(ctx context.Context, in *RegisterProducerRequest, opts ...grpc.CallOption) (*RegisterProducerResponse, error)
	// 注销
	UnRegisterConsumer(ctx context.Context, in *UnRegisterConsumerRequest, opts ...grpc.CallOption) (*UnRegisterConsumerResponse, error)
	UnRegisterProducer(ctx context.Context, in *UnRegisterProducerRequest, opts ...grpc.CallOption) (*UnRegisterProducerResponse, error)
	// 客户端和server之间的心跳
	Heartbeat(ctx context.Context, in *MQHeartBeatData, opts ...grpc.CallOption) (*HeartBeatResponseData, error)
	ConsumerDisConnect(ctx context.Context, in *DisConnectInfo, opts ...grpc.CallOption) (*Response, error)
	ProducerDisConnect(ctx context.Context, in *DisConnectInfo, opts ...grpc.CallOption) (*Response, error)
	// 拉取消息
	PullMessage(ctx context.Context, in *PullMessageRequest, opts ...grpc.CallOption) (*PullMessageResponse, error)
	// 推送消息
	PushMessage(ctx context.Context, in *PushMessageRequest, opts ...grpc.CallOption) (*PushMessageResponse, error)
	ConfirmIdentity(ctx context.Context, in *ConfirmIdentityRequest, opts ...grpc.CallOption) (*ConfirmIdentityResponse, error)
	RegisterBroker(ctx context.Context, in *RegisterBrokerRequest, opts ...grpc.CallOption) (*RegisterBrokerResponse, error)
}

type mqServerCallClient struct {
	cc grpc.ClientConnInterface
}

func NewMqServerCallClient(cc grpc.ClientConnInterface) MqServerCallClient {
	return &mqServerCallClient{cc}
}

func (c *mqServerCallClient) JoinConsumerGroup(ctx context.Context, in *JoinConsumerGroupRequest, opts ...grpc.CallOption) (*JoinConsumerGroupResponse, error) {
	out := new(JoinConsumerGroupResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/JoinConsumerGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) LeaveConsumerGroup(ctx context.Context, in *LeaveConsumerGroupRequest, opts ...grpc.CallOption) (*LeaveConsumerGroupResponse, error) {
	out := new(LeaveConsumerGroupResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/LeaveConsumerGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) CheckSourceTerm(ctx context.Context, in *CheckSourceTermRequest, opts ...grpc.CallOption) (*CheckSourceTermResponse, error) {
	out := new(CheckSourceTermResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/CheckSourceTerm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) SubscribeTopic(ctx context.Context, in *SubscribeTopicRequest, opts ...grpc.CallOption) (*SubscribeTopicResponse, error) {
	out := new(SubscribeTopicResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/SubscribeTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) UnSubscribeTopic(ctx context.Context, in *UnSubscribeTopicRequest, opts ...grpc.CallOption) (*UnSubscribeTopicResponse, error) {
	out := new(UnSubscribeTopicResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/UnSubscribeTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) RegisterConsumerGroup(ctx context.Context, in *RegisterConsumerGroupRequest, opts ...grpc.CallOption) (*RegisterConsumerGroupResponse, error) {
	out := new(RegisterConsumerGroupResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/RegisterConsumerGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicResponse, error) {
	out := new(CreateTopicResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/CreateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) QueryTopic(ctx context.Context, in *QueryTopicRequest, opts ...grpc.CallOption) (*QueryTopicResponse, error) {
	out := new(QueryTopicResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/QueryTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) AddPart(ctx context.Context, in *AddPartRequest, opts ...grpc.CallOption) (*AddPartResponse, error) {
	out := new(AddPartResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/AddPart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) RemovePart(ctx context.Context, in *RemovePartRequest, opts ...grpc.CallOption) (*RemovePartResponse, error) {
	out := new(RemovePartResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/RemovePart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) RegisterConsumer(ctx context.Context, in *RegisterConsumerRequest, opts ...grpc.CallOption) (*RegisterConsumerResponse, error) {
	out := new(RegisterConsumerResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/RegisterConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) RegisterProducer(ctx context.Context, in *RegisterProducerRequest, opts ...grpc.CallOption) (*RegisterProducerResponse, error) {
	out := new(RegisterProducerResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/RegisterProducer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) UnRegisterConsumer(ctx context.Context, in *UnRegisterConsumerRequest, opts ...grpc.CallOption) (*UnRegisterConsumerResponse, error) {
	out := new(UnRegisterConsumerResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/UnRegisterConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) UnRegisterProducer(ctx context.Context, in *UnRegisterProducerRequest, opts ...grpc.CallOption) (*UnRegisterProducerResponse, error) {
	out := new(UnRegisterProducerResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/UnRegisterProducer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) Heartbeat(ctx context.Context, in *MQHeartBeatData, opts ...grpc.CallOption) (*HeartBeatResponseData, error) {
	out := new(HeartBeatResponseData)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) ConsumerDisConnect(ctx context.Context, in *DisConnectInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/ConsumerDisConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) ProducerDisConnect(ctx context.Context, in *DisConnectInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/ProducerDisConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) PullMessage(ctx context.Context, in *PullMessageRequest, opts ...grpc.CallOption) (*PullMessageResponse, error) {
	out := new(PullMessageResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/PullMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) PushMessage(ctx context.Context, in *PushMessageRequest, opts ...grpc.CallOption) (*PushMessageResponse, error) {
	out := new(PushMessageResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/PushMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) ConfirmIdentity(ctx context.Context, in *ConfirmIdentityRequest, opts ...grpc.CallOption) (*ConfirmIdentityResponse, error) {
	out := new(ConfirmIdentityResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/ConfirmIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqServerCallClient) RegisterBroker(ctx context.Context, in *RegisterBrokerRequest, opts ...grpc.CallOption) (*RegisterBrokerResponse, error) {
	out := new(RegisterBrokerResponse)
	err := c.cc.Invoke(ctx, "/MqServer.MqServerCall/RegisterBroker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MqServerCallServer is the server API for MqServerCall service.
// All implementations must embed UnimplementedMqServerCallServer
// for forward compatibility
type MqServerCallServer interface {
	JoinConsumerGroup(context.Context, *JoinConsumerGroupRequest) (*JoinConsumerGroupResponse, error)
	LeaveConsumerGroup(context.Context, *LeaveConsumerGroupRequest) (*LeaveConsumerGroupResponse, error)
	CheckSourceTerm(context.Context, *CheckSourceTermRequest) (*CheckSourceTermResponse, error)
	// 订阅
	SubscribeTopic(context.Context, *SubscribeTopicRequest) (*SubscribeTopicResponse, error)
	UnSubscribeTopic(context.Context, *UnSubscribeTopicRequest) (*UnSubscribeTopicResponse, error)
	// 注册消费者组
	RegisterConsumerGroup(context.Context, *RegisterConsumerGroupRequest) (*RegisterConsumerGroupResponse, error)
	// 创建话题
	CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicResponse, error)
	QueryTopic(context.Context, *QueryTopicRequest) (*QueryTopicResponse, error)
	// Add / Del Part
	AddPart(context.Context, *AddPartRequest) (*AddPartResponse, error)
	RemovePart(context.Context, *RemovePartRequest) (*RemovePartResponse, error)
	// 注册
	RegisterConsumer(context.Context, *RegisterConsumerRequest) (*RegisterConsumerResponse, error)
	RegisterProducer(context.Context, *RegisterProducerRequest) (*RegisterProducerResponse, error)
	// 注销
	UnRegisterConsumer(context.Context, *UnRegisterConsumerRequest) (*UnRegisterConsumerResponse, error)
	UnRegisterProducer(context.Context, *UnRegisterProducerRequest) (*UnRegisterProducerResponse, error)
	// 客户端和server之间的心跳
	Heartbeat(context.Context, *MQHeartBeatData) (*HeartBeatResponseData, error)
	ConsumerDisConnect(context.Context, *DisConnectInfo) (*Response, error)
	ProducerDisConnect(context.Context, *DisConnectInfo) (*Response, error)
	// 拉取消息
	PullMessage(context.Context, *PullMessageRequest) (*PullMessageResponse, error)
	// 推送消息
	PushMessage(context.Context, *PushMessageRequest) (*PushMessageResponse, error)
	ConfirmIdentity(context.Context, *ConfirmIdentityRequest) (*ConfirmIdentityResponse, error)
	RegisterBroker(context.Context, *RegisterBrokerRequest) (*RegisterBrokerResponse, error)
	mustEmbedUnimplementedMqServerCallServer()
}

// UnimplementedMqServerCallServer must be embedded to have forward compatible implementations.
type UnimplementedMqServerCallServer struct {
}

func (UnimplementedMqServerCallServer) JoinConsumerGroup(context.Context, *JoinConsumerGroupRequest) (*JoinConsumerGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinConsumerGroup not implemented")
}
func (UnimplementedMqServerCallServer) LeaveConsumerGroup(context.Context, *LeaveConsumerGroupRequest) (*LeaveConsumerGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveConsumerGroup not implemented")
}
func (UnimplementedMqServerCallServer) CheckSourceTerm(context.Context, *CheckSourceTermRequest) (*CheckSourceTermResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSourceTerm not implemented")
}
func (UnimplementedMqServerCallServer) SubscribeTopic(context.Context, *SubscribeTopicRequest) (*SubscribeTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeTopic not implemented")
}
func (UnimplementedMqServerCallServer) UnSubscribeTopic(context.Context, *UnSubscribeTopicRequest) (*UnSubscribeTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnSubscribeTopic not implemented")
}
func (UnimplementedMqServerCallServer) RegisterConsumerGroup(context.Context, *RegisterConsumerGroupRequest) (*RegisterConsumerGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterConsumerGroup not implemented")
}
func (UnimplementedMqServerCallServer) CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedMqServerCallServer) QueryTopic(context.Context, *QueryTopicRequest) (*QueryTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTopic not implemented")
}
func (UnimplementedMqServerCallServer) AddPart(context.Context, *AddPartRequest) (*AddPartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPart not implemented")
}
func (UnimplementedMqServerCallServer) RemovePart(context.Context, *RemovePartRequest) (*RemovePartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePart not implemented")
}
func (UnimplementedMqServerCallServer) RegisterConsumer(context.Context, *RegisterConsumerRequest) (*RegisterConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterConsumer not implemented")
}
func (UnimplementedMqServerCallServer) RegisterProducer(context.Context, *RegisterProducerRequest) (*RegisterProducerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterProducer not implemented")
}
func (UnimplementedMqServerCallServer) UnRegisterConsumer(context.Context, *UnRegisterConsumerRequest) (*UnRegisterConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegisterConsumer not implemented")
}
func (UnimplementedMqServerCallServer) UnRegisterProducer(context.Context, *UnRegisterProducerRequest) (*UnRegisterProducerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegisterProducer not implemented")
}
func (UnimplementedMqServerCallServer) Heartbeat(context.Context, *MQHeartBeatData) (*HeartBeatResponseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedMqServerCallServer) ConsumerDisConnect(context.Context, *DisConnectInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsumerDisConnect not implemented")
}
func (UnimplementedMqServerCallServer) ProducerDisConnect(context.Context, *DisConnectInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProducerDisConnect not implemented")
}
func (UnimplementedMqServerCallServer) PullMessage(context.Context, *PullMessageRequest) (*PullMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullMessage not implemented")
}
func (UnimplementedMqServerCallServer) PushMessage(context.Context, *PushMessageRequest) (*PushMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMessage not implemented")
}
func (UnimplementedMqServerCallServer) ConfirmIdentity(context.Context, *ConfirmIdentityRequest) (*ConfirmIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmIdentity not implemented")
}
func (UnimplementedMqServerCallServer) RegisterBroker(context.Context, *RegisterBrokerRequest) (*RegisterBrokerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterBroker not implemented")
}
func (UnimplementedMqServerCallServer) mustEmbedUnimplementedMqServerCallServer() {}

// UnsafeMqServerCallServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MqServerCallServer will
// result in compilation errors.
type UnsafeMqServerCallServer interface {
	mustEmbedUnimplementedMqServerCallServer()
}

func RegisterMqServerCallServer(s grpc.ServiceRegistrar, srv MqServerCallServer) {
	s.RegisterService(&MqServerCall_ServiceDesc, srv)
}

func _MqServerCall_JoinConsumerGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinConsumerGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).JoinConsumerGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/JoinConsumerGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).JoinConsumerGroup(ctx, req.(*JoinConsumerGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_LeaveConsumerGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveConsumerGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).LeaveConsumerGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/LeaveConsumerGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).LeaveConsumerGroup(ctx, req.(*LeaveConsumerGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_CheckSourceTerm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSourceTermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).CheckSourceTerm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/CheckSourceTerm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).CheckSourceTerm(ctx, req.(*CheckSourceTermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_SubscribeTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).SubscribeTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/SubscribeTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).SubscribeTopic(ctx, req.(*SubscribeTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_UnSubscribeTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnSubscribeTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).UnSubscribeTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/UnSubscribeTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).UnSubscribeTopic(ctx, req.(*UnSubscribeTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_RegisterConsumerGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterConsumerGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).RegisterConsumerGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/RegisterConsumerGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).RegisterConsumerGroup(ctx, req.(*RegisterConsumerGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).CreateTopic(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_QueryTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).QueryTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/QueryTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).QueryTopic(ctx, req.(*QueryTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_AddPart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).AddPart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/AddPart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).AddPart(ctx, req.(*AddPartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_RemovePart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).RemovePart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/RemovePart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).RemovePart(ctx, req.(*RemovePartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_RegisterConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).RegisterConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/RegisterConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).RegisterConsumer(ctx, req.(*RegisterConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_RegisterProducer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterProducerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).RegisterProducer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/RegisterProducer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).RegisterProducer(ctx, req.(*RegisterProducerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_UnRegisterConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnRegisterConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).UnRegisterConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/UnRegisterConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).UnRegisterConsumer(ctx, req.(*UnRegisterConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_UnRegisterProducer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnRegisterProducerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).UnRegisterProducer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/UnRegisterProducer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).UnRegisterProducer(ctx, req.(*UnRegisterProducerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MQHeartBeatData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).Heartbeat(ctx, req.(*MQHeartBeatData))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_ConsumerDisConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisConnectInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).ConsumerDisConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/ConsumerDisConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).ConsumerDisConnect(ctx, req.(*DisConnectInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_ProducerDisConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisConnectInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).ProducerDisConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/ProducerDisConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).ProducerDisConnect(ctx, req.(*DisConnectInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_PullMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).PullMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/PullMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).PullMessage(ctx, req.(*PullMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_PushMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).PushMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/PushMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).PushMessage(ctx, req.(*PushMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_ConfirmIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).ConfirmIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/ConfirmIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).ConfirmIdentity(ctx, req.(*ConfirmIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqServerCall_RegisterBroker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterBrokerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServerCallServer).RegisterBroker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqServer.MqServerCall/RegisterBroker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServerCallServer).RegisterBroker(ctx, req.(*RegisterBrokerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MqServerCall_ServiceDesc is the grpc.ServiceDesc for MqServerCall service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MqServerCall_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MqServer.MqServerCall",
	HandlerType: (*MqServerCallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinConsumerGroup",
			Handler:    _MqServerCall_JoinConsumerGroup_Handler,
		},
		{
			MethodName: "LeaveConsumerGroup",
			Handler:    _MqServerCall_LeaveConsumerGroup_Handler,
		},
		{
			MethodName: "CheckSourceTerm",
			Handler:    _MqServerCall_CheckSourceTerm_Handler,
		},
		{
			MethodName: "SubscribeTopic",
			Handler:    _MqServerCall_SubscribeTopic_Handler,
		},
		{
			MethodName: "UnSubscribeTopic",
			Handler:    _MqServerCall_UnSubscribeTopic_Handler,
		},
		{
			MethodName: "RegisterConsumerGroup",
			Handler:    _MqServerCall_RegisterConsumerGroup_Handler,
		},
		{
			MethodName: "CreateTopic",
			Handler:    _MqServerCall_CreateTopic_Handler,
		},
		{
			MethodName: "QueryTopic",
			Handler:    _MqServerCall_QueryTopic_Handler,
		},
		{
			MethodName: "AddPart",
			Handler:    _MqServerCall_AddPart_Handler,
		},
		{
			MethodName: "RemovePart",
			Handler:    _MqServerCall_RemovePart_Handler,
		},
		{
			MethodName: "RegisterConsumer",
			Handler:    _MqServerCall_RegisterConsumer_Handler,
		},
		{
			MethodName: "RegisterProducer",
			Handler:    _MqServerCall_RegisterProducer_Handler,
		},
		{
			MethodName: "UnRegisterConsumer",
			Handler:    _MqServerCall_UnRegisterConsumer_Handler,
		},
		{
			MethodName: "UnRegisterProducer",
			Handler:    _MqServerCall_UnRegisterProducer_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _MqServerCall_Heartbeat_Handler,
		},
		{
			MethodName: "ConsumerDisConnect",
			Handler:    _MqServerCall_ConsumerDisConnect_Handler,
		},
		{
			MethodName: "ProducerDisConnect",
			Handler:    _MqServerCall_ProducerDisConnect_Handler,
		},
		{
			MethodName: "PullMessage",
			Handler:    _MqServerCall_PullMessage_Handler,
		},
		{
			MethodName: "PushMessage",
			Handler:    _MqServerCall_PushMessage_Handler,
		},
		{
			MethodName: "ConfirmIdentity",
			Handler:    _MqServerCall_ConfirmIdentity_Handler,
		},
		{
			MethodName: "RegisterBroker",
			Handler:    _MqServerCall_RegisterBroker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "MqServerRpc.proto",
}
