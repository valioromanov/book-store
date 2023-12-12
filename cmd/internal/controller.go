package internal

import (
	"book-store/design/gen/book"
	"book-store/pkg/repository"
	"context"
	"fmt"
	"time"
)

type BookRepository interface {
	FindById(id int) (repository.BookDB, error)
	InsertBook(repository.BookDB) (int, error)
	UpdateBook(repository.BookDB) error
}

type Controller struct {
	repo BookRepository
}

func NewController(repository BookRepository) *Controller {
	return &Controller{
		repo: repository,
	}
}

func (c *Controller) GetBook(context context.Context, payload *book.GetBookPayload) (*book.BookResult, error) {
	if payload == nil {
		return nil, book.MakeBadRequest(fmt.Errorf("no id for book provided"))
	}

	bookResult, err := c.repo.FindById(payload.BookID)
	if err != nil {
		return nil, book.MakeNotFound(fmt.Errorf("error while fetching book: %w", err))
	}
	timeString := bookResult.PublishedAt.Format("2006-01-02")
	res := book.BookResult{
		ID:          &bookResult.ID,
		Title:       &bookResult.Title,
		Author:      &bookResult.Author,
		BookCover:   bookResult.BookCover,
		PublishedAt: &timeString,
	}
	return &res, nil
}

func (c *Controller) PostBook(context context.Context, request *book.BookReq) (*book.BookResult, error) {
	if request.Title == "" || request.Author == "" {
		return nil, book.MakeBadRequest(fmt.Errorf("requeired property not provided"))
	}

	bookDB := repository.BookDB{
		Title:     request.Title,
		Author:    request.Author,
		BookCover: request.BookCover,
	}

	if request.PublishedAt != nil {
		publishedAt, err := time.Parse("2006-01-02", *request.PublishedAt)
		if err != nil {
			return nil, book.MakeBadRequest(fmt.Errorf("wrong time provided"))
		}
		bookDB.PublishedAt = publishedAt
	}
	id, err := c.repo.InsertBook(bookDB)

	if err != nil {
		return nil, fmt.Errorf("error while inserting in db: %w", err)
	}
	res := book.BookResult{
		ID: &id,
	}
	return &res, nil
}

func (c *Controller) PatchBook(context context.Context, request *book.BookPathcReq) error {
	fmt.Println("request: ", request)
	bookDB, err := c.repo.FindById(*request.ID)
	if err != nil {
		return book.MakeNotFound(fmt.Errorf("error while fetching book: %w", err))
	}

	if request.Author != nil {
		bookDB.Author = *request.Author
	}
	if request.Title != nil {
		bookDB.Title = *request.Title
	}
	if len(request.BookCover) > 0 {
		bookDB.BookCover = request.BookCover
	}
	if request.PublishedAt != nil {
		publishedAt, err := time.Parse("2006-01-02", *request.PublishedAt)
		if err != nil {
			return book.MakeBadRequest(fmt.Errorf("wrong time provided"))
		}
		bookDB.PublishedAt = publishedAt
	}
	return c.repo.UpdateBook(bookDB)
}
