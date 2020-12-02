package service

import (
	"github.com/ganselmo/go-book-store-api/internal/config"
	"github.com/ganselmo/go-book-store-api/internal/models"
	"github.com/jmoiron/sqlx"
)

func getAuthors() {

}

type AuthorService interface {
	FindAllAuthors() []*models.Author

	FindAuthorId(int) *models.Author

	AddAuthor(models.Author) error

	PatchAuthor(int) error

	DeleteAuthor(int) error
}

type authorService struct {
	db     *sqlx.DB
	config *config.Config
}

func NewAuthorService(db *sqlx.DB, c *config.Config) (AuthorService, error) {
	return authorService{db, c}, nil
}

func (s authorService) AddAuthor(models.Author) error {
	return nil
}
func (s authorService) FindAuthorId(int) *models.Author {
	return nil
}
func (s authorService) FindAllAuthors() []*models.Author {
	var list []*models.Author
	if err := s.db.Select(&list, "SELECT * FROM authors"); err != nil {
		panic(err)
	}
	return list
}

func (s authorService) PatchAuthor(int) error {
	return nil
}

func (s authorService) DeleteAuthor(int) error {
	return nil
}
