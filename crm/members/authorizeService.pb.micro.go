// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: authorizeService.proto

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

// Api Endpoints for AuthorizeService service

func NewAuthorizeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AuthorizeService service

type AuthorizeService interface {
	//客户注册
	Register(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error)
	//账号登录
	Login(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*common.TokenResponse, error)
	//短信登录
	SmsLogin(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*common.TokenResponse, error)
	//获得当前客户
	Info(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*CustomerResponse, error)
	//安全退出
	Logout(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*AuthorizeResponse, error)
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

func (c *authorizeService) Register(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.Register", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) Login(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*common.TokenResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.Login", in)
	out := new(common.TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) SmsLogin(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*common.TokenResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.SmsLogin", in)
	out := new(common.TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) Info(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.Info", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) Logout(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*AuthorizeResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.Logout", in)
	out := new(AuthorizeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthorizeService service

type AuthorizeServiceHandler interface {
	//客户注册
	Register(context.Context, *Customer, *CustomerResponse) error
	//账号登录
	Login(context.Context, *AuthorizeRequest, *common.TokenResponse) error
	//短信登录
	SmsLogin(context.Context, *AuthorizeRequest, *common.TokenResponse) error
	//获得当前客户
	Info(context.Context, *common.Empty, *CustomerResponse) error
	//安全退出
	Logout(context.Context, *common.Empty, *AuthorizeResponse) error
}

func RegisterAuthorizeServiceHandler(s server.Server, hdlr AuthorizeServiceHandler, opts ...server.HandlerOption) error {
	type authorizeService interface {
		Register(ctx context.Context, in *Customer, out *CustomerResponse) error
		Login(ctx context.Context, in *AuthorizeRequest, out *common.TokenResponse) error
		SmsLogin(ctx context.Context, in *AuthorizeRequest, out *common.TokenResponse) error
		Info(ctx context.Context, in *common.Empty, out *CustomerResponse) error
		Logout(ctx context.Context, in *common.Empty, out *AuthorizeResponse) error
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

func (h *authorizeServiceHandler) Register(ctx context.Context, in *Customer, out *CustomerResponse) error {
	return h.AuthorizeServiceHandler.Register(ctx, in, out)
}

func (h *authorizeServiceHandler) Login(ctx context.Context, in *AuthorizeRequest, out *common.TokenResponse) error {
	return h.AuthorizeServiceHandler.Login(ctx, in, out)
}

func (h *authorizeServiceHandler) SmsLogin(ctx context.Context, in *AuthorizeRequest, out *common.TokenResponse) error {
	return h.AuthorizeServiceHandler.SmsLogin(ctx, in, out)
}

func (h *authorizeServiceHandler) Info(ctx context.Context, in *common.Empty, out *CustomerResponse) error {
	return h.AuthorizeServiceHandler.Info(ctx, in, out)
}

func (h *authorizeServiceHandler) Logout(ctx context.Context, in *common.Empty, out *AuthorizeResponse) error {
	return h.AuthorizeServiceHandler.Logout(ctx, in, out)
}