package routing

import (
	. "golang-restful-json-api/v8/restapi/handlers"
	. "golang-restful-json-api/v8/restapi/logging"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	// swagger:route GET /todos todos listTodos
	//
	// Handler returning list of todos.
	//
	// Responses:
	// 	200: todoList
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},

	// swagger:route GET /todos/{todoId} todos getTodoById
	//
	// Handler returning information about todo.
	//
	// Information about todo
	//
	// Responses:
	// 	200: todoResp
	//  404: badReq

	Route{
		"GetTodoById",
		"GET",
		"/todos/{todoId}",
		GetTodoById,
	},

	// swagger:route POST /todos todos todoCreate
	//
	// Handler creating a todo.
	//
	// Responses:
	// 	200: todoResp
	//  400: badReq
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},

	// swagger:route DELETE /todos/{todoId} todos todoDelete
	//
	// Handler deleting a todo.
	//
	// Responses:
	// 	200: ok
	//  404: badReq
	//  500: internal
	Route{
		"TodoDelete",
		"DELETE",
		"/todos/{todoId}",
		TodoDelete,
	},
}

func RegisterTodoRoutes(r *mux.Router) {

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
}
