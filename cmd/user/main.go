package main

import (
	userrpc "github.com/hh02/minimal-douyin/kitex_gen/userrpc/userservice"
	"log"
)

func main() {
	svr := userrpc.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
