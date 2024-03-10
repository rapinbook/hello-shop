package main

import (
	"context"
	"log"
	"os"

	"github.com/rapinbook/hello-shop/config"
	database "github.com/rapinbook/hello-shop/pkg/database/script"
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
	db := database.DbConn(ctx, &cfg)

	defer db.Disconnect(ctx)
	log.Println(db)
}
