package repository

import "time"

type BookDB struct {
	ID          int       `gorm:"primarykey column:id"`
	Title       string    `gorm:"column:title"`
	Author      string    `gorm:"column:author"`
	BookCover   []byte    `gorm:"colums:book_cover"`
	PublishedAt time.Time `gorm:"colums:published_at,omitempty"`
}

func (bdb BookDB) TableName() string {
	return "books"
}
