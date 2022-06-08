package constants

const (
	IdentityKey = "id"
	Total       = "total"
	User        = "user"

	ApiServiceName     = "demoapi"
	UserServiceName    = "userservice"
	VideoServiceName   = "videoservice"
	FollowServiceName  = "followservice"
	CommentServiceName = "commentservice"
	LikeServiceName    = "likeservice"

	UserTableName    = "users"
	VideoTableName   = "videos"
	FollowTableName  = "follows"
	CommentTableName = "comments"
	LikeTableName    = "likes"

	UserServerAddress    = "0.0.0.0:8888"
	VideoServerAddress   = "0.0.0.0:8889"
	FollowServerAddress  = "0.0.0.0:8890"
	CommentServerAddress = "0.0.0.0:8891"
	LikeServerAdress     = "0.0.0.0:8892"

	MySQLDefaultDSN         = "gorm:gorm@tcp(120.46.179.205:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress             = "120.46.179.205:2379"
	CPURateLimit    float64 = 80.0
	DefaultLimit            = 10

	SecretKey        = "secret key"
	StaticServerPath = "/static"
	StaticFolder     = "./static"

	VideoFolder     = "/root/public/video/"
	CoverFolder     = "/root/public/cover/"
	FileServer      = "120.46.179.205"
	TotalFeedNumber = 20
)
