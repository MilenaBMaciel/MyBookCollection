package api

import "github.com/MilenaBMaciel/MyBookCollection/internal/domain/gateway"

type Handler struct {
	booksHandler gateway.Books
}

func NewHandler(books gateway.Books) *Handler{
	return &Handler{booksHandler: books}
}