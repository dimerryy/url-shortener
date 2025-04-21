package main

import (
	"example/Go/projects/url-shortener/internal/config"
	"fmt"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)
}
