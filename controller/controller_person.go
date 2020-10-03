package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
	"github.com/ramawidrap/goevent/model"
	"github.com/ramawidrap/goevent/utils"
)

func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var (
		person     model.Person
		arr_person []model.Person
	)
	db := utils.InitDB()
	fmt.Println("CEK BEBI")

	rows, err := db.Query("select * from users")
	if err != nil {
		fmt.Println("Error")

		if err == sql.ErrNoRows {
			fmt.Println("No result")
		}
	}

	
	defer db.Close()
	fmt.Println(rows)

	for rows.Next() {
		if err := rows.Scan(&person.ID, &person.Name, &person.Token, &person.Role); err != nil {
		} else {
			arr_person = append(arr_person, person)
		}

	}
	response := make(map[string]interface{})
	response["status"] = 1
	response["message"] = "ok"
	response["data"] = arr_person
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func InsertPerson(name string) {
	db := utils.InitDB()
	sqlStatement := `INSERT INTO PERSON (name) VALUES ($1)`
	_, err := db.Exec(sqlStatement, name)
	if err != nil {
		panic(err)
	}

}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	var user model.Person
	token:= r.FormValue("token")
	db := utils.InitDB()
	sqlStatement:= "SELECT * FROM users WHERE = $1"
	response := make(map[string]interface{})
	err:= db.QueryRow(sqlStatement,token).Scan(&user)
	if err != nil {
		response["status"] = 2
		response["message"] = "error"
		response["data"] = "failed get user"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	response["status"] = 1
	response["message"] = "ok"
	response["data"] = user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

