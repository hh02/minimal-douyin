// Code generated by Kitex v0.3.1. DO NOT EDIT.

package userservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/hh02/minimal-douyin/kitex_gen/douyin/core"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UserRegister(ctx context.Context, Req *core.UserRegisterRequest, callOptions ...callopt.Option) (r *core.UserRegisterResponse, err error)
	UserLogin(ctx context.Context, Req *core.UserLoginRequest, callOptions ...callopt.Option) (r *core.UserLoginResponse, err error)
	User(ctx context.Context, Req *core.UserRequest, callOptions ...callopt.Option) (r *core.UserResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) UserRegister(ctx context.Context, Req *core.UserRegisterRequest, callOptions ...callopt.Option) (r *core.UserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserRegister(ctx, Req)
}

func (p *kUserServiceClient) UserLogin(ctx context.Context, Req *core.UserLoginRequest, callOptions ...callopt.Option) (r *core.UserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserLogin(ctx, Req)
}

func (p *kUserServiceClient) User(ctx context.Context, Req *core.UserRequest, callOptions ...callopt.Option) (r *core.UserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.User(ctx, Req)
}
