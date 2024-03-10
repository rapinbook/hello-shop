package playerHandler

import (
	"github.com/rapinbook/hello-shop/config"
	"github.com/rapinbook/hello-shop/modules/player/playerUsecase"
)

type (
	PlayerQueueHandlerService interface{}
	playerQueueHandler        struct {
		ctg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewAuthQueueHandler(ctg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerQueueHandlerService {
	return &playerQueueHandler{
		ctg:           ctg,
		playerUsecase: playerUsecase,
	}
}
