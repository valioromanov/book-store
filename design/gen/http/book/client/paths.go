// Code generated by goa v3.14.1, DO NOT EDIT.
//
// HTTP request path constructors for the book service.
//
// Command:
// $ goa gen book-store/design

package client

import (
	"fmt"
)

// GetBookBookPath returns the URL path to the book service getBook HTTP endpoint.
func GetBookBookPath(bookID int) string {
	return fmt.Sprintf("/book/%v", bookID)
}

// PostBookBookPath returns the URL path to the book service postBook HTTP endpoint.
func PostBookBookPath() string {
	return "/book"
}

// PatchBookBookPath returns the URL path to the book service patchBook HTTP endpoint.
func PatchBookBookPath(id int) string {
	return fmt.Sprintf("/book/%v", id)
}

// DeleteBookBookPath returns the URL path to the book service deleteBook HTTP endpoint.
func DeleteBookBookPath(bookID int) string {
	return fmt.Sprintf("/book/%v", bookID)
}
