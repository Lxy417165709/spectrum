package ers

// ------------------------------ 通用错误码 ------------------------------
const (
	CodeOfMysqlError = 1 + iota
	CodeOfRedisError
)

// ------------------------------ 选项相关错误码 ------------------------------
const (
	CodeOfBlankOptionClassName = 100000 + iota
	CodeOfEmptyOption
	CodeOfBlankOptionName
)
