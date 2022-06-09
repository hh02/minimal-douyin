package rpc

import (
	"context"
	"github.com/hh02/minimal-douyin/kitex_gen/likerpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc/videoservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/hh02/minimal-douyin/kitex_gen/likerpc/likeservice"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var likeClient likeservice.Client
var videoClient videoservice.Client

func initLikeRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := likeservice.NewClient(
		constants.LikeServiceName,
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

	likeClient = c
}

func MGetUser(ctx context.Context, req *videorpc.MGetVideoRequest) ([]*videorpc.Video, error) {
	resp, err := videoClient.MGetVideo(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Status.StatusCode != errno.SuccessCode {
		return nil, errno.Status2ErrorNo(resp.Status)
	}

	return resp.Videos, nil
}

func AddLikeCont(ctx context.Context, req *likerpc.CreateLikeRequest) error {
	resp, err := likeClient.CreateLike(ctx, req)
	if resp.Status.StatusCode != errno.SuccessCode {
		return errno.Status2ErrorNo(resp.Status)
	}
	if err != nil {
		return err
	}

	return nil
}
