// Code generated by Kitex v0.3.1. DO NOT EDIT.

package favoriteservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/hh02/minimal-douyin/kitex_gen/favoriterpc"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateFavorite(ctx context.Context, Req *favoriterpc.CreateFavoriteRequest, callOptions ...callopt.Option) (r *favoriterpc.CreateFavoriteResponse, err error)
	DeleteFavorite(ctx context.Context, Req *favoriterpc.DeleteFavoriteRequest, callOptions ...callopt.Option) (r *favoriterpc.DeleteFavoriteResponse, err error)
	QueryFavoriteByUserId(ctx context.Context, Req *favoriterpc.QueryFavoriteByUserIdRequest, callOptions ...callopt.Option) (r *favoriterpc.QueryFavoriteByUserIdResponse, err error)
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
	return &kFavoriteServiceClient{
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

type kFavoriteServiceClient struct {
	*kClient
}

func (p *kFavoriteServiceClient) CreateFavorite(ctx context.Context, Req *favoriterpc.CreateFavoriteRequest, callOptions ...callopt.Option) (r *favoriterpc.CreateFavoriteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateFavorite(ctx, Req)
}

func (p *kFavoriteServiceClient) DeleteFavorite(ctx context.Context, Req *favoriterpc.DeleteFavoriteRequest, callOptions ...callopt.Option) (r *favoriterpc.DeleteFavoriteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteFavorite(ctx, Req)
}

func (p *kFavoriteServiceClient) QueryFavoriteByUserId(ctx context.Context, Req *favoriterpc.QueryFavoriteByUserIdRequest, callOptions ...callopt.Option) (r *favoriterpc.QueryFavoriteByUserIdResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryFavoriteByUserId(ctx, Req)
}