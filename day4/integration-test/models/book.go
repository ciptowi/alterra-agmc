package models

type Book struct {
	ID     uint `gorm:"primary_key;"`
	UserId int
	Title  string `json:"title"`
	Isbn   string `json:"isbn"`
	Writer string `json:"writer"`
}

type TestBook struct {
	ID     uint `gorm:"primary_key;"`
	UserId int
	Title  string `json:"title"`
	Isbn   string `json:"isbn"`
	Writer string `json:"writer"`
}

type Books []Book

type TestBooks []TestBook
