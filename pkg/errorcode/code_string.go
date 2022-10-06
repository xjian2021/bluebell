// Code generated by "stringer -type=Code -linecomment"; DO NOT EDIT.

package errorcode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CodeUnknownError-0]
	_ = x[CodeSuccess-1]
	_ = x[CodeInvalidParam-1002]
	_ = x[CodeUserExist-1003]
	_ = x[CodeUserNotExist-1004]
	_ = x[CodeInvalidPassword-1005]
	_ = x[CodeInvalidToken-1006]
}

const (
	_Code_name_0 = "未知错误成功"
	_Code_name_1 = "无效参数用户已存在用户不存在密码错误无效token"
)

var (
	_Code_index_0 = [...]uint8{0, 12, 18}
	_Code_index_1 = [...]uint8{0, 12, 27, 42, 54, 65}
)

func (i Code) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _Code_name_0[_Code_index_0[i]:_Code_index_0[i+1]]
	case 1002 <= i && i <= 1006:
		i -= 1002
		return _Code_name_1[_Code_index_1[i]:_Code_index_1[i+1]]
	default:
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
