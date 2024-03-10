package authHandler

import (
	"github.com/rapinbook/hello-shop/config"
	"github.com/rapinbook/hello-shop/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		ctg         *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(ctg *config.Config, authUsecase authUsecase.AuthUsecaseService) authUsecase.AuthUsecaseService {
	return &authGrpcHandler{ctg: ctg, authUsecase: authUsecase}
}
