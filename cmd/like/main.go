package main

import (
	likerpc "github.com/hh02/minimal-douyin/kitex_gen/likerpc/userservice"
	"log"
)

func main() {
	svr := likerpc.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
