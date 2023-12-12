package repository

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type repository struct {
	db gorm.DB
}

func newRepository() *repository {
	repository := repository{}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"root", "root", "localhost", "3306", "books")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Error while connection to the database: ", err)
		panic(err)
	}
	logrus.Info("Successfull connection to the database")
	repository.db = *db
	return &repository
}
