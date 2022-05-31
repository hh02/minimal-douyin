package service

import (
	"context"

)

type QueryVideoService struct {
	ctx context.Context
}

func NewQueryVideoService(ctx context.Context) *QueryVideoService {
	return &QueryVideoService{ctx: ctx}
}
