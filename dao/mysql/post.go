package mysql

import "github.com/xjian2021/bluebell/models"

func CreatePost(post *models.Post) (newID int64, err error) {
	sqlStr := "insert into post( post_id, title, content, author_id, community_id) VALUES (?,?,?,?,?)"
	result, err := db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
