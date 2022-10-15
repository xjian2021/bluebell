package logic

import (
	"github.com/xjian2021/bluebell/dao/mysql"
	"github.com/xjian2021/bluebell/models"
)

func GetCommunityList() (output []*models.Community, err error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(communityID int64) (output *models.Community, err error) {
	return mysql.GetCommunityDetail(communityID)
}
