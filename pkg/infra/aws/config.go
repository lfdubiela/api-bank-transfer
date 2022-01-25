package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/lfdubiela/api-bank-transfer/pkg/infra/env"
	"github.com/lfdubiela/api-bank-transfer/pkg/infra/logger"
)

func GetConfig(ctx context.Context) aws.Config {
	environment := env.Get().Env
	endpoint := env.Get().AwsEndpoint

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if environment == "dev" && endpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           endpoint,
				SigningRegion: region,
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := getSessionConfig(ctx, customResolver)

	if err != nil {
		logger.Fatal("unable to load SDK config", "startup", logger.Fields{"error": err.Error()})
	}

	return cfg
}

func getSessionConfig(ctx context.Context, customResolver aws.EndpointResolverWithOptionsFunc) (aws.Config, error) {
	region := env.Get().AwsRegion
	key := env.Get().AwsKey
	secret := env.Get().AwsSecret
	tokenFile := env.Get().AwsWebTokenFile

	if tokenFile != "" {
		return config.LoadDefaultConfig(ctx, config.WithRegion(region))
	}

	credentialsProvider := credentials.NewStaticCredentialsProvider(key, secret, "")
	return config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentialsProvider),
		config.WithRegion(region),
		config.WithEndpointResolverWithOptions(customResolver))
}
