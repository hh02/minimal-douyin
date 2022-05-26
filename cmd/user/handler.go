package main

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/user/kitex_gen/douyin/core"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *core.UserRegisterRequest) (resp *core.UserRegisterResponse, err error) {
	// TODO: Your code here...
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
