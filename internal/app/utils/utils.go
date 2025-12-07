package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseBody struct {
	Error string `json:"error"`
}

func ResponseError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&ResponseBody{Error: message})
}
