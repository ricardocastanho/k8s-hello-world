package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	fmt.Println("server running on http://localhost")

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
