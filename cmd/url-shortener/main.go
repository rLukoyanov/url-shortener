package main

import (
	"fmt"

	"example.com/url-shorterner/internal/config"
)

func main() {
	// init config - cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// init logger - slog
	// init storage - sqlite
	// init router - chi
	// run server
}
