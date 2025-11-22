package postgres

import (
	"context"

	"github.com/MilenaBMaciel/MyBookCollection/db/sqlc"
	"github.com/MilenaBMaciel/MyBookCollection/internal/domain/entity"
	"github.com/MilenaBMaciel/MyBookCollection/internal/utils"
	"github.com/oklog/ulid/v2"
)

func (h *PostgresHandler) PostBook(ctx context.Context, book entity.Book) error {
	return h.queries.PostBook(ctx, sqlc.PostBookParams{
		ID:       ulid.Make().String(),
		Title:    book.Title,
		Author:   book.Author,
		Category: book.Category,
		Lang:     book.Language,
		BoughtBy: book.BoughtBy,
		ReadBy:   utils.ToNullString(book.ReadBy),
	})
}

func (h *PostgresHandler) GetBooks(ctx context.Context) ([]entity.Book, error) {
	dbBook, err := h.queries.GetBooks(ctx)
	if err != nil || len(dbBook) <= 0 {
		return nil, err
	}

	var books []entity.Book = make([]entity.Book, len(dbBook))

	for i, book := range dbBook {
		books[i] = entity.Book{
			ID:       book.ID,
			Title:    book.Title,
			Author:   book.Author,
			Language: book.Lang,
			Category: book.Category,
			BoughtBy: book.BoughtBy,
			ReadBy:   utils.FromNullStr(book.ReadBy),
		}
	}
	return books, nil
}

func (h *PostgresHandler) UpdateBook(ctx context.Context, book entity.Book) error{
	return h.queries.UpdateBook(ctx, sqlc.UpdateBookParams{
		ID:       book.ID,
		Title:    book.Title,
		Author:   book.Author,
		Category: book.Category,
		Lang:     book.Language,
		BoughtBy: book.BoughtBy,
		ReadBy:   utils.ToNullString(book.ReadBy),
	})
}

func (h *PostgresHandler) DeleteBook(ctx context.Context, book entity.Book) error{
	return h.queries.DeleteBook(ctx, book.ID)
}