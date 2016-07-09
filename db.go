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

func dbOpenLinks(offset, count int) []Message {
	var err error
	var result []Message
	query := "select id, user, link, date from links limit ?, ?;"

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Println("Error while selecting links: " + err.Error())
		return nil
	}
	defer db.Close()

	rows, err := db.Query(query, offset, count)
	if err != nil {
		log.Println("Error while query: " + err.Error())
		return nil
	}

	for rows.Next() {
		tmp := Message{}
		err = rows.Scan(&tmp.Id, &tmp.User, &tmp.Link, &tmp.Timestamp)
		if err != nil {
			log.Println("dbOpenLinks: Scanning error: " + err.Error())
			continue
		}

		result = append(result, tmp)
	}

	return result
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
