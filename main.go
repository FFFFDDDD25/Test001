package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	//FFFFFDDDDOOOO~~!!!!))))))
	db, err1 := sql.Open("mysql", "dave.gan:12345678@/movie_database")
	if err1 != nil {
		w.Write([]byte("sql open error"))
		return
	}

	_, err2 := db.Exec("INSERT INTO TestMovie (name) VALUES ('Dave Gan')")

	if err2 != nil {
		w.Write([]byte("sql insert error"))
		return
	}

	w.Write([]byte("Hello!"))
	return
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
