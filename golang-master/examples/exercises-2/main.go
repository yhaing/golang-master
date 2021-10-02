package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var count int = 1

func main() {
	fmt.Printf("====== This is http server ======\n")
	flag.Set("v", "4")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/getstatuscode", Getstatuscode)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("------ This is roothandler[%d] ------\n", count)
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("Hello [%s]\n", user))
	} else {
		io.WriteString(w, "Hello [Stranger]\n")
	}
	io.WriteString(w, "========== Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s = %s\n", k, v))
	}
	count++
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("------ This is healthz[%d] ------\n", count)
	io.WriteString(w, "200\n")
	count++
}

func Getstatuscode(w http.ResponseWriter, r *http.Request) {
	rsp, err := http.Get("http://127.0.0.1:8090")
	if err != nil {
		fmt.Println("httpget rsp error is: ", err)
		return
	}
	io.WriteString(w, fmt.Sprintf("getstatuscode is: %d", rsp.StatusCode))
	fmt.Printf("statuscode is: %d", rsp.StatusCode)
}
