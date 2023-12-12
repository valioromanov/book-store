package internal

import (
	"book-store/design/gen/book"
	"book-store/pkg/repository"
	"context"
	"fmt"
)

type BookRepository interface {
	FindById(id int) (repository.BookDB, error)
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
	timeString := bookResult.PublishedAt.Format("2006-01-02 15:04:05")
	res := book.BookResult{
		ID:          &bookResult.ID,
		Title:       &bookResult.Title,
		Author:      &bookResult.Author,
		BookCover:   bookResult.BookCover,
		PublishedAt: &timeString,
	}
	return &res, nil
}

func (c *Controller) PostBook(context.Context, *book.BookResult) (res *book.BookResult, err error) {
	return nil, nil
}
