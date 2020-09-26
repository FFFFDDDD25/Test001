package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	_ "github.com/go-sql-driver/mysql"
)

//gcloud sql connect movie-english-database --user=dave.gan --quiet

func HelloHandler_3(w http.ResponseWriter, req *http.Request) {

	ctx := context.Background()

	projectID := "movieenglish"
	bucketName := "my-new-bucket"
	fileName := "測試檔名.txt"

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		w.Write([]byte("0000" + err.Error()))
		return
	}

	bucket := client.Bucket(bucketName)

	// Creates the new bucket.
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		w.Write([]byte("1111" + err.Error()))
		return
	}

	fmt.Printf("Bucket %v created.\n", bucketName)

	io.WriteString(w, "\nAbbreviated file content (first line and last 1K):\n")

	rc, err := bucket.Object(fileName).NewReader(ctx)
	if err != nil {
		w.Write([]byte("2222" + err.Error()))
		return
	}
	defer rc.Close()
	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		w.Write([]byte("3333" + err.Error()))
		return
	}

	w.Write([]byte("4444"))
	w.Write(slurp)
	//fmt.Fprintf(w, "%s\n", bytes.SplitN(slurp, []byte("\n"), 2)[0])
}

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
	r.HandleFunc("/h3", HelloHandler_3)
	r.HandleFunc("/world", WorldHandler)
	r.HandleFunc("/", MainHandler)

	http.ListenAndServe(":8080", r)
}
