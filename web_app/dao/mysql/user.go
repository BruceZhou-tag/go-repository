package mysql

import (
	"database/sql"
	"encoding/base64"
	"web_app/models"

	"go.uber.org/zap"
	"golang.org/x/crypto/scrypt"
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return 
}

func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into user(user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(Pasword string) string {
	salt := []byte{0xc6, 0x28, 0xf7, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
	dk, err := scrypt.Key([]byte(Pasword), salt, 32768, 8, 1, 32)
	if err != nil {
		zap.L().Error("加密失败", zap.String("encryptPassword", "密码加密失败"))
	}
	return base64.StdEncoding.EncodeToString(dk)
}

func Login(user *models.User) (err error){
	oPassword := user.Password
	sqlStr := `select user_id,username,password from user where username = ?`
	err = db.Get(user,sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return 
}