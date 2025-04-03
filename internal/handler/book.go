package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vasujain275/bookbridge-api/internal/service"
	"github.com/vasujain275/bookbridge-api/internal/util"
)

// BookHandler handles HTTP requests for books.
type BookHandler struct {
	service service.BookService
}

// NewBookHandler creates a new BookHandler.
func NewBookHandler(s service.BookService) *BookHandler {
	return &BookHandler{
		service: s,
	}
}

// CreateBookRequest represents the expected request payload for creating a book.
type CreateBookRequest struct {
	ISBN string `json:"isbn" binding:"required"`
}

// GetBook godoc
// @Summary Get book by ID
// @Description Get a book by its ID.
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} util.Response "Book found"
// @Failure 400 {object} util.Response "Invalid ID supplied"
// @Failure 404 {object} util.Response "Book not found"
// @Router /books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		util.SendBadRequest(c, "Invalid book ID", err.Error())
		return
	}
	book, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		util.SendNotFound(c, err.Error())
		return
	}
	util.SendOK(c, "Book found", book)
}

// GetBookByISBN godoc
// @Summary Get book by ISBN
// @Description Get a book by its ISBN.
// @Tags books
// @Accept json
// @Produce json
// @Param isbn path string true "Book ISBN"
// @Success 200 {object} util.Response "Book found"
// @Failure 400 {object} util.Response "Invalid ISBN supplied"
// @Failure 404 {object} util.Response "Book not found"
// @Router /books/isbn/{isbn} [get]
func (h *BookHandler) GetBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")
	if isbn == "" {
		util.SendBadRequest(c, "Invalid ISBN", "ISBN cannot be empty")
		return
	}
	book, err := h.service.GetByISBN(c.Request.Context(), isbn)
	if err != nil {
		util.SendNotFound(c, err.Error())
		return
	}
	util.SendOK(c, "Book found", book)
}

// ListBooks godoc
// @Summary List books
// @Description Get a list of books with pagination.
// @Tags books
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} util.Response "Books retrieved successfully"
// @Failure 400 {object} util.Response "Invalid request"
// @Failure 500 {object} util.Response "Internal server error"
// @Router /books [get]
func (h *BookHandler) ListBooks(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		util.SendBadRequest(c, "Invalid limit parameter", err.Error())
		return
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 32)
	if err != nil {
		util.SendBadRequest(c, "Invalid offset parameter", err.Error())
		return
	}

	books, err := h.service.List(c.Request.Context(), int32(limit), int32(offset))
	if err != nil {
		util.SendInternalServerError(c, err.Error())
		return
	}

	util.SendOK(c, "Books retrieved successfully", books)
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book by fetching details from OpenLibrary using ISBN
// @Tags books
// @Accept json
// @Produce json
// @Param book body CreateBookRequest true "Book ISBN"
// @Success 201 {object} util.Response "Book created successfully"
// @Failure 400 {object} util.Response "Invalid request"
// @Failure 500 {object} util.Response "Internal server error"
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendBadRequest(c, "Invalid request payload", err.Error())
		return
	}

	book, err := h.service.Create(c.Request.Context(), req.ISBN)
	if err != nil {
		util.SendInternalServerError(c, "Failed to create book: "+err.Error())
		return
	}

	util.SendCreated(c, "Book created successfully", book)
}
