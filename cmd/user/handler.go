package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"
	"github.com/hh02/minimal-douyin/cmd/user/pack"
	"github.com/hh02/minimal-douyin/cmd/user/service"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userrpc.CreateUserRequest) (resp *userrpc.CreateUserResponse, err error) {
	// TODO: Your code here...
	resp = new(userrpc.CreateUserResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.StatusCode, resp.StatusMessage = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err, Id := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.StatusCode, resp.StatusMessage = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.StatusCode, resp.StatusMessage = pack.BuildBaseResp(errno.Success)
	resp.UserId = Id
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userrpc.GetUserRequest) (resp *userrpc.GetUserResponse, err error) {
	// TODO: Your code here...
	resp := new(userrpc.GetUserResponse)

	if len(req.UserId) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil

	return
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *userrpc.MGetUserRequest) (resp *userrpc.MGetUserResponse, err error) {
	// TODO: Your code here...\
	if len(req.UserIds) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userrpc.CheckUserRequest) (resp *userrpc.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.CheckUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
