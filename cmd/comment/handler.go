package main

import (
	"context"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CreateComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CreateComment(ctx context.Context, req *commentrpc.CreateCommentRequest) (resp *commentrpc.CreateCommentResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *commentrpc.DeleteCommentRequest) (resp *commentrpc.DeleteCommentResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryCommentByVideo implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) QueryCommentByVideo(ctx context.Context, req *commentrpc.QueryCommentByVideoIdRequest) (resp *commentrpc.QueryCommentByVideoIdResponse, err error) {
	// TODO: Your code here...
	return
}
