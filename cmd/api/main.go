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
		Unauthorized: func(c *gin.Context, code int, message string) {
			statusCode := errno.AuthErr.ErrCode
			if message == errno.LoginErr.ErrMsg {
				statusCode = errno.LoginErr.ErrCode
			}
			c.JSON(code, handlers.AuthResponse{
				StatusCode: statusCode,
				StatusMsg:  message,
			})
		},
		HTTPStatusMessageFunc: func(e error, c *gin.Context) string {
			err := errno.ConvertErr(e)
			return err.ErrMsg
		},
		LoginResponse: func(c *gin.Context, code int, tokenString string, time time.Time) {
			if code != http.StatusOK {
				c.JSON(code, handlers.AuthResponse{
					StatusCode: int32(code),
					StatusMsg:  "登陆失败",
					UserId:     0,
					Token:      "",
				})
			}
			c.AddParam("token", tokenString)
			c.Next()
		},
		TokenLookup:   "query: token, param: token",
		TokenHeadName: "",
		TimeFunc:      time.Now,
	})

	r.Static(constants.StaticRelativePath, constants.StaticLocalPath)

	douyin := r.Group("/douyin")

	// 登录获取token，中间件解析token，UserLogin返回结果
	douyin.POST("/user/login/", authMiddleware.LoginHandler,
		authMiddleware.MiddlewareFunc(),
		handlers.UserLoginResponse,
	)

	// 注册后登录
	douyin.POST("/user/register/", handlers.UserRegister,
		authMiddleware.LoginHandler,
		authMiddleware.MiddlewareFunc(),
		handlers.UserLoginResponse,
	)

	// 要根据是否有token来决定是否调用鉴权
	douyin.GET("/feed", handlers.Feed, authMiddleware.MiddlewareFunc())

	// 因为token在body中，要先取出来
	douyin.POST("/publish/action/", handlers.PublishAction, authMiddleware.MiddlewareFunc())

	douyin.Use(authMiddleware.MiddlewareFunc())
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
