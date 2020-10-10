package model

import "fmt"

type TodoList struct {
	Model

	Title string `gorm:"column:title;" json:"title"`
	Status bool `gorm:"status;default:false" json:"status"`
}

func (td *TodoList) CreateTodo() error {
	return DB.Create(&td).Error
}

func (td *TodoList) UpdateTodo() error {
	return DB.Save(&td).Error
}

func GetTodo(id uint) (*TodoList, error) {
	var tl TodoList
	err := DB.Where("id = ?", id).First(&tl).Error
	return &tl, err
}

func GetTodoList(title string, offset, limit int) ([]*TodoList, error) {
	var tls []*TodoList
	where := fmt.Sprintf("title like %%%s%%", title)
	err := DB.Where(where).Offset(offset).Limit(limit).Find(&tls).Error
	return tls, err
}

func DeleteTodo(id uint) error {
	return DB.Delete(id).Error
}

func ExistTodoByID(id uint) bool {
	var tl TodoList
	DB.Where("id = ? ", id).First(&tl)
	if tl.ID > 0 {
		return true
	}
	return false
}