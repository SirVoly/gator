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

	state := state{
		cfg: &cfg,
	}

	commands := generateCommands()

	if len(os.Args) < 2 {
		fmt.Println("no command was passed")
		os.Exit(1)
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = commands.run(&state, cmd)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type state struct {
	cfg *config.Config
}
