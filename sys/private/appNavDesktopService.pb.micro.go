// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: appNavDesktopService.proto

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

// Api Endpoints for AppNavDesktopService service

func NewAppNavDesktopServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AppNavDesktopService service

type AppNavDesktopService interface {
	Create(ctx context.Context, in *AppNavDesktop, opts ...client.CallOption) (*AppNavDesktopResponse, error)
	Delete(ctx context.Context, in *AppNavDesktopWhere, opts ...client.CallOption) (*AppNavDesktopResponse, error)
	List(ctx context.Context, in *AppNavDesktopWhere, opts ...client.CallOption) (*AppNavDesktopResponse, error)
	Search(ctx context.Context, in *AppNavDesktopWhere, opts ...client.CallOption) (*AppNavDesktopResponse, error)
	ChangeSort(ctx context.Context, in *AppNavDesktop, opts ...client.CallOption) (*AppNavDesktopResponse, error)
}

type appNavDesktopService struct {
	c    client.Client
	name string
}

func NewAppNavDesktopService(name string, c client.Client) AppNavDesktopService {
	return &appNavDesktopService{
		c:    c,
		name: name,
	}
}

func (c *appNavDesktopService) Create(ctx context.Context, in *AppNavDesktop, opts ...client.CallOption) (*AppNavDesktopResponse, error) {
	req := c.c.NewRequest(c.name, "AppNavDesktopService.Create", in)
	out := new(AppNavDesktopResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appNavDesktopService) Delete(ctx context.Context, in *AppNavDesktopWhere, opts ...client.CallOption) (*AppNavDesktopResponse, error) {
	req := c.c.NewRequest(c.name, "AppNavDesktopService.Delete", in)
	out := new(AppNavDesktopResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appNavDesktopService) List(ctx context.Context, in *AppNavDesktopWhere, opts ...client.CallOption) (*AppNavDesktopResponse, error) {
	req := c.c.NewRequest(c.name, "AppNavDesktopService.List", in)
	out := new(AppNavDesktopResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appNavDesktopService) Search(ctx context.Context, in *AppNavDesktopWhere, opts ...client.CallOption) (*AppNavDesktopResponse, error) {
	req := c.c.NewRequest(c.name, "AppNavDesktopService.Search", in)
	out := new(AppNavDesktopResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appNavDesktopService) ChangeSort(ctx context.Context, in *AppNavDesktop, opts ...client.CallOption) (*AppNavDesktopResponse, error) {
	req := c.c.NewRequest(c.name, "AppNavDesktopService.ChangeSort", in)
	out := new(AppNavDesktopResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppNavDesktopService service

type AppNavDesktopServiceHandler interface {
	Create(context.Context, *AppNavDesktop, *AppNavDesktopResponse) error
	Delete(context.Context, *AppNavDesktopWhere, *AppNavDesktopResponse) error
	List(context.Context, *AppNavDesktopWhere, *AppNavDesktopResponse) error
	Search(context.Context, *AppNavDesktopWhere, *AppNavDesktopResponse) error
	ChangeSort(context.Context, *AppNavDesktop, *AppNavDesktopResponse) error
}

func RegisterAppNavDesktopServiceHandler(s server.Server, hdlr AppNavDesktopServiceHandler, opts ...server.HandlerOption) error {
	type appNavDesktopService interface {
		Create(ctx context.Context, in *AppNavDesktop, out *AppNavDesktopResponse) error
		Delete(ctx context.Context, in *AppNavDesktopWhere, out *AppNavDesktopResponse) error
		List(ctx context.Context, in *AppNavDesktopWhere, out *AppNavDesktopResponse) error
		Search(ctx context.Context, in *AppNavDesktopWhere, out *AppNavDesktopResponse) error
		ChangeSort(ctx context.Context, in *AppNavDesktop, out *AppNavDesktopResponse) error
	}
	type AppNavDesktopService struct {
		appNavDesktopService
	}
	h := &appNavDesktopServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AppNavDesktopService{h}, opts...))
}

type appNavDesktopServiceHandler struct {
	AppNavDesktopServiceHandler
}

func (h *appNavDesktopServiceHandler) Create(ctx context.Context, in *AppNavDesktop, out *AppNavDesktopResponse) error {
	return h.AppNavDesktopServiceHandler.Create(ctx, in, out)
}

func (h *appNavDesktopServiceHandler) Delete(ctx context.Context, in *AppNavDesktopWhere, out *AppNavDesktopResponse) error {
	return h.AppNavDesktopServiceHandler.Delete(ctx, in, out)
}

func (h *appNavDesktopServiceHandler) List(ctx context.Context, in *AppNavDesktopWhere, out *AppNavDesktopResponse) error {
	return h.AppNavDesktopServiceHandler.List(ctx, in, out)
}

func (h *appNavDesktopServiceHandler) Search(ctx context.Context, in *AppNavDesktopWhere, out *AppNavDesktopResponse) error {
	return h.AppNavDesktopServiceHandler.Search(ctx, in, out)
}

func (h *appNavDesktopServiceHandler) ChangeSort(ctx context.Context, in *AppNavDesktop, out *AppNavDesktopResponse) error {
	return h.AppNavDesktopServiceHandler.ChangeSort(ctx, in, out)
}
