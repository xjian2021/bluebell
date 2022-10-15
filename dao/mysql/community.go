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

func GetCommunityDetail(id int64) (output *models.Community, err error) {
	output = &models.Community{}
	sqlStr := "select community_id,community_name,introduction from community where community_id = ?"
	err = db.Get(output, sqlStr, id)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
