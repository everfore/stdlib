package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("stdlib/")))
	err := http.ListenAndServe(":80", nil)
	if check_err(err) {
		return
	}
}

func check_err(err error) bool {
	if nil != err {
		return true
	}
	return false
}
