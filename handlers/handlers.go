package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	data2 "github.com/vishwaspuri/ecoCodes/data"
	"github.com/vishwaspuri/ecoCodes/utils"
	"log"
	"net/http"
	"strings"
)

type Result struct {
	Cached bool        `json:"cached,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func GetAllCodes() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := data2.GetAllOpenings()
		payload := Result{
			Data: data,
		}
		writePayload(w, http.StatusOK, payload)
	})
}

func GetCode(rdb *redis.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		code := strings.TrimPrefix(path, "/")

		// Checking redis cache
		data, err := utils.GetCache(path, rdb)
		if err == nil {
			payload := Result{
				Data:   data,
				Cached: true,
			}
			writePayload(w, http.StatusOK, payload)
			return
		}

		// If cache isn't found, web scraper is used
		data, err = data2.GetOpeningByCode(code)
		var payload Result
		if err != nil {
			payload = Result{
				Error: err.Error(),
			}
			writePayload(w, http.StatusBadRequest, payload)
		} else {
			payload = Result{
				Data:   data,
				Cached: false,
			}
			writePayload(w, http.StatusOK, payload)
		}

		// Web scrapped data is cached
		err = utils.InsertCache(path, data, rdb)
		if err != nil {
			fmt.Println(err)
		}
	})
}

func writePayload(w http.ResponseWriter, responseCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Fatalln(err)
	}
}
