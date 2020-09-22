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

    // listen and serve
    http.ListenAndServe(":8080/Hello", helloHandler)
    http.ListenAndServe(":8080/World", worldHandler)

}
