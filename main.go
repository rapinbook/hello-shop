package main

import (
	"context"
	"log"
	"os"

	"github.com/rapinbook/hello-shop/config"
)

func main() {
	ctx := context.Background()
	_ = ctx
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Required an input args")
		}
		return os.Args[1]
	}())
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	log.Println(cfg)
}
