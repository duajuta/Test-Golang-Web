package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"title": "Belajar Golang Web",
			"name": "Wawin Fauzani"}
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/member", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"title": "Belajar Golang Web",
			"name": "Wawin Fauzani"}
		err = tmpl.ExecuteTemplate(w, "member", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("asset"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
