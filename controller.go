package main

import (
	"fmt"
	u "github.com/featherr-engineering/rest-api/utils"
	"net/http"
)

type Controller struct {
	Data map[interface{}]interface{}
}

type ControllerInterface interface {
	Init()
	Handle()
	Get()    //method = GET processing
	Post()   //method = POST processing
	Delete() //method = DELETE processing
	Put()    //method = PUT handling
}

func (c *Controller) Init() {
	c.Data = make(map[interface{}]interface{})
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, u.Message(http.StatusMethodNotAllowed, "Method Not Allowed"))
}

func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, u.Message(http.StatusMethodNotAllowed, "Method Not Allowed"))
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, u.Message(http.StatusMethodNotAllowed, "Method Not Allowed"))
}

func (c *Controller) Put(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, u.Message(http.StatusMethodNotAllowed, "Method Not Allowed"))
}

func (c *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		c.Get(w, r)
		return
	} else if r.Method == "POST" {
		c.Post(w, r)
		return
	} else if r.Method == "DELETE" {
		c.Delete(w, r)
		return
	} else if r.Method == "PUT" {
		c.Put(w, r)
		return
	} else {
		u.Respond(w, u.Message(http.StatusMethodNotAllowed, "Method Not Allowed"))
	}
}
