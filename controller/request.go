package controller

import (
	"bytes"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	validatorPkg "github.com/xjian2021/bluebell/pkg/validator-trans"
)

const (
	ReqKey    = "req-key"
	UserIDKey = "userID-key"
)

func AuthBindJson(c *gin.Context, input interface{}) (err error) {
	return authErrorHandler(c.ShouldBindJSON(input))
}

func AuthBindQuery(c *gin.Context, input interface{}) (err error) {
	return authErrorHandler(c.ShouldBindQuery(input))
}

func authErrorHandler(err error) error {
	if err != nil {
		ves, ok := err.(validator.ValidationErrors)
		if ok {
			buff := bytes.NewBufferString("")
			for key, tranErr := range ves.Translate(validatorPkg.Trans) {
				buff.WriteString(key)
				buff.WriteString(": ")
				buff.WriteString(tranErr)
				buff.WriteString("; ")
			}
			err = errors.New(strings.TrimSpace(buff.String()))
		}
	}
	return err
}

func GetCurrenUserID(c *gin.Context) (userID int64) {
	return c.Value(UserIDKey).(int64)
}
