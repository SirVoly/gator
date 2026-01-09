package cli

import (
	"github.com/SirVoly/gator/internal/config"
	"github.com/SirVoly/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}