package helpers

import "time"

type PublicUser struct {
	ID        uint       `gorm:"primary_key;"`
	Name      string     `json:"name" xml:"name" form:"name" query:"name"`
	Email     string     `json:"email" xml:"email" form:"email" query:"email"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"autoDeleteTime"`
}

type TestPublicUser struct {
	ID        uint       `gorm:"primary_key;"`
	Name      string     `json:"name" xml:"name" form:"name" query:"name"`
	Email     string     `json:"email" xml:"email" form:"email" query:"email"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"autoDeleteTime"`
}

type PublicUsers []PublicUser
type TestPublicUsers []TestPublicUser
