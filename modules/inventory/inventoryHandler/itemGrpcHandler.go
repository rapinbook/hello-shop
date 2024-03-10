package inventoryHandler

import (
	"github.com/rapinbook/hello-shop/config"
	inventoryUsecase "github.com/rapinbook/hello-shop/modules/inventory/inventoryUsecase"
)

type (
	inventoryGrpcHandler struct {
		ctg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewinventoryGrpcHandler(ctg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{ctg: ctg, inventoryUsecase: inventoryUsecase}
}
