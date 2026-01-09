package cli

import (
	"database/sql"
	"log"
	"os"

	"github.com/SirVoly/gator/internal/config"
	"github.com/SirVoly/gator/internal/database"
	_ "github.com/lib/pq"
)

func Run(){

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	state := state{
		cfg: &cfg,
		db:  database.New(db),
	}

	commands := generateCommands()

	if len(os.Args) < 2 {
		log.Fatal("Usage: gator <command> [args...]")
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = commands.run(&state, cmd)

	if err != nil {
		log.Fatal(err)
	}
}
