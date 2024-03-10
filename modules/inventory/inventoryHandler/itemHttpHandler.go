package inventoryHandler

import (
	"github.com/rapinbook/hello-shop/config"
	"github.com/rapinbook/hello-shop/modules/inventory/inventoryUsecase"
)

type (
	InventoryHttpHandlerService interface{}
	inventoryHttpHandler        struct {
		ctg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewinventoryHttpHandler(ctg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryHttpHandlerService {
	return &inventoryHttpHandler{
		ctg:              ctg,
		inventoryUsecase: inventoryUsecase,
	}
}
