package main

import (
	. "golang-restful-json-api/v5/config"
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
