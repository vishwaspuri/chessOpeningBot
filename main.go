package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/vishwaspuri/ecoCodes/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.Handle("/", handlers.GetAllCodes())
	r.Handle("/{code}", handlers.GetCode())

	server := &http.Server{
		Handler:      r,
		Addr:         ":" + os.Getenv("PORT_FOR_WEBAPP"),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("HTTP server starting on :%s", os.Getenv("PORT_FOR_WEBAPP"))
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("HTTP server started on :%s", os.Getenv("PORT_FOR_WEBAPP"))
}
