package model

import (
	"time"
)

type User struct {
	Id        int64     `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"unique;notnull"`
	Password  string    `gorm:"notnull"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
