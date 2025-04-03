package types

// OpenLibraryBook represents the data structure returned by the Open Library API
type OpenLibraryBook struct {
	Identifiers struct {
		Goodreads    []string `json:"goodreads,omitempty"`
		Librarything []string `json:"librarything,omitempty"`
	} `json:"identifiers,omitempty"`
	Title   string `json:"title,omitempty"`
	Authors []struct {
		Key string `json:"key,omitempty"`
	} `json:"authors,omitempty"`
	PublishDate   string   `json:"publish_date,omitempty"`
	Bio           string   `json:"bio,omitempty"`
	Publishers    []string `json:"publishers,omitempty"`
	Covers        []int    `json:"covers,omitempty"`
	Contributions []string `json:"contributions,omitempty"`
	Languages     []struct {
		Key string `json:"key,omitempty"`
	} `json:"languages,omitempty"`
	SourceRecords []string `json:"source_records,omitempty"`
	LocalID       []string `json:"local_id,omitempty"`
	Type          struct {
		Key string `json:"key,omitempty"`
	} `json:"type,omitempty"`
	FirstSentence struct {
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"first_sentence,omitempty"`
	Key           string `json:"key,omitempty"`
	NumberOfPages int    `json:"number_of_pages,omitempty"`
	Works         []struct {
		Key string `json:"key,omitempty"`
	} `json:"works,omitempty"`
	Classifications map[string]interface{} `json:"classifications,omitempty"`
	OCAID           string                 `json:"ocaid,omitempty"`
	ISBN10          []string               `json:"isbn_10,omitempty"`
	ISBN13          []string               `json:"isbn_13,omitempty"`
	LatestRevision  int                    `json:"latest_revision,omitempty"`
	Revision        int                    `json:"revision,omitempty"`
	Created         struct {
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"created,omitempty"`
	LastModified struct {
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"last_modified,omitempty"`
}
