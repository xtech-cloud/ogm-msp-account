// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/account/profile.proto

package account

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Profile service

func NewProfileEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Profile service

type ProfileService interface {
	// 获取
	Query(ctx context.Context, in *QueryProfileRequest, opts ...client.CallOption) (*QueryProfileResponse, error)
	// 更新
	Update(ctx context.Context, in *UpdateProfileRequest, opts ...client.CallOption) (*UpdateProfileResponse, error)
}

type profileService struct {
	c    client.Client
	name string
}

func NewProfileService(name string, c client.Client) ProfileService {
	return &profileService{
		c:    c,
		name: name,
	}
}

func (c *profileService) Query(ctx context.Context, in *QueryProfileRequest, opts ...client.CallOption) (*QueryProfileResponse, error) {
	req := c.c.NewRequest(c.name, "Profile.Query", in)
	out := new(QueryProfileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileService) Update(ctx context.Context, in *UpdateProfileRequest, opts ...client.CallOption) (*UpdateProfileResponse, error) {
	req := c.c.NewRequest(c.name, "Profile.Update", in)
	out := new(UpdateProfileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profile service

type ProfileHandler interface {
	// 获取
	Query(context.Context, *QueryProfileRequest, *QueryProfileResponse) error
	// 更新
	Update(context.Context, *UpdateProfileRequest, *UpdateProfileResponse) error
}

func RegisterProfileHandler(s server.Server, hdlr ProfileHandler, opts ...server.HandlerOption) error {
	type profile interface {
		Query(ctx context.Context, in *QueryProfileRequest, out *QueryProfileResponse) error
		Update(ctx context.Context, in *UpdateProfileRequest, out *UpdateProfileResponse) error
	}
	type Profile struct {
		profile
	}
	h := &profileHandler{hdlr}
	return s.Handle(s.NewHandler(&Profile{h}, opts...))
}

type profileHandler struct {
	ProfileHandler
}

func (h *profileHandler) Query(ctx context.Context, in *QueryProfileRequest, out *QueryProfileResponse) error {
	return h.ProfileHandler.Query(ctx, in, out)
}

func (h *profileHandler) Update(ctx context.Context, in *UpdateProfileRequest, out *UpdateProfileResponse) error {
	return h.ProfileHandler.Update(ctx, in, out)
}
