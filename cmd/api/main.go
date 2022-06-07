package main

import (
	"context"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/handlers"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
	"github.com/hh02/minimal-douyin/pkg/tracer"
)

func Init() {
	rpc.InitRPC()
	tracer.InitJaeger(constants.ApiServiceName)
}

func main() {
	Init()
	r := gin.New()
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar handlers.UserParam
			if err := c.ShouldBind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.Username) == 0 || len(loginVar.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &userrpc.CheckUserRequest{Username: loginVar.Username, Password: loginVar.Password})
		},
		TokenLookup:   "query: token, cookie: jwt",
		TokenHeadName: "",
		TimeFunc:      time.Now,
	})

	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.POST("/user/register/", func(c *gin.Context) {
		var userVar handlers.UserParam
		if err := c.BindQuery(&userVar); err != nil {
			handlers.SendRegisterResponse(c, errno.ConvertErr(err))
			return
		}
		if len(userVar.Username) == 0 || len(userVar.Password) == 0 {
			handlers.SendRegisterResponse(c, errno.ParamErr)
			return
		}

		err := rpc.CreateUser(context.Background(), &userrpc.CreateUserRequest{
			Username: userVar.Username,
			Password: userVar.Password,
		})
		if err != nil {
			handlers.SendRegisterResponse(c, errno.ConvertErr(err))
			return
		}

		c.Request.URL.Path = "/douyin/user/login"
		r.HandleContext(c)
	})
	apiRouter.GET("/feed/", handlers.Feed)
	apiRouter.POST("/user/login/", authMiddleware.LoginHandler)

	apiRouter.Use(authMiddleware.MiddlewareFunc())
	{
		apiRouter.GET("/user/", handlers.UserInfo)
	}

	apiRouter.POST("/publish/action/", handlers.PublishAction)
	apiRouter.GET("/publish/list/", handlers.PublishList)

	// extra apis - I
	//apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", handlers.CommentAction)
	apiRouter.GET("/comment/list/", handlers.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", handlers.RelationAction)
	apiRouter.GET("/relation/follow/list/", handlers.FollowList)
	apiRouter.GET("/relation/follower/list/", handlers.FollowerList)

	if err := http.ListenAndServe(":80", r); err != nil {
		klog.Fatal(err)
	}
}
