package http

import (
	"encoding/json"
	"net/http"
)

func renderJSON(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
