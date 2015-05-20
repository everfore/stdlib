package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("ready...")
	http.HandleFunc("/callback", callback)
	http.HandleFunc("/", root)
	err := http.ListenAndServe(":80", nil)
	if check_err(err) {
		return
	}
}

func callback(rw http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	rw.Write([]byte("[CALLBACK]" + time.Now().String()))
}

func root(rw http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	rw.Write([]byte("[ROOT]" + time.Now().String()))
}

func check_err(err error) bool {
	if nil != err {
		return true
	}
	return false
}
