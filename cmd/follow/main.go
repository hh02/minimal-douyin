package main

import (
	"log"

	"github.com/hh02/minimal-douyin/cmd/follow/rpc"
	"github.com/hh02/minimal-douyin/cmd/user/dal"
	followrpc "github.com/hh02/minimal-douyin/kitex_gen/followrpc/followservice"
)

func Init() {
	dal.Init()
	rpc.InitRPC()
}

func main() {
	svr := followrpc.NewServer(new(FollowServiceImpl))

	Init()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
