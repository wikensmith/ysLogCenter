package ysMongo

import (
	"fmt"
	"io/ioutil"

	"labix.org/v2/mgo"

	"github.com/gin-gonic/gin"
)

type MongoOpt struct {
	Session        *mgo.Session
	CollectionName string
	Collection     *mgo.Collection
	Context        *gin.Context
}

func (m *MongoOpt) Add() {
	body := m.Context.Request.Body
	ioutil.ReadAll(body)

}

func (m *MongoOpt) AddToMongo(context *gin.Context) error {
	fmt.Println("MongoMain begin")
	CollectionName := context.Param("group") + context.Param("project")
	fmt.Println(CollectionName)
	msg := context.Param("msg")
	fmt.Println(msg)

	fmt.Println("MongoMain end")
}
func (m *MongoOpt) New(context *gin.Context) *MongoOpt {
	m.Session = <-MongoChain
	m.Session.SetMode(mgo.Monotonic, true)
	db := m.Session.DB(DBName)
	m.CollectionName = context.Param("Group") + context.Param("Project")
	m.Collection = db.C(m.CollectionName)
	return m
}

func NewMongoOpt(context *gin.Context) *MongoOpt {
	return new(MongoOpt).New(context)
}
