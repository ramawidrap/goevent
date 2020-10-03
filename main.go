package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramawidrap/goevent/controller"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.HelloWorld).Methods("GET")
	router.HandleFunc("/user", controller.GetAllPerson).Methods("GET")
	router.HandleFunc("/channel/create", controller.CreateChannel).Methods("POST")
	router.HandleFunc("/channel/join", controller.JoinChannel).Methods("POST")
	router.HandleFunc("/channel/delete/{token}", controller.EndChannel).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

	err := http.ListenAndServe(":"+"8080", nil)
	if err != nil {
		panic(err)
	}
}
