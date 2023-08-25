package xerr

const OK uint32 = 200

// 全局错误码
const (
	SERVER_COMMON_ERROR  uint32 = 1001
	REUQEST_PARAM_ERROR  uint32 = 1002
	TOKEN_EXPIRE_ERROR   uint32 = 1003
	TOKEN_GENERATE_ERROR uint32 = 1004
	DB_ERROR             uint32 = 1005
	MONGODB_ERROR        uint32 = 1006
	REDIS_ERROR          uint32 = 1007

	NETWORK_ERROR uint32 = 2001
)
