package handlerutil

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gkatanacio/sample-lambda-app/pkg/errs"
)

func Response(status int, body interface{}) (events.APIGatewayProxyResponse, error) {
	if body == nil {
		return events.APIGatewayProxyResponse{StatusCode: status}, nil
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBody),
		StatusCode: status,
	}, nil
}

type errorResponseBody struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

func ErrorResponse(err error) (events.APIGatewayProxyResponse, error) {
	body := &errorResponseBody{}
	body.Error.Message = err.Error()

	var status int
	switch e := err.(type) {
	case errs.HttpError:
		status = e.StatusCode()
	default:
		status = http.StatusInternalServerError
	}

	return Response(status, body)
}
