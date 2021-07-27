package movie

import (
	"github.com/google/uuid"
)

type Movie struct {
	Id       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Director string    `json:"director"`
	Year     int       `json:"year"`
}
