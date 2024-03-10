package middlewareHandler

import (
	"github.com/rapinbook/hello-shop/config"
	"github.com/rapinbook/hello-shop/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface{}
	middlewareHandler        struct {
		ctg               *config.Config
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(ctg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{
		ctg:               ctg,
		middlewareUsecase: middlewareUsecase,
	}
}
