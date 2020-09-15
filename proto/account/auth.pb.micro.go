// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/account/auth.proto

package account

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthService interface {
	// 注册
	Signup(ctx context.Context, in *SignupRequest, opts ...client.CallOption) (*SignupResponse, error)
	// 登录
	Signin(ctx context.Context, in *SigninRequest, opts ...client.CallOption) (*SigninResponse, error)
	// 登出
	Signout(ctx context.Context, in *SignoutRequest, opts ...client.CallOption) (*SignoutResponse, error)
	// 重置密码
	ResetPasswd(ctx context.Context, in *ResetPasswdRequest, opts ...client.CallOption) (*ResetPasswdResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Signup(ctx context.Context, in *SignupRequest, opts ...client.CallOption) (*SignupResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Signup", in)
	out := new(SignupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Signin(ctx context.Context, in *SigninRequest, opts ...client.CallOption) (*SigninResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Signin", in)
	out := new(SigninResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Signout(ctx context.Context, in *SignoutRequest, opts ...client.CallOption) (*SignoutResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Signout", in)
	out := new(SignoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ResetPasswd(ctx context.Context, in *ResetPasswdRequest, opts ...client.CallOption) (*ResetPasswdResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ResetPasswd", in)
	out := new(ResetPasswdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	// 注册
	Signup(context.Context, *SignupRequest, *SignupResponse) error
	// 登录
	Signin(context.Context, *SigninRequest, *SigninResponse) error
	// 登出
	Signout(context.Context, *SignoutRequest, *SignoutResponse) error
	// 重置密码
	ResetPasswd(context.Context, *ResetPasswdRequest, *ResetPasswdResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		Signup(ctx context.Context, in *SignupRequest, out *SignupResponse) error
		Signin(ctx context.Context, in *SigninRequest, out *SigninResponse) error
		Signout(ctx context.Context, in *SignoutRequest, out *SignoutResponse) error
		ResetPasswd(ctx context.Context, in *ResetPasswdRequest, out *ResetPasswdResponse) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Signup(ctx context.Context, in *SignupRequest, out *SignupResponse) error {
	return h.AuthHandler.Signup(ctx, in, out)
}

func (h *authHandler) Signin(ctx context.Context, in *SigninRequest, out *SigninResponse) error {
	return h.AuthHandler.Signin(ctx, in, out)
}

func (h *authHandler) Signout(ctx context.Context, in *SignoutRequest, out *SignoutResponse) error {
	return h.AuthHandler.Signout(ctx, in, out)
}

func (h *authHandler) ResetPasswd(ctx context.Context, in *ResetPasswdRequest, out *ResetPasswdResponse) error {
	return h.AuthHandler.ResetPasswd(ctx, in, out)
}
