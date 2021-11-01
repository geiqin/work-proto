// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: authorizeService.proto

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

// Api Endpoints for AuthorizeService service

func NewAuthorizeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AuthorizeService service

type AuthorizeService interface {
	//获得当前客户
	Info(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*CustomerResponse, error)
	//绑定微信通过小程序
	BindWxByMini(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error)
	//绑定微信通过APP
	BindWxByApp(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error)
	//绑定手机号
	BindMobile(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error)
	//绑定邮箱账号
	BindEmail(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error)
}

type authorizeService struct {
	c    client.Client
	name string
}

func NewAuthorizeService(name string, c client.Client) AuthorizeService {
	return &authorizeService{
		c:    c,
		name: name,
	}
}

func (c *authorizeService) Info(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.Info", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) BindWxByMini(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.BindWxByMini", in)
	out := new(AuthorizeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) BindWxByApp(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.BindWxByApp", in)
	out := new(AuthorizeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) BindMobile(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.BindMobile", in)
	out := new(AuthorizeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) BindEmail(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.BindEmail", in)
	out := new(AuthorizeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthorizeService service

type AuthorizeServiceHandler interface {
	//获得当前客户
	Info(context.Context, *AuthorizeRequest, *CustomerResponse) error
	//绑定微信通过小程序
	BindWxByMini(context.Context, *AuthorizeRequest, *AuthorizeResponse) error
	//绑定微信通过APP
	BindWxByApp(context.Context, *AuthorizeRequest, *AuthorizeResponse) error
	//绑定手机号
	BindMobile(context.Context, *AuthorizeRequest, *AuthorizeResponse) error
	//绑定邮箱账号
	BindEmail(context.Context, *AuthorizeRequest, *AuthorizeResponse) error
}

func RegisterAuthorizeServiceHandler(s server.Server, hdlr AuthorizeServiceHandler, opts ...server.HandlerOption) error {
	type authorizeService interface {
		Info(ctx context.Context, in *AuthorizeRequest, out *CustomerResponse) error
		BindWxByMini(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error
		BindWxByApp(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error
		BindMobile(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error
		BindEmail(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error
	}
	type AuthorizeService struct {
		authorizeService
	}
	h := &authorizeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AuthorizeService{h}, opts...))
}

type authorizeServiceHandler struct {
	AuthorizeServiceHandler
}

func (h *authorizeServiceHandler) Info(ctx context.Context, in *AuthorizeRequest, out *CustomerResponse) error {
	return h.AuthorizeServiceHandler.Info(ctx, in, out)
}

func (h *authorizeServiceHandler) BindWxByMini(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error {
	return h.AuthorizeServiceHandler.BindWxByMini(ctx, in, out)
}

func (h *authorizeServiceHandler) BindWxByApp(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error {
	return h.AuthorizeServiceHandler.BindWxByApp(ctx, in, out)
}

func (h *authorizeServiceHandler) BindMobile(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error {
	return h.AuthorizeServiceHandler.BindMobile(ctx, in, out)
}

func (h *authorizeServiceHandler) BindEmail(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error {
	return h.AuthorizeServiceHandler.BindEmail(ctx, in, out)
}
