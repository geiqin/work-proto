// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: customerSecurityService.proto

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

// Api Endpoints for CustomerSecurityService service

func NewCustomerSecurityServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CustomerSecurityService service

type CustomerSecurityService interface {
	//显示信息
	Display(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error)
	//判断有无密码保护
	HasPwd(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error)
	//创建密码(未设置密码前可用)
	CreatePwd(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error)
	//修改密码
	ModifyPwd(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error)
	//绑定手机号
	BindMobile(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error)
	//绑定邮箱
	BindEmail(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error)
	//注销账户
	Destroy(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error)
}

type customerSecurityService struct {
	c    client.Client
	name string
}

func NewCustomerSecurityService(name string, c client.Client) CustomerSecurityService {
	return &customerSecurityService{
		c:    c,
		name: name,
	}
}

func (c *customerSecurityService) Display(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerSecurityService.Display", in)
	out := new(CustomerSecurityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerSecurityService) HasPwd(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerSecurityService.HasPwd", in)
	out := new(CustomerSecurityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerSecurityService) CreatePwd(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerSecurityService.CreatePwd", in)
	out := new(CustomerSecurityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerSecurityService) ModifyPwd(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerSecurityService.ModifyPwd", in)
	out := new(CustomerSecurityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerSecurityService) BindMobile(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerSecurityService.BindMobile", in)
	out := new(CustomerSecurityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerSecurityService) BindEmail(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerSecurityService.BindEmail", in)
	out := new(CustomerSecurityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerSecurityService) Destroy(ctx context.Context, in *CustomerSecurityRequest, opts ...client.CallOption) (*CustomerSecurityResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerSecurityService.Destroy", in)
	out := new(CustomerSecurityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomerSecurityService service

type CustomerSecurityServiceHandler interface {
	//显示信息
	Display(context.Context, *CustomerSecurityRequest, *CustomerSecurityResponse) error
	//判断有无密码保护
	HasPwd(context.Context, *CustomerSecurityRequest, *CustomerSecurityResponse) error
	//创建密码(未设置密码前可用)
	CreatePwd(context.Context, *CustomerSecurityRequest, *CustomerSecurityResponse) error
	//修改密码
	ModifyPwd(context.Context, *CustomerSecurityRequest, *CustomerSecurityResponse) error
	//绑定手机号
	BindMobile(context.Context, *CustomerSecurityRequest, *CustomerSecurityResponse) error
	//绑定邮箱
	BindEmail(context.Context, *CustomerSecurityRequest, *CustomerSecurityResponse) error
	//注销账户
	Destroy(context.Context, *CustomerSecurityRequest, *CustomerSecurityResponse) error
}

func RegisterCustomerSecurityServiceHandler(s server.Server, hdlr CustomerSecurityServiceHandler, opts ...server.HandlerOption) error {
	type customerSecurityService interface {
		Display(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error
		HasPwd(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error
		CreatePwd(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error
		ModifyPwd(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error
		BindMobile(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error
		BindEmail(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error
		Destroy(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error
	}
	type CustomerSecurityService struct {
		customerSecurityService
	}
	h := &customerSecurityServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CustomerSecurityService{h}, opts...))
}

type customerSecurityServiceHandler struct {
	CustomerSecurityServiceHandler
}

func (h *customerSecurityServiceHandler) Display(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error {
	return h.CustomerSecurityServiceHandler.Display(ctx, in, out)
}

func (h *customerSecurityServiceHandler) HasPwd(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error {
	return h.CustomerSecurityServiceHandler.HasPwd(ctx, in, out)
}

func (h *customerSecurityServiceHandler) CreatePwd(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error {
	return h.CustomerSecurityServiceHandler.CreatePwd(ctx, in, out)
}

func (h *customerSecurityServiceHandler) ModifyPwd(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error {
	return h.CustomerSecurityServiceHandler.ModifyPwd(ctx, in, out)
}

func (h *customerSecurityServiceHandler) BindMobile(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error {
	return h.CustomerSecurityServiceHandler.BindMobile(ctx, in, out)
}

func (h *customerSecurityServiceHandler) BindEmail(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error {
	return h.CustomerSecurityServiceHandler.BindEmail(ctx, in, out)
}

func (h *customerSecurityServiceHandler) Destroy(ctx context.Context, in *CustomerSecurityRequest, out *CustomerSecurityResponse) error {
	return h.CustomerSecurityServiceHandler.Destroy(ctx, in, out)
}
