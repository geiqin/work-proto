// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: accessLogisticsService.proto

package services

import (
	fmt "fmt"
	_ "github.com/geiqin/micro-kit/protobuf/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for AccessLogisticsService service

func NewAccessLogisticsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AccessLogisticsService service

type AccessLogisticsService interface {
	// 快递下单
	Submit(ctx context.Context, in *AccessLogisticsReq, opts ...client.CallOption) (*AccessLogisticsResponse, error)
}

type accessLogisticsService struct {
	c    client.Client
	name string
}

func NewAccessLogisticsService(name string, c client.Client) AccessLogisticsService {
	return &accessLogisticsService{
		c:    c,
		name: name,
	}
}

func (c *accessLogisticsService) Submit(ctx context.Context, in *AccessLogisticsReq, opts ...client.CallOption) (*AccessLogisticsResponse, error) {
	req := c.c.NewRequest(c.name, "AccessLogisticsService.Submit", in)
	out := new(AccessLogisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccessLogisticsService service

type AccessLogisticsServiceHandler interface {
	// 快递下单
	Submit(context.Context, *AccessLogisticsReq, *AccessLogisticsResponse) error
}

func RegisterAccessLogisticsServiceHandler(s server.Server, hdlr AccessLogisticsServiceHandler, opts ...server.HandlerOption) error {
	type accessLogisticsService interface {
		Submit(ctx context.Context, in *AccessLogisticsReq, out *AccessLogisticsResponse) error
	}
	type AccessLogisticsService struct {
		accessLogisticsService
	}
	h := &accessLogisticsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AccessLogisticsService{h}, opts...))
}

type accessLogisticsServiceHandler struct {
	AccessLogisticsServiceHandler
}

func (h *accessLogisticsServiceHandler) Submit(ctx context.Context, in *AccessLogisticsReq, out *AccessLogisticsResponse) error {
	return h.AccessLogisticsServiceHandler.Submit(ctx, in, out)
}