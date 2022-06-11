package main

import (
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/handlers"
	"github.com/hh02/minimal-douyin/cmd/api/middleware"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/tracer"
)

func Init() {
	middleware.InitMiddleware()
	rpc.InitRPC()
	tracer.InitJaeger(constants.ApiServiceName)
}

func main() {
	Init()
	r := gin.New()

	r.Static(constants.StaticRelativePath, constants.StaticLocalPath)

	douyin := r.Group("/douyin")
	douyin.POST("/user/login/", handlers.UserLogin)

	// 注册后登录
	douyin.POST("/user/register/", handlers.UserRegister)

	// 要根据是否有token来决定是否调用鉴权
	douyin.GET("/feed", handlers.Feed, middleware.AuthMiddleware.MiddlewareFunc())

	// 因为token在body中，要先取出来
	douyin.POST("/publish/action/", handlers.PublishAction, middleware.AuthMiddleware.MiddlewareFunc())

	douyin.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		douyin.GET("/user/", handlers.UserInfo)

		douyin.POST("/favorite/action/", handlers.FavoriteAction)
		douyin.GET("/favorite/list/", handlers.FavoriteList)

		douyin.POST("/relation/action/", handlers.RelationAction)
		douyin.GET("/relation/follow/list/", handlers.FollowList)
		douyin.GET("/relation/follower/list/", handlers.FollowerList)

		douyin.GET("/publish/list/", handlers.PublishList)

		douyin.POST("/comment/action/", handlers.CommentAction)
		douyin.GET("/comment/list/", handlers.CommentList)
	}

	if err := http.ListenAndServe(":80", r); err != nil {
		klog.Fatal(err)
	}
}
