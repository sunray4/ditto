package main

import (
	"net/http"
)

// url shoudl follow the format:
// http://ditto.io/<servercode>/<command>


func main() {

	http.HandleFunc("/createMock", CreateMock)
	

	http.ListenAndServe(":8080", nil)
}
