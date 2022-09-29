package mysql

import "bluebell/models"

func CheckUserExist(username string) (bool, error) {
	sqlStr := "select count(*) from users where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

func AddUser(user models.User) (newID int64, err error) {
	sqlStr := "insert into users(user_id, username, password, email) values (?,?,?,?)"
	result, err := db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.Email)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
