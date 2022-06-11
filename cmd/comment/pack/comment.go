package pack

import (
	"context"
	"fmt"
	"github.com/hh02/minimal-douyin/cmd/comment/dal/db"
	"github.com/hh02/minimal-douyin/cmd/comment/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
)

// Comment pack db.Comment to commentrpc.Comment
func Comment(ctx context.Context, c *db.Comment) *commentrpc.Comment {
	if c == nil {
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
	fmt.Println("commentid:", c.ID)
	return &commentrpc.Comment{
		CommentId:  int64(c.ID),
		Content:    c.Content,
		User:       user,
		CreateDate: c.CreatedAt.Format("January-02"),
	}
}

// MComment pack many db.Comment to commentrpc.Comment
func MComment(ctx context.Context, comments []*db.Comment, tokenId int64) ([]*commentrpc.Comment, error) {
	if len(comments) == 0 {
		return nil, nil
	}
	res := make([]*commentrpc.Comment, 0)
	ids := make([]int64, 0)
	// 建立 id 对应查询出来的 users 的 id
	// 因为 MGetUser 函数对重复的 id 只会查询出一个结果
	// 自带去重功能，因此想要得到重复的  user 就先建立 id 对应 user 的 map
	id2user := make(map[int64]*userrpc.User, 0)

	for _, comment := range comments {
		ids = append(ids, comment.UserId)
	}
	// 批量查询 user
	users, err := rpc.MGetUser(ctx, &userrpc.MGetUserRequest{
		UserIds:        ids,
		TokenUserId:    tokenId,
		ReturnIsFollow: true,
	})
	// 获取 id 对应 user 的 map
	for _, user := range users {
		id2user[user.Id] = user
	}
	if err != nil {
		return nil, err
	}

	for i, comment := range comments {
		temp := &commentrpc.Comment{
			CommentId:  int64(comment.ID),
			Content:    comment.Content,
			User:       id2user[ids[i]],
			CreateDate: comment.CreatedAt.Format("January-02"),
		}
		res = append(res, temp)
	}
	return res, nil
}
