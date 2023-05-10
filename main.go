package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	service := flag.String("service", "random", "response of a service you'd like to use")
	flag.Parse()

	r.HandleFunc("/{path:.*}", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, *service)
	}).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions)

	addr := "127.0.0.1:8080"
	log.Printf("Server listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func handler(w http.ResponseWriter, r *http.Request, s string) {
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

	switch s {
	case "random":
		var response = []string{
			"Why did the tomato turn red? Because it saw the salad dressing!",
			"Why couldn't the bicycle stand up by itself? Because it was two-tired!",
			"Why don't scientists trust atoms? Because they make up everything!",
			"Why did the hipster burn his tongue? He drank his coffee before it was cool.",
			"Why don't seagulls fly by the bay? Because then they'd be bagels.",
			"Why do cows have hooves instead of feet? Because they lactose!",
			"Why did the chicken cross the playground? To get to the other slide!",
			"Why don't skeletons fight each other? They don't have the guts.",
			"Why don't ants get sick? Because they have little ant-bodies.",
		}
		r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		fmt.Fprintln(w, response[r.Intn(len(response))])
	default:
		fmt.Fprintln(w, "👋")
	}
}
