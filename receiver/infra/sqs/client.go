package sqs

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SqsClient struct{}

func (c *SqsClient) Create() *sqs.Client {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(os.Getenv("AWS_REGION")),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           os.Getenv("SQS_ENDPOINT"),
					SigningRegion: region,
				}, nil
			}),
		),
	)
	if err != nil {
		log.Fatalf("erro carregando config: %v", err)
	}

	return sqs.NewFromConfig(cfg)
}
