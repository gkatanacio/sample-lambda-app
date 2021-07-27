package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gkatanacio/sample-lambda-app/pkg/errs"
	"github.com/gkatanacio/sample-lambda-app/pkg/handlerutil"
	"github.com/gkatanacio/sample-lambda-app/pkg/movie"
)

var movieService *movie.Service

func init() {
	cfg := movie.NewConfig()
	db := movie.NewDatabase()
	sqsClient := movie.NewSqsClient(cfg.ProcessMovieQueueUrl)

	movieService = movie.NewService(cfg.SomeSharedEnvVar, db, sqsClient)

	log.Printf("sample env var: %d", cfg.SomeFxnEnvVar)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("received: %s", request.Body)

	var movie movie.Movie
	if err := json.Unmarshal([]byte(request.Body), &movie); err != nil {
		log.Printf("error: %s", err.Error())
		return handlerutil.ErrorResponse(errs.NewBadRequest("invalid request body"))
	}

	created, err := movieService.CreateMovie(movie)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return handlerutil.ErrorResponse(err)
	}

	return handlerutil.Response(http.StatusCreated, created)
}

func main() {
	lambda.Start(handlerutil.CorsMiddleware(handler))
}
