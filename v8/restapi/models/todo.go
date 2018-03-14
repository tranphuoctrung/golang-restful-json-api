package models

import "gopkg.in/mgo.v2/bson"

// swagger:model todo
type Todo struct {
	// the id for this todo
	//
	// required: true
	Id bson.ObjectId `bson:"_id" json:"id"`

	// the name for this todo
	// required: true
	// min length: 3
	Name string `json:"name"`

	// the status of this todo
	//
	// required: true
	Completed bool `json:"completed"`
}

// swagger:model todoList
type TodoList []Todo

// CreateTodoReq contains request data for create todo API
type CreateTodoReq struct {
	// Name of the repository
	Name string `json:"name"`
	// Completed defines whether created todo had been completed or not
	Completed bool `json:"completed"`
}

type TodoResponse struct {
	Code    int    `json:"code"`
	Data    Todo   `json:"data"`
	Message string `json:"msg,omitempty"`
}

type TodosResponse struct {
	Code    int      `json:"code"`
	Data    TodoList `json:"data"`
	Message string   `json:"msg,omitempty"`
}
