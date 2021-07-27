package movie

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type ProcessMovieQueuer interface {
	SendMovieForProcessing(movie Movie) error
}

type SqsClient struct {
	*sqs.SQS
	processMovieQueueUrl string
}

func NewSqsClient(processMovieQueueUrl string) *SqsClient {
	sess := session.Must(session.NewSession())
	return &SqsClient{sqs.New(sess), processMovieQueueUrl}
}

func (c *SqsClient) SendMovieForProcessing(movie Movie) error {
	message, err := json.Marshal(movie)
	if err != nil {
		return err
	}

	_, err = c.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(message)),
		QueueUrl:    aws.String(c.processMovieQueueUrl),
	})

	return err
}
