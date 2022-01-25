package web

import (
	"github.com/labstack/echo/v4"
	"github.com/lfdubiela/api-bank-transfer/pkg/adapter"
	"github.com/lfdubiela/api-bank-transfer/pkg/adapter/web/action"
)

const (
	HealthPath string = "/health"
)

type Router struct {
	echo *echo.Echo
	di   adapter.DependencyInjector
}

func NewRouter(echo *echo.Echo, di adapter.DependencyInjector) *Router {
	return &Router{
		echo: echo,
		di:   di,
	}
}

func (r *Router) RegisterHandlers() {
	//HEALTH
	r.echo.GET(HealthPath, action.GetHealth(r.di))
}
