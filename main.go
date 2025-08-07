package main

import (
	"fmt"
	"os"

	"github.com/SirVoly/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg.SetUser("mino")

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(cfg.SPrint())
}
