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

// 调用  follow 服务中的查询是否关注
func IsFollow(ctx context.Context, req *followrpc.CheckFollowRequest) (bool, error) {
	resp, err := userClient.CheckFollow(ctx, req)
	if err != nil {
		return false, err
	}

	if resp.Status.StatusCode != errno.SuccessCode {
		return false, errno.Status2ErrorNo(resp.Status)
	}

	return resp.IsFollow, nil
}
