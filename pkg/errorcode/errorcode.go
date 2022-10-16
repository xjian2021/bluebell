//go:generate stringer -type=Code -linecomment
package errorcode

const (
	CodeUnknownError    Code = iota        // 未知错误
	CodeSuccess                            // 成功
	CodeInvalidParam    Code = 1000 + iota // 无效参数
	CodeUserExist                          // 用户已存在
	CodeUserNotExist                       // 用户不存在
	CodeInvalidPassword                    // 密码错误
	CodeInvalidToken                       // 无效token
	CodeInvalidID                          // 无效id
)

type (
	Code int64
)

func (i Code) Error() string {
	return i.String()
}

func (i Code) Code() int64 {
	return int64(i)
}
