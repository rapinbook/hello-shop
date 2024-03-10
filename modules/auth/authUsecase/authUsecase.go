package authUsecase

import "github.com/rapinbook/hello-shop/modules/auth/authRepository"

type (
	AuthUsecaseService interface{}
	authUsecase        struct {
		authRepository authRepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authRepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository}
}
