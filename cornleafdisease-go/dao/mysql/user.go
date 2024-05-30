package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"project/cornleafdisease/models"
)

const secret = "Honkai Star Railway"

// CheckUserExist 检测是否重名
func CheckUserExist(username string) (err error) {
	var user models.User
	DB.Where("username = ?", username).First(&user)
	if user.ID != 0 {
		return ErrorUserExit
	}
	return nil
}

// InsertUser 向数据库插入一条新的用户数据
func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	create := DB.Create(&user)
	err = create.Error
	return
}

// encryptPassword 加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	res := DB.Where("username = ?", user.Username).Find(&user)
	if res.RowsAffected == 0 {
		return ErrorUserNotExist
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	res := DB.Where("user_id = ?", uid).Find(&user)
	if res.RowsAffected == 0 {
		err = ErrorUserNotExist
	}
	return
}
