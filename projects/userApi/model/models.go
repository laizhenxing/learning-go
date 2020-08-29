package model

import (
	"time"
)

type Model struct {
	ID        uint64     `gorm:"column:id;type:bigint(20);primary_key;auto_increment" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;" json:"deleted_at"`
}
