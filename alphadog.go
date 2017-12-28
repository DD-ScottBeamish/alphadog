package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

// Initialize Counter
var count=0

type Counter struct {
     Count int `json:"count,omitempty"`
}

// Health Check
type Health struct {
     Health string `json:"health,omitempty"`
}

// Encode Health Check as json
func GetHealthCheck(w http.ResponseWriter, r *http.Request)
    json.NewEncoder(w).Encode(Health{"ok"})
}

// Encode count as json
func GetCount(w http.ResponseWriter, r *http.Request) {
    count++
    json.NewEncoder(w).Encode(Counter{count})
}

// main function
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/getcount", GetCount).Methods("GET")
    router.HandleFunc("/healthcheck", GetHealthCheck).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}
