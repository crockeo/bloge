package database

import (
	"github.com/crockeo/bloge/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	gorm.DB
}

func OpenDB() DB {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic(err)
	}

	db.CreateTable(models.Post{})
	db.CreateTable(models.Auth{})

	return DB{db}
}
