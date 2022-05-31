package main

import (
	"github.com/hh02/minimal-douyin/cmd/follow/rpc"
	"github.com/hh02/minimal-douyin/cmd/user/dal"
	userrpc "github.com/hh02/minimal-douyin/kitex_gen/userrpc/userservice"
	"log"
)

func Init() {
	dal.Init()
	rpc.InitRPC()
}

func main() {
	svr := userrpc.NewServer(new(UserServiceImpl))

	Init()
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
