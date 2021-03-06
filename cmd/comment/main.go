package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/hh02/minimal-douyin/cmd/comment/dal"
	"github.com/hh02/minimal-douyin/cmd/comment/rpc"
	commentrpc "github.com/hh02/minimal-douyin/kitex_gen/commentrpc/commentservice"
	"github.com/hh02/minimal-douyin/pkg/constants"
	tracer2 "github.com/hh02/minimal-douyin/pkg/tracer"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"net"
)

func Init() {
	tracer2.InitJaeger(constants.FollowServiceName)
	dal.Init()
	rpc.InitRPC()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.CommentServerAddress)
	if err != nil {
		panic(err)
	}

	Init()

	svr := commentrpc.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.CommentServiceName}), // server name
		//server.WithMiddleware(middleware.CommonMiddleware),                                             // middleware
		//server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		//server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
