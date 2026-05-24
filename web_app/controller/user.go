package controller

import (
	"errors"
	"fmt"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {

	// var u models.SignupUser
	u := new(models.SignupUser)
	if err := c.ShouldBindJSON(u); err != nil {
		zap.L().Error("SignUpHandler with invalid params err", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, CodeInvalidParam.Msg())
		return
	}
	// if len(u.Username) == 0 || len(u.Password) == 0 || len(u.ConfirmPassword) == 0 || u.Password != u.ConfirmPassword {
	// 	zap.L().Error("SignUpHandler with invalid params err")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "请求参数有误",
	// 	})
	// 	return
	// }
	// fmt.Println(u)
	if err := logic.SignUp(u); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, CodeServerBusy.Msg())
		return
	}

	ResponseSuccess(c, CodeSuccess.Msg())
}

func LoginHandler(c *gin.Context) {
	u := new(models.ParamLogin)

	if err := c.ShouldBindJSON(u); err != nil {
		zap.L().Error("LoginHandler with invalid params err", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, CodeInvalidParam.Msg())
		return
	}
	user, err := logic.Login(u)
	if err != nil {
		zap.L().Error("login user does not exists", zap.String("username", u.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
