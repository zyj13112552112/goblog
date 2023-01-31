package model

import (
	"github.com/jinzhu/gorm"
	"goblog/utils/errmsg"
)

// 文章
type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"` //cid外键

	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArticle(data *Article) (code int) {
	err_ := db.Create(&data).Error
	if err_ != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //200
}

// 编辑文章
func EditArticle(id int, data *Article) (code int) {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	maps["cid"] = data.Cid

	Err := db.Model(&art).Where("id=?", id).Update(maps).Error
	if Err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArticle(id int) (code int) {
	var art Article
	Err := db.Where("id=?", id).Delete(&art).Error
	if Err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //200
}

// 查询分类下的所有文章
func GetCateArt(id, pageSize, pageNum int) ([]Article, int, int) { //id:分类的id
	var CatArtList []Article
	var total int
	ERR := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).
		Where("cid=?", id).Find(&CatArtList).Count(&total).Error
	if ERR != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0 //分类不存在
	}
	return CatArtList, errmsg.SUCCSE, total
}

// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	ERR := db.Preload("Category").Where("id=?", id).First(&art).Error
	if ERR != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCSE
}

// 查询文章列表
func GetArticle(pageSize, pageNum int) ([]Article, int, int) {
	var art []Article
	var total int
	ERR := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Count(&total).Error
	//ERR := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Error //得不到category
	if ERR != nil && ERR != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0 //500
	}
	return art, errmsg.SUCCSE, total //200
}
