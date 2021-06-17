package helper

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, msg string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	r := make(map[string]string)
	r["msg"] = msg
	jR, _ := json.Marshal(r)
	w.Write(jR)
}
