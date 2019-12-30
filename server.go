// Server for example electronic medical record using Graph Database.

package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const (
	ServerAddress = "localhost:8080"
	DGraphAddress = "localhost:9080"
)

func main() {
	router := getRouter()
	log.Printf("serving on %s", ServerAddress)
	log.Fatal(http.ListenAndServe(ServerAddress, router))
}

func getRouter() *httprouter.Router {
	r := httprouter.New()
	r.GET("/", GetIndex)
	r.POST("/patients", AddPatient)
	r.GET("/patients", GetPatients)
	r.GET("/patients/:id", GetPatient)
	r.PATCH("/patients/:id", PatchPatient)
	return r
}
