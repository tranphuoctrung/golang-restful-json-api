package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	. "golang-restful-json-api/v7/config"
	. "golang-restful-json-api/v7/models"
	. "golang-restful-json-api/v7/repo"

	"github.com/gorilla/mux"
)

var config = Config{}
var dbContext = DBContext{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dbContext.Server = config.Server
	dbContext.Database = config.Database
	dbContext.DbUser = config.DbUser
	dbContext.DbPassword = config.DbPassword

	dbContext.Connect()
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	todos, err := dbContext.FindAll()

	if err != nil {
		panic(err)
	}

	if err = json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	todo, err := dbContext.FindById(todoID)

	if err != nil {
		panic(err)
	}

	if todo.Id != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(JsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos

*/
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	err = dbContext.Insert(todo)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode("created"); err != nil {
		panic(err)
	}
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	todo, err := dbContext.FindById(todoID)

	if err != nil {
		panic(err)
	}

	if todo.Id != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(JsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
		return
	}

	err = dbContext.Delete(todo)

	if err != nil {
		panic(err)
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(JsonErr{Code: http.StatusOK, Text: "Todo deleted"}); err != nil {
		panic(err)
	}
}
