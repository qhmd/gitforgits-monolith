package book

import "context"

type BookRepository interface {
	GetBookByID(ctx context.Context, id int) (*Book, error)
	ListBook(ctx context.Context) ([]*Book, error)
	CreateBook(ctx context.Context, book *Book) error
	UpdateBook(ctx context.Context, book *Book) error
	DeleteBookByID(ctx context.Context, id int) error

	GetBookByTitle(ctx context.Context, title string) (*Book, error)
}
