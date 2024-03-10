package paymentHandler

import (
	"github.com/rapinbook/hello-shop/config"
	"github.com/rapinbook/hello-shop/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface{}
	paymentHttpHandler        struct {
		ctg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewAuthHttpHandler(ctg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{
		ctg:            ctg,
		paymentUsecase: paymentUsecase,
	}
}
