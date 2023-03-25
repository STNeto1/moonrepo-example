package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Data struct {
	Timestamp string `json:"timestamp"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Timestamp: time.Now().Format(time.RFC3339),
		}

		serialized, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(serialized)
	})

	log.Println("Listening on :4000")
	http.ListenAndServe(":4000", r)
}
