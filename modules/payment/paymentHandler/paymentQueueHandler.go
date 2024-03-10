package paymentHandler

import (
	"github.com/rapinbook/hello-shop/config"
	"github.com/rapinbook/hello-shop/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface{}
	paymentQueueHandler        struct {
		ctg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewAuthQueueHandler(ctg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{
		ctg:            ctg,
		paymentUsecase: paymentUsecase,
	}
}
