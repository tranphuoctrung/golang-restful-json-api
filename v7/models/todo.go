package models

import "gopkg.in/mgo.v2/bson"

type Todo struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `json:"name"`
	Completed bool          `json:"completed"`
}

type TodoList []Todo
