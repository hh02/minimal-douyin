package pack

import (
	"errors"

	"github.com/hh02/minimal-douyin/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) (int32, string) {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)

}

func baseResp(err errno.ErrNo) (int32, string) {
	return err.ErrCode, err.ErrMsg
}
