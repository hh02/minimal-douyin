package pack

import (
	"github.com/hh02/minimal-douyin/cmd/comment/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
)

func Comment(c *db.Comment, userId int64) *commentrpc.Comment {
	if c != nil {
		return nil
	}
	return &commentrpc.Comment{
		CommentId: int64(c.ID),
	}
}
