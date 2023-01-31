package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

//文章模块接口

// 查询分类下的所有文章
func GetCateArt(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	if pageNum == 0 {
		pageNum = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, total := model.GetCateArt(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个文章
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询文章列表
func GetArticle(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))

	if pageNum == 0 {
		pageNum = -1
	} //gorm取消这个参数
	if pageSize == 0 {
		pageSize = -1
	} //gorm取消这个参数

	data, code, total := model.GetArticle(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 添加文章
func AddArticle(ctx *gin.Context) {
	var data model.Article
	Err := ctx.ShouldBindJSON(&data) //解析

	if Err != nil {
		fmt.Println("解析json出错", Err)
	}

	code = model.CreateArticle(&data)

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 编辑文章
func EditArticle(ctx *gin.Context) {
	var data model.Article
	ctx.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(ctx.Param("id"))
	code = model.EditArticle(id, &data)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id")) //分类id
	code = model.DeleteArticle(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
