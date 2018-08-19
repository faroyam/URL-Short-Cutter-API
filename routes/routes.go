package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/faroyam/url-short-cutter-API/config"
	"github.com/faroyam/url-short-cutter-API/response"
	"github.com/faroyam/url-short-cutter-API/shortcutter"
)

// Converter router for /v1?...
func Converter(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url != "" {
		toConvert := strings.TrimPrefix(strings.TrimPrefix(url, "https://"), "http://")

		shortURL, err := shortcutter.Converter(toConvert)
		if err != nil {
			log.Println("db error", err)
			w.WriteHeader(http.StatusForbidden)
			response.NewResponse(w, "db error")
		}
		log.Println(r.RemoteAddr, "converted", url, shortURL)

		w.WriteHeader(http.StatusOK)
		response.NewResponse(w, config.C.Host+"/"+shortURL)
	} else {
		w.WriteHeader(http.StatusForbidden)
		response.NewResponse(w, "invalid request")
	}

}

// Redirecter router for short URLs
func Redirecter(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimLeft(r.RequestURI, "/")
	longURL, err := shortcutter.ReConverter(shortURL)

	if err != nil {
		log.Println("not redirected:", r.RemoteAddr, shortURL, longURL, "err:", err)
		w.WriteHeader(http.StatusForbidden)
		response.NewResponse(w, "invalid request")

	} else {
		log.Println("redirected:", r.RemoteAddr, shortURL, longURL, "err:", err)
		http.Redirect(w, r, "http://"+longURL, http.StatusSeeOther)
	}
}
