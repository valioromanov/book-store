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
	Method("getBook", func() {
		// Payload describes the method payload.
		// Here the payload is an object that consists of two fields.
		Payload(func() {
			// Field describes an object field given a field index, a field
			// name, a type and a description.
			Field(1, "bookID", Int, "Book ID")
			// Required list the names of fields that are required.
			Required("bookID")
		})

		Error("NotFound")
		Error("BadRequest")

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

			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("postBook", func() {
		// Payload describes the method payload.
		// Here the payload is an object that consists of two fields.
		Payload(Book, func() {
			View("inserting")
		})

		Error("NotFound")
		Error("BadRequest")

		// Result describes the method result.
		// Here the result is a simple integer value.
		Result(Book, func() {
			View("resultOperation")
		})

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests.
			// The payload fields are encoded as path parameters.
			POST("/book")

			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
		})
	})

	// Serve the file with relative path ./gen/http/openapi.json for
	// requests sent to /swagger.json.
	Files("/swagger.json", "./gen/http/openapi.json")
})

var Book = ResultType("application/vnd.goa.example.book", "BookResult", func() {
	Description("A bottle of wine")

	Attributes(func() {
		Attribute("id", Int, "ID of the book")
		Attribute("title", String, " of the book")
		Attribute("author", String, "author of the book")
		Attribute("bookCover", Bytes, "cover of the book")
		Attribute("publishedAt", String, "cover of the book")
	})

	View("default", func() { // Explicitly define default view
		Attribute("id", Int)
		Attribute("title", String)
		Attribute("author", String)
		Attribute("bookCover", Bytes)
		Attribute("publishedAt", String)
	})

	View("inserting", func() {
		Attribute("title", String)
		Attribute("author", String)
		Attribute("bookCover", Bytes)
		Attribute("publishedAt", String)
		Required("title", "author")
	})

	View("resultOperation", func() {
		Attribute("id", Int)
		Required("id")
	})

})
