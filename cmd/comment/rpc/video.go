package rpc

import (
	"context"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc/videoservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
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

	videoClient = c
}

// AddCommentCount 在 video 中让点赞数加一减一
func AddCommentCount(ctx context.Context, req *videorpc.AddCommentCountRequest) error {
	resp, err := videoClient.AddCommentCount(ctx, req)
	if err != nil {
		return err
	}
	if resp.Status.StatusCode != errno.SuccessCode {
		return errno.Status2ErrorNo(resp.Status)
	}
	return nil
}
