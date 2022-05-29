package main

import (
	video_rpc "github.com/hh02/minimal-douyin/kitex_gen/video_rpc/videoservice"
	"log"
)

func main() {
	svr := video_rpc.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
