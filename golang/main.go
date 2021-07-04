package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUwa(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("uwa")
}

func main() {
	fmt.Println("lets go")
	r := mux.NewRouter()
	r.HandleFunc("/v1/", getUwa).Methods("GET")
	log.Fatal(http.ListenAndServe(":4201", r))
}
