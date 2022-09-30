package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/xjian2021/bluebell/logic"
	"github.com/xjian2021/bluebell/models"
	validatorPkg "github.com/xjian2021/bluebell/pkg/validator-trans"
)

func SignUpHandler(c *gin.Context) {
	input := &models.SignUpInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		zap.L().Error("BindJSON fail", zap.Error(err))
		errors, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": errors.Translate(validatorPkg.Trans)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	zap.S().Debugf("SignUpInput:%+v", input)
	if err := logic.SignUp(input); err != nil {
		zap.L().Error("SignUp fail", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1})
}
