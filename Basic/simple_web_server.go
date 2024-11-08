package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "hello")
}

func mai1n() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
