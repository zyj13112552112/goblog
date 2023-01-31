package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

//文章类型模块接口

// 添加分类
func AddCate(ctx *gin.Context) {
	var data model.Category
	Err := ctx.ShouldBindJSON(&data) //解析

	if Err != nil {
		fmt.Println("解析json出错", Err)
	}

	code = model.CheckCategory(data.Name) //检查用户是否已经存在

	//fmt.Println(code)

	if code == errmsg.SUCCSE {
		model.CreateCategory(&data)
	}

	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询分类列表
func GetCate(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))

	if pageNum == 0 {
		pageNum = -1
	} //gorm取消这个参数
	if pageSize == 0 {
		pageSize = -1
	} //gorm取消这个参数

	data, code, total := model.GetCategory(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类
func EditCate(ctx *gin.Context) {
	var data model.Category
	ctx.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name) //查询更新的用户名是否已经存在
	if code == errmsg.SUCCSE {
		id, _ := strconv.Atoi(ctx.Param("id"))
		model.EditCategory(id, &data) // 修改用户信息
	}
	if code == errmsg.ERROR_CATENAME_USED {
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCate(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id")) //分类id
	code = model.DeleteCategory(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo 查询分类下的所有文章
