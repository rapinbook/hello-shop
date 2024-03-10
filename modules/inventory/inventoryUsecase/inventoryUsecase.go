package inventoryUsecase

import (
	"github.com/rapinbook/hello-shop/modules/inventory/inventoryRepository"
)

type (
	InventoryUsecaseService interface{}
	inventoryUsecase        struct {
		inventoryRepository inventoryRepository.InventoryRepositoryService
	}
)

func NewinventoryUsecase(inventoryRepository inventoryRepository.InventoryRepositoryService) InventoryUsecaseService {
	return &inventoryUsecase{inventoryRepository: inventoryRepository}
}
