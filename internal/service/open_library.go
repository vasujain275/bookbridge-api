package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vasujain275/bookbridge-api/internal/types"
)

type OpenLibraryServiceImpl struct{}

// NewOpenLibraryService creates a new OpenLibrary service
func NewOpenLibraryService() OpenLibraryService {
	return &OpenLibraryServiceImpl{}
}

// GetByISBN gets a book by ISBN from OpenLibrary
func (s *OpenLibraryServiceImpl) GetByISBN(isbn string) (*types.OpenLibraryBook, error) {
	url := fmt.Sprintf("https://openlibrary.org/isbn/%s.json", isbn)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch book from OpenLibrary: %w", err)
	}
	defer resp.Body.Close()

	// Encode the response into the OpenLibraryBook struct
	var book types.OpenLibraryBook
	if err := json.NewDecoder(resp.Body).Decode(&book); err != nil {
		return nil, fmt.Errorf("failed to decode OpenLibrary response: %w", err)
	}

	return &book, nil
}
