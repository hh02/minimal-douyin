package main

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/publish/kitex_gen/douyin/core"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *core.PublishListRequest) (resp *core.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *core.PublishActionRequest) (resp *core.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}
