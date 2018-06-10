package main

import (
	urlshortcutter "UrlShortCutterApi/shortcutter"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type url struct {
	ShortURL string
	URL      string
	Host     string
	Title    string
}

const host = "localhost:8081"
const title = "URL Short-Cutter"

func main() {
	muxer := mux.NewRouter()

	muxer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		if r.ContentLength > 0 {
			toConvert := strings.TrimPrefix(strings.TrimPrefix(r.Form.Get("URL"), "https://"), "http://")
			result := urlshortcutter.Converter(toConvert)
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
			toConvert := strings.TrimPrefix(strings.TrimPrefix(url, "https://"), "http://")
			result := host + "/" + urlshortcutter.Converter(toConvert)
			w.Write([]byte(result))
		})

	muxer.HandleFunc("/{[A-Za-z0-9]+$}",
		func(w http.ResponseWriter, r *http.Request) {
			shortURL := strings.TrimLeft(r.RequestURI, "/")
			longURL := "https://" + urlshortcutter.GetShortURL(shortURL)
			if longURL != "" {
				http.Redirect(w, r, longURL, http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		})

	http.Handle("/", muxer)
	fmt.Printf("Starting %s at %s\n", title, host)
	http.ListenAndServe(host, nil)

}
