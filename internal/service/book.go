package service

import (
	"github.com/ganselmo/go-book-store-api/internal/config"
	"github.com/ganselmo/go-book-store-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type BookService interface {
	FindAllBooks() []*models.Book

	FindBooksById(int) *models.Book

	AddBook(models.Book) error

	PatchBook(int) error

	DeleteBook(int) error
}

type bookService struct {
	db     *sqlx.DB
	config *config.Config
}

func NewBookService(db *sqlx.DB, c *config.Config) (BookService, error) {
	return bookService{db, c}, nil
}

func (s bookService) AddBook(models.Book) error {
	return nil
}
func (s bookService) FindBooksById(int) *models.Book {
	return nil
}
func (s bookService) FindAllBooks() []*models.Book {
	var list []*models.Book
	if err := s.db.Select(&list, "SELECT * FROM books"); err != nil {
		panic(err)
	}
	return list
}

func (s bookService) PatchBook(int) error {
	return nil
}

func (s bookService) DeleteBook(int) error {
	return nil
}
