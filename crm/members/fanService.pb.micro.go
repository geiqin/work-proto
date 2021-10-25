// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: fanService.proto

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

// Api Endpoints for FanService service

func NewFanServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FanService service

type FanService interface {
	Update(ctx context.Context, in *Fan, opts ...client.CallOption) (*FanResponse, error)
	Delete(ctx context.Context, in *Fan, opts ...client.CallOption) (*FanResponse, error)
	Get(ctx context.Context, in *Fan, opts ...client.CallOption) (*FanResponse, error)
	Search(ctx context.Context, in *FanRequest, opts ...client.CallOption) (*FanResponse, error)
}

type fanService struct {
	c    client.Client
	name string
}

func NewFanService(name string, c client.Client) FanService {
	return &fanService{
		c:    c,
		name: name,
	}
}

func (c *fanService) Update(ctx context.Context, in *Fan, opts ...client.CallOption) (*FanResponse, error) {
	req := c.c.NewRequest(c.name, "FanService.Update", in)
	out := new(FanResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanService) Delete(ctx context.Context, in *Fan, opts ...client.CallOption) (*FanResponse, error) {
	req := c.c.NewRequest(c.name, "FanService.Delete", in)
	out := new(FanResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanService) Get(ctx context.Context, in *Fan, opts ...client.CallOption) (*FanResponse, error) {
	req := c.c.NewRequest(c.name, "FanService.Get", in)
	out := new(FanResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanService) Search(ctx context.Context, in *FanRequest, opts ...client.CallOption) (*FanResponse, error) {
	req := c.c.NewRequest(c.name, "FanService.Search", in)
	out := new(FanResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FanService service

type FanServiceHandler interface {
	Update(context.Context, *Fan, *FanResponse) error
	Delete(context.Context, *Fan, *FanResponse) error
	Get(context.Context, *Fan, *FanResponse) error
	Search(context.Context, *FanRequest, *FanResponse) error
}

func RegisterFanServiceHandler(s server.Server, hdlr FanServiceHandler, opts ...server.HandlerOption) error {
	type fanService interface {
		Update(ctx context.Context, in *Fan, out *FanResponse) error
		Delete(ctx context.Context, in *Fan, out *FanResponse) error
		Get(ctx context.Context, in *Fan, out *FanResponse) error
		Search(ctx context.Context, in *FanRequest, out *FanResponse) error
	}
	type FanService struct {
		fanService
	}
	h := &fanServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FanService{h}, opts...))
}

type fanServiceHandler struct {
	FanServiceHandler
}

func (h *fanServiceHandler) Update(ctx context.Context, in *Fan, out *FanResponse) error {
	return h.FanServiceHandler.Update(ctx, in, out)
}

func (h *fanServiceHandler) Delete(ctx context.Context, in *Fan, out *FanResponse) error {
	return h.FanServiceHandler.Delete(ctx, in, out)
}

func (h *fanServiceHandler) Get(ctx context.Context, in *Fan, out *FanResponse) error {
	return h.FanServiceHandler.Get(ctx, in, out)
}

func (h *fanServiceHandler) Search(ctx context.Context, in *FanRequest, out *FanResponse) error {
	return h.FanServiceHandler.Search(ctx, in, out)
}
