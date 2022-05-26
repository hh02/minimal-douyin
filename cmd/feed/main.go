package main

import (
	core "github.com/hh02/minimal-douyin/cmd/feed/kitex_gen/douyin/core/feedservice"
	"log"
)

func main() {
	svr := core.NewServer(new(FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
