package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vasujain275/bookbridge-api/internal/repository"
	"github.com/vasujain275/bookbridge-api/internal/util"
)

type BookServiceImpl struct {
	repo               *repository.Queries
	openLibraryService OpenLibraryService
}

// NewBookService creates a new book service
func NewBookService(repo *repository.Queries, openLibraryService OpenLibraryService) BookService {
	return &BookServiceImpl{
		repo:               repo,
		openLibraryService: openLibraryService,
	}
}

// GetByID gets a book by ID
func (s *BookServiceImpl) GetByID(ctx context.Context, id uuid.UUID) (*repository.Book, error) {
	book, err := s.repo.GetBook(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get book: %w", err)
	}
	return &book, nil
}

// GetByISBN gets a book by ISBN
func (s *BookServiceImpl) GetByISBN(ctx context.Context, isbn string) (*repository.Book, error) {
	book, err := s.repo.GetBookByISBN(ctx, isbn)
	if err != nil {
		return nil, fmt.Errorf("failed to get book by ISBN: %w", err)
	}
	return &book, nil
}

// List gets a list of books
func (s *BookServiceImpl) List(ctx context.Context, limit, offset int32) ([]*repository.Book, error) {
	books, err := s.repo.ListBooks(ctx, repository.ListBooksParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list books: %w", err)
	}
	bookPtrs := make([]*repository.Book, len(books))
	for i := range books {
		bookPtrs[i] = &books[i]
	}

	return bookPtrs, nil
}

// Create creates a new book
func (s *BookServiceImpl) Create(ctx context.Context, isbn string) (*repository.Book, error) {
	fetchedBook, err := s.openLibraryService.GetByISBN(isbn)

	if err != nil {
		fmt.Println("Error fetching book from OpenLibrary:", err)
	}

	book, err := s.repo.CreateBook(ctx, repository.CreateBookParams{
		Isbn10:        util.StringToPgText(fetchedBook.ISBN10[0]),
		Isbn13:        fetchedBook.ISBN13[0],
		Title:         fetchedBook.Title,
		Publisher:     util.StringToPgText(fetchedBook.Publishers[0]),
		PublishedDate: util.StringToPgText(fetchedBook.PublishDate),
		Description:   util.StringToPgText(fetchedBook.Bio),
		PageCount:     util.Int32ToPgInt(int32(fetchedBook.NumberOfPages)),
		Language:      util.StringToPgText(fetchedBook.Languages[0].Key),
		ThumbnailUrl:  util.StringToPgText(fmt.Sprintf("https://covers.openlibrary.org/b/id/%d-L.jpg", fetchedBook.Covers[0])),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create book: %w", err)
	}

	return &book, nil

}
