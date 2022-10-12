package mysql

import (
	"database/sql"

	"github.com/xjian2021/bluebell/models"
)

func GetCommunityList() (output []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	err = db.Select(&output, sqlStr)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
