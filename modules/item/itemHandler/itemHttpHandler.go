package itemHandler

import (
	"github.com/rapinbook/hello-shop/config"
	itemUsecase "github.com/rapinbook/hello-shop/modules/Item/ItemUsecase"
)

type (
	ItemHttpHandlerService interface{}
	itemHttpHandler        struct {
		ctg         *config.Config
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(ctg *config.Config, itemUsecase itemUsecase.ItemUsecaseService) ItemHttpHandlerService {
	return &itemHttpHandler{
		ctg:         ctg,
		itemUsecase: itemUsecase,
	}
}
