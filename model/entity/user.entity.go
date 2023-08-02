package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey; column:id; autoIncrement"`
	Username  string         `json:"username" gorm:"index; unique; column:username; not null"`
	Email     string         `json:"email" gorm:"index; unique; column:email; not null"`
	Password  string         `json:"-" gorm:"column:password; not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;column:deleted_at"`
}
