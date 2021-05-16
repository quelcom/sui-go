package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/quelcom/sui-go/config"
	"github.com/quelcom/sui-go/handlers"
)

//go:embed templates
var indexHTML embed.FS

//go:embed assets
var staticFiles embed.FS

var tpl = template.Must(template.ParseFS(indexHTML, "templates/index.html.tmpl"))

func main() {
	conf, err := config.ParseFromFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/assets/", http.FileServer(http.FS(staticFiles)))

	page := handlers.NewPage(&conf, tpl)
	http.HandleFunc("/", page.FrontpageHandler)

	log.Printf("Listening on :%d\n", conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil))
}
