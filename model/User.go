package model

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"goblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"log"
)

//用户

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null " json:"password" validate:"required,min=6,max=20" label:"用户密码"`
	//角色类型，管理or读者
	Role int `gorm:"type:int;default:2 " json:"role" validate:"required,gte=2" label:"角色码"`
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username=?", name).First(&user)
	if user.ID >= 1 { //gorm.Model自带
		return errmsg.ERROR_USERNAME_USED //用户名已存在 1001
	}
	return errmsg.SUCCSE //用户名不存在
}

// 新增用户
func CreateUser(data *User) (code int) {
	data.Password = ScryptPw(data.Password) //方式一 ：直接加密
	err_ := db.Create(&data).Error
	if err_ != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //200
}

//方式二 ：使用钩子函数BeforeSave() === 方式一
//func (u *User)BeforeSave(){
//	u.Password = ScryptPw(u.Password)
//}

// 查询用户列表，分页查询
func GetUsers(pageSize, pageNum int) ([]User, int, int) {
	var users []User
	var total int
	ERR := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if ERR != nil && ERR != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0 //500
	}
	return users, errmsg.SUCCSE, total //200
}

// 编辑用户(修改用户密码独立出去，不在这里)
func EditUser(id int, data *User) (code int) {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	Err := db.Model(&user).Where("id=?", id).Update(maps).Error
	if Err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 用户密码加密
func ScryptPw(Password string) string {
	//使用scrypt库进行加密
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	HashPw, ERR := scrypt.Key([]byte(Password), salt, 16384, 8, 1, KeyLen)
	if ERR != nil {
		log.Fatal(ERR)
	}
	//HashPw转字符串返回
	return base64.StdEncoding.EncodeToString(HashPw)
}

// 删除用户，gorm采用的是软删除
func DeleteUser(id int) (code int) {
	var usr User
	Err := db.Where("id=?", id).Delete(&usr).Error
	if Err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //200
}

// 登录验证
func CheckLogin(username string, password string) int {
	var user User
	//var code int
	db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST //用户不存在
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG //密码错误
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NOT_RIGHT //该用户无权限
	}
	return errmsg.SUCCSE
}
