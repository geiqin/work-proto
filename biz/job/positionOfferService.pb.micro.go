// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: positionOfferService.proto

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

// Api Endpoints for PositionOfferService service

func NewPositionOfferServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PositionOfferService service

type PositionOfferService interface {
	Create(ctx context.Context, in *PositionOfferRequest, opts ...client.CallOption) (*PositionOfferResponse, error)
	Update(ctx context.Context, in *PositionOffer, opts ...client.CallOption) (*PositionOfferResponse, error)
	Check(ctx context.Context, in *PositionOffer, opts ...client.CallOption) (*PositionOfferResponse, error)
	Delete(ctx context.Context, in *PositionOffer, opts ...client.CallOption) (*PositionOfferResponse, error)
	Get(ctx context.Context, in *PositionOffer, opts ...client.CallOption) (*PositionOfferResponse, error)
	Search(ctx context.Context, in *PositionOfferRequest, opts ...client.CallOption) (*PositionOfferResponse, error)
}

type positionOfferService struct {
	c    client.Client
	name string
}

func NewPositionOfferService(name string, c client.Client) PositionOfferService {
	return &positionOfferService{
		c:    c,
		name: name,
	}
}

func (c *positionOfferService) Create(ctx context.Context, in *PositionOfferRequest, opts ...client.CallOption) (*PositionOfferResponse, error) {
	req := c.c.NewRequest(c.name, "PositionOfferService.Create", in)
	out := new(PositionOfferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionOfferService) Update(ctx context.Context, in *PositionOffer, opts ...client.CallOption) (*PositionOfferResponse, error) {
	req := c.c.NewRequest(c.name, "PositionOfferService.Update", in)
	out := new(PositionOfferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionOfferService) Check(ctx context.Context, in *PositionOffer, opts ...client.CallOption) (*PositionOfferResponse, error) {
	req := c.c.NewRequest(c.name, "PositionOfferService.Check", in)
	out := new(PositionOfferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionOfferService) Delete(ctx context.Context, in *PositionOffer, opts ...client.CallOption) (*PositionOfferResponse, error) {
	req := c.c.NewRequest(c.name, "PositionOfferService.Delete", in)
	out := new(PositionOfferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionOfferService) Get(ctx context.Context, in *PositionOffer, opts ...client.CallOption) (*PositionOfferResponse, error) {
	req := c.c.NewRequest(c.name, "PositionOfferService.Get", in)
	out := new(PositionOfferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionOfferService) Search(ctx context.Context, in *PositionOfferRequest, opts ...client.CallOption) (*PositionOfferResponse, error) {
	req := c.c.NewRequest(c.name, "PositionOfferService.Search", in)
	out := new(PositionOfferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PositionOfferService service

type PositionOfferServiceHandler interface {
	Create(context.Context, *PositionOfferRequest, *PositionOfferResponse) error
	Update(context.Context, *PositionOffer, *PositionOfferResponse) error
	Check(context.Context, *PositionOffer, *PositionOfferResponse) error
	Delete(context.Context, *PositionOffer, *PositionOfferResponse) error
	Get(context.Context, *PositionOffer, *PositionOfferResponse) error
	Search(context.Context, *PositionOfferRequest, *PositionOfferResponse) error
}

func RegisterPositionOfferServiceHandler(s server.Server, hdlr PositionOfferServiceHandler, opts ...server.HandlerOption) error {
	type positionOfferService interface {
		Create(ctx context.Context, in *PositionOfferRequest, out *PositionOfferResponse) error
		Update(ctx context.Context, in *PositionOffer, out *PositionOfferResponse) error
		Check(ctx context.Context, in *PositionOffer, out *PositionOfferResponse) error
		Delete(ctx context.Context, in *PositionOffer, out *PositionOfferResponse) error
		Get(ctx context.Context, in *PositionOffer, out *PositionOfferResponse) error
		Search(ctx context.Context, in *PositionOfferRequest, out *PositionOfferResponse) error
	}
	type PositionOfferService struct {
		positionOfferService
	}
	h := &positionOfferServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PositionOfferService{h}, opts...))
}

type positionOfferServiceHandler struct {
	PositionOfferServiceHandler
}

func (h *positionOfferServiceHandler) Create(ctx context.Context, in *PositionOfferRequest, out *PositionOfferResponse) error {
	return h.PositionOfferServiceHandler.Create(ctx, in, out)
}

func (h *positionOfferServiceHandler) Update(ctx context.Context, in *PositionOffer, out *PositionOfferResponse) error {
	return h.PositionOfferServiceHandler.Update(ctx, in, out)
}

func (h *positionOfferServiceHandler) Check(ctx context.Context, in *PositionOffer, out *PositionOfferResponse) error {
	return h.PositionOfferServiceHandler.Check(ctx, in, out)
}

func (h *positionOfferServiceHandler) Delete(ctx context.Context, in *PositionOffer, out *PositionOfferResponse) error {
	return h.PositionOfferServiceHandler.Delete(ctx, in, out)
}

func (h *positionOfferServiceHandler) Get(ctx context.Context, in *PositionOffer, out *PositionOfferResponse) error {
	return h.PositionOfferServiceHandler.Get(ctx, in, out)
}

func (h *positionOfferServiceHandler) Search(ctx context.Context, in *PositionOfferRequest, out *PositionOfferResponse) error {
	return h.PositionOfferServiceHandler.Search(ctx, in, out)
}