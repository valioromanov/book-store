package repository

import (
	"book-store/cmd/env"
	"fmt"

	"github.com/sirupsen/logrus"
)

type BookRepo struct {
	*repository
}

func NewBookRepo(config env.AppConfig) *BookRepo {
	repo := newRepository(config)
	return &BookRepo{
		repo,
	}
}

func (br *BookRepo) FindById(id int) (BookDB, error) {
	book := BookDB{}
	tx := br.db.Table("books").Select("*").Where("id = ?", id).Find(&book)

	if tx.Error != nil {
		logrus.Error(fmt.Sprintf("error while fetching book by id '%d' from database: %s ", id, tx.Error.Error()))
		return BookDB{}, fmt.Errorf("error while fetching from database: %s", tx.Error.Error())
	}

	if tx.RowsAffected == 0 {
		logrus.Error("No data found for book with id: ", id)
		return BookDB{}, fmt.Errorf("no data found")
	}

	return book, nil
}

func (br *BookRepo) InsertBook(book BookDB) (int, error) {
	result := br.db.Create(&book)

	if result.Error != nil {
		logrus.Error(fmt.Sprintf("error while inserting book: %s", result.Error.Error()))
		return 0, fmt.Errorf("error while inserting book: %w", result.Error)
	}

	return book.ID, nil
}

func (br *BookRepo) UpdateBook(book BookDB) error {
	if book.ID == 0 {
		logrus.Error("error while updating book: no book id provided")
		return fmt.Errorf("no book id provided")
	}

	tx := br.db.Save(&book)

	if tx.Error != nil {
		logrus.Error(fmt.Sprintf("error while updating book: %s", tx.Error.Error()))
		return fmt.Errorf("error while updating a record: %w", tx.Error)
	}

	if tx.RowsAffected == 0 {
		return fmt.Errorf("no data found")
	}
	return nil
}

func (br *BookRepo) DeleteBook(bookID int) error {
	if bookID == 0 {
		logrus.Error("error while deleting book: no book id provided")
		return fmt.Errorf("no book id provided")
	}

	bookForDeletion := BookDB{
		ID: bookID,
	}

	tx := br.db.Delete(&bookForDeletion)
	if tx.Error != nil {
		logrus.Error(fmt.Sprintf("error while deleting book: %s", tx.Error.Error()))
		return fmt.Errorf("error while deleting a record: %w", tx.Error)
	}

	if tx.RowsAffected == 0 {
		return fmt.Errorf("no rows deleted")
	}

	return nil
}
