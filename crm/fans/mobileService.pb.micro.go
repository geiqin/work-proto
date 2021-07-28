// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: mobileService.proto

package services

import (
	fmt "fmt"
	_ "github.com/geiqin/microkit/protobuf/common"
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

// Api Endpoints for MobileService service

func NewMobileServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MobileService service

type MobileService interface {
	Get(ctx context.Context, in *Mobile, opts ...client.CallOption) (*MobileResponse, error)
	Search(ctx context.Context, in *MobileRequest, opts ...client.CallOption) (*MobileResponse, error)
}

type mobileService struct {
	c    client.Client
	name string
}

func NewMobileService(name string, c client.Client) MobileService {
	return &mobileService{
		c:    c,
		name: name,
	}
}

func (c *mobileService) Get(ctx context.Context, in *Mobile, opts ...client.CallOption) (*MobileResponse, error) {
	req := c.c.NewRequest(c.name, "MobileService.Get", in)
	out := new(MobileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) Search(ctx context.Context, in *MobileRequest, opts ...client.CallOption) (*MobileResponse, error) {
	req := c.c.NewRequest(c.name, "MobileService.Search", in)
	out := new(MobileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MobileService service

type MobileServiceHandler interface {
	Get(context.Context, *Mobile, *MobileResponse) error
	Search(context.Context, *MobileRequest, *MobileResponse) error
}

func RegisterMobileServiceHandler(s server.Server, hdlr MobileServiceHandler, opts ...server.HandlerOption) error {
	type mobileService interface {
		Get(ctx context.Context, in *Mobile, out *MobileResponse) error
		Search(ctx context.Context, in *MobileRequest, out *MobileResponse) error
	}
	type MobileService struct {
		mobileService
	}
	h := &mobileServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MobileService{h}, opts...))
}

type mobileServiceHandler struct {
	MobileServiceHandler
}

func (h *mobileServiceHandler) Get(ctx context.Context, in *Mobile, out *MobileResponse) error {
	return h.MobileServiceHandler.Get(ctx, in, out)
}

func (h *mobileServiceHandler) Search(ctx context.Context, in *MobileRequest, out *MobileResponse) error {
	return h.MobileServiceHandler.Search(ctx, in, out)
}
