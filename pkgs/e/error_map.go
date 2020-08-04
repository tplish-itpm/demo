package e

const (
	Success          = 200
	Fail             = 500
	ErrInvalidParams = iota + 2001
	ErrUserExist
	ErrUserNotExist
	ErrNamePwd
	ErrNameFormat
	ErrPwdFormat
	ErrEmailFormat
	ErrPermission
	ErrEncrypt
)

var ErrorMap = map[int]string{
	Success:          "ok",
	Fail:             "fail",
	ErrInvalidParams: "请求参数错误",
	ErrUserExist:     "用户已存在",
	ErrUserNotExist:  "用户不存在",
	ErrNamePwd:       "用户名或密码错误",
	ErrNameFormat:    "用户名格式错误",
	ErrPwdFormat:     "密码格式错误",
	ErrEmailFormat:   "邮箱格式错误",
	ErrPermission:    "权限不足",
	ErrEncrypt:       "加密失败",
}

func GetMsg(code int) string {
	msg, ok := ErrorMap[code]
	if ok {
		return msg
	}
	return ErrorMap[Fail]
}
