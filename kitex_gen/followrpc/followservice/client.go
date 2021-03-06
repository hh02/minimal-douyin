// Code generated by Kitex v0.3.1. DO NOT EDIT.

package followservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateFollow(ctx context.Context, Req *followrpc.CreateFollowRequest, callOptions ...callopt.Option) (r *followrpc.CreateFollowResponse, err error)
	DeleteFollow(ctx context.Context, Req *followrpc.DeleteFollowRequest, callOptions ...callopt.Option) (r *followrpc.DeleteFollowResponse, err error)
	QueryFollow(ctx context.Context, Req *followrpc.QueryFollowRequest, callOptions ...callopt.Option) (r *followrpc.QueryFollowResponse, err error)
	QueryFollower(ctx context.Context, Req *followrpc.QueryFollowerRequest, callOptions ...callopt.Option) (r *followrpc.QueryFollowerResponse, err error)
	CheckFollow(ctx context.Context, Req *followrpc.CheckFollowRequest, callOptions ...callopt.Option) (r *followrpc.CheckFollowResponse, err error)
	MCheckFollow(ctx context.Context, Req *followrpc.MCheckFollowRequest, callOptions ...callopt.Option) (r *followrpc.MCheckFollowResponse, err error)
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
	return &kFollowServiceClient{
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

type kFollowServiceClient struct {
	*kClient
}

func (p *kFollowServiceClient) CreateFollow(ctx context.Context, Req *followrpc.CreateFollowRequest, callOptions ...callopt.Option) (r *followrpc.CreateFollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateFollow(ctx, Req)
}

func (p *kFollowServiceClient) DeleteFollow(ctx context.Context, Req *followrpc.DeleteFollowRequest, callOptions ...callopt.Option) (r *followrpc.DeleteFollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteFollow(ctx, Req)
}

func (p *kFollowServiceClient) QueryFollow(ctx context.Context, Req *followrpc.QueryFollowRequest, callOptions ...callopt.Option) (r *followrpc.QueryFollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryFollow(ctx, Req)
}

func (p *kFollowServiceClient) QueryFollower(ctx context.Context, Req *followrpc.QueryFollowerRequest, callOptions ...callopt.Option) (r *followrpc.QueryFollowerResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryFollower(ctx, Req)
}

func (p *kFollowServiceClient) CheckFollow(ctx context.Context, Req *followrpc.CheckFollowRequest, callOptions ...callopt.Option) (r *followrpc.CheckFollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckFollow(ctx, Req)
}

func (p *kFollowServiceClient) MCheckFollow(ctx context.Context, Req *followrpc.MCheckFollowRequest, callOptions ...callopt.Option) (r *followrpc.MCheckFollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MCheckFollow(ctx, Req)
}
