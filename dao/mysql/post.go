package mysql

import (
	"database/sql"

	"github.com/xjian2021/bluebell/models"
)

func CreatePost(post *models.Post) (newID int64, err error) {
	sqlStr := "insert into post( post_id, title, content, author_id, community_id) VALUES (?,?,?,?,?)"
	result, err := db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetPostDetail(postID int64) {

}

func PostList(lastPostID, limit int64) (output []*models.Post, err error) {
	sqlStr := "select post_id,community_id,title,content from post where post_id > ? limit ?"
	err = db.Select(&output, sqlStr, lastPostID, limit)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
