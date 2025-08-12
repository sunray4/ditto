package main

import (
	"net/http"
)

// url shoudl follow the format:
// http://ditto.io/<servercode>/<command>


func main() {
	// http.HandleFunc("/createUser", createUser)
	http.HandleFunc("/createMock", CreateMock)
	http.HandleFunc("/mock", MockRequest)  // url shoudl be {url}/<mock>?servercode=<servercode>
	

	http.ListenAndServe(":8080", nil)
}
