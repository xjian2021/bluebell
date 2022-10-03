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
)
