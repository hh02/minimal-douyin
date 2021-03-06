package errno

import (
	"errors"
	"fmt"

	"github.com/hh02/minimal-douyin/kitex_gen/common"
)

const (
	SuccessCode                 = 0
	ServiceErrCode              = 10001
	ParamErrCode                = 10002
	LoginErrCode                = 10003
	UserNotExistErrCode         = 10004
	UserAlreadyExistErrCode     = 10005
	FollowNotExistErrCode       = 10006
	PermissionErrCode           = 10007
	AuthErrCode                 = 10008 //鉴权失败
	FavoriteAlreadyExistErrCode = 10009
	DeleteErrCode               = 10010
	GetVideoErrCode             = 10011
	GetCoverErrCode             = 10012
	GetIdFromClaimsErrCode      = 10013
	OpenVideoFileErrCode        = 10014
	SaveVideoErrCode            = 10015
	OssPutObjectErrCode         = 10016
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                 = NewErrNo(SuccessCode, "Success")
	ServiceErr              = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr                = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	LoginErr                = NewErrNo(LoginErrCode, "Wrong username or password")
	UserNotExistErr         = NewErrNo(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr     = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	FollowNotExistErr       = NewErrNo(FollowNotExistErrCode, "Follow does not exist")
	PermissionErr           = NewErrNo(PermissionErrCode, "User does not have permission")
	AuthErr                 = NewErrNo(AuthErrCode, "Authentication failed")
	FavoriteAlreadyExistErr = NewErrNo(FavoriteAlreadyExistErrCode, "Favorite already exist")
	DeleteErr               = NewErrNo(DeleteErrCode, "Delete has Wrong")
	GetVideoErr             = NewErrNo(GetVideoErrCode, "No more videos")
	GetCoverErr             = NewErrNo(GetCoverErrCode, "Get cover failed")
	GetIdFromClaimsErr      = NewErrNo(GetIdFromClaimsErrCode, "Failed to get ID from claims")
	OpenVideoFileErr        = NewErrNo(OpenVideoFileErrCode, "Can not open the upload video")
	SaveVideoErr            = NewErrNo(SaveVideoErrCode, "Failed to save video")
	OssPutObjectErr         = NewErrNo(OssPutObjectErrCode, "Failed to Put object to OSS")
)

// BuildStatus build baseResp from error
func BuildStatus(err error) *common.Status {
	if err == nil {
		return ErrNo2Status(Success)
	}

	e := ErrNo{}

	// 如果为ErrNo类型
	if errors.As(err, &e) {
		return ErrNo2Status(e)
	}

	s := ServiceErr.WithMessage(err.Error())
	return ErrNo2Status(s)

}

func ErrNo2Status(err ErrNo) *common.Status {
	return &common.Status{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg}
}

func Status2ErrorNo(status *common.Status) ErrNo {
	return ErrNo{
		ErrCode: status.StatusCode,
		ErrMsg:  status.StatusMessage,
	}
}

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
