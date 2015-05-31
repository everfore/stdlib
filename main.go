package main

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("templates")))
	http.HandleFunc("/r", root)
	http.ListenAndServe(":80", nil)
}

func root(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("AngularJS"))
	io.WriteString(rw, "Angular.js")
	fmt.Fprintln(rw, "A n g u l a r . j s")
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
