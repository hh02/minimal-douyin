package utils

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/pkg/constants"
)

// 如果返回0则表示失败
func GetIdFromClaims(c *gin.Context) int64 {
	claims := jwt.ExtractClaims(c)
	if val, ok := claims[constants.IdentityKey]; ok {
		return int64(val.(float64))
	}
	return 0
}
