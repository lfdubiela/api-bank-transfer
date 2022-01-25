package usecase

import (
	"context"
	"github.com/lfdubiela/api-bank-transfer/pkg/adapter/web/port"
	"github.com/lfdubiela/api-bank-transfer/pkg/infra/database"
)

type (
	HealthUseCase interface {
		Execute(ctx context.Context) port.HealthOutput
	}

	healthInteractor struct {
		mysql     database.Client
		presenter port.HealthPresenter
	}
)

func NewHealthInteractor(
	mysql database.Client,
	presenter port.HealthPresenter) HealthUseCase {
	return healthInteractor{mysql: mysql, presenter: presenter}
}

func (b healthInteractor) Execute(ctx context.Context) port.HealthOutput {
	result := b.mysql.CheckDB(ctx)
	return b.presenter.Output(result)
}
