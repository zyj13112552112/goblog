package model

import (
	"github.com/jinzhu/gorm"
	"goblog/utils/errmsg"
)

//文章类型

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name=?", name).First(&cate)
	if cate.ID >= 1 { //gorm.Model自带
		return errmsg.ERROR_CATENAME_USED //分类已存在 2001
	}
	return errmsg.SUCCSE //分类不存在
}

// 新增分类
func CreateCategory(cate *Category) (code int) {
	err_ := db.Create(&cate).Error
	if err_ != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //200
}

// 查询分类列表，分页查询
func GetCategory(pageSize, pageNum int) ([]Category, int, int) {
	var cate []Category
	var total int
	ERR := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	if ERR != nil && ERR != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0 //500
	}
	return cate, errmsg.SUCCSE, total //200
}

// 编辑分类
func EditCategory(id int, data *Category) (code int) {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	Err := db.Model(&cate).Where("id=?", id).Update(maps).Error
	if Err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteCategory(id int) (code int) {
	var cate Category
	Err := db.Where("id=?", id).Delete(&cate).Error
	if Err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //200
}
