package main

import (
	"fmt"
	"io"
	"net/http"
)

type Endpoint struct {
	Method string
	Path string
	Expectation any
}


func MockRequest(w http.ResponseWriter, r *http.Request) {
	// get server object from database with server code
	servercode := r.URL.Query().Get("servercode")

	server := mockServer.GetServer(servercode)

	if server == nil {
		http.Error(w, "Server not found", http.StatusNotFound)
		return
	}

	// handle the request
	
	switch r.Method {

	case http.MethodGet: // returns all endpoints
		fmt.Printf("received GET request\n")
		serverInfo, err := server.GetServer()
		if err != nil {
			http.Error(w, "Error getting server info", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(serverInfo)
		return

	case http.MethodPost: // adds a new endpoint
		fmt.Printf("received POST request\n")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		fmt.Printf("body: %v\n", string(body))


	case http.MethodPut: // updates an endpoint
		fmt.Printf("received PUT request\n")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "body: %v\n", string(body))
	}


	
	
	// save to database






}