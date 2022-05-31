package main

import (
	"context"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userrpc.CreateUserRequest) (resp *userrpc.CreateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userrpc.GetUserRequest) (resp *userrpc.GetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *userrpc.MGetUserRequest) (resp *userrpc.MGetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userrpc.CheckUserRequest) (resp *userrpc.CheckUserResponse, err error) {
	// TODO: Your code here...
	return
}
