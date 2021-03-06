// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/myTest.proto

package myTest

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for MyTest service

func NewMyTestEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MyTest service

type MyTestService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (MyTest_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (MyTest_PingPongService, error)
}

type myTestService struct {
	c    client.Client
	name string
}

func NewMyTestService(name string, c client.Client) MyTestService {
	return &myTestService{
		c:    c,
		name: name,
	}
}

func (c *myTestService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "MyTest.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myTestService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (MyTest_StreamService, error) {
	req := c.c.NewRequest(c.name, "MyTest.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &myTestServiceStream{stream}, nil
}

type MyTest_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type myTestServiceStream struct {
	stream client.Stream
}

func (x *myTestServiceStream) Close() error {
	return x.stream.Close()
}

func (x *myTestServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *myTestServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *myTestServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *myTestServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *myTestService) PingPong(ctx context.Context, opts ...client.CallOption) (MyTest_PingPongService, error) {
	req := c.c.NewRequest(c.name, "MyTest.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &myTestServicePingPong{stream}, nil
}

type MyTest_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type myTestServicePingPong struct {
	stream client.Stream
}

func (x *myTestServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *myTestServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *myTestServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *myTestServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *myTestServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *myTestServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MyTest service

type MyTestHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, MyTest_StreamStream) error
	PingPong(context.Context, MyTest_PingPongStream) error
}

func RegisterMyTestHandler(s server.Server, hdlr MyTestHandler, opts ...server.HandlerOption) error {
	type myTest interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type MyTest struct {
		myTest
	}
	h := &myTestHandler{hdlr}
	return s.Handle(s.NewHandler(&MyTest{h}, opts...))
}

type myTestHandler struct {
	MyTestHandler
}

func (h *myTestHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.MyTestHandler.Call(ctx, in, out)
}

func (h *myTestHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.MyTestHandler.Stream(ctx, m, &myTestStreamStream{stream})
}

type MyTest_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type myTestStreamStream struct {
	stream server.Stream
}

func (x *myTestStreamStream) Close() error {
	return x.stream.Close()
}

func (x *myTestStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *myTestStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *myTestStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *myTestStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *myTestHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.MyTestHandler.PingPong(ctx, &myTestPingPongStream{stream})
}

type MyTest_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type myTestPingPongStream struct {
	stream server.Stream
}

func (x *myTestPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *myTestPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *myTestPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *myTestPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *myTestPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *myTestPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
