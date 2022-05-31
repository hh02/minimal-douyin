package rpc

import (
	"context"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc/followservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient followservice.Client

func initFollowRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := followservice.NewClient(
		constants.FollowServiceName,
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)

	if err != nil {
		panic(err)
	}

	userClient = c
}

// 调用  follow 服务中的查询关注函数
func QueryFollow(ctx context.Context, req *followrpc.QueryFollowRequest) (int, error) {
	resp, err := userClient.QueryFollow(ctx, req)
	if err != nil {
		return 0, err
	}

	if resp.Status.StatusCode != errno.SuccessCode {
		return 0, errno.Status2ErrorNo(resp.Status)
	}

	return len(resp.Users), nil
}

// 调用 follow 中的查询粉丝函数
func QueryFollower(ctx context.Context, req *followrpc.QueryFollowerRequest) (int, error) {
	resp, err := userClient.QueryFollower(ctx, req)
	if err != nil {
		return 0, err
	}

	if resp.Status.StatusCode != errno.SuccessCode {
		return 0, errno.Status2ErrorNo(resp.Status)
	}

	return len(resp.Users), nil
}
