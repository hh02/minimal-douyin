package main

import (
	core "github.com/hh02/minimal-douyin/kitex_gen/douyin/core/userservice"
	"log"
)

func main() {
	svr := core.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
