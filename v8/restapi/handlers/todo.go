package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	. "golang-restful-json-api/v8/config"
	. "golang-restful-json-api/v8/restapi/models"
	. "golang-restful-json-api/v8/restapi/repo"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
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

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	todo, err := dbContext.FindById(todoID)

	if err != nil {
		panic(err)
	}

	if todo.Id != "" {
		res := TodoResponse{
			Code:    http.StatusOK,
			Data:    todo,
			Message: "",
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(res.Code)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(JsonErr{Code: http.StatusNotFound, Message: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos

*/
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todoReq CreateTodoReq
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todoReq); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(JsonErr{Code: 400, Message: err.Error()}); err != nil {
			panic(err)
		}
	}
	todo := Todo{
		Id:        bson.NewObjectId(),
		Name:      todoReq.Name,
		Completed: todoReq.Completed,
	}
	err = dbContext.Insert(todo)

	res := TodoResponse{
		Code:    http.StatusOK,
		Data:    todo,
		Message: "",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(res.Code)
	if err := json.NewEncoder(w).Encode(res); err != nil {
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var errRes JsonErr

	if todo.Id != "" {
		errRes.Code = http.StatusNotFound
		errRes.Message = "Not Found"
		w.WriteHeader(errRes.Code)
		if err := json.NewEncoder(w).Encode(errRes); err != nil {
			panic(err)
		}
		return
	}

	err = dbContext.Delete(todo)
	if err != nil {

		errRes.Code = http.StatusInternalServerError
		errRes.Message = "StatusInternalServerError"
		w.WriteHeader(errRes.Code)
		if er := json.NewEncoder(w).Encode(errRes); er != nil {
			panic(er)
		}
		return
	}

	scsRes := struct {
		Code int
	}{
		Code: 200,
	}

	// If we didn't find it, 404
	w.WriteHeader(scsRes.Code)
	if err := json.NewEncoder(w).Encode(scsRes); err != nil {
		panic(err)
	}
}
