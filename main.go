package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, World! This is Tony here :)"))
		if err != nil {
			log.Println("could not write response to client", err)
		}
	})
	log.Println("Running")
	log.Fatal(http.ListenAndServe(":80", nil))
}
