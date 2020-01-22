package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type jsonMap map[string]interface{}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/hello", handleHello())
	r.Get("/ready", handleReady())
	r.Get("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Fatalf("Oh well. Time to say goodbye!")
	})

	log.Printf("Starting service on port 8000 ...")
	http.ListenAndServe(":8000", r)
}

func handleHello() http.HandlerFunc {
	hostname, _ := os.Hostname()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling /hello on: %s\n", hostname)
		resp, _ := json.Marshal(jsonMap{
			"message": "hello world",
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	})
}

func handleReady() http.HandlerFunc {
	state := http.StatusServiceUnavailable
	go func() {
		time.Sleep(time.Second * 15)
		state = http.StatusOK
	}()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, _ := json.Marshal(jsonMap{
			"state": http.StatusText(state),
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(state)
		w.Write(resp)
	})
}
