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
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

//gcloud sql connect movie-english-database --user=dave.gan --quiet

func HelloHandler_2(w http.ResponseWriter, req1 *http.Request) {

	var err error
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		},
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	_, err = selenium.NewChromeDriverService("/opt/google/chrome/chromedriver", 9515, opts...)
	if err != nil {
		w.Write([]byte("1111" + err.Error()))
		return
	}
	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		w.Write([]byte("8888" + err.Error()))
		return
	}
	// 这是目标网站留下的坑，不加这个在linux系统中会显示手机网页，每个网站的策略不一样，需要区别处理。
	webDriver.AddCookie(&selenium.Cookie{
		Name:  "defaultJumpDomain",
		Value: "www",
	})
	// 导航到目标网站
	err = webDriver.Get("https://www.xvideos.com/")
	if err != nil {
		w.Write([]byte("7777" + err.Error()))
		return
	}

	title, err := webDriver.Title()
	if err != nil {
		w.Write([]byte("6666" + title))
		return
	}
	w.Write([]byte("5555" + title))
}

func HelloHandler_3(w http.ResponseWriter, req *http.Request) {

	ctx := context.Background()

	projectID := "movieenglish"
	bucketName := projectID + ".appspot.com"
	fileName := "測試檔名.txt"

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		w.Write([]byte("0000" + err.Error()))
		return
	}

	bucket := client.Bucket(bucketName)

	// Creates the new bucket.
	FALSE := false
	if FALSE {
		ctx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()
		if err := bucket.Create(ctx, projectID, nil); err != nil {
			w.Write([]byte("1111" + err.Error()))
			return
		}
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

/*
<html><head>
<meta http-equiv="content-type" content="text/html;charset=utf-8">
<err>500 Server Error</err>
</head>
<body text=#000000 bgcolor=#ffffff>
<h1>Error: Server Error</h1>
<h2>The server encountered an error and could not complete your request.<p>Please try again in 30 seconds.</h2>
<h2></h2>
</body></html>
*/

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
	r.HandleFunc("/h2", HelloHandler_2)
	r.HandleFunc("/world", WorldHandler)
	r.HandleFunc("/", MainHandler)

	http.ListenAndServe(":8080", r)
}
