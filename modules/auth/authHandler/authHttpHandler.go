package authHandler

import (
	"github.com/rapinbook/hello-shop/config"
	"github.com/rapinbook/hello-shop/modules/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface{}
	authHttpHandler        struct {
		ctg         *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(ctg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{
		ctg:         ctg,
		authUsecase: authUsecase,
	}
}
