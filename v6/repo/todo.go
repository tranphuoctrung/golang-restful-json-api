package repo

import (
	"fmt"
	. "golang-restful-json-api/v6/models"
)

var currentId int

var Todos TodoList

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range Todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	Todos = append(Todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range Todos {
		if t.Id == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
