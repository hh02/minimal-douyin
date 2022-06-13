package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

func CheckAuth(c *gin.Context) error {
	data, exist := c.Get(constants.AuthErrKey)
	if exist {
		authErr := data.(string)
		return errno.AuthErr.WithMessage(authErr)
	}
	return nil
}
