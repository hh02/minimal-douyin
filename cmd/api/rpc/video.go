package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc/videoservice"
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
	videoClient = c
}
func QueryVideoByUserId(ctx context.Context, req *videorpc.QueryVideoByUserIdRequest) ([]*videorpc.Video, error) {
	resp, err := videoClient.QueryVideoByUserId(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Status.StatusCode != errno.SuccessCode {
		return nil, errno.Status2ErrorNo(resp.Status)
	}
	return resp.Videos, nil
}

func QueryVideoByTime(ctx context.Context, req *videorpc.QueryVideoByTimeRequest) (int64, []*videorpc.Video, error) {
	resp, err := videoClient.QueryVideoByTime(ctx, req)
	if err != nil {
		return 0, nil, err
	}

	if resp.Status.StatusCode != errno.SuccessCode {
		return 0, nil, errno.Status2ErrorNo(resp.Status)
	}
	return resp.NextTime, resp.Videos, nil
}

func CreateVideo(ctx context.Context, req *videorpc.CreateVideoRequest) error {
	resp, err := videoClient.CreateVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.Status.StatusCode != errno.SuccessCode {
		return errno.Status2ErrorNo(resp.Status)
	}
	return nil
}
