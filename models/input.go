package models

type (
	SignUpInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Age      int8   `json:"age" binding:"gte=1,lte=130" `
	}

	LoginInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	CreatePostInput struct {
		Title       string `json:"title" binding:"required"`
		Content     string `json:"content" binding:"required"`
		CommunityID int64  `json:"community_id" binding:"required"`
		AuthorID    int64
	}

	PostListInput struct {
		Order      string `json:"order" form:"order" binding:"required,oneof=time score"`
		LastPostID int64  `json:"last_post_id" form:"last_post_id" binding:"gte=0"`
		Limit      int64  `json:"limit" form:"limit" binding:"gt=0,lte=20"`
	}

	VoteInput struct {
		PostID    string `json:"post_id" binging:"required"`
		UserID    int64
		Operating int8 `json:"direction,string" binding:"oneof=-1 0 1"`
	}
)
