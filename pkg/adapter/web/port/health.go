package port

import (
	"github.com/lfdubiela/api-bank-transfer/pkg/infra/database"
)

type (
	HealthPresenter interface {
		Output(check database.DBChecker) HealthOutput
	}

	HealthOutput struct {
		Database Database `json:"database"`
	}

	Database struct {
		Read  bool `json:"read"`
		Write bool `json:"write"`
	}

	healthPresenter struct{}
)

func NewHealthPresenter() HealthPresenter {
	return healthPresenter{}
}

func (p healthPresenter) Output(check database.DBChecker) HealthOutput {
	return HealthOutput{
		Database: Database{
			Read:  check.CanRead(),
			Write: check.CanWrite(),
		},
	}
}
