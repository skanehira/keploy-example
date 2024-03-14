package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	dsn := "root:root@tcp(0.0.0.0:3306)/keploy?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4,utf8"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect db: %s", err)
	}
}
