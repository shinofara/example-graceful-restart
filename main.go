package main

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"io"
	"net/http"
	"os"
	"time"
)

var now = time.Now()

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	//	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	//	http.ListenAndServe(":8080", nil)

	s := &http.Server{Addr: ":8080", Handler: &myHandler{}}
	gracehttp.Serve(s)
}

var mux map[string]func(http.ResponseWriter, *http.Request)

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	io.WriteString(w, fmt.Sprintf("pid is %d", os.Getpid()))
}
