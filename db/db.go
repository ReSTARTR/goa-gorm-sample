package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func init() {
	config := &mysql.Config{
		User:   "root",
		Addr:   "127.0.0.1",
		DBName: "goa_sample",
		ParseTime: true, // ref: https://github.com/Go-SQL-Driver/MySQL/issues/9
	}
	var err error
	DB, err = gorm.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
	DB.LogMode(true)
}
