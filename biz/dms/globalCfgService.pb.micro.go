// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: globalCfgService.proto

package services

import (
	fmt "fmt"
	common "github.com/geiqin/micro-kit/protobuf/common"
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

// Api Endpoints for GlobalCfgService service

func NewGlobalCfgServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GlobalCfgService service

type GlobalCfgService interface {
	//设置销分销配置信息
	Set(ctx context.Context, in *GlobalCfg, opts ...client.CallOption) (*GlobalCfgResponse, error)
	//获取分销配置信息
	Get(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*GlobalCfgResponse, error)
}

type globalCfgService struct {
	c    client.Client
	name string
}

func NewGlobalCfgService(name string, c client.Client) GlobalCfgService {
	return &globalCfgService{
		c:    c,
		name: name,
	}
}

func (c *globalCfgService) Set(ctx context.Context, in *GlobalCfg, opts ...client.CallOption) (*GlobalCfgResponse, error) {
	req := c.c.NewRequest(c.name, "GlobalCfgService.Set", in)
	out := new(GlobalCfgResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalCfgService) Get(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*GlobalCfgResponse, error) {
	req := c.c.NewRequest(c.name, "GlobalCfgService.Get", in)
	out := new(GlobalCfgResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GlobalCfgService service

type GlobalCfgServiceHandler interface {
	//设置销分销配置信息
	Set(context.Context, *GlobalCfg, *GlobalCfgResponse) error
	//获取分销配置信息
	Get(context.Context, *common.Empty, *GlobalCfgResponse) error
}

func RegisterGlobalCfgServiceHandler(s server.Server, hdlr GlobalCfgServiceHandler, opts ...server.HandlerOption) error {
	type globalCfgService interface {
		Set(ctx context.Context, in *GlobalCfg, out *GlobalCfgResponse) error
		Get(ctx context.Context, in *common.Empty, out *GlobalCfgResponse) error
	}
	type GlobalCfgService struct {
		globalCfgService
	}
	h := &globalCfgServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GlobalCfgService{h}, opts...))
}

type globalCfgServiceHandler struct {
	GlobalCfgServiceHandler
}

func (h *globalCfgServiceHandler) Set(ctx context.Context, in *GlobalCfg, out *GlobalCfgResponse) error {
	return h.GlobalCfgServiceHandler.Set(ctx, in, out)
}

func (h *globalCfgServiceHandler) Get(ctx context.Context, in *common.Empty, out *GlobalCfgResponse) error {
	return h.GlobalCfgServiceHandler.Get(ctx, in, out)
}
