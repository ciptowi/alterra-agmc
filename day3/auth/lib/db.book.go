package lib

import (
	"auth/models"
)

// var db, _ = from /db.user.go

func CreateBook(book *models.Book) error {
	if err := db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func FindBooks(books *models.Books) error {
	if err := db.Find(&books).Error; err != nil {
		return err
	}
	return nil
}

func FindBookById(id int, book *models.Book) error {
	if err := db.First(&book, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBookById(id int, user *models.Book) error {
	if err := db.Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBookById(id int, user *models.Book) error {
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
