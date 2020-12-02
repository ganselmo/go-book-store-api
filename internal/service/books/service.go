package book

import (
	"github.com/ganselmo/go-book-store-api/internal/config"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	Id          int
	Name        string
	Description string
	Rating      int
	Pages       int
	Author      string
}

type Service interface {
	FindAll() []*Book

	FindById(int) *Book

	AddBook(Book)

	PatchBook(Book)

	DeleteBook(int)
}

type service struct {
	db     *sqlx.DB
	config *config.Config
}

func NewBookService(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) FindAll() []*Book {
	var list []*Book
	if err := s.db.Select(&list, "SELECT * FROM books"); err != nil {
		panic(err)
	}
	return list
}

func (s service) FindById(id int) *Book {
	var b []*Book
	if err := s.db.Select(&b, "SELECT * FROM books WHERE id = ?", id); err != nil {
		panic(err)
	}

	if len(b) > 0 {
		return b[0]
	} else {
		return nil
	}
}

func (s service) AddBook(b Book) {
	query := `INSERT INTO books (
		name, description, rating, 
		pages, author) VALUES (?, ?, ?, ?, ?)`

	_, err := s.db.Exec(query,
		b.Name, b.Description, b.Rating, b.Pages, b.Author)

	if err != nil {
		panic(err)
	}
}

func (s service) PatchBook(b Book) {
	query := `UPDATE books 
	SET name = ?, description = ?, rating = ?, 
	pages = ?, author = ?
	WHERE id = ?`
	_, err := s.db.Exec(query, b.Name,
		b.Description, b.Rating, b.Pages, b.Author, b.Id)

	if err != nil {
		panic(err)
	}
}

func (s service) DeleteBook(id int) {
	query := "DELETE FROM books WHERE id = ?"

	_, err := s.db.Exec(query, id)

	if err != nil {
		panic(err)
	}
}
