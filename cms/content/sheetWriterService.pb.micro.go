// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sheetWriterService.proto

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

// Api Endpoints for SheetWriterService service

func NewSheetWriterServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SheetWriterService service

type SheetWriterService interface {
	Get(ctx context.Context, in *SheetWriter, opts ...client.CallOption) (*SheetWriterResponse, error)
	Search(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error)
	List(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error)
	Submit(ctx context.Context, in *SheetWriter, opts ...client.CallOption) (*SheetWriterResponse, error)
	Update(ctx context.Context, in *SheetWriter, opts ...client.CallOption) (*SheetWriterResponse, error)
	Delete(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error)
}

type sheetWriterService struct {
	c    client.Client
	name string
}

func NewSheetWriterService(name string, c client.Client) SheetWriterService {
	return &sheetWriterService{
		c:    c,
		name: name,
	}
}

func (c *sheetWriterService) Get(ctx context.Context, in *SheetWriter, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "SheetWriterService.Get", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetWriterService) Search(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "SheetWriterService.Search", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetWriterService) List(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "SheetWriterService.List", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetWriterService) Submit(ctx context.Context, in *SheetWriter, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "SheetWriterService.Submit", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetWriterService) Update(ctx context.Context, in *SheetWriter, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "SheetWriterService.Update", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetWriterService) Delete(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "SheetWriterService.Delete", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SheetWriterService service

type SheetWriterServiceHandler interface {
	Get(context.Context, *SheetWriter, *SheetWriterResponse) error
	Search(context.Context, *SheetWriterWhere, *SheetWriterResponse) error
	List(context.Context, *SheetWriterWhere, *SheetWriterResponse) error
	Submit(context.Context, *SheetWriter, *SheetWriterResponse) error
	Update(context.Context, *SheetWriter, *SheetWriterResponse) error
	Delete(context.Context, *SheetWriterWhere, *SheetWriterResponse) error
}

func RegisterSheetWriterServiceHandler(s server.Server, hdlr SheetWriterServiceHandler, opts ...server.HandlerOption) error {
	type sheetWriterService interface {
		Get(ctx context.Context, in *SheetWriter, out *SheetWriterResponse) error
		Search(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error
		List(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error
		Submit(ctx context.Context, in *SheetWriter, out *SheetWriterResponse) error
		Update(ctx context.Context, in *SheetWriter, out *SheetWriterResponse) error
		Delete(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error
	}
	type SheetWriterService struct {
		sheetWriterService
	}
	h := &sheetWriterServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SheetWriterService{h}, opts...))
}

type sheetWriterServiceHandler struct {
	SheetWriterServiceHandler
}

func (h *sheetWriterServiceHandler) Get(ctx context.Context, in *SheetWriter, out *SheetWriterResponse) error {
	return h.SheetWriterServiceHandler.Get(ctx, in, out)
}

func (h *sheetWriterServiceHandler) Search(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error {
	return h.SheetWriterServiceHandler.Search(ctx, in, out)
}

func (h *sheetWriterServiceHandler) List(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error {
	return h.SheetWriterServiceHandler.List(ctx, in, out)
}

func (h *sheetWriterServiceHandler) Submit(ctx context.Context, in *SheetWriter, out *SheetWriterResponse) error {
	return h.SheetWriterServiceHandler.Submit(ctx, in, out)
}

func (h *sheetWriterServiceHandler) Update(ctx context.Context, in *SheetWriter, out *SheetWriterResponse) error {
	return h.SheetWriterServiceHandler.Update(ctx, in, out)
}

func (h *sheetWriterServiceHandler) Delete(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error {
	return h.SheetWriterServiceHandler.Delete(ctx, in, out)
}

// Api Endpoints for MySheetWriterService service

func NewMySheetWriterServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MySheetWriterService service

type MySheetWriterService interface {
	Submit(ctx context.Context, in *SheetWriter, opts ...client.CallOption) (*SheetWriterResponse, error)
	Update(ctx context.Context, in *SheetWriter, opts ...client.CallOption) (*SheetWriterResponse, error)
	Delete(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error)
	Get(ctx context.Context, in *SheetWriter, opts ...client.CallOption) (*SheetWriterResponse, error)
	Search(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error)
	List(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error)
	GetForm(ctx context.Context, in *SheetWhere, opts ...client.CallOption) (*SheetResponse, error)
}

type mySheetWriterService struct {
	c    client.Client
	name string
}

func NewMySheetWriterService(name string, c client.Client) MySheetWriterService {
	return &mySheetWriterService{
		c:    c,
		name: name,
	}
}

func (c *mySheetWriterService) Submit(ctx context.Context, in *SheetWriter, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "MySheetWriterService.Submit", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySheetWriterService) Update(ctx context.Context, in *SheetWriter, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "MySheetWriterService.Update", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySheetWriterService) Delete(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "MySheetWriterService.Delete", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySheetWriterService) Get(ctx context.Context, in *SheetWriter, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "MySheetWriterService.Get", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySheetWriterService) Search(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "MySheetWriterService.Search", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySheetWriterService) List(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "MySheetWriterService.List", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySheetWriterService) GetForm(ctx context.Context, in *SheetWhere, opts ...client.CallOption) (*SheetResponse, error) {
	req := c.c.NewRequest(c.name, "MySheetWriterService.GetForm", in)
	out := new(SheetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MySheetWriterService service

type MySheetWriterServiceHandler interface {
	Submit(context.Context, *SheetWriter, *SheetWriterResponse) error
	Update(context.Context, *SheetWriter, *SheetWriterResponse) error
	Delete(context.Context, *SheetWriterWhere, *SheetWriterResponse) error
	Get(context.Context, *SheetWriter, *SheetWriterResponse) error
	Search(context.Context, *SheetWriterWhere, *SheetWriterResponse) error
	List(context.Context, *SheetWriterWhere, *SheetWriterResponse) error
	GetForm(context.Context, *SheetWhere, *SheetResponse) error
}

func RegisterMySheetWriterServiceHandler(s server.Server, hdlr MySheetWriterServiceHandler, opts ...server.HandlerOption) error {
	type mySheetWriterService interface {
		Submit(ctx context.Context, in *SheetWriter, out *SheetWriterResponse) error
		Update(ctx context.Context, in *SheetWriter, out *SheetWriterResponse) error
		Delete(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error
		Get(ctx context.Context, in *SheetWriter, out *SheetWriterResponse) error
		Search(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error
		List(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error
		GetForm(ctx context.Context, in *SheetWhere, out *SheetResponse) error
	}
	type MySheetWriterService struct {
		mySheetWriterService
	}
	h := &mySheetWriterServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MySheetWriterService{h}, opts...))
}

type mySheetWriterServiceHandler struct {
	MySheetWriterServiceHandler
}

func (h *mySheetWriterServiceHandler) Submit(ctx context.Context, in *SheetWriter, out *SheetWriterResponse) error {
	return h.MySheetWriterServiceHandler.Submit(ctx, in, out)
}

func (h *mySheetWriterServiceHandler) Update(ctx context.Context, in *SheetWriter, out *SheetWriterResponse) error {
	return h.MySheetWriterServiceHandler.Update(ctx, in, out)
}

func (h *mySheetWriterServiceHandler) Delete(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error {
	return h.MySheetWriterServiceHandler.Delete(ctx, in, out)
}

func (h *mySheetWriterServiceHandler) Get(ctx context.Context, in *SheetWriter, out *SheetWriterResponse) error {
	return h.MySheetWriterServiceHandler.Get(ctx, in, out)
}

func (h *mySheetWriterServiceHandler) Search(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error {
	return h.MySheetWriterServiceHandler.Search(ctx, in, out)
}

func (h *mySheetWriterServiceHandler) List(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error {
	return h.MySheetWriterServiceHandler.List(ctx, in, out)
}

func (h *mySheetWriterServiceHandler) GetForm(ctx context.Context, in *SheetWhere, out *SheetResponse) error {
	return h.MySheetWriterServiceHandler.GetForm(ctx, in, out)
}

// Api Endpoints for FrontSheetWriterService service

func NewFrontSheetWriterServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FrontSheetWriterService service

type FrontSheetWriterService interface {
	// ????????????????????????(??????????????????????????????)
	Search(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error)
}

type frontSheetWriterService struct {
	c    client.Client
	name string
}

func NewFrontSheetWriterService(name string, c client.Client) FrontSheetWriterService {
	return &frontSheetWriterService{
		c:    c,
		name: name,
	}
}

func (c *frontSheetWriterService) Search(ctx context.Context, in *SheetWriterWhere, opts ...client.CallOption) (*SheetWriterResponse, error) {
	req := c.c.NewRequest(c.name, "FrontSheetWriterService.Search", in)
	out := new(SheetWriterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FrontSheetWriterService service

type FrontSheetWriterServiceHandler interface {
	// ????????????????????????(??????????????????????????????)
	Search(context.Context, *SheetWriterWhere, *SheetWriterResponse) error
}

func RegisterFrontSheetWriterServiceHandler(s server.Server, hdlr FrontSheetWriterServiceHandler, opts ...server.HandlerOption) error {
	type frontSheetWriterService interface {
		Search(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error
	}
	type FrontSheetWriterService struct {
		frontSheetWriterService
	}
	h := &frontSheetWriterServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FrontSheetWriterService{h}, opts...))
}

type frontSheetWriterServiceHandler struct {
	FrontSheetWriterServiceHandler
}

func (h *frontSheetWriterServiceHandler) Search(ctx context.Context, in *SheetWriterWhere, out *SheetWriterResponse) error {
	return h.FrontSheetWriterServiceHandler.Search(ctx, in, out)
}
