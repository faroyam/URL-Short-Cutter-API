package response

import (
	"encoding/json"
	"net/http"

	"github.com/faroyam/url-short-cutter-API/config"
)

type response struct {
	Service string `json:"service"`
	Answer  string `json:"answer"`
}

// NewResponse write json to http.ResponseWriter
func NewResponse(w http.ResponseWriter, s string) {
	var resp = response{}
	resp.Service = config.C.Service
	resp.Answer = s
	json.NewEncoder(w).Encode(resp)
}
