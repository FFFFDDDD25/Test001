package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func HelloHandler_4(w http.ResponseWriter, req *http.Request) {

	dbURI := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", "dave.gan", "12345678", "/cloudsql", "movieenglish:us-central1:movie-english-database", "movie_database")

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	_, err = db.Exec("INSERT INTO TestMovie (name) VALUES ('Dave Gan')")
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	var name string
	row := db.QueryRow("SELECT name FROM TestMovie WHERE id = 2", name)
	err1 := row.Scan(&name)
	if err1 != nil {
		log.Fatal(err1)
		w.Write([]byte(err1.Error()))
	} else {
		log.Fatal("沒事啦")
		w.Write([]byte("沒事啦"))
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
	r.HandleFunc("/h4", HelloHandler_4)
	r.HandleFunc("/world", WorldHandler)
	r.HandleFunc("/", MainHandler)

	http.ListenAndServe(":8080", r)
}
