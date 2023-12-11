//go:generate goa gen book-store/design

package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("example", func() {
	Title("Book Service")
	Description("Service that provides CRUD operations for a book")

	Server("book", func() {
		Description("book hosts the Book Service.")

		// List the services hosted by this server.
		Services("book")

		// List the Hosts and their transport URLs.
		Host("development", func() {
			Description("Development hosts.")
			// Transport specific URLs, supported schemes are:
			// 'http', 'https', 'grpc' and 'grpcs' with the respective default
			// ports: 80, 443, 8080, 8443.
			URI("http://localhost:8000/book")
		})
	})
})

var _ = Service("book", func() {
	Description("The book service performs operations for books")

	// Method describes a service method (endpoint)
	Method("get", func() {
		// Payload describes the method payload.
		// Here the payload is an object that consists of two fields.
		Payload(func() {
			// Field describes an object field given a field index, a field
			// name, a type and a description.
			Field(1, "bookID", Int, "Book ID")
			// Required list the names of fields that are required.
			Required("bookID")
		})

		// Result describes the method result.
		// Here the result is a simple integer value.
		Result(Book, func() {
			View("default")
		})

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests.
			// The payload fields are encoded as path parameters.
			GET("/book/{bookID}")
			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body (default).
			Response(StatusOK)
		})
	})

	// Serve the file with relative path ./gen/http/openapi.json for
	// requests sent to /swagger.json.
	Files("/swagger.json", "./gen/http/openapi.json")
})

var Book = ResultType("application/vnd.goa.example.bottle", "BottleResult", func() {
	Description("A bottle of wine")

	Attributes(func() {
		Attribute("id", Int, "ID of bottle")
		Attribute("book", String, "name of book")
		Required("id", "book")
	})

	View("default", func() { // Explicitly define default view
		Attribute("id")
		Attribute("book")
	})
})
