package main

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/api/rpc"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"
	"github.com/hh02/minimal-douyin/pkg/constants"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/handlers"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

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

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &userdemo.CheckUserRequest{UserName: loginVar.UserName, Password: loginVar.PassWord})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	apiRouter := r.Group("/douyin")

	// basic apis
	//apiRouter.GET("/feed/", handlers.Feed)
	apiRouter.GET("/user/", handlers.UserInfo)
	apiRouter.POST("/user/register/", handlers.Register)
	apiRouter.POST("/user/login/", authMiddleware.LoginHandler)
	//apiRouter.POST("/publish/action/", handlers.Publish)
	//apiRouter.GET("/publish/list/", handlers.PublishList)

	// extra apis - I
	//apiRouter.POST("/favorite/action/", handlers.FavoriteAction)
	//apiRouter.GET("/favorite/list/", handlers.FavoriteList)
	//apiRouter.POST("/comment/action/", handlers.CommentAction)
	//apiRouter.GET("/comment/list/", handlers.CommentList)

	// extra apis - II
	//apiRouter.POST("/relation/action/", handlers.RelationAction)
	//apiRouter.GET("/relation/follow/list/", handlers.FollowList)
	//apiRouter.GET("/relation/follower/list/", handlers.FollowerList)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
