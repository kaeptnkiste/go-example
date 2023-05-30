package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, ", r.URL.Query().Get("name"))

}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}
	http.HandleFunc("/hello",hello)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
