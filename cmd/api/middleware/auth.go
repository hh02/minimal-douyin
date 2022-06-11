package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/hh02/minimal-douyin/pkg/constants"
)

var AuthMiddleware *jwt.GinJWTMiddleware

func initAuthMiddleware() {
	// t := AuthMiddleware.RefreshHandler
	AuthMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
		Key:           []byte(constants.SecretKey),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "query: token, param: token",
		TokenHeadName: "",
		TimeFunc:      time.Now,
		DisabledAbort: true,
	})
}
