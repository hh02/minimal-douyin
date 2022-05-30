package constants

const (
	IdentityKey               = "id"
	Total                     = "total"
	Notes                     = "notes"
	NoteID                    = "note_id"
	ApiServiceName            = "demoapi"
	UserServiceName           = "userservice"
	VideoServiceName          = "videoservice"
	FollowServiceName         = "followservice"
	MySQLDefaultDSN           = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress               = "127.0.0.1:2379"
	CPURateLimit      float64 = 80.0
	DefaultLimit              = 10
	VideoTableName            = "videos"
	UserTableName             = "users"
	FollowTableName           = "follows"
	SecretKey                 = "secret key"
)
