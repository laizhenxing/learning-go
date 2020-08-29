package model

// 默认表名是 `users`
type User struct {
	UserId   int    `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	UserName string `gorm:"column:username;type:varchar(50);unique_index"`
	Password string `gorm:"column:password;type:varchar(50)`
}

// 设置表名
func (u *User) TableName() string {
	return "user"
}