package adapter

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/lfdubiela/api-bank-transfer/pkg/infra/database"
)

type DependencyInjector struct {
	AwsCfg   aws.Config
	DBClient database.Client
}

func NewDI(cfg aws.Config, dbClient database.Client) DependencyInjector {
	return DependencyInjector{
		AwsCfg:   cfg,
		DBClient: dbClient,
	}
}
