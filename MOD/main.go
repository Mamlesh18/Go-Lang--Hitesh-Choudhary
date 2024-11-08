package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("HEY I AM NEW")
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	http.ListenAndServe(":8080", r)

}

func serverHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Hey i am sending data</h1>"))
}
