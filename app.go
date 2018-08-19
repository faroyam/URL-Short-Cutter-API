package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faroyam/url-short-cutter-API/config"
	"github.com/faroyam/url-short-cutter-API/routes"
	"github.com/faroyam/url-short-cutter-API/shortcutter"
	"github.com/gorilla/mux"
)

type url struct {
	ShortURL string
	URL      string
	Host     string
}

func main() {
	defer shortcutter.DB.Close()
	var Muxer = mux.NewRouter()

	Muxer.HandleFunc("/v1", routes.Converter)
	Muxer.HandleFunc("/{[A-Za-z0-9]+$}", routes.Redirecter)
	http.Handle("/", Muxer)

	fmt.Printf("Starting %s %s at %s\nMongoDB at %s\n\n", config.C.Service, config.C.Version, config.C.Host, config.C.MongoIP)

	err := http.ListenAndServe(config.C.Host, nil)
	if err != nil {
		log.Fatal(err)
	}
}
