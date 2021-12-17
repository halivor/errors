package errors

type ErrorCode int64

var _busiCode ErrorCode = 0

const Succ = 0

func InitBusiCode(busiCode int64) {
	_busiCode = ErrorCode(busiCode)
}
