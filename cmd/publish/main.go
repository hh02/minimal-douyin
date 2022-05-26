package main

import (
	core "github.com/hh02/minimal-douyin/cmd/publish/kitex_gen/douyin/core/publishservice"
	"log"
)

func main() {
	svr := core.NewServer(new(PublishServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
