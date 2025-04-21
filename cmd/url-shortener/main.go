package main

import (
	"fmt"
	"log/slog"
	"os"
	"url-shortener/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)

	log.Info("starting url-shortener")
}
