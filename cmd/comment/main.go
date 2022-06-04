package main

import (
	commentrpc "github.com/hh02/minimal-douyin/kitex_gen/commentrpc/commentservice"
	"log"
)

func main() {
	svr := commentrpc.NewServer(new(CommentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
