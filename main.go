package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Running")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, World! This is Tony here :)"))
		if err != nil {
			log.Println("could not write response to client", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
