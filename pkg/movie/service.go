package movie

import (
	"log"

	"github.com/google/uuid"
)

type Service struct {
	sharedEnvVar string
	db           Repository
	sqsClient    ProcessMovieQueuer
}

func NewService(sharedEnvVar string, db Repository, sqsClient ProcessMovieQueuer) *Service {
	return &Service{sharedEnvVar, db, sqsClient}
}

func (s *Service) CreateMovie(movie Movie) (*Movie, error) {
	s.logSharedEnvVar()

	created, err := s.db.CreateMovie(movie)
	if err != nil {
		return nil, err
	}

	if err := s.sqsClient.SendMovieForProcessing(*created); err != nil {
		return nil, err
	}

	return created, nil
}

func (s *Service) GetMovie(id uuid.UUID) (*Movie, error) {
	s.logSharedEnvVar()

	return s.db.GetMovie(id)
}

func (s *Service) ProcessMovie(movie Movie) error {
	s.logSharedEnvVar()

	log.Printf("processing movie: %v", movie)

	return nil
}

func (s *Service) logSharedEnvVar() {
	log.Printf("test: %s", s.sharedEnvVar)
}
