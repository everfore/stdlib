package main

import (
	"github.com/Unknwon/macaron"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("templates")))
	http.HandleFunc("/r", root)
	http.ListenAndServe(":80", nil)
}

func root(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("AngularJS"))
}

func v1() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/", func(ctx *macaron.Context) {
		// ctx.Data["Name"] = "jeremy"
		ctx.HTML(200, "hello") // 200 为响应码
	})
	m.Run(80)
}
