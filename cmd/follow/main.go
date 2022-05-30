package main

import (
	followrpc "github.com/hh02/minimal-douyin/kitex_gen/followrpc/followservice"
	"log"
)

func main() {
	svr := followrpc.NewServer(new(FollowServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
