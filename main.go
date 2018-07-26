package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/faroyam/url-short-cutter-API/shortcutter"

	"github.com/gorilla/mux"
)

type url struct {
	ShortURL string
	URL      string
	Host     string
	Title    string
}

const title = "URL Short-Cutter"

func main() {
	if len(os.Args) < 3 {
		log.Fatal("\nfirst arguement: <localAddr:port>\nsecond arguement: <mongoAddr:port>")
	}
	var host = os.Args[1]
	var mongo = os.Args[2]
	muxer := mux.NewRouter()

	muxer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		if r.ContentLength > 0 {
			toConvert := strings.TrimPrefix(strings.TrimPrefix(r.Form.Get("URL"), "https://"), "http://")
			result := urlshortcutter.Converter(toConvert, mongo)
			u := url{ShortURL: result, URL: toConvert, Host: host, Title: title}
			t := template.Must(template.ParseFiles("tmpl/result.html"))
			t.Execute(w, u)
		} else {
			u := url{Title: title}
			t := template.Must(template.ParseFiles("tmpl/index.html"))
			t.Execute(w, u)
		}
	})

	muxer.HandleFunc("/v1",
		func(w http.ResponseWriter, r *http.Request) {
			url := r.URL.Query().Get("url")
			if url != "" {
				toConvert := strings.TrimPrefix(strings.TrimPrefix(url, "https://"), "http://")
				result := host + "/" + urlshortcutter.Converter(toConvert, mongo)
				w.Write([]byte(result))
			} else {
				w.Write([]byte("invalid request"))
			}

		})

	muxer.HandleFunc("/{[A-Za-z0-9]+$}",
		func(w http.ResponseWriter, r *http.Request) {
			shortURL := strings.TrimLeft(r.RequestURI, "/")
			longURL := "https://" + urlshortcutter.ReConverter(shortURL, mongo)
			if longURL != "" {
				http.Redirect(w, r, longURL, http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		})

	http.Handle("/", muxer)
	fmt.Printf("Starting %s at %s\nMongoDB at %s\n", title, host, mongo)
	http.ListenAndServe(host, nil)

}
