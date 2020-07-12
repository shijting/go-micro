// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: Courses.proto

package courses

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// Api Endpoints for CoursesService service

func NewCoursesServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CoursesService service

type CoursesService interface {
	ListForTop(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	GetDetail(ctx context.Context, in *DetailRequest, opts ...client.CallOption) (*DetailResponse, error)
}

type coursesService struct {
	c    client.Client
	name string
}

func NewCoursesService(name string, c client.Client) CoursesService {
	return &coursesService{
		c:    c,
		name: name,
	}
}

func (c *coursesService) ListForTop(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "CoursesService.ListForTop", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesService) GetDetail(ctx context.Context, in *DetailRequest, opts ...client.CallOption) (*DetailResponse, error) {
	req := c.c.NewRequest(c.name, "CoursesService.GetDetail", in)
	out := new(DetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CoursesService service

type CoursesServiceHandler interface {
	ListForTop(context.Context, *ListRequest, *ListResponse) error
	GetDetail(context.Context, *DetailRequest, *DetailResponse) error
}

func RegisterCoursesServiceHandler(s server.Server, hdlr CoursesServiceHandler, opts ...server.HandlerOption) error {
	type coursesService interface {
		ListForTop(ctx context.Context, in *ListRequest, out *ListResponse) error
		GetDetail(ctx context.Context, in *DetailRequest, out *DetailResponse) error
	}
	type CoursesService struct {
		coursesService
	}
	h := &coursesServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CoursesService{h}, opts...))
}

type coursesServiceHandler struct {
	CoursesServiceHandler
}

func (h *coursesServiceHandler) ListForTop(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.CoursesServiceHandler.ListForTop(ctx, in, out)
}

func (h *coursesServiceHandler) GetDetail(ctx context.Context, in *DetailRequest, out *DetailResponse) error {
	return h.CoursesServiceHandler.GetDetail(ctx, in, out)
}