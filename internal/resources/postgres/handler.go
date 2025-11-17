package postgres

import (
	"github.com/MilenaBMaciel/MyBookCollection/db/sqlc"
	"github.com/MilenaBMaciel/MyBookCollection/internal/domain/gateway"
	"github.com/jmoiron/sqlx"
)

type PostgresHandler struct {
	db *sqlx.DB
	queries *sqlc.Queries
}

var _ gateway.Books = (*PostgresHandler)(nil)

func New(db *sqlx.DB) *PostgresHandler{
	return &PostgresHandler{db: db, queries: sqlc.New(db)}
}