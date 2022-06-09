package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc/commentservice"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var commentClient commentservice.Client

func initCommentRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		constants.CommentServiceName,
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
	commentClient = c
}

// CreateComment create comment
func CreateComment(ctx context.Context, req *commentrpc.CreateCommentRequest) (*commentrpc.Comment, error) {
	resp, err := commentClient.CreateComment(ctx, req)
	fmt.Println(req.UserId, req.VideoId, req.CommentText)
	fmt.Println(resp.Comment, resp.Status)
	if err != nil {
		return nil, err
	}
	if resp.Status.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Status.StatusCode, resp.Status.StatusMessage)
	}
	return resp.Comment, nil
}

// DeleteComment delete comment
func DeleteComment(ctx context.Context, req *commentrpc.DeleteCommentRequest) error {
	resp, err := commentClient.DeleteComment(ctx, req)
	if err != nil {
		return err
	}
	if resp.Status.StatusCode != 0 {
		return errno.NewErrNo(resp.Status.StatusCode, resp.Status.StatusMessage)
	}
	return nil
}

// QueryCommentByVideo query comment list by video id
func QueryCommentByVideo(ctx context.Context, req *commentrpc.QueryCommentByVideoIdRequest) ([]*commentrpc.Comment, error) {
	resp, err := commentClient.QueryCommentByVideo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Status.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Status.StatusCode, resp.Status.StatusMessage)
	}
	return resp.CommentList, nil
}
