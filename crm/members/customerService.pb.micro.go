// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: customerService.proto

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

// Api Endpoints for CustomerService service

func NewCustomerServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CustomerService service

type CustomerService interface {
	//手动添加用户
	Create(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error)
	//手动添加单位用户
	CreateCompany(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error)
	//手动修改单位用户
	UpdateCompany(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error)
	//从粉丝添加用户
	CreateByFan(ctx context.Context, in *Fan, opts ...client.CallOption) (*CustomerResponse, error)
	//修改客户
	Update(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error)
	//删除客户
	Delete(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error)
	//锁定客户
	Lock(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error)
	//解锁客户
	Unlock(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error)
	//获得客户
	Get(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error)
	//获得客户详情
	Detail(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error)
	//根据ids获得客户
	List(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error)
	//查询客户
	Search(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error)
	//设置会员标签
	SetTags(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error)
	//设置会员卡
	SetCards(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error)
	//获取已绑定手机用户(SRV专用)
	GetByMobile(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error)
	//获取已绑定邮箱用户(SRV专用)
	GetByEmail(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error)
}

type customerService struct {
	c    client.Client
	name string
}

func NewCustomerService(name string, c client.Client) CustomerService {
	return &customerService{
		c:    c,
		name: name,
	}
}

func (c *customerService) Create(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.Create", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) CreateCompany(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.CreateCompany", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) UpdateCompany(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.UpdateCompany", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) CreateByFan(ctx context.Context, in *Fan, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.CreateByFan", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) Update(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.Update", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) Delete(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.Delete", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) Lock(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.Lock", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) Unlock(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.Unlock", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) Get(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.Get", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) Detail(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.Detail", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) List(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.List", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) Search(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.Search", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) SetTags(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.SetTags", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) SetCards(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.SetCards", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetByMobile(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.GetByMobile", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetByEmail(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.GetByEmail", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomerService service

type CustomerServiceHandler interface {
	//手动添加用户
	Create(context.Context, *Customer, *CustomerResponse) error
	//手动添加单位用户
	CreateCompany(context.Context, *Customer, *CustomerResponse) error
	//手动修改单位用户
	UpdateCompany(context.Context, *Customer, *CustomerResponse) error
	//从粉丝添加用户
	CreateByFan(context.Context, *Fan, *CustomerResponse) error
	//修改客户
	Update(context.Context, *Customer, *CustomerResponse) error
	//删除客户
	Delete(context.Context, *CustomerRequest, *CustomerResponse) error
	//锁定客户
	Lock(context.Context, *CustomerRequest, *CustomerResponse) error
	//解锁客户
	Unlock(context.Context, *CustomerRequest, *CustomerResponse) error
	//获得客户
	Get(context.Context, *Customer, *CustomerResponse) error
	//获得客户详情
	Detail(context.Context, *Customer, *CustomerResponse) error
	//根据ids获得客户
	List(context.Context, *CustomerRequest, *CustomerResponse) error
	//查询客户
	Search(context.Context, *CustomerRequest, *CustomerResponse) error
	//设置会员标签
	SetTags(context.Context, *Customer, *CustomerResponse) error
	//设置会员卡
	SetCards(context.Context, *Customer, *CustomerResponse) error
	//获取已绑定手机用户(SRV专用)
	GetByMobile(context.Context, *CustomerRequest, *CustomerResponse) error
	//获取已绑定邮箱用户(SRV专用)
	GetByEmail(context.Context, *CustomerRequest, *CustomerResponse) error
}

func RegisterCustomerServiceHandler(s server.Server, hdlr CustomerServiceHandler, opts ...server.HandlerOption) error {
	type customerService interface {
		Create(ctx context.Context, in *Customer, out *CustomerResponse) error
		CreateCompany(ctx context.Context, in *Customer, out *CustomerResponse) error
		UpdateCompany(ctx context.Context, in *Customer, out *CustomerResponse) error
		CreateByFan(ctx context.Context, in *Fan, out *CustomerResponse) error
		Update(ctx context.Context, in *Customer, out *CustomerResponse) error
		Delete(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error
		Lock(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error
		Unlock(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error
		Get(ctx context.Context, in *Customer, out *CustomerResponse) error
		Detail(ctx context.Context, in *Customer, out *CustomerResponse) error
		List(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error
		Search(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error
		SetTags(ctx context.Context, in *Customer, out *CustomerResponse) error
		SetCards(ctx context.Context, in *Customer, out *CustomerResponse) error
		GetByMobile(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error
		GetByEmail(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error
	}
	type CustomerService struct {
		customerService
	}
	h := &customerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CustomerService{h}, opts...))
}

type customerServiceHandler struct {
	CustomerServiceHandler
}

func (h *customerServiceHandler) Create(ctx context.Context, in *Customer, out *CustomerResponse) error {
	return h.CustomerServiceHandler.Create(ctx, in, out)
}

func (h *customerServiceHandler) CreateCompany(ctx context.Context, in *Customer, out *CustomerResponse) error {
	return h.CustomerServiceHandler.CreateCompany(ctx, in, out)
}

func (h *customerServiceHandler) UpdateCompany(ctx context.Context, in *Customer, out *CustomerResponse) error {
	return h.CustomerServiceHandler.UpdateCompany(ctx, in, out)
}

func (h *customerServiceHandler) CreateByFan(ctx context.Context, in *Fan, out *CustomerResponse) error {
	return h.CustomerServiceHandler.CreateByFan(ctx, in, out)
}

func (h *customerServiceHandler) Update(ctx context.Context, in *Customer, out *CustomerResponse) error {
	return h.CustomerServiceHandler.Update(ctx, in, out)
}

func (h *customerServiceHandler) Delete(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error {
	return h.CustomerServiceHandler.Delete(ctx, in, out)
}

func (h *customerServiceHandler) Lock(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error {
	return h.CustomerServiceHandler.Lock(ctx, in, out)
}

func (h *customerServiceHandler) Unlock(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error {
	return h.CustomerServiceHandler.Unlock(ctx, in, out)
}

func (h *customerServiceHandler) Get(ctx context.Context, in *Customer, out *CustomerResponse) error {
	return h.CustomerServiceHandler.Get(ctx, in, out)
}

func (h *customerServiceHandler) Detail(ctx context.Context, in *Customer, out *CustomerResponse) error {
	return h.CustomerServiceHandler.Detail(ctx, in, out)
}

func (h *customerServiceHandler) List(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error {
	return h.CustomerServiceHandler.List(ctx, in, out)
}

func (h *customerServiceHandler) Search(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error {
	return h.CustomerServiceHandler.Search(ctx, in, out)
}

func (h *customerServiceHandler) SetTags(ctx context.Context, in *Customer, out *CustomerResponse) error {
	return h.CustomerServiceHandler.SetTags(ctx, in, out)
}

func (h *customerServiceHandler) SetCards(ctx context.Context, in *Customer, out *CustomerResponse) error {
	return h.CustomerServiceHandler.SetCards(ctx, in, out)
}

func (h *customerServiceHandler) GetByMobile(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error {
	return h.CustomerServiceHandler.GetByMobile(ctx, in, out)
}

func (h *customerServiceHandler) GetByEmail(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error {
	return h.CustomerServiceHandler.GetByEmail(ctx, in, out)
}
