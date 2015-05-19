package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("ready...")
	http.HandleFunc("/callback", callback)
	err := http.ListenAndServe(":80", nil)
	if check_err(err) {
		return
	}
}

func callback(rw http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	var ret interface{}
	err := json.NewDecoder(req.Body).Decode(&ret)
	if check_err(err) {
		rw.Write([]byte(err.Error()))
		return
	}
	log.Printf("%#v\n", ret)
}

func check_err(err error) bool {
	if nil != err {
		return true
	}
	return false
}
