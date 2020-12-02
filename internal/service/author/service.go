package author

import (
	"github.com/ganselmo/go-book-store-api/internal/config"
	"github.com/jmoiron/sqlx"
)

type Author struct {
	id       int
	name     string
	lastname string
	location string
	raiting  int
}

func getAuthors() {

}

type Service interface {
	FindAll() []*Author

	FindById(int) *Author

	AddAuthor(Author) error

	PatchAuthor(int) error

	DeleteAuthor(int) error
}

type service struct {
	db     *sqlx.DB
	config *config.Config
}

func NewAuthorService(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddAuthor(Author) error {
	return nil
}
func (s service) FindById(int) *Author {
	return nil
}
func (s service) FindAll() []*Author {
	var list []*Author
	if err := s.db.Select(&list, "SELECT * FROM authors"); err != nil {
		panic(err)
	}
	return list
}

func (s service) PatchAuthor(int) error {
	return nil
}

func (s service) DeleteAuthor(int) error {
	return nil
}
