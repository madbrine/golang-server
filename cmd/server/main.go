package main

import (
	"fmt"
	"servv/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}