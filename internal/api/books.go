package api

import (
	"context"

	"github.com/MilenaBMaciel/MyBookCollection/gen/openapi"
	"github.com/MilenaBMaciel/MyBookCollection/internal/domain/entity"
)

func (h *Handler) PostBook(ctx context.Context, req *openapi.PostBookReq) error{
	book := entity.Book{
		Title: req.Title,
		Author: req.Author,
		Language: req.Language,
		Category: req.Category,
		BoughtBy: req.BoughtBy,
		ReadBy: req.ReadBy,
	}
	err := h.booksHandler.PostBook(ctx, book)
	if err != nil{
		return err
	}

	return nil

}

func (h *Handler) GetBooks(ctx context.Context) ([]openapi.Book, error){
	books, err := h.booksHandler.GetBooks(ctx)
	if err != nil{
		return nil, err
	}

	resBooks := make([]openapi.Book, len(books))

	for i, book := range books {
		resBooks[i] = openapi.Book{
			ID: book.ID,
			Title: book.Title,
			Author: book.Author,
			Language: book.Language,
			Category: book.Category,
			BoughtBy: book.BoughtBy,
			ReadBy: book.ReadBy,
		}
	}
	return resBooks, nil
}

func (h *Handler) UpdateBook(ctx context.Context, req *openapi.UpdateBookReq, params openapi.UpdateBookParams) error{
	book := entity.Book{
		ID: params.ID,
		Title: req.Title,
		Author: req.Author,
		Language: req.Language,
		Category: req.Category,
		BoughtBy: req.BoughtBy,
		ReadBy: req.ReadBy,
	}
	return h.booksHandler.UpdateBook(ctx, book)
}

func (h *Handler) DeleteBook(ctx context.Context, params openapi.DeleteBookParams) error{
	return h.booksHandler.DeleteBook(ctx, entity.Book{ID: params.ID})
}