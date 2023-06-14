package xerr

const OK uint32 = 200

// 全局错误码
const (
	SERVER_COMMON_ERROR  uint32 = 100001
	REUQEST_PARAM_ERROR  uint32 = 100002
	TOKEN_EXPIRE_ERROR   uint32 = 100003
	TOKEN_GENERATE_ERROR uint32 = 100004
	DB_ERROR             uint32 = 100005
	MONGODB_ERROR        uint32 = 100006
	REDIS_ERROR          uint32 = 100007
)
