package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = "0.0.0.0:8000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("https://%s%s", r.Host, r.URL.Path)
		http.Redirect(w, r, url, http.StatusPermanentRedirect)
	})

	log.Fatal(http.ListenAndServe(addr, nil))
}
