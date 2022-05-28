package main

import (
	"context"
	"github.com/hh02/minimal-douyin/kitex_gen/douyin/core"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *core.UserRegisterRequest) (resp *core.UserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(core.UserRegisterResponse)

	//当用户名密码长度都为 0 时，返回
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *core.UserLoginRequest) (resp *core.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// User implements the UserServiceImpl interface.
func (s *UserServiceImpl) User(ctx context.Context, req *core.UserRequest) (resp *core.UserResponse, err error) {
	// TODO: Your code here...
	return
}
