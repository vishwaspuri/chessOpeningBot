package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/vishwaspuri/ecoCodes/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	_ = godotenv.Load(".env")
	//Adding redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	fmt.Println("Redis Connected")

	r := mux.NewRouter()
	r.Handle("/", handlers.GetAllCodes())
	r.Handle("/{code}", handlers.GetCode(rdb))

	server := &http.Server{
		Handler:      r,
		Addr:         ":" + os.Getenv("PORT_FOR_WEBAPP"),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("HTTP server starting on :%s", os.Getenv("PORT_FOR_WEBAPP"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("HTTP server started on :%s", os.Getenv("PORT_FOR_WEBAPP"))
}
