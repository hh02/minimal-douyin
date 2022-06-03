package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc/followservice"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var followClient followservice.Client

func initFollowRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := followservice.NewClient(
		constants.FollowServiceName,
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	followClient = c
}


func CreateFollow(ctx context.Context, req *followrpc.CreateFollowRequest) error {
	resp, err := followClient.CreateFollow(ctx, req)
	if err != nil {
		return err
	}
	if resp.Status.StatusCode != errno.SuccessCode {
		return errno.Status2ErrorNo(resp.Status)
	}
	return nil
}

func DeleteFollow(ctx context.Context, req *followrpc.DeleteFollowRequest) error {
	resp, err := followClient.DeleteFollow(ctx, req)
	if err != nil {
		return err
	}
	if resp.Status.StatusCode != errno.SuccessCode {
		return errno.Status2ErrorNo(resp.Status)
	}
	return nil
}

func QueryFollow(ctx context.Context, req *followrpc.QueryFollowRequest) ([]*userrpc.User, error){
	resp, err := followClient.QueryFollow(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Status.StatusCode != errno.SuccessCode {
		return nil, errno.Status2ErrorNo(resp.Status)
	}
	return resp.Users, nil
}

func QueryFollower(ctx context.Context, req *followrpc.QueryFollowerRequest) ([]*userrpc.User, error){
	resp, err := followClient.QueryFollower(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Status.StatusCode != errno.SuccessCode {
		return nil, errno.Status2ErrorNo(resp.Status)
	}
	return resp.Users, nil


}