package main

import (
	. "golang-restful-json-api/v7/routing"
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
