package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gkatanacio/sample-lambda-app/pkg/errs"
	"github.com/gkatanacio/sample-lambda-app/pkg/handlerutil"
	"github.com/gkatanacio/sample-lambda-app/pkg/movie"
	"github.com/google/uuid"
)

var movieService *movie.Service

func init() {
	cfg := movie.NewConfig()
	db := movie.NewDatabase()

	movieService = movie.NewService(cfg.SomeSharedEnvVar, db, nil)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	idParam := request.PathParameters["id"]
	log.Printf("id: %s", idParam)

	id, err := uuid.Parse(idParam)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return handlerutil.ErrorResponse(errs.NewBadRequest("id param not a valid uuid"))
	}

	movie, err := movieService.GetMovie(id)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return handlerutil.ErrorResponse(err)
	}

	return handlerutil.Response(http.StatusOK, movie)
}

func main() {
	lambda.Start(handlerutil.CorsMiddleware(handler))
}
