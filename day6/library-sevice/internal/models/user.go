package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key;"`
	Name      string     `json:"name" xml:"name" form:"name" query:"name"`
	Email     string     `json:"email" xml:"email" form:"email" query:"email"`
	Password  string     `json:"password" xml:"password" form:"password" query:"password"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"autoDeleteTime"`
}
