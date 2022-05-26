package main

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/feed/kitex_gen/douyin/core"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *core.FeedRequest) (resp *core.FeedResponse, err error) {
	// TODO: Your code here...
	return
}
