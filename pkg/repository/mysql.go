package repository

import (
	"book-store/cmd/env"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type repository struct {
	db gorm.DB
}

func newRepository(config env.AppConfig) *repository {
	repository := repository{}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Error while connection to the database: ", err)
		panic(err)
	}
	logrus.Info("Successfull connection to the database")
	repository.db = *db
	return &repository
}
