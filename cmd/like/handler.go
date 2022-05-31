package main

import (
	"context"
	"github.com/hh02/minimal-douyin/kitex_gen/likerpc"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateLike implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateLike(ctx context.Context, req *likerpc.CreateLikeRequest) (resp *likerpc.CreateLikeResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *likerpc.DeleteLikeRequest) (resp *likerpc.DeleteLikeResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *likerpc.QueryLikeByUserIdRequest) (resp *likerpc.QueryLikeByUserIdResponse, err error) {
	// TODO: Your code here...
	return
}
