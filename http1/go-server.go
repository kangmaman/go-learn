package main

import (
	"net/http"
)

// create a handler struct
type HTTPHandler struct{}

//implement ServeHTTP
func (h HTTPHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//create response binary data
	data := []byte("Hello world!") //clise of bytes

	//write data to response
	res.Write(data)
}

func main() {
	//create a new handler
	handler := HTTPHandler{}

	//listen and server
	http.ListenAndServe(":9000", handler)
}
