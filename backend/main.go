package main

import (
	"fmt"
	"handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//url routes
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/register", handlers.RegisterAccount)
	mux.HandleFunc("/item", handlers.GetItemFromName)
	//mux.HandleFunc("/item", handlers.GetItemFromID)
	mux.HandleFunc("/items", handlers.GetItems)

	//show the user where the server is running
	serverAddress := "127.0.0.1:8080"

	fmt.Printf("Server running on: %s\n", serverAddress)
	fmt.Printf("Access file system on: %s/files/\n", serverAddress)

	//will be removed later
	fmt.Printf("Test item get request on: %s/item?=ITEM_NAME where you type the item's name in ITEM_NAME ", serverAddress)

	//setup the file directory for static files
	dir := http.Dir("../static")
	fs := http.FileServer(dir)

	//url route of static files in static folder
	mux.Handle("/files/", http.StripPrefix("/files/", fs))

	//run the server
	err := http.ListenAndServe(serverAddress, mux)
	if err != nil {
		log.Fatal(err)
	}
}
