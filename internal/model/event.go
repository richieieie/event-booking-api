package model

import (
	"time"
)

type Event struct {
	Id          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(255);not null"` // Specify varchar length, and not null
	Description string    `gorm:"type:text;not null"`
	Location    string    `gorm:"type:varchar(255);not null"`
	DateTime    time.Time `gorm:"not null"` // By default, will be stored as timestamp
	UserId      int64     `gorm:"not null"`
	User        User      `gorm:"not null;foreignKey:UserId;"`
}
