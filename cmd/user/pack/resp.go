package pack

import (
	"errors"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

// BuildBaseResp 判断错误，返回相应错误码和错误信息。
func BuildBaseResp(err error) (Code int32, Msg string) {
	if err == nil {
		return errno.Success.ErrCode, errno.Success.ErrMsg
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) (Code int32, Msg string) {
	return err.ErrCode, err.ErrMsg
}
