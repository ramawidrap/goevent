package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ramawidrap/goevent/controller"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", controller.GetAllPerson).Methods("GET")
	router.HandleFunc("/channel/create", controller.CreateChannel).Methods("POST")
	router.HandleFunc("/channel/join",controller.JoinChannel).Methods("POST")
	router.HandleFunc("/channel/delete/{token}",controller.EndChannel).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
