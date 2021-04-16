package errcode

import "fmt"

// 这是gRPC的错误码
type Error struct {
	code int
	msg  string
}

// 定义map,将code码与错误信息对应起来
var _codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := _codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在,请更换一个", code))
	}
	_codes[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

// 实现Error接口
func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.code, e.msg)
}

// code字段的get方法
func (e *Error) Code() int {
	return e.code
}

// msg字段的get方法
func (e *Error) Msg() string {
	return e.msg
}
