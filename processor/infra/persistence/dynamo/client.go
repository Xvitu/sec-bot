package dynamo

import (
	"context"
	"fmt"
	"xvitu/sec-bot/shared/env"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewClient(ctx context.Context) (*dynamodb.Client, error) {
	envs := env.Get()

	endpoint := envs.DynamoDBEndpoint
	region := envs.AwsRegion
	if region == "" {
		region = "us-east-1"
	}

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				"dummy",
				"dummy",
				"",
			),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("falha ao carregar configuração AWS: %w", err)
	}

	if envs.Env == "local" {
		cfg.BaseEndpoint = aws.String(endpoint)
	}

	return dynamodb.NewFromConfig(cfg), nil
}
