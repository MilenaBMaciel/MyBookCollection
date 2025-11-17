package postgres

import (
	"context"

	"github.com/MilenaBMaciel/MyBookCollection/db/sqlc"
	"github.com/MilenaBMaciel/MyBookCollection/internal/domain/entity"
	"github.com/MilenaBMaciel/MyBookCollection/internal/utils"
	"github.com/oklog/ulid/v2"
)

func (h *PostgresHandler) PostBook(ctx context.Context, book entity.Book) error{
	return h.queries.PostBook(ctx, sqlc.PostBookParams{
		ID: ulid.Make().String(),
		Title: book.Title,
		Author: book.Author,
		Category: book.Category,
		Lang: book.Language,
		BoughtBy: book.Bought_by,
		ReadBy: utils.ToNullString(book.Read_by),
	})
}