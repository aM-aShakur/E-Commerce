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

	//show the user where the server is running
	serverAddress := "127.0.0.1:8080"
	fmt.Printf("Server running on: %s\n", serverAddress)
	fmt.Printf("Access file system on: %s/files/\n", serverAddress)

	//setup the file directory for static files
	dir := http.Dir("../frontend")
	fs := http.FileServer(dir)

	//url route of static files in frontend folder
	mux.Handle("/files/", http.StripPrefix("/files/", fs))

	//run the server
	err := http.ListenAndServe(serverAddress, mux)
	if err != nil {
		log.Fatal(err)
	}
}
