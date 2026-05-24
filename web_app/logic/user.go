package logic

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

var (
	ErrorUserExist = errors.New("用户已存在")
	ErrorUserNotExist = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

func SignUp(u *models.SignupUser) (err error) {
	if err = mysql.CheckUserExist(u.Username); err != nil {
		return ErrorUserExist
	}

	userID := snowflake.GenID()

	user := &models.User{
		UserID:   userID,
		Username: u.Username,
		Password: u.Password,
	}
	return mysql.InsertUser(user)
}

func Login(u *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: u.Username,
		Password: u.Password,
	}
	if err = mysql.Login(user); err != nil {
		return nil, err
	}
	token,_ := jwt.GenToken(user.UserID,user.Username)
	
	user.Token = token 
	return
}
