// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderPromotionService.proto

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

// Api Endpoints for OrderPromotionService service

func NewOrderPromotionServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderPromotionService service

type OrderPromotionService interface {
	Create(ctx context.Context, in *OrderPromotion, opts ...client.CallOption) (*OrderPromotionResponse, error)
	Update(ctx context.Context, in *OrderPromotion, opts ...client.CallOption) (*OrderPromotionResponse, error)
	Delete(ctx context.Context, in *OrderPromotion, opts ...client.CallOption) (*OrderPromotionResponse, error)
	Get(ctx context.Context, in *OrderPromotion, opts ...client.CallOption) (*OrderPromotionResponse, error)
	Search(ctx context.Context, in *OrderPromotionParams, opts ...client.CallOption) (*OrderPromotionResponse, error)
	List(ctx context.Context, in *OrderPromotionParams, opts ...client.CallOption) (*OrderPromotionResponse, error)
}

type orderPromotionService struct {
	c    client.Client
	name string
}

func NewOrderPromotionService(name string, c client.Client) OrderPromotionService {
	return &orderPromotionService{
		c:    c,
		name: name,
	}
}

func (c *orderPromotionService) Create(ctx context.Context, in *OrderPromotion, opts ...client.CallOption) (*OrderPromotionResponse, error) {
	req := c.c.NewRequest(c.name, "OrderPromotionService.Create", in)
	out := new(OrderPromotionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderPromotionService) Update(ctx context.Context, in *OrderPromotion, opts ...client.CallOption) (*OrderPromotionResponse, error) {
	req := c.c.NewRequest(c.name, "OrderPromotionService.Update", in)
	out := new(OrderPromotionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderPromotionService) Delete(ctx context.Context, in *OrderPromotion, opts ...client.CallOption) (*OrderPromotionResponse, error) {
	req := c.c.NewRequest(c.name, "OrderPromotionService.Delete", in)
	out := new(OrderPromotionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderPromotionService) Get(ctx context.Context, in *OrderPromotion, opts ...client.CallOption) (*OrderPromotionResponse, error) {
	req := c.c.NewRequest(c.name, "OrderPromotionService.Get", in)
	out := new(OrderPromotionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderPromotionService) Search(ctx context.Context, in *OrderPromotionParams, opts ...client.CallOption) (*OrderPromotionResponse, error) {
	req := c.c.NewRequest(c.name, "OrderPromotionService.Search", in)
	out := new(OrderPromotionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderPromotionService) List(ctx context.Context, in *OrderPromotionParams, opts ...client.CallOption) (*OrderPromotionResponse, error) {
	req := c.c.NewRequest(c.name, "OrderPromotionService.List", in)
	out := new(OrderPromotionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderPromotionService service

type OrderPromotionServiceHandler interface {
	Create(context.Context, *OrderPromotion, *OrderPromotionResponse) error
	Update(context.Context, *OrderPromotion, *OrderPromotionResponse) error
	Delete(context.Context, *OrderPromotion, *OrderPromotionResponse) error
	Get(context.Context, *OrderPromotion, *OrderPromotionResponse) error
	Search(context.Context, *OrderPromotionParams, *OrderPromotionResponse) error
	List(context.Context, *OrderPromotionParams, *OrderPromotionResponse) error
}

func RegisterOrderPromotionServiceHandler(s server.Server, hdlr OrderPromotionServiceHandler, opts ...server.HandlerOption) error {
	type orderPromotionService interface {
		Create(ctx context.Context, in *OrderPromotion, out *OrderPromotionResponse) error
		Update(ctx context.Context, in *OrderPromotion, out *OrderPromotionResponse) error
		Delete(ctx context.Context, in *OrderPromotion, out *OrderPromotionResponse) error
		Get(ctx context.Context, in *OrderPromotion, out *OrderPromotionResponse) error
		Search(ctx context.Context, in *OrderPromotionParams, out *OrderPromotionResponse) error
		List(ctx context.Context, in *OrderPromotionParams, out *OrderPromotionResponse) error
	}
	type OrderPromotionService struct {
		orderPromotionService
	}
	h := &orderPromotionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderPromotionService{h}, opts...))
}

type orderPromotionServiceHandler struct {
	OrderPromotionServiceHandler
}

func (h *orderPromotionServiceHandler) Create(ctx context.Context, in *OrderPromotion, out *OrderPromotionResponse) error {
	return h.OrderPromotionServiceHandler.Create(ctx, in, out)
}

func (h *orderPromotionServiceHandler) Update(ctx context.Context, in *OrderPromotion, out *OrderPromotionResponse) error {
	return h.OrderPromotionServiceHandler.Update(ctx, in, out)
}

func (h *orderPromotionServiceHandler) Delete(ctx context.Context, in *OrderPromotion, out *OrderPromotionResponse) error {
	return h.OrderPromotionServiceHandler.Delete(ctx, in, out)
}

func (h *orderPromotionServiceHandler) Get(ctx context.Context, in *OrderPromotion, out *OrderPromotionResponse) error {
	return h.OrderPromotionServiceHandler.Get(ctx, in, out)
}

func (h *orderPromotionServiceHandler) Search(ctx context.Context, in *OrderPromotionParams, out *OrderPromotionResponse) error {
	return h.OrderPromotionServiceHandler.Search(ctx, in, out)
}

func (h *orderPromotionServiceHandler) List(ctx context.Context, in *OrderPromotionParams, out *OrderPromotionResponse) error {
	return h.OrderPromotionServiceHandler.List(ctx, in, out)
}