package gateway

import (
	"context"

	"github.com/MilenaBMaciel/MyBookCollection/internal/domain/entity"
)

type Books interface {
	PostBook(context.Context, entity.Book) error
}