package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"models"
	"net/http"
)

// send json response by passing the responseWriter and a struct
func SendJSONResponse(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(v)
}

// route: /search
func SearchItem(w http.ResponseWriter, r *http.Request) {
	//read response body (post request data) bytes
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Can't read request body:", err)
	}

	//create post request struct instance
	var itemSearch models.ItemSearch
	//get post request data into loginData struct
	err = json.Unmarshal(body, &itemSearch)

	if err != nil {
		fmt.Println("Unmarshal error:", err)
	} else {
		//use post request data
		fmt.Printf("Item: %s\n", itemSearch.Item)
	}

	//item struct for json response
	//later will come from database query
	item := models.Item{ID: "1", Name: "laptop", Price: 499.99}

	SendJSONResponse(w, item)
}

// route: /login
func Login(w http.ResponseWriter, r *http.Request) {
	//read response body (post request data) bytes
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "Can't read request body")
	}

	//create post request struct instance
	var loginData models.UserLogin
	//get post request data into loginData struct
	err = json.Unmarshal(body, &loginData)

	if err != nil {
		fmt.Println("Unmarshal error:", err)
	} else {
		//use post request data
		fmt.Println("login data:", loginData)
	}
	loginData.Password = ""
	SendJSONResponse(w, loginData)
}
