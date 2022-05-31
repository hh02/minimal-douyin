package rpc

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/hh02/minimal-douyin/kitex_gen/douyin/core"
	"github.com/hh02/minimal-douyin/kitex_gen/douyin/core/userservice"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
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
	userClient = c
}

// UserInfo 返回用户信息
func UsreInfo(ctx context.Context, req *core.UserRequest) (*core.User , error) {
	// 调用 user 服务
	resp, err := userClient.User(ctx, req)
	if err != nil {
		return nil , err
	}
	if resp.StatusCode != 0 {
		return nil,errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.User,nil
}

// Register check user info
func Register(ctx context.Context, req *core.UserRegisterRequest) (int64, error) {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.UserId, nil
}

// Login check user info
func Login(ctx context.Context, req *core.UserLoginRequest) (int64, error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.UserId, nil
}
