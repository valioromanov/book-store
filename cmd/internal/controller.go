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
	fmt.Println("ID: ", payload.BookID)
	bookResult, err := c.repo.FindById(payload.BookID)
	if err != nil {
		return nil, fmt.Errorf("error while fetching book: ", err)
	}
	timeString := bookResult.PublishedAt.Format("2006-01-02 15:04:05")
	res := book.BookResult{
		ID:          &bookResult.ID,
		Title:       &bookResult.Title,
		Author:      &bookResult.Author,
		PublishedAt: &timeString,
	}
	return &res, nil
}
