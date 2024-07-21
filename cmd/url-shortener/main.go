package main

import (
	"cthulhu/internal/config"
	"fmt"
)

func main() {
	// init config/ cleanenv (умеет читать с разных расширений. struct tag)
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger/ slog

	// TODO: init storage/ sqlite

	// TODO: init router/ chi, "chi render"

	// TODO: run server
}
