package model

import (
	"time"
)

type User struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Email     string `gorm:"column:email;uniqueIndex;not null" json:"email"`
	Name      string `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}
