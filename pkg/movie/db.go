package movie

import (
	"github.com/google/uuid"
)

type Repository interface {
	CreateMovie(movie Movie) (*Movie, error)
	GetMovie(id uuid.UUID) (*Movie, error)
}

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) CreateMovie(movie Movie) (*Movie, error) {
	created := movie
	created.Id = uuid.New()

	return &created, nil
}

func (db *Database) GetMovie(id uuid.UUID) (*Movie, error) {
	fetched := &Movie{
		Id:       id,
		Title:    "Forrest Gump",
		Director: "Robert Zemeckis",
		Year:     1994,
	}

	return fetched, nil
}
