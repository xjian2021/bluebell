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

	CreatePost struct {
		Title       string `json:"title" binding:"required"`
		Content     string `json:"content" binding:"required"`
		CommunityID int64  `json:"community_id" binding:"required"`
		AuthorID    int64
	}
)
