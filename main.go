package main

import (
    "net/http"
)

// create a handler struct
type HelloHandler struct{}
type WorldHandler struct{}

// implement `ServeHTTP` method on `HelloHandler` struct
func (h HelloHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

    // create response binary data
    data := []byte("Hello") // slice of bytes

    // write `data` to response
    res.Write(data)
}


func (h WorldHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

    // create response binary data
    data := []byte("World") // slice of bytes

    // write `data` to response
    res.Write(data)
}

func main() {

    // create a new handler
    helloHandler := HelloHandler{}
    worldHandler := WorldHandler{}

    http.HandleFunc("/Hello", helloHandler)
    http.HandleFunc("/World", worldHandler)
    http.ListenAndServe(":8080", helloHandler)

}
