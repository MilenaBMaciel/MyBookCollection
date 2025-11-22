package gateway

import (
	"context"

	"github.com/MilenaBMaciel/MyBookCollection/internal/domain/entity"
)

type Books interface {
	PostBook(context.Context, entity.Book) error
	GetBooks(context.Context) ([]entity.Book, error)
	UpdateBook(context.Context, entity.Book) error
	DeleteBook(context.Context, entity.Book) error
}