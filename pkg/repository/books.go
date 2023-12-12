package repository

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type BookRepo struct {
	*repository
}

func NewBookRepo() *BookRepo {
	repo := newRepository()
	return &BookRepo{
		repo,
	}
}

func (br *BookRepo) FindById(id int) (BookDB, error) {
	book := BookDB{}
	tx := br.db.Table("books").Select("*").Where("id = ?", id).Find(&book)

	if tx.Error != nil {
		logrus.Error(fmt.Sprintf("error while fetching book by id '%s' from database: %s ", id, tx.Error.Error()))
		return BookDB{}, fmt.Errorf("error while fetching from database: %w", tx.Error)
	}

	if tx.RowsAffected == 0 {
		logrus.Error("No data found for book with id: ", id)
		return BookDB{}, fmt.Errorf("no rows affected")
	}

	return book, nil
}
