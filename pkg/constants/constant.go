package constants

const (
	UserTableName           = "user"
	SecretKey               = "secret key"
	IdentityKey             = "id"
	Total                   = "total"
	Notes                   = "notes"
	NoteID                  = "note_id"
	ApiServiceName          = "demoapi"
	NoteServiceName         = "demonote"
	UserServiceName         = "demouser"
	MySQLDefaultDSN         = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress             = "127.0.0.1:2379"
	CPURateLimit    float64 = 80.0
	DefaultLimit            = 10
)
