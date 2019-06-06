// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: showservice/proto/show/show.proto

package show

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

import (
	context "context"

	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShowService service

type ShowService interface {
	GetAllShows(ctx context.Context, in *GetAllShowsRequest, opts ...client.CallOption) (*GetAllShowsResponse, error)
	GetShow(ctx context.Context, in *GetShowRequest, opts ...client.CallOption) (*GetShowResponse, error)
	AddShow(ctx context.Context, in *AddShowRequest, opts ...client.CallOption) (*AddShowResponse, error)
	RemoveShow(ctx context.Context, in *RemoveShowRequest, opts ...client.CallOption) (*RemoveShowResponse, error)
}

type showService struct {
	c    client.Client
	name string
}

func NewShowService(name string, c client.Client) ShowService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "show"
	}
	return &showService{
		c:    c,
		name: name,
	}
}

func (c *showService) GetAllShows(ctx context.Context, in *GetAllShowsRequest, opts ...client.CallOption) (*GetAllShowsResponse, error) {
	req := c.c.NewRequest(c.name, "ShowService.GetAllShows", in)
	out := new(GetAllShowsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showService) GetShow(ctx context.Context, in *GetShowRequest, opts ...client.CallOption) (*GetShowResponse, error) {
	req := c.c.NewRequest(c.name, "ShowService.GetShow", in)
	out := new(GetShowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showService) AddShow(ctx context.Context, in *AddShowRequest, opts ...client.CallOption) (*AddShowResponse, error) {
	req := c.c.NewRequest(c.name, "ShowService.AddShow", in)
	out := new(AddShowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showService) RemoveShow(ctx context.Context, in *RemoveShowRequest, opts ...client.CallOption) (*RemoveShowResponse, error) {
	req := c.c.NewRequest(c.name, "ShowService.RemoveShow", in)
	out := new(RemoveShowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShowService service

type ShowServiceHandler interface {
	GetAllShows(context.Context, *GetAllShowsRequest, *GetAllShowsResponse) error
	GetShow(context.Context, *GetShowRequest, *GetShowResponse) error
	AddShow(context.Context, *AddShowRequest, *AddShowResponse) error
	RemoveShow(context.Context, *RemoveShowRequest, *RemoveShowResponse) error
}

func RegisterShowServiceHandler(s server.Server, hdlr ShowServiceHandler, opts ...server.HandlerOption) error {
	type showService interface {
		GetAllShows(ctx context.Context, in *GetAllShowsRequest, out *GetAllShowsResponse) error
		GetShow(ctx context.Context, in *GetShowRequest, out *GetShowResponse) error
		AddShow(ctx context.Context, in *AddShowRequest, out *AddShowResponse) error
		RemoveShow(ctx context.Context, in *RemoveShowRequest, out *RemoveShowResponse) error
	}
	type ShowService struct {
		showService
	}
	h := &showServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ShowService{h}, opts...))
}

type showServiceHandler struct {
	ShowServiceHandler
}

func (h *showServiceHandler) GetAllShows(ctx context.Context, in *GetAllShowsRequest, out *GetAllShowsResponse) error {
	return h.ShowServiceHandler.GetAllShows(ctx, in, out)
}

func (h *showServiceHandler) GetShow(ctx context.Context, in *GetShowRequest, out *GetShowResponse) error {
	return h.ShowServiceHandler.GetShow(ctx, in, out)
}

func (h *showServiceHandler) AddShow(ctx context.Context, in *AddShowRequest, out *AddShowResponse) error {
	return h.ShowServiceHandler.AddShow(ctx, in, out)
}

func (h *showServiceHandler) RemoveShow(ctx context.Context, in *RemoveShowRequest, out *RemoveShowResponse) error {
	return h.ShowServiceHandler.RemoveShow(ctx, in, out)
}
