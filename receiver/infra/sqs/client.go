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

func (c *SqsClient) Create(ctx context.Context) *sqs.Client {
	if ctx == nil {
		ctx = context.Background()
	}

	customEndpoint := os.Getenv("SQS_ENDPOINT")
	region := os.Getenv("AWS_REGION")

	var opts []func(*config.LoadOptions) error

	if customEndpoint != "" {
		opts = append(opts, config.WithRegion(region))
		resolver := aws.EndpointResolverWithOptionsFunc(
			func(service, r string, _ ...interface{}) (aws.Endpoint, error) {
				if service == sqs.ServiceID {
					return aws.Endpoint{
						URL:           customEndpoint,
						SigningRegion: region,
					}, nil
				}
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			},
		)

		opts = append(opts, config.WithEndpointResolverWithOptions(resolver))
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		log.Fatalf("erro carregando config: %v", err)
	}

	return sqs.NewFromConfig(cfg)
}
