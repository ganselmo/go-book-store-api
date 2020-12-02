package book

import (
	"github.com/ganselmo/go-book-store-api/internal/config"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	id          int
	name        string
	description string
	rating      int
	pages       int
}

type Service interface {
	FindAll() []*Book

	FindById(int) *Book

	AddBook(Book) error

	PatchBook(int) error

	DeleteBook(int) error
}

type service struct {
	db     *sqlx.DB
	config *config.Config
}

func NewBookService(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddBook(Book) error {
	return nil
}
func (s service) FindById(int) *Book {
	return nil
}
func (s service) FindAll() []*Book {
	var list []*Book
	if err := s.db.Select(&list, "SELECT * FROM books"); err != nil {
		panic(err)
	}
	return list
}

func (s service) PatchBook(int) error {
	return nil
}

func (s service) DeleteBook(int) error {
	return nil
}
