package playerHandler

import (
	"github.com/rapinbook/hello-shop/config"
	"github.com/rapinbook/hello-shop/modules/player/playerUsecase"
)

type (
	PlayerHttpHandlerService interface{}
	playerHttpHandler        struct {
		ctg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewAuthHttpHandler(ctg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerHttpHandlerService {
	return &playerHttpHandler{
		ctg:           ctg,
		playerUsecase: playerUsecase,
	}
}
