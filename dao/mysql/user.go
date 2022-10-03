package mysql

import (
	"github.com/xjian2021/bluebell/models"
	"github.com/xjian2021/bluebell/pkg/errorcode"
)

func CheckUserExist(username string) error {
	sqlStr := "select count(*) from users where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errorcode.CodeUserExist
	}
	return nil
}

func AddUser(user *models.User) (newID int64, err error) {
	sqlStr := "insert into users(user_id, username, password, email) values (?,?,?,?)"
	result, err := db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.Email)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetUserByUsername(username string) (user *models.User, err error) {
	user = &models.User{}
	sqlStr := "select email,user_id, password from users where username = ?"
	err = db.Get(user, sqlStr, username)
	return user, err
}
