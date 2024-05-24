package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/SeaRoll/mini-auth/internal/api"
	"github.com/SeaRoll/mini-auth/internal/config"
	"github.com/SeaRoll/mini-auth/internal/db"
)
	
func main() {
	config := config.NewConfig("config.yml")
	dbo := db.NewService()
	defer dbo.Close()
	s := api.NewServer(config)

	stopCh := make(chan bool)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go s.Run(ctx, stopCh)

	<-ctx.Done()
	stopCh <- true

	log.Println("Shutting down...")

	time.Sleep(5 * time.Second)
}
