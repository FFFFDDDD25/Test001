package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func HelloHandler_4(w http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("mysql", "dave.gan:12345678@tcp(35.194.153.230)/movie_database")
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	_, err = db.Exec("INSERT INTO TestMovie (name) VALUES ('Dave Gan')")
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	return
}

func HelloHandler_3(w http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("mysql", "dave.gan:12345678@tcp(34.66.219.20)/movie_database")
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	_, err = db.Exec("INSERT INTO TestMovie (name) VALUES ('Dave Gan')")
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	return
}

func HelloHandler_2(w http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("mysql", "dave.gan:12345678@tcp(35.194.153.230:3306)/movie_database")
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	_, err = db.Exec("INSERT INTO TestMovie (name) VALUES ('Dave Gan')")
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	return
}

func HelloHandler_1(w http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("mysql", "dave.gan:12345678@tcp(34.66.219.20:3306)/movie_database")
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	_, err = db.Exec("INSERT INTO TestMovie (name) VALUES ('Dave Gan')")
	if err != nil {
		w.Write([]byte(err.Error()))
	}

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
	r.HandleFunc("/h1", HelloHandler_1)
	r.HandleFunc("/h2", HelloHandler_2)
	r.HandleFunc("/h3", HelloHandler_3)
	r.HandleFunc("/h4", HelloHandler_4)
	r.HandleFunc("/world", WorldHandler)
	r.HandleFunc("/", MainHandler)

	http.ListenAndServe(":8080", r)
}
