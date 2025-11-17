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
		Bought_by: req.BoughtBy,
		Read_by: req.ReadBy,
		ID: "1",
	}
	err := h.booksHandler.PostBook(ctx, book)
	if err != nil{
		return err
	}

	return nil

}