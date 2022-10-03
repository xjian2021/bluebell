package models

type (
	User struct {
		UserID   int64  `db:"user_id,omitempty"`
		Username string `db:"username,omitempty"`
		Password string `db:"password,omitempty"`
		Email    string `db:"email,omitempty"`
		Gender   string `db:"gender,omitempty"`
	}
)
