package movie_test

import (
	"testing"

	mock_movie "github.com/gkatanacio/sample-lambda-app/mocks/pkg/movie"
	"github.com/gkatanacio/sample-lambda-app/pkg/movie"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Service_CreateMovie(t *testing.T) {
	// given
	testMovie := movie.Movie{
		Title:    "some-movie",
		Director: "some-director",
		Year:     2021,
	}

	createdMovie := &movie.Movie{
		Id:       uuid.New(),
		Title:    testMovie.Title,
		Director: testMovie.Director,
		Year:     testMovie.Year,
	}

	mockDb := new(mock_movie.Repository)
	mockDb.On("CreateMovie", testMovie).Return(createdMovie, nil)

	mockSqsClient := new(mock_movie.ProcessMovieQueuer)
	mockSqsClient.On("SendMovieForProcessing", *createdMovie).Return(nil)

	service := movie.NewService("testing only", mockDb, mockSqsClient)

	// when
	result, err := service.CreateMovie(testMovie)

	// then
	assert.NoError(t, err)
	assert.Equal(t, createdMovie, result)
	mockDb.AssertExpectations(t)
	mockSqsClient.AssertExpectations(t)
}

func Test_Service_GetMovie(t *testing.T) {
	// given
	id := uuid.New()

	fetchedMovie := &movie.Movie{
		Id:       id,
		Title:    "some-movie",
		Director: "some-director",
		Year:     2021,
	}

	mockDb := new(mock_movie.Repository)
	mockDb.On("GetMovie", id).Return(fetchedMovie, nil)

	service := movie.NewService("testing only", mockDb, nil)

	// when
	result, err := service.GetMovie(id)

	// then
	assert.NoError(t, err)
	assert.Equal(t, fetchedMovie, result)
	mockDb.AssertExpectations(t)
}

func Test_Service_ProcessMovie(t *testing.T) {
	// given
	testMovie := movie.Movie{
		Id:       uuid.New(),
		Title:    "some-movie",
		Director: "some-director",
		Year:     2021,
	}

	service := movie.NewService("testing only", nil, nil)

	// when
	err := service.ProcessMovie(testMovie)

	// then
	assert.NoError(t, err)
}
