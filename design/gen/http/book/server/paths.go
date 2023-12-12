// Code generated by goa v3.14.1, DO NOT EDIT.
//
// HTTP request path constructors for the book service.
//
// Command:
// $ goa gen book-store/design

package server

import (
	"fmt"
)

// GetBookBookPath returns the URL path to the book service getBook HTTP endpoint.
func GetBookBookPath(bookID int) string {
	return fmt.Sprintf("/book/%v", bookID)
}
