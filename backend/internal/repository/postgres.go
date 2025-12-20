package repository

import (
	"database/sql"

	"github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/database"
)

type Postgres struct {
	primary   *database.Database
	secondary *database.Database
}

func New(db *sql.DB) *Postgres {
	logger := database.StdLogger{}

	database := database.New(db, logger)

	return &Postgres{
		primary:   database,
		secondary: database,
	}
}
