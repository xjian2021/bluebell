package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

func SignUp(input *models.SignUpInput) (err error) {
	exist, err := mysql.CheckUserExist(input.Username)
	if err != nil {
		return fmt.Errorf("CheckUserExist username:%s err:%s", input.Username, err.Error())
	}
	if exist {
		return fmt.Errorf("user:%s exist", input.Username)
	}
	userID := snowflake.GenID()
	newUser := models.User{
		UserID:   userID,
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}
	id, err := mysql.AddUser(newUser)
	if err != nil {
		return fmt.Errorf("AddUser newUser:%+v err:%s", newUser, err.Error())
	}
	zap.S().Infof("new user id:%d", id)
	return nil
}
