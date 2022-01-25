package action

import (
	"github.com/labstack/echo/v4"
	"github.com/lfdubiela/api-bank-transfer/pkg/adapter"
	"github.com/lfdubiela/api-bank-transfer/pkg/adapter/web/port"
	"github.com/lfdubiela/api-bank-transfer/pkg/domain/usecase"
	"net/http"
)

func GetHealth(di adapter.DependencyInjector) echo.HandlerFunc {
	return func(c echo.Context) error {
		presenter := port.NewHealthPresenter()

		output := usecase.NewHealthInteractor(di.DBClient, presenter).
			Execute(c.Request().Context())

		return c.JSON(http.StatusOK, output)
	}
}
