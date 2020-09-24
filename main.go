package main

import (
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello!"))
}
func WorldHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("World!"))
}
func MainHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Main Page!"))
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/hello", HelloHandler)
	r.HandleFunc("/world", WorldHandler)
	r.HandleFunc("/", MainHandler)

	http.ListenAndServe(":8080", r)
}
