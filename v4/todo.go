package main

type Todo struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type Todos []Todo
