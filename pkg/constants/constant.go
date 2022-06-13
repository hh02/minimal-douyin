package constants

const (
	IdentityKey = "id"
	AuthErrKey  = "AuthErr"
	Total       = "total"
	User        = "user"

	ApiServiceName      = "demoapi"
	UserServiceName     = "userservice"
	VideoServiceName    = "videoservice"
	FollowServiceName   = "followservice"
	CommentServiceName  = "commentservice"
	FavoriteServiceName = "favoriteservice"

	UserTableName     = "users"
	VideoTableName    = "videos"
	FollowTableName   = "follows"
	CommentTableName  = "comments"
	FavoriteTableName = "favorites"

	UserServerAddress     = "0.0.0.0:8888"
	VideoServerAddress    = "0.0.0.0:8889"
	FollowServerAddress   = "0.0.0.0:8890"
	CommentServerAddress  = "0.0.0.0:8891"
	FavoriteServerAddress = "0.0.0.0:8892"

	MySQLDefaultDSN         = "gorm:gorm@tcp(120.46.179.205:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress             = "120.46.179.205:2379"
	CPURateLimit    float64 = 80.0
	DefaultLimit            = 10

	StaticRelativePath = "/static"
	StaticServerPath   = "http://120.46.179.205" + StaticRelativePath
	StaticLocalPath    = "./static"
	VideoLocalPath     = StaticLocalPath + "/video"
	CoverLocalPath     = StaticLocalPath + "/cover"
	VideoServerPath    = StaticServerPath + "/video"
	CoverServerPath    = StaticServerPath + "/cover"

	OssEndpoint        = "http://oss-cn-beijing.aliyuncs.com"
	OssAccessKeyId     = "LTAI5t9ceDjAisE3Rc6YDXLQ"
	OssAccessKeySecret = "4EKkDRLeTYcMeY3CtdSnw0CUnmAcrg"
	OssBucketName      = "minimal-douyin"
	OssUrlPrefix       = "https://minimal-douyin.oss-cn-beijing.aliyuncs.com/"
	TempFoler          = "./temp/" // 用来存放临时文件
	TotalFeedNumber    = 20
	SecretKey          = "secret key"
)
