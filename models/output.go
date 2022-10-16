package models

import "time"

type (
	ResponseData struct {
		Code int64       `json:"code"`
		Msg  string      `json:"msg,omitempty"`
		Data interface{} `json:"data,omitempty"`
	}

	LoginResData struct {
		Token    string `json:"token,omitempty"`
		Username string `json:"username,omitempty"`
		Email    string `json:"email,omitempty"`
		UserID   int64  `json:"user_id,omitempty"`
	}

	PostListResData struct {
		Data []*Post `json:"data"`
	}

	PostDetailResData struct {
		ID            int64     `json:"id,string" db:"post_id"`
		Status        int32     `json:"status" db:"status"`
		Title         string    `json:"title" db:"title"`
		Content       string    `json:"content" db:"content"`
		Username      string    `json:"author_username" db:"username"`
		CommunityName string    `json:"community_name" db:"community_name"`
		Introduction  string    `json:"introduction" db:"introduction"`
		CreateTime    time.Time `json:"create_time" db:"create_time"`
	}
)
