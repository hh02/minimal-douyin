package main

import (
	videorpc "github.com/hh02/minimal-douyin/kitex_gen/videorpc/videoservice"
	"log"
)

func main() {
	svr := videorpc.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
