// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: attendanceService.proto

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

// Api Endpoints for AttendanceService service

func NewAttendanceServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AttendanceService service

type AttendanceService interface {
	Create(ctx context.Context, in *Attendance, opts ...client.CallOption) (*AttendanceResponse, error)
	Update(ctx context.Context, in *Attendance, opts ...client.CallOption) (*AttendanceResponse, error)
	Delete(ctx context.Context, in *Attendance, opts ...client.CallOption) (*AttendanceResponse, error)
	Get(ctx context.Context, in *Attendance, opts ...client.CallOption) (*AttendanceResponse, error)
	Search(ctx context.Context, in *AttendanceRequest, opts ...client.CallOption) (*AttendanceResponse, error)
}

type attendanceService struct {
	c    client.Client
	name string
}

func NewAttendanceService(name string, c client.Client) AttendanceService {
	return &attendanceService{
		c:    c,
		name: name,
	}
}

func (c *attendanceService) Create(ctx context.Context, in *Attendance, opts ...client.CallOption) (*AttendanceResponse, error) {
	req := c.c.NewRequest(c.name, "AttendanceService.Create", in)
	out := new(AttendanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendanceService) Update(ctx context.Context, in *Attendance, opts ...client.CallOption) (*AttendanceResponse, error) {
	req := c.c.NewRequest(c.name, "AttendanceService.Update", in)
	out := new(AttendanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendanceService) Delete(ctx context.Context, in *Attendance, opts ...client.CallOption) (*AttendanceResponse, error) {
	req := c.c.NewRequest(c.name, "AttendanceService.Delete", in)
	out := new(AttendanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendanceService) Get(ctx context.Context, in *Attendance, opts ...client.CallOption) (*AttendanceResponse, error) {
	req := c.c.NewRequest(c.name, "AttendanceService.Get", in)
	out := new(AttendanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendanceService) Search(ctx context.Context, in *AttendanceRequest, opts ...client.CallOption) (*AttendanceResponse, error) {
	req := c.c.NewRequest(c.name, "AttendanceService.Search", in)
	out := new(AttendanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AttendanceService service

type AttendanceServiceHandler interface {
	Create(context.Context, *Attendance, *AttendanceResponse) error
	Update(context.Context, *Attendance, *AttendanceResponse) error
	Delete(context.Context, *Attendance, *AttendanceResponse) error
	Get(context.Context, *Attendance, *AttendanceResponse) error
	Search(context.Context, *AttendanceRequest, *AttendanceResponse) error
}

func RegisterAttendanceServiceHandler(s server.Server, hdlr AttendanceServiceHandler, opts ...server.HandlerOption) error {
	type attendanceService interface {
		Create(ctx context.Context, in *Attendance, out *AttendanceResponse) error
		Update(ctx context.Context, in *Attendance, out *AttendanceResponse) error
		Delete(ctx context.Context, in *Attendance, out *AttendanceResponse) error
		Get(ctx context.Context, in *Attendance, out *AttendanceResponse) error
		Search(ctx context.Context, in *AttendanceRequest, out *AttendanceResponse) error
	}
	type AttendanceService struct {
		attendanceService
	}
	h := &attendanceServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AttendanceService{h}, opts...))
}

type attendanceServiceHandler struct {
	AttendanceServiceHandler
}

func (h *attendanceServiceHandler) Create(ctx context.Context, in *Attendance, out *AttendanceResponse) error {
	return h.AttendanceServiceHandler.Create(ctx, in, out)
}

func (h *attendanceServiceHandler) Update(ctx context.Context, in *Attendance, out *AttendanceResponse) error {
	return h.AttendanceServiceHandler.Update(ctx, in, out)
}

func (h *attendanceServiceHandler) Delete(ctx context.Context, in *Attendance, out *AttendanceResponse) error {
	return h.AttendanceServiceHandler.Delete(ctx, in, out)
}

func (h *attendanceServiceHandler) Get(ctx context.Context, in *Attendance, out *AttendanceResponse) error {
	return h.AttendanceServiceHandler.Get(ctx, in, out)
}

func (h *attendanceServiceHandler) Search(ctx context.Context, in *AttendanceRequest, out *AttendanceResponse) error {
	return h.AttendanceServiceHandler.Search(ctx, in, out)
}
