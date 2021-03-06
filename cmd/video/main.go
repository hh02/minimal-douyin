package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/hh02/minimal-douyin/cmd/video/dal"
	"github.com/hh02/minimal-douyin/cmd/video/rpc"
	videorpc "github.com/hh02/minimal-douyin/kitex_gen/videorpc/videoservice"
	"github.com/hh02/minimal-douyin/pkg/constants"
	tracer2 "github.com/hh02/minimal-douyin/pkg/tracer"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	tracer2.InitJaeger(constants.VideoServiceName)
	rpc.InitRPC()
	dal.Init()
}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.VideoServerAddress)
	if err != nil {
		panic(err)
	}
	Init()
	svr := videorpc.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.VideoServiceName}), // server name
		// server.WithMiddleware(middleware.CommonMiddleware),                                             // middleWare
		// server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		// server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry)
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
