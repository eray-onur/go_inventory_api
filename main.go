package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go_inventory_api/controllers"
	"github.com/go_inventory_api/middlewares"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func main() {

	var gracefulShutdownTime time.Duration

	r := mux.NewRouter()
	r.Use(middlewares.LoggingMiddleware)

	r.HandleFunc("/", homeHandler)

	r.HandleFunc("/products", controllers.GetProductsHandler).Methods("GET", "OPTIONS")

	http.Handle("/", r)
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("/assets"))),
	)

	server := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTime)
	defer cancel()

	server.Shutdown(ctx)
	log.Println("Shutting down...")

	os.Exit(0)

}
