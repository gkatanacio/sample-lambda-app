package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gkatanacio/sample-lambda-app/pkg/movie"
)

var movieService *movie.Service

func init() {
	cfg := movie.NewConfig()

	movieService = movie.NewService(cfg.SomeSharedEnvVar, nil, nil)
}

func handler(ctx context.Context, event events.SQSEvent) error {
	for _, record := range event.Records {
		log.Printf("received: %s", record.Body)

		var movie movie.Movie
		if err := json.Unmarshal([]byte(record.Body), &movie); err != nil {
			log.Printf("error: %s", err.Error())
			continue
		}

		if err := movieService.ProcessMovie(movie); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
