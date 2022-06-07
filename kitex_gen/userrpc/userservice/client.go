// Code generated by Kitex v0.3.1. DO NOT EDIT.

package userservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateUser(ctx context.Context, Req *userrpc.CreateUserRequest, callOptions ...callopt.Option) (r *userrpc.CreateUserResponse, err error)
	GetUser(ctx context.Context, Req *userrpc.GetUserRequest, callOptions ...callopt.Option) (r *userrpc.GetUserResponse, err error)
	MGetUser(ctx context.Context, Req *userrpc.MGetUserRequest, callOptions ...callopt.Option) (r *userrpc.MGetUserResponse, err error)
	CheckUser(ctx context.Context, Req *userrpc.CheckUserRequest, callOptions ...callopt.Option) (r *userrpc.CheckUserResponse, err error)
	AddFollowCount(ctx context.Context, Req *userrpc.AddFollowCountRequest, callOptions ...callopt.Option) (r *userrpc.AddFollowCountResponse, err error)
	AddFollowerCount(ctx context.Context, Req *userrpc.AddFollowerCountRequest, callOptions ...callopt.Option) (r *userrpc.AddFollowerCountResponse, err error)
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

func (p *kUserServiceClient) CreateUser(ctx context.Context, Req *userrpc.CreateUserRequest, callOptions ...callopt.Option) (r *userrpc.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, Req)
}

func (p *kUserServiceClient) GetUser(ctx context.Context, Req *userrpc.GetUserRequest, callOptions ...callopt.Option) (r *userrpc.GetUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUser(ctx, Req)
}

func (p *kUserServiceClient) MGetUser(ctx context.Context, Req *userrpc.MGetUserRequest, callOptions ...callopt.Option) (r *userrpc.MGetUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetUser(ctx, Req)
}

func (p *kUserServiceClient) CheckUser(ctx context.Context, Req *userrpc.CheckUserRequest, callOptions ...callopt.Option) (r *userrpc.CheckUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckUser(ctx, Req)
}

func (p *kUserServiceClient) AddFollowCount(ctx context.Context, Req *userrpc.AddFollowCountRequest, callOptions ...callopt.Option) (r *userrpc.AddFollowCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddFollowCount(ctx, Req)
}

func (p *kUserServiceClient) AddFollowerCount(ctx context.Context, Req *userrpc.AddFollowerCountRequest, callOptions ...callopt.Option) (r *userrpc.AddFollowerCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddFollowerCount(ctx, Req)
}
