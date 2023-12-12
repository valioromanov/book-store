//go:generate goa gen book-store/design

package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("book", func() {
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
			Field(1, "bookID", Int, "Book ID")
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
		Payload(BookReq)
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
			Body(BookReq)
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("patchBook", func() {
		Payload(BookPathcReq)
		Error("NotFound")
		Error("BadRequest")
		Result(Empty)

		HTTP(func() {
			PATCH("/book/{id}")
			Body(BookPathcReq)
			Response(StatusNoContent)
			Response("BadRequest", StatusBadRequest)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("deleteBook", func() {
		Payload(func() {
			Field(1, "bookID", Int, "Book ID")
			Required("bookID")
		})
		Error("NotFound")
		Error("BadRequest")
		Result(Empty)

		HTTP(func() {
			DELETE("/book/{bookID}")
			Response(StatusNoContent)
			Response("BadRequest", StatusBadRequest)
			Response("NotFound", StatusNotFound)
		})
	})

	// Serve the file with relative path ./gen/http/openapi.json for
	// requests sent to /swagger.json.
	Files("/swagger.json", "./gen/http/openapi.json")
})

var BookReq = Type("BookReq", func() {
	Description("A single book request")
	Attribute("id", Int)
	Attribute("title", String)
	Attribute("author", String)
	Attribute("bookCover", Bytes)
	Attribute("publishedAt", String)
	Required("title", "author")
})

var BookPathcReq = Type("BookPathcReq", func() {
	Description("A single book request for patching")
	Attribute("id", Int)
	Attribute("title", String)
	Attribute("author", String)
	Attribute("bookCover", Bytes)
	Attribute("publishedAt", String)
})

var Book = ResultType("application/vnd.goa.example.book", "BookResult", func() {
	Description("A single book response")

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

	View("resultOperation", func() {
		Attribute("id", Int)
		Required("id")
	})

})
