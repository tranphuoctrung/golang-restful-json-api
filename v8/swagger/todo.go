package swagger

import (
	model "golang-restful-json-api/v8/restapi/models"
)

// A UserNameParams parameter model.
//
// This is used for operations that want the ID of a todo in the path
// swagger:parameters getTodoById todoDelete
type TodoIdParams struct {
	// The id of todo
	//
	// in: path
	// required: true
	TodoId string `json:"todoId"`
}

// HTTP status code 200 and todo model in data
// swagger:response todoResp
type SwaggTodoResp struct {
	// in:body
	Body struct {
		// HTTP status code 200
		Code int `json:"code"`
		// User model
		Data model.Todo `json:"data"`
	}
}

// HTTP status code 200 and an array of todo models in data
// swagger:response todosResp
type SwaggTodosResp struct {
	// in:body
	Body struct {
		// HTTP status code 200 - Status OK
		Code int `json:"code"`
		// Array of user models
		Data model.TodoList `json:"data"`
	}
}

// swagger:parameters todoCreate
type SwaggerCreateTodoReq struct {
	// in:body
	Body model.CreateTodoReq
}
