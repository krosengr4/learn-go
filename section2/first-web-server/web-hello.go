package main

// Importing is how we access packages in the Go library
import (
	"fmt"
	"net/http"
)

// To run this try the url http://localhost:8080?name=<yourName>
func main() {
	// Handles requests to the /hello path
	http.HandleFunc("/hello", Handler)

	// Listening to requests on port 8080
	http.ListenAndServe(":8080", nil)

}

func Handler(rw http.ResponseWriter, req *http.Request) {
	// Using the url to get the value of name (8080?name=Bob)
	name := req.URL.Query().Get("name")
	// Writing to the web page
	rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
}
