// Code generated by Kitex v0.3.1. DO NOT EDIT.
package feedservice

import (
	"github.com/cloudwego/kitex/server"
	"github.com/hh02/minimal-douyin/kitex_gen/douyin/core"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler core.FeedService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
