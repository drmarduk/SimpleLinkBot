package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func RunHttpHandler() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+*flagPort, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	links := dbOpenLinks(0, 100)
	log.Println(links)

	io.WriteString(w, `<!DOCTYPE html>
	<html><head><title>Linky</title></head><body>
	<ul>`)
	for _, l := range links {
		io.WriteString(w, fmt.Sprintf(`<li><a href="%s">%s</a> - %s (%s)</li>`, l.Link, l.Link, l.User, l.Timestamp))
	}
	io.WriteString(w, "</ul")
}
