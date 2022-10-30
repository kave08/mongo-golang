package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)


func main() {

	r := httprouter.New()
	
	r.GET("",)
	r.POST("",)
	r.DELETE("",)

}