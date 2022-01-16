package handlers

import (
	"encoding/json"
	data2 "github.com/vishwaspuri/ecoCodes/data"
	"net/http"
	"strings"
)

type Result struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func GetAllCodes() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := data2.GetAllOpenings()
		payload := Result{
			Data: data,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(payload)
		if err != nil {
			panic(err)
		}
	})
}

func GetCode() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := strings.TrimPrefix(r.URL.Path, "/")
		data, err := data2.GetOpeningByCode(code)
		var payload Result
		if err != nil {
			payload = Result{
				Error: err.Error(),
			}
		} else {
			payload = Result{
				Data: data,
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(payload)
		if err != nil {
			panic(err)
		}
	})
}
