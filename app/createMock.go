package main

import (
	"ditto-backend/mockServer"
	"fmt"
	"net/http"
)

func CreateMock(w http.ResponseWriter, req *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "text/plain")

	// Validate HTTP method
	if req.Method != http.MethodGet && req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get and validate required parameters
	username := req.URL.Query().Get("username")
	servername := req.URL.Query().Get("servername")

	if username == "" {
		http.Error(w, "username parameter is required", http.StatusBadRequest)
		return
	}

	if servername == "" {
		http.Error(w, "servername parameter is required", http.StatusBadRequest)
		return
	}

	// Create the mock server
	newMockServer := mockServer.CreateMockServer(username, servername)

	// save to database  of servers

	// save to user's list of servers (just server code)

	fmt.Fprintf(w, "mock server created: %v\n", newMockServer)
	
}