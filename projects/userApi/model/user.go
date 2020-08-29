package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"

	"userApi/pkg/auth"
	"userApi/pkg/constvar"
)

type User struct {
	Model

	Username string `json:"username" gorm:"column:username;type:varchar(255)" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;type:varchar(255)" binding:"required" validate:"min=5,max=32"`
}

// 创建用户
func (u *User) Create() error {
	return DB.Self.Create(&u).Error
}

// 更新用户信息
func (u *User) Update() error {
	// update: 更新已更改的字段
	return DB.Self.Model(&User{}).Updates(u).Error
}

// 密码加密
func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// 参数校验
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// 获取一个用户信息
func GetUser(username string) (*User, error) {
	u := &User{}
	db := DB.Self.Where("username =? ", username).First(&u)
	return u, db.Error
}

// 获取用户列表
func GetUserList(username string, offset, limit int) (users []*User, count uint64, err error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err = DB.Self.Model(&User{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err = DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// 删除用户
func DeleteUser(id uint64) error {
	// 启用SQL调式
	//DB.Self.LogMode(true)
	// DB.Self.Debug().Where("id = ?", id).Delete(User{})
	// 软删除1
	return DB.Self.Where("id = ?", id).Delete(User{}).Error
	// 软删除2
	//u := User{}
	//u.Model.ID = id
	//return DB.Self.Delete(&u).Error
	// 使用Unscoped永久删除
	//return DB.Self.Unscoped().Where("id=?", id).Delete(&User{}).Error
}

// 查询用户是否存在
//func CheckExist(username, password string) (bool, error) {
//	var user User
//	DB.Self.Where("username=?", username).Find(&user)
//	if user.ID > 0 {
//		hashedPass, err := auth.Encrypt(password)
//		if err != nil {
//			return false, errno.ErrEncrypt
//		}
//		err = auth.Compare(user.Password, hashedPass)
//		if err != nil {
//			return false, errno.ErrPasswordIncorrect
//		}
//
//		return true, nil
//	}
//
//	return false, errno.ErrUserNotFound
//}
