package ers

// ------------------------------ 通用错误 ------------------------------
var (
	MysqlError = New(CodeOfMysqlError, "数据库操作时发生错误，请联系 小悦悦 解决～")
	RedisError = New(CodeOfRedisError, "缓存操作时发生错误，请联系 小悦悦 解决～")
)
