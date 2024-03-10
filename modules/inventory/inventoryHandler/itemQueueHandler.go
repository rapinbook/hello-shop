package inventoryHandler

import (
	"github.com/rapinbook/hello-shop/config"
	"github.com/rapinbook/hello-shop/modules/inventory/inventoryUsecase"
)

type (
	InventoryQueueHandlerService interface{}
	inventoryQueueHandler        struct {
		ctg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewinventoryQueueHandler(ctg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryQueueHandlerService {
	return &inventoryQueueHandler{
		ctg:              ctg,
		inventoryUsecase: inventoryUsecase,
	}
}
