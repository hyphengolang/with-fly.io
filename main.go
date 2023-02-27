package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// read query param for name
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}

		// write response
		w.Write([]byte("Hello, " + name + "!\n"))
	})

	http.ListenAndServe(":"+port, nil)
}
