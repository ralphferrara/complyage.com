package responses

import (
	"encoding/json"
	"net/http"
)

//||------------------------------------------------------------------------------------------------||
//|| Success Response
//||------------------------------------------------------------------------------------------------||

func Success(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]any{
		"success": true,
		"data":    data,
	}

	json.NewEncoder(w).Encode(resp)
}

//||------------------------------------------------------------------------------------------------||
//|| Error Response
//||------------------------------------------------------------------------------------------------||

func Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]any{
		"success": false,
		"error":   message,
	}

	json.NewEncoder(w).Encode(resp)
}
