package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{path:.*}", handler).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions)

	addr := "127.0.0.1:8080"
	log.Printf("Server listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request path: %s", r.URL.Path)

	log.Println("Request query parameters:")
	for name, values := range r.URL.Query() {
		for _, value := range values {
			log.Printf("%s: %s", name, value)
		}
	}

	log.Println("Request headers:")
	for name, values := range r.Header {
		for _, value := range values {
			log.Printf("%s: %s", name, value)
		}
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
	} else {
		log.Printf("Request body: %s", string(body))
	}

	// Echo the request back to the client
	fmt.Fprintln(w, "👋")
}
