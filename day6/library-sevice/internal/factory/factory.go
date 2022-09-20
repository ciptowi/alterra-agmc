package factory

import (
	"library-sevice/database"
	"library-sevice/internal/repository"
)

type Factory struct {
	BookRepository repository.Book
	UserRepository repository.User
}

func NewFactory() *Factory {
	db, _ := database.CreateConnection()
	return &Factory{
		BookRepository: repository.NewBook(db),
		UserRepository: repository.NewUser(db),
	}
}
