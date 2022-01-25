package infra

import (
	"context"
	"github.com/lfdubiela/api-bank-transfer/pkg/adapter"
	"github.com/lfdubiela/api-bank-transfer/pkg/adapter/web"
	"github.com/lfdubiela/api-bank-transfer/pkg/infra/aws"
	"github.com/lfdubiela/api-bank-transfer/pkg/infra/database"
	"github.com/lfdubiela/api-bank-transfer/pkg/infra/server"
)

func Start() {
	echo := server.Config()

	// setup dependencies
	db := database.Config()
	cfg := aws.GetConfig(context.TODO())
	di := adapter.NewDI(cfg, db)

	web.NewRouter(echo, di).RegisterHandlers()
}
