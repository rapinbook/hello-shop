package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App      App
		Db       Db
		Kafka    Kafka
		Jwt      Jwt
		Grpc     Grpc
		Paginate Paginate
	}
	App struct {
		Stage string
		Name  string
		Url   string
	}
	Db struct {
		Url string
	}
	Jwt struct {
		AccessSecretKey  string
		AccessDuration   int64
		RefreshSecretKey string
		RefreshDuration  string
		ApiSecretKey     string
	}
	Kafka struct {
		Url       string
		ApiKey    string
		ApiSecret string
	}
	Grpc struct {
		AuthUrl      string
		ItemUrl      string
		PlayerUrl    string
		InventoryUrl string
		PaymentUrl   string
	}
	Paginate struct {
		ItemNextPageBasedUrl      string
		InventoryNextPageBasedUrl string
	}
)

func LoadConfig(path string) Config {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config{
		App: App{
			Stage: os.Getenv("APP_STAGE"),
			Name:  os.Getenv("APP_NAME"),
			Url:   os.Getenv("APP_URL"),
		},
		Db: Db{Url: os.Getenv("DB_URL")},
		Kafka: Kafka{
			Url:       os.Getenv("KAFKA_URL"),
			ApiKey:    os.Getenv("KAFKA_API_KEY"),
			ApiSecret: os.Getenv("KAFKA_API_SECRET"),
		},
		Grpc: Grpc{
			AuthUrl:      os.Getenv("GRPC_AUTH_URL"),
			ItemUrl:      os.Getenv("GRPC_ITEM_URL"),
			PlayerUrl:    os.Getenv("GRPC_PLAYER_URL"),
			InventoryUrl: os.Getenv("GRPC_INVENTORY_URL"),
			PaymentUrl:   os.Getenv("GRPC_PAYMENT_URL"),
		},
		Paginate: Paginate{
			ItemNextPageBasedUrl:      os.Getenv("PAGINATE_ITEM_NEXT_PAGE_BASED_URL"),
			InventoryNextPageBasedUrl: os.Getenv("PAGINATE_INVENTORY_NEXT_PAGE_BASED_URL"),
		},
	}
}
