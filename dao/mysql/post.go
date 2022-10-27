package mysql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"

	"github.com/xjian2021/bluebell/models"
	"github.com/xjian2021/bluebell/pkg/errorcode"
)

func CreatePost(post *models.Post) (newID int64, err error) {
	sqlStr := "insert into post( post_id, title, content, author_id, community_id) VALUES (?,?,?,?,?)"
	result, err := db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetPostDetail(postID int64) (output *models.PostDetailResData, err error) {
	output = new(models.PostDetailResData)
	sqlStr := "select p.post_id,p.status,p.title,p.content,p.create_time,c.community_name,c.introduction,u.username from post as p join community c on c.community_id = p.community_id join users u on p.author_id = u.user_id where p.post_id = ?"
	err = db.Get(output, sqlStr, postID)
	if err == sql.ErrNoRows {
		err = errorcode.CodeInvalidID
	}
	return
}

func PostList(postIDs []string, limit int64) (output []*models.Post, err error) {
	sqlStr := "select post_id,community_id,title,content from post where post_id in (?) order by find_in_set(post_id,?) limit ?"
	query, args, err := sqlx.In(sqlStr, postIDs, strings.Join(postIDs, ","), limit)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	err = db.Select(&output, query, args...)
	return
}
