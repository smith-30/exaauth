package rdb

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ORM interface {
	Conn() *gorm.DB
}

type rdb struct {
	orm *gorm.DB
}

var RDB ORM

func InitRDB(options ...func(*rdb) error) error {
	a := &rdb{}
	for _, option := range options {
		if err := option(a); err != nil {
			return err
		}
	}

	if a.orm == nil {
		return errors.New("failure connecting db")
	}

	RDB = a

	return nil
}

func Mysql(username, password, host, port, dbName string) func(*rdb) error {
	return func(a *rdb) error {
		address := host + ":" + port
		setting := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", username, password, address, dbName)
		db, err := gorm.Open("mysql", setting)
		a.orm = db
		return err
	}
}

func Postgres(username, password, host, port, dbName string) func(*rdb) error {
	return func(a *rdb) error {
		setting := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, username, dbName, password)
		db, err := gorm.Open(setting)
		a.orm = db
		return err
	}
}

func Sqlite(username, password, host, port, dbName string) func(*rdb) error {
	return func(a *rdb) error {
		db, err := gorm.Open("sqlite3", "./gorm.db")
		a.orm = db
		return err
	}
}

func (a *rdb) Conn() *gorm.DB {
	return a.orm.New()
}
