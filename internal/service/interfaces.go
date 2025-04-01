package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vasujain275/bookbridge-api/internal/repository"
)

// UserService defines the interface for user operations
type UserService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*repository.User, error)
	GetByUsername(ctx context.Context, username string) (*repository.User, error)
	GetByEmail(ctx context.Context, email string) (*repository.User, error)
	List(ctx context.Context, limit, offset int32) ([]*repository.User, error)
	Create(ctx context.Context, params repository.CreateUserParams) (*repository.User, error)
	Update(ctx context.Context, params repository.UpdateUserParams) (*repository.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// BookService defines the interface for book operations
type BookService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*repository.Book, error)
	GetByGoogleID(ctx context.Context, googleID string) (*repository.Book, error)
	List(ctx context.Context, limit, offset int32) ([]*repository.Book, error)
	Search(ctx context.Context, query string, limit, offset int32) ([]*repository.Book, error)
	Create(ctx context.Context, params repository.CreateBookParams) (*repository.Book, error)
	Update(ctx context.Context, params repository.UpdateBookParams) (*repository.Book, error)
	UpdateCopies(ctx context.Context, params repository.UpdateBookCopiesParams) (*repository.Book, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetFullBookDetails(ctx context.Context, id uuid.UUID) (*BookDetails, error)
}

// AuthorService defines the interface for author operations
type AuthorService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*repository.Author, error)
	GetByName(ctx context.Context, name string) (*repository.Author, error)
	List(ctx context.Context, limit, offset int32) ([]*repository.Author, error)
	Create(ctx context.Context, name string) (*repository.Author, error)
	Update(ctx context.Context, id uuid.UUID, name string) (*repository.Author, error)
	Delete(ctx context.Context, id uuid.UUID) error
	ListByBookID(ctx context.Context, bookID uuid.UUID) ([]*repository.Author, error)
	AddBookAuthor(ctx context.Context, bookID, authorID uuid.UUID) error
	RemoveBookAuthor(ctx context.Context, bookID, authorID uuid.UUID) error
	RemoveAllBookAuthors(ctx context.Context, bookID uuid.UUID) error
}

// CategoryService defines the interface for category operations
type CategoryService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*repository.Category, error)
	GetByName(ctx context.Context, name string) (*repository.Category, error)
	List(ctx context.Context, limit, offset int32) ([]*repository.Category, error)
	Create(ctx context.Context, name string) (*repository.Category, error)
	Update(ctx context.Context, id uuid.UUID, name string) (*repository.Category, error)
	Delete(ctx context.Context, id uuid.UUID) error
	ListByBookID(ctx context.Context, bookID uuid.UUID) ([]*repository.Category, error)
	AddBookCategory(ctx context.Context, bookID, categoryID uuid.UUID) error
	RemoveBookCategory(ctx context.Context, bookID, categoryID uuid.UUID) error
	RemoveAllBookCategories(ctx context.Context, bookID uuid.UUID) error
}

// LoanService defines the interface for loan operations
type LoanService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*repository.Loan, error)
	List(ctx context.Context, limit, offset int32) ([]*repository.Loan, error)
	ListByUserID(ctx context.Context, userID uuid.UUID, limit, offset int32) ([]*repository.Loan, error)
	ListByBookID(ctx context.Context, bookID uuid.UUID, limit, offset int32) ([]*repository.Loan, error)
	ListActive(ctx context.Context, limit, offset int32) ([]*repository.Loan, error)
	ListOverdue(ctx context.Context, limit, offset int32) ([]*repository.Loan, error)
	Create(ctx context.Context, params repository.CreateLoanParams) (*repository.Loan, error)
	Update(ctx context.Context, params repository.UpdateLoanParams) (*repository.Loan, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string, returnedDate *time.Time) (*repository.Loan, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// ReviewService defines the interface for review operations
type ReviewService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*repository.BookReview, error)
	GetByUserAndBook(ctx context.Context, userID, bookID uuid.UUID) (*repository.BookReview, error)
	List(ctx context.Context, limit, offset int32) ([]*repository.BookReview, error)
	ListByBookID(ctx context.Context, bookID uuid.UUID, limit, offset int32) ([]*repository.BookReview, error)
	ListByUserID(ctx context.Context, userID uuid.UUID, limit, offset int32) ([]*repository.BookReview, error)
	Create(ctx context.Context, params repository.CreateReviewParams) (*repository.BookReview, error)
	Update(ctx context.Context, id uuid.UUID, rating int32, reviewText *string) (*repository.BookReview, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// BookDetails contains all information about a book including its related entities
type BookDetails struct {
	Book       *repository.Book         `json:"book"`
	Authors    []*repository.Author     `json:"authors"`
	Categories []*repository.Category   `json:"categories"`
	Reviews    []*repository.BookReview `json:"reviews,omitempty"`
}
