package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//gcloud sql connect movie-english-database --user=dave.gan --quiet

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

	var id string = req.URL.Query().Get("id")
	var name string = req.URL.Query().Get("name")
	var sfsfd string = req.URL.Query().Get("sfsfd")
	if sfsfd != "" {
		w.Write([]byte("sfsfd非空白"))
	} else {
		w.Write([]byte("sfsfd為空白"))
	}

	if id != "" {
		err1 := db.QueryRow("select name from TestMovie where id = ?", id).Scan(&name)
		if err1 != nil {
			w.Write([]byte(err1.Error()))
		} else {
			w.Write([]byte("沒事啦," + name))
		}
	} else if name != "" {
		var n string
		err1 := db.QueryRow("select name from TestMovie where name = ?", name).Scan(&n)
		if err1 != nil {
			w.Write([]byte(err1.Error()))
		} else {
			w.Write([]byte("沒事啦," + n))
		}
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
