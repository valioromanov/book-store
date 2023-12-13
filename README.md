# book-store

# Overview
The Book CRUD Service is a RESTful API that allows you to perform CRUD operations (Create, Read, Update, Delete) on a collection of books. Each book has properties such as title, author, book cover, and publishedAt.

# Run the application
In this repo there is a `docker-compose.yml` file which create a container in docker for the MySQL Database. Before running the application you should run `docker-compose up` from db directory(`cd db`). After running this command you will have one record in the database, then you need to set the following env vars:</br>
`- DB_HOST=...` </br>
`- DB_PORT=...` </br>
`- DB_USER=...` </br>
`- DB_PASSWORD=...` </br>
`- DB_NAME=...` </br>
There is a file `.env` where you can find the values of the env vars.
Then you navigate to cmd directory `cd cmd` and run `go run main.go`.

# Endpoints
- GET /book/{id}: Retrieve details of a specific book.
- POST /book: Create a new book.
- PUT /book/{id}: Update details of a specific book.
- DELETE /book/{id}: Delete a specific book.
# Get a single book
- endpoint: /book/{book_id}
- Request: GET
- Path param: book_id - the id of the book from the database.
- Response: `{
"title": "string",
"author": "string",
"bookCover": "bytes",
"publishedAt": "timestamp"
}`
- Example:
`curl --location --request GET 'http://localhost:8081/book/4'`
- Response: Returns found book
`{
    "id": 4,
    "title": "Harry Potter",
    "author": "J. K. Rowling",
    "bookCover": "/9j/4AAQSkZJRgABAQE...",
    "publishedAt": "2023-01-01"
}`
# Create a book
- endpoint: /book
- Request: POST
- Body: application/json </br>
`{
"title": "string",
"author": "string",
"bookCover": "bytes",
"publishedAt": "timestamp",
}`
- Response: Returns the id of the newly created book
`{
  "id": "string"
}`
-Example: 
`curl --location --request POST 'http://localhost:8081/book' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Harry Potter",
    "Author":"J. K. Rowling",
    "publishedAt": "2023-01-01",
    "bookCover": "/9j/4AAQSkZJRgABAQEAYABgAA..."
  }
'`

# Update a book
- endpoint: /book/{book_id}
- Request: PATCH
- Path param: book_id - the id of the book you want to update
- Body: application/json </br>
`{
"title": "string",
"author": "string",
"bookCover": "bytes",
"publishedAt": "timestamp",
}`
- Reponse: 204NoContent
- Example: `curl --location --request PATCH 'http://localhost:8081/book/4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Harry Potter and changing name"    
}
'`

# Delete a book
- endpoint: /book/{book_id}
- Request: DELETE
- Path param: book_id - the id of the book you want to delete
- Reponse: 204NoContent
- Example: `curl --location --request DELETE 'http://localhost:8081/book/3'`

FYI: In the repo you will find a postman collection you can test with.
