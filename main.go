package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	port := GetPort()
	log.Println("[-] Listening on...", port)
    http.HandleFunc("/", func (res http.ResponseWriter, req *http.Request) {
        fmt.Fprintln(res, "hello, world")
    })

    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
