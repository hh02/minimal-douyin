package main

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/user/pack"
	"github.com/hh02/minimal-douyin/cmd/user/service"
	"github.com/hh02/minimal-douyin/kitex_gen/douyin/core"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	ctx context.Context
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *core.UserRegisterRequest) (resp *core.UserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(core.UserRegisterResponse)

	// 当用户名密码长度都为 0 时，返回参数错误的状态
	if len(req.Username) == 0 || len(req.Password) == 0 {
		//resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		resp.StatusCode = errno.ParamErr.ErrCode
		resp.StatusMsg = errno.ParamErr.ErrMsg

		return resp, nil
	}

	// 用户注册
	UserId, token, err := service.NewUserRegiaterService(ctx).UserRegister(req)
	// 如果注册发生错误
	if err != nil {
		resp.StatusCode, resp.StatusMsg = pack.BuildBaseResp(err)
		return resp, nil
	}
	// 注册成功
	resp.StatusCode, resp.StatusMsg = pack.BuildBaseResp(errno.Success)
	resp.UserId = int64(UserId)
	resp.Token = token
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *core.UserLoginRequest) (resp *core.UserLoginResponse, err error) {
	// TODO: Your code here...
	resp = new(core.UserLoginResponse)

	// 当用户名密码长度都为 0 时，返回参数错误的状态
	if len(req.Username) == 0 || len(req.Password) == 0 {
		//resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		resp.StatusCode = errno.ParamErr.ErrCode
		resp.StatusMsg = errno.ParamErr.ErrMsg

		return resp, nil
	}

	// 用户登录
	UserId, token, err := service.NewUserLoginService(ctx).UserLogin(req)
	// 如果登录发生错误
	if err != nil {
		resp.StatusCode, resp.StatusMsg = pack.BuildBaseResp(err)
		return resp, nil
	}
	// 登录成功
	resp.StatusCode, resp.StatusMsg = pack.BuildBaseResp(errno.Success)
	resp.UserId = int64(UserId)
	resp.Token = token
	return resp, nil
}

// User implements the UserServiceImpl interface.
func (s *UserServiceImpl) User(ctx context.Context, req *core.UserRequest) (resp *core.UserResponse, err error) {
	// TODO: Your code here...
	resp = new(core.UserResponse)

	// 当用户名密码长度都为 0 时，返回参数错误的状态
	if req.UserId == 0 || len(req.Token) == 0 {
		//resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		resp.StatusCode = errno.ParamErr.ErrCode
		resp.StatusMsg = errno.ParamErr.ErrMsg

		return resp, nil
	}

	// 用户登录
	user, err := service.NewUserService(ctx).User(req)
	// 如果登录发生错误
	if err != nil {
		resp.StatusCode, resp.StatusMsg = pack.BuildBaseResp(err)
		return resp, nil
	}
	// 登录成功
	resp.User = pack.User(user)
	return resp, nil
}
