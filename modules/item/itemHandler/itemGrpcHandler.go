package itemHandler

import (
	"github.com/rapinbook/hello-shop/config"
	itemUsecase "github.com/rapinbook/hello-shop/modules/Item/ItemUsecase"
)

type (
	itemGrpcHandler struct {
		ctg         *config.Config
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemGrpcHandler(ctg *config.Config, itemUsecase itemUsecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{ctg: ctg, itemUsecase: itemUsecase}
}
