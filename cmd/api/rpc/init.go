package rpc

// InitRPC init rpc client
func InitRPC() {
	initUserRpc()
	initFollowRpc()
	initCommentRpc()
	initVideoRpc()
	initFavoriteRpc()
}
