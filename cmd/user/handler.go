package main

import (
	"context"
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
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	err, Id := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.UserId = Id
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userrpc.GetUserRequest) (resp *userrpc.GetUserResponse, err error) {
	// TODO: Your code here...
	resp = new(userrpc.GetUserResponse)

	if req.UserId == 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	user, err := service.NewGetUserService(ctx).GetUser(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.User = user
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *userrpc.MGetUserRequest) (resp *userrpc.MGetUserResponse, err error) {
	// TODO: Your code here...\

	resp = new(userrpc.MGetUserResponse)

	if len(req.UserIds) == 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.Users = users
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userrpc.CheckUserRequest) (resp *userrpc.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = new(userrpc.CheckUserResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.Status = errno.BuildStatus(errno.Success)
	return resp, nil
}

// AddFollowCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddFollowCount(ctx context.Context, req *userrpc.AddFollowCountRequest) (resp *userrpc.AddFollowCountResponse, err error) {
	// TODO: Your code here...
	resp = new(userrpc.AddFollowCountResponse)

	if req.UserId == 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}
	err = service.NewAddFollowCountService(ctx).AddFollowCount(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	return resp, nil
}

// AddFollowerCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddFollowerCount(ctx context.Context, req *userrpc.AddFollowerCountRequest) (resp *userrpc.AddFollowerCountResponse, err error) {
	// TODO: Your code here...
	resp = new(userrpc.AddFollowerCountResponse)

	if req.UserId == 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	err = service.NewAddFollowerCountService(ctx).AddFollowerCount(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	return resp, nil
}
