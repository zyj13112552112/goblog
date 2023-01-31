package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/middleware"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
)

//登录模块接口

func Login(c *gin.Context) {
	var user model.User
	var token string
	var code int
	c.ShouldBindJSON(&user)
	code = model.CheckLogin(user.Username, user.Password)
	if code == errmsg.SUCCSE { //密码正确
		//生成token
		token, code = middleware.SetToken(user.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})

}
