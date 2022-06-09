package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/hh02/minimal-douyin/cmd/like/dal"
	"github.com/hh02/minimal-douyin/cmd/like/rpc"
	likerpc "github.com/hh02/minimal-douyin/kitex_gen/likerpc/likeservice"
	"github.com/hh02/minimal-douyin/pkg/constants"
	tracer2 "github.com/hh02/minimal-douyin/pkg/tracer"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	tracer2.InitJaeger(constants.FollowServiceName)
	rpc.InitRPC()
	dal.Init()

}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.LikeServerAddress)
	if err != nil {
		panic(err)
	}
	Init()
	svr := likerpc.NewServer(new(LikeServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.LikeServiceName}), // server name
		// server.WithMiddleware(middleware.CommonMiddleware),                                             // middleWare
		// server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		// server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
