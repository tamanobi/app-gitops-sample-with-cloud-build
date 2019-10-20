package main

import (
	"log"
	"net/http"
)

func Hello() string {
	return "Hello"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(Hello()))
	})
	log.Fatal(http.ListenAndServe(":1333", nil))
}
