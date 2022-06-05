package pack

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/comment/dal/db"
	"github.com/hh02/minimal-douyin/cmd/comment/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
)

// Comment pack db.Comment to commentrpc.Comment
func Comment(ctx context.Context, c *db.Comment) *commentrpc.Comment {
	if c != nil {
		return nil
	}
	user, err := rpc.GetUser(ctx, &userrpc.GetUserRequest{
		UserId:         c.UserId,
		TokenUserId:    c.UserId,
		ReturnIsFollow: false,
	})
	if err != nil {
		return nil
	}
	user.IsFollow = false
	return &commentrpc.Comment{
		CommentId:  int64(c.ID),
		Content:    c.Content,
		User:       user,
		CreateDate: c.CreatedAt.Format("2006-January-02 03:04:05.999 pm"),
	}
}

// MComment pack many db.Comment to commentrpc.Comment
func MComment(ctx context.Context, comments []*db.Comment, tokenId int64) ([]*commentrpc.Comment, error) {
	res := make([]*commentrpc.Comment, 0)
	ids := make([]int64, 0)

	for _, comment := range comments {
		ids = append(ids, comment.UserId)
	}

	// 批量查询 user
	users, err := rpc.MGetUser(ctx, &userrpc.MGetUserRequest{
		UserIds:        ids,
		TokenUserId:    tokenId,
		ReturnIsFollow: true,
	})
	if err != nil {
		return nil, err
	}

	for i, comment := range comments {
		temp := &commentrpc.Comment{
			CommentId:  int64(comment.ID),
			Content:    comment.Content,
			User:       users[i],
			CreateDate: comment.CreatedAt.Format("2006-January-02 03:04:05.999 pm"),
		}
		res = append(res, temp)
	}
	return res, nil
}
