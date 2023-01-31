package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"goblog/utils/validate"
	"net/http"
	"strconv"
)

//用户模块接口

var code int

// 添加用户
func AddUser(ctx *gin.Context) {
	var data model.User
	var msg string
	Err := ctx.ShouldBindJSON(&data) //解析

	msg, code = validate.Validate(data)

	if code != errmsg.SUCCSE {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}

	if Err != nil {
		fmt.Println("解析json出错", Err)
	}

	code = model.CheckUser(data.Username) //检查用户是否已经存在

	//fmt.Println(code)

	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}

	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

//查询单个用户

// 查询用户列表，分页查询
func GetUsers(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))

	if pageNum == 0 {
		pageNum = -1
	} //gorm取消这个参数
	if pageSize == 0 {
		pageSize = -1
	} //gorm取消这个参数

	data, code, total := model.GetUsers(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(ctx *gin.Context) {
	var data model.User
	ctx.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username) //查询更新的用户名是否已经存在
	if code == errmsg.SUCCSE {
		id, _ := strconv.Atoi(ctx.Param("id"))
		model.EditUser(id, &data) // 修改用户信息
	}
	if code == errmsg.ERROR_USERNAME_USED {
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户,gorm采用的是软删除
func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id")) //用户id
	code = model.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
