package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", handler)
	s := http.Server{
		Addr:    ":8989",
		Handler: m,
	}
	log.Fatal(s.ListenAndServe())
}
