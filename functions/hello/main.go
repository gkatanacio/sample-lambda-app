package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
)

func handler(ctx context.Context, event events.CloudWatchEvent) error {
	log.Printf("hello %s!", os.Getenv("HELLO_WHO"))
	return nil
}

func main() {
	lambda.Start(handler)
}
