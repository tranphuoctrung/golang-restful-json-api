package main

import (
	. "golang-restful-json-api/v8/restapi/routing"
	. "golang-restful-json-api/v8/swagger"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// go-swagger examples.
//
// The purpose of this application is to provide some
// use cases describing how to generate docs for your API
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
func main() {
	router := NewRouter()
	registerV8Routes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func registerV8Routes(r *mux.Router) {
	RegisterTodoRoutes(r)
	RegisterDocRoutes(r)
}
