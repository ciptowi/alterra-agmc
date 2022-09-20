package models

type Book struct {
	ID     uint   `gorm:"primary_key;"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Isbn   string `json:"isbn"`
	Writer string `json:"writer"`
}

type Books []Books
