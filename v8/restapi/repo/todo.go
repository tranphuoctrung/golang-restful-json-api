package repo

import (
	. "golang-restful-json-api/v8/restapi/models"
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DBContext struct {
	Server     string
	Database   string
	DbUser     string
	DbPassword string
}

var todoDB *mgo.Database

const (
	COLLECTION = "todos"
)

func (m *DBContext) Connect() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{m.Server},
		Timeout:  60 * time.Second,
		Database: m.Database,
		Username: m.DbUser,
		Password: m.DbPassword,
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatal(err)
	}
	todoDB = session.DB(m.Database)
}

// Find list of todos
func (m *DBContext) FindAll() (TodoList, error) {
	var todos TodoList
	err := todoDB.C(COLLECTION).Find(bson.M{}).All(&todos)
	return todos, err
}

// Find a todo by its id
func (m *DBContext) FindById(id string) (Todo, error) {
	var todo Todo
	err := todoDB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&todo)
	return todo, err
}

// Insert a todo into database
func (m *DBContext) Insert(todo Todo) error {
	err := todoDB.C(COLLECTION).Insert(&todo)
	return err
}

// Delete an existing todo
func (m *DBContext) Delete(todo Todo) error {
	err := todoDB.C(COLLECTION).Remove(&todo)
	return err
}

// Update an existing todo
func (m *DBContext) Update(todo Todo) error {
	err := todoDB.C(COLLECTION).UpdateId(todo.Id, &todo)
	return err
}
