package common

import (
	"errors"

	"github.com/hh02/minimal-douyin/kitex_gen/common"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildStatus(err error) *common.Status {
	if err == nil {
		return ErrorNo2Status(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return ErrorNo2Status(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return ErrorNo2Status(s)

}

func ErrorNo2Status(err errno.ErrNo) *common.Status {
	return &common.Status{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg}
}
