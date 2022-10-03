package models

type (
	ResponseData struct {
		Code int64       `json:"code"`
		Msg  interface{} `json:"msg,omitempty"`
		Data interface{} `json:"data,omitempty"`
	}

	LoginResData struct {
		Token    string `json:"token,omitempty"`
		Username string `json:"username,omitempty"`
		Email    string `json:"email,omitempty"`
		UserID   int64  `json:"user_id,omitempty"`
	}
)
