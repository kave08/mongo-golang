package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)


func main() {

	r := httprouter.New()
	uc := controller.NewUserController(getSession())

	r.GET("",)
	r.POST("",)
	r.DELETE("",)

}

func getSession() *mgo.Session{

	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil{
	}
}