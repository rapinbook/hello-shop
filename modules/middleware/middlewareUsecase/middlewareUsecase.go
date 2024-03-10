package middlewareUsecase

import "github.com/rapinbook/hello-shop/modules/middleware/middlewareRepository"

type (
	MiddlewareUsecaseService interface{}
	middlewareUsecase        struct {
		middlewareRepository middlewareRepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewareRepository.MiddlewareRepositoryService) MiddlewareUsecaseService {
	return &middlewareUsecase{middlewareRepository: middlewareRepository}
}
