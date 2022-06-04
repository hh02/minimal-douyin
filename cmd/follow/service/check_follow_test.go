package service

import (
	"context"
	"testing"

	"github.com/hh02/minimal-douyin/cmd/follow/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/stretchr/testify/assert"
)

func TestCheckFollow(t *testing.T) {

	// 初始化工作
	db.Init()
	_, err := db.CreateFollow(context.Background(), &db.Follow{
		UserId:   1234,
		FollowId: 5678,
	})
	if err != nil {
		t.Log(err.Error())
	}
	_, err = db.DeleteFollow(context.Background(), &db.Follow{
		UserId:   1234,
		FollowId: 12345,
	})
	if err != nil {
		t.Log(err.Error())
	}

	tests := map[string]struct {
		userId   int64
		followId int64
		want     bool
	}{
		"true case":   {userId: 1234, followId: 5678, want: true},
		"false case":  {userId: 1234, followId: 12345, want: false},
		"follow self": {userId: 1234, followId: 1234, want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewCheckFollowService(context.Background()).CheckFollow(&followrpc.CheckFollowRequest{
				UserId:   tc.userId,
				FollowId: tc.followId,
			})
			if err != nil {
				t.Error(err.Error())
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
