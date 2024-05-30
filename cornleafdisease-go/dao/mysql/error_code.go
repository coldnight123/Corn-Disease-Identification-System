package mysql

import "errors"

var (
	ErrorUserExit        = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
	ErrorDiseaseNotExist = errors.New("病害不存在")
)