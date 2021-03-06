package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc/userservice"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
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

// GetUser 远程调用 userservice 获取 User
func GetUser(ctx context.Context, req *userrpc.GetUserRequest) (*userrpc.User, error) {
	resp, err := userClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Status.StatusCode != errno.SuccessCode {
		return nil, errno.Status2ErrorNo(resp.Status)
	}

	return resp.User, nil
}

func MGetUser(ctx context.Context, req *userrpc.MGetUserRequest) ([]*userrpc.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Status.StatusCode != errno.SuccessCode {
		return nil, errno.Status2ErrorNo(resp.Status)
	}

	return resp.Users, nil
}
