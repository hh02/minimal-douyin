package main

import (
	"github.com/hh02/minimal-douyin/cmd/like/rpc"
	"github.com/hh02/minimal-douyin/cmd/like/dal"
	likerpc "github.com/hh02/minimal-douyin/kitex_gen/likerpc/likeservice"
	"log"
)

func init() {
	dal.Init()
	rpc.InitRPC()
}

func main() {
	svr := likerpc.NewServer(new(LikeServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
