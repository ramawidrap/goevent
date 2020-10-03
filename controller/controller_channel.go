package controller

import (
	"fmt"
	"database/sql"
	"encoding/json"
	"net/http"
	"math/rand"
	"github.com/ramawidrap/goevent/model"
	"github.com/ramawidrap/goevent/utils"
	"github.com/gorilla/mux"

)

func CreateChannel(w http.ResponseWriter, r *http.Request) {
	var channel model.Channel
	token:= createToken()
	_ = json.NewDecoder(r.Body).Decode(&channel)
	channel.Token = token
	db := utils.InitDB()
	var name string
	sqlCheckStatement:= "SELECT name FROM CHANNEL WHERE token = $1"
	errCheck:= db.QueryRow(sqlCheckStatement,token).Scan(&name)
	for errCheck != sql.ErrNoRows {
		token = createToken()
		sqlCheckStatement:= "SELECT name FROM CHANNEL WHERE token = $1"
		errCheck= db.QueryRow(sqlCheckStatement,token).Scan(&name)
	}
	sqlStatement := `INSERT INTO CHANNEL (token,name) VALUES ($1,$2)`
	
	_, err := db.Exec(sqlStatement, token,channel.Name)
	if err != nil {
		panic(err)
	}
	sqlStatement = "INSERT INTO users (name,token,role) VALUES ($1,$2,$3) returning id"
	var id int
	err = db.QueryRow(sqlStatement,channel.User,token,"owner").Scan(&id)
	fmt.Println(channel.User)
	if err != nil {
		panic(err)
	}
	var user model.Person
	user.ID = id
	user.Name = channel.User
	user.Token = token
	user.Role = "owner"
	response := make(map[string]interface{})
	response["status"] = 1
	response["message"] = "ok"
	response["data"] = user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func createToken() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 6)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func JoinChannel(w http.ResponseWriter, r *http.Request) {
	db := utils.InitDB()
	var channel model.Channel
	sqlCheckStatement:= "SELECT name FROM CHANNEL WHERE token = $1"
	_ = json.NewDecoder(r.Body).Decode(&channel)
	var name string
	errCheck:= db.QueryRow(sqlCheckStatement,channel.Token).Scan(&name)
	response := make(map[string]interface{})
	if errCheck == sql.ErrNoRows {	
	response["status"] = 2
	response["message"] = "error"
	response["data"] = "failed join channel"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	} else {
		sqlCheckStatement = "INSERT INTO users (name,token,role) VALUES ($1, $2, $3) returning id"
		var id int
		err := db.QueryRow(sqlCheckStatement,channel.User,channel.Token,"participant").Scan(&id)
		if err != nil {
			response["status"] = 2
			response["message"] = "error"
			response["data"] = "failed join channel"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
		var user model.Person
		user.ID = id
		user.Token = channel.Token
		user.Role = "participant"
		user.Name = channel.User
		response["status"] = 1
			response["message"] = "ok"
			response["data"] = user
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
	}
}

func EndChannel(w http.ResponseWriter, r *http.Request) {
	db := utils.InitDB()
	var channel model.Channel
	_ = json.NewDecoder(r.Body).Decode(&channel)
	response := make(map[string]interface{})
	params := mux.Vars(r)
	sqlStatement:= "DELETE FROM CHANNEL WHERE token = $1"
	_, err := db.Exec(sqlStatement,params["token"])
	if err != nil {
		response["status"] = 2
		response["message"] = "error"
		response["data"] = "failed end channel"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	response["status"] = 1
	response["message"] = "ok"
	response["data"] = "success end channel"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
