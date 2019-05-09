package rdb

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func InitFakeRDB() {
	RDB = &FakeORM{}
}

type FakeORM struct{}

func (a *FakeORM) Conn() *gorm.DB {
	_, _, err := sqlmock.NewWithDSN("sqlmock_db_0")
	if err != nil {
		panic("Got an unexpected error.")
	}
	db, err := gorm.Open("sqlmock", "sqlmock_db_0")
	if err != nil {
		panic("Got an unexpected error.")
	}
	return db
}
