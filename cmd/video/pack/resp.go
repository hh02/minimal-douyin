package pack

import (
	"errors"

	"github.com/hh02/minimal-douyin/kitex_gen/resp"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *resp.BaseResp {
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

func baseResp(err errno.ErrNo) *resp.BaseResp {
	return &resp.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg}
}
