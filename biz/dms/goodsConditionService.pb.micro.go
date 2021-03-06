// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: goodsConditionService.proto

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

// Api Endpoints for GoodsConditionService service

func NewGoodsConditionServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GoodsConditionService service

type GoodsConditionService interface {
	List(ctx context.Context, in *GoodsCondition, opts ...client.CallOption) (*GoodsConditionResponse, error)
}

type goodsConditionService struct {
	c    client.Client
	name string
}

func NewGoodsConditionService(name string, c client.Client) GoodsConditionService {
	return &goodsConditionService{
		c:    c,
		name: name,
	}
}

func (c *goodsConditionService) List(ctx context.Context, in *GoodsCondition, opts ...client.CallOption) (*GoodsConditionResponse, error) {
	req := c.c.NewRequest(c.name, "GoodsConditionService.List", in)
	out := new(GoodsConditionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoodsConditionService service

type GoodsConditionServiceHandler interface {
	List(context.Context, *GoodsCondition, *GoodsConditionResponse) error
}

func RegisterGoodsConditionServiceHandler(s server.Server, hdlr GoodsConditionServiceHandler, opts ...server.HandlerOption) error {
	type goodsConditionService interface {
		List(ctx context.Context, in *GoodsCondition, out *GoodsConditionResponse) error
	}
	type GoodsConditionService struct {
		goodsConditionService
	}
	h := &goodsConditionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GoodsConditionService{h}, opts...))
}

type goodsConditionServiceHandler struct {
	GoodsConditionServiceHandler
}

func (h *goodsConditionServiceHandler) List(ctx context.Context, in *GoodsCondition, out *GoodsConditionResponse) error {
	return h.GoodsConditionServiceHandler.List(ctx, in, out)
}
