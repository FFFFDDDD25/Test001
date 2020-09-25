package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	//FFFFFDDDDOOOO~~!!!!))))))AAAAA

	var err error

	_, err = sql.Open("mysql", "dave.gan:12345678@tcp(34.66.219.20:3306)/movie_database")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	_, err = sql.Open("mysql", "dave.gan:12345678@tcp(35.194.153.230:3306)/movie_database")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	_, err = sql.Open("mysql", "dave.gan:12345678@tcp(34.66.219.20)/movie_database")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	db, err := sql.Open("mysql", "dave.gan:12345678@tcp(35.194.153.230)/movie_database")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	//35.194.153.230
	_, err = db.Exec("INSERT INTO TestMovie (name) VALUES ('Dave Gan')")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Hello!  sql insert done!!"))
	return
}
func WorldHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("World!~~~@@@@@"))
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
