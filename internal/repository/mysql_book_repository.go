package repository

import (
	"context"
	"errors"

	"github.com/qhmd/gitforgits/internal/domain/book"
	"gorm.io/gorm"
)

type mysqlBookRepository struct {
	db *gorm.DB
}

// CreateBook implements book.BookRepository.
func (m *mysqlBookRepository) CreateBook(ctx context.Context, book *book.Book) error {
	return m.db.WithContext(ctx).Create(book).Error
}

// DeleteBookByID implements book.BookRepository.
func (m *mysqlBookRepository) DeleteBookByID(ctx context.Context, id int) error {
	result := m.db.WithContext(ctx).Delete(&book.Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetBookByID implements book.BookRepository.
func (m *mysqlBookRepository) GetBookByID(ctx context.Context, id int) (*book.Book, error) {
	var b book.Book
	err := m.db.WithContext(ctx).First(&b, id).Error
	return &b, err
}

// ListBook implements book.BookRepository.
func (m *mysqlBookRepository) ListBook(ctx context.Context) ([]*book.Book, error) {
	var books []*book.Book
	err := m.db.WithContext(ctx).Find(&books).Error
	return books, err
}

// UpdateBook implements book.BookRepository.
func (m *mysqlBookRepository) UpdateBook(ctx context.Context, book *book.Book) error {
	return m.db.WithContext(ctx).Where("id = ?", book.ID).Updates(book).Error
}

func (m *mysqlBookRepository) GetBookByTitle(ctx context.Context, title string) (*book.Book, error) {
	var b book.Book
	if err := m.db.WithContext(ctx).Where("title = ?", title).First(&b).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &b, nil
}

func NewMySQLBookRepository(db *gorm.DB) book.BookRepository {
	return &mysqlBookRepository{db: db}
}
