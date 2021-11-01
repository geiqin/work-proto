// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: creditService.proto

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

// Api Endpoints for CreditService service

func NewCreditServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CreditService service

type CreditService interface {
	Set(ctx context.Context, in *Credit, opts ...client.CallOption) (*CreditResponse, error)
	Detail(ctx context.Context, in *Credit, opts ...client.CallOption) (*CreditResponse, error)
	Search(ctx context.Context, in *CreditWhere, opts ...client.CallOption) (*CreditResponse, error)
}

type creditService struct {
	c    client.Client
	name string
}

func NewCreditService(name string, c client.Client) CreditService {
	return &creditService{
		c:    c,
		name: name,
	}
}

func (c *creditService) Set(ctx context.Context, in *Credit, opts ...client.CallOption) (*CreditResponse, error) {
	req := c.c.NewRequest(c.name, "CreditService.Set", in)
	out := new(CreditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditService) Detail(ctx context.Context, in *Credit, opts ...client.CallOption) (*CreditResponse, error) {
	req := c.c.NewRequest(c.name, "CreditService.Detail", in)
	out := new(CreditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditService) Search(ctx context.Context, in *CreditWhere, opts ...client.CallOption) (*CreditResponse, error) {
	req := c.c.NewRequest(c.name, "CreditService.Search", in)
	out := new(CreditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CreditService service

type CreditServiceHandler interface {
	Set(context.Context, *Credit, *CreditResponse) error
	Detail(context.Context, *Credit, *CreditResponse) error
	Search(context.Context, *CreditWhere, *CreditResponse) error
}

func RegisterCreditServiceHandler(s server.Server, hdlr CreditServiceHandler, opts ...server.HandlerOption) error {
	type creditService interface {
		Set(ctx context.Context, in *Credit, out *CreditResponse) error
		Detail(ctx context.Context, in *Credit, out *CreditResponse) error
		Search(ctx context.Context, in *CreditWhere, out *CreditResponse) error
	}
	type CreditService struct {
		creditService
	}
	h := &creditServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CreditService{h}, opts...))
}

type creditServiceHandler struct {
	CreditServiceHandler
}

func (h *creditServiceHandler) Set(ctx context.Context, in *Credit, out *CreditResponse) error {
	return h.CreditServiceHandler.Set(ctx, in, out)
}

func (h *creditServiceHandler) Detail(ctx context.Context, in *Credit, out *CreditResponse) error {
	return h.CreditServiceHandler.Detail(ctx, in, out)
}

func (h *creditServiceHandler) Search(ctx context.Context, in *CreditWhere, out *CreditResponse) error {
	return h.CreditServiceHandler.Search(ctx, in, out)
}
