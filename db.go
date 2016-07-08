package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func dbLinkSave(user, link string, timestamp time.Time) error {
	var err error
	db, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Println("Error while opening database: " + err.Error())
		return err
	}
	defer db.Close()

	query := "insert into links(user, link, date) values(?, ?, ?);"
	_, err = db.Exec(query, user, link, timestamp)
	return nil
}

func CheckDbTables() error {
	var err error
	query := "create table if not exists links(id int primary_key , user text, link text, date timestamp);"

	db, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	return nil
}
