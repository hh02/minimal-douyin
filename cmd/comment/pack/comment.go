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
	return &commentrpc.Comment{
		CommentId:  int64(c.ID),
		Content:    c.Content,
		User:       user,
		CreateDate: c.CreatedAt.Format("January-02"),
	}
}

// MComment pack many db.Comment to commentrpc.Comment
func MComment(ctx context.Context, comments []*db.Comment, tokenId int64) ([]*commentrpc.Comment, error) {
	res := make([]*commentrpc.Comment, 0)
	ids := make([]int64, 0)
	// 建立 id 对应查询出来的 users 的 id
	// 因为 MGetUser 函数对重复的 id 只会查询出一个结果
	// 自带去重功能，因此想要得到重复的  user 就先建立一个去重前 id
	// 与去重后 id 的对应关系。就是 id2userid
	id2userid := make([]int64, 0)
	it := int64(0)

	for i, comment := range comments {
		if i != 0 {
			if comment.UserId != ids[i-1] {
				it++
			}
		}
		// 得到去重前 id 与去重后 id 对应关系

		id2userid = append(id2userid, it)
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
			User:       users[id2userid[i]],
			CreateDate: comment.CreatedAt.Format("January-02"),
		}
		res = append(res, temp)
	}
	return res, nil
}
