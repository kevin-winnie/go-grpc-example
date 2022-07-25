package core

import (
	"flag"
	"fmt"
	"go-grpc/config"
	_ "go-grpc/controller/rpc/service"
	"go-grpc/internal/service"
	"go-grpc/pkg/logger"
	"go-grpc/pkg/mysql"
	_ "google.golang.org/grpc"
	"html"
	"log"
	"net/http"
)

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`
var (
	listenAddr = flag.String("http", ":8080", "http listen address")
	dataFile = flag.String("file", "store.json", "data store file name")
	hostname = flag.String("host", "localhost:8080", "host name and port")
)

func Run(cfg *config.Config) {
	flag.Parse()

	l := logger.New(cfg.Log.Level)
	//初始化mysql
	dbObject := mysql.Init()
	db, err := dbObject.New("default", cfg.Mysql)
	if err != nil {
		log.Fatalf("db connect err:%s", err)
	}

	//urlStore := service.NewUrlStore()
	//urlStore.PutKey("https://www.google.com")
	//aa := urlStore.Get("aa1")
	//fmt.Println(aa)
	//rs, err := model.QueryOrder(db)
	//if err != nil {
	//	log.Fatalf("db query err:%s", err)
	//}
	//fmt.Println(rs)

	//grpc注册
	//grpcServer := grpc.NewServer()
	//service.Register(grpcServer)

	//crontab注册

	//http服务 gin框架
	// HTTP
	//handler := gin.New()
	//var dataFile = flag.String("file", "store.json", "data store file name")

	http.HandleFunc("/add", add)
	http.HandleFunc("/", redirect)
	err = http.ListenAndServe(":8111", nil)
	if err != nil {
		return
	}

	// Mysql Close
	if err := db.Close(); err != nil {
		l.Error(fmt.Errorf("app - Run - MysqlMap.Close: %w", err))
	}

}

func redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := service.NewUrlStore(*dataFile).Get(key)
	if url == "" {
		http.NotFound(w, r)
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func add(w http.ResponseWriter, r *http.Request) {
	//url := r.FormValue("url")
	url := "https://www.baidu.com"
	//if url == "" {
	//	fmt.Fprint(w, AddForm)
	//	return
	//}

	key := service.NewUrlStore(*dataFile).PutKey(url)
	_, err := fmt.Fprintf(w, "http://localhost:8080/%s", key)
	if err != nil {
		return
	}
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	if err != nil {
		return
	}
}
