package handlerutil

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type APIGatewayLambdaHandler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func CorsMiddleware(handler APIGatewayLambdaHandler) APIGatewayLambdaHandler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		response, err := handler(ctx, request)
		if response.Headers == nil {
			response.Headers = make(map[string]string)
		}
		response.Headers["Access-Control-Allow-Origin"] = "*"
		response.Headers["Access-Control-Allow-Credentials"] = "true"

		return response, err
	}
}
