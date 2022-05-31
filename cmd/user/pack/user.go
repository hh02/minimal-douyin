package pack

import (
	"github.com/hh02/minimal-douyin/cmd/user/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
)

// User pack user info
func User(u *db.User) *userrpc.User {
	if u == nil {
		return nil
	}

	return &core.User{
		Id:   int64(u.ID),
		Name: u.Name,
	}
}
