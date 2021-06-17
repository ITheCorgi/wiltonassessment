package cond

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"wiltonassessment/helper"
)

func GetInputData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			helper.ErrorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
			return
		}
		var parseErr *json.UnmarshalTypeError
		var i Input
		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&i)
		if err != nil {
			if errors.As(err, &parseErr) {
				helper.ErrorResponse(w, fmt.Sprint("Bad Request. Wrong Type provided for field ", parseErr.Field), http.StatusBadRequest)
			} else {
				helper.ErrorResponse(w, fmt.Sprint("Bad Request ", err.Error()), http.StatusBadRequest)
			}
			return
		}
		log.Println(i)

		if err = dataProceeding(&i); err != nil {
			log.Fatal(err)
		}

		helper.ErrorResponse(w, "Success", http.StatusOK)
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
