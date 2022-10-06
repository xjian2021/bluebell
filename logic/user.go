package logic

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/xjian2021/bluebell/dao/mysql"
	"github.com/xjian2021/bluebell/models"
	"github.com/xjian2021/bluebell/pkg/errorcode"
	"github.com/xjian2021/bluebell/pkg/jwt"
	"github.com/xjian2021/bluebell/pkg/snowflake"
	"github.com/xjian2021/bluebell/pkg/utils"
)

func SignUp(input *models.SignUpInput) (err error) {
	if err = mysql.CheckUserExist(input.Username); err != nil {
		return err
	}
	userID := snowflake.GenID()
	newUser := &models.User{
		UserID:   userID,
		Username: input.Username,
		Password: utils.Md5(input.Password),
		Email:    input.Email,
	}
	id, err := mysql.AddUser(newUser)
	if err != nil {
		return fmt.Errorf("AddUser newUser:%+v err:%s", newUser, err.Error())
	}
	zap.S().Infof("new user id:%d", id)
	return nil
}

func Login(input *models.LoginInput) (output *models.LoginResData, err error) {
	user, err := mysql.GetUserByUsername(input.Username)
	if err != nil {
		return nil, err
	}
	if user.Password != utils.Md5(input.Password) {
		return nil, errorcode.CodeInvalidPassword
	}

	token, err := jwt.GenToken(user.UserID, input.Username)
	if err != nil {
		return nil, err
	}

	return &models.LoginResData{
		Token:    token,
		Username: input.Username,
		Email:    user.Email,
		UserID:   user.UserID,
	}, nil
}
