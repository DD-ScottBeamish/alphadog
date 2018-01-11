package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "fmt"
    "strconv"
    ddtrace "github.com/DataDog/dd-trace-go/opentracing"
    opentracing "github.com/opentracing/opentracing-go"
    logs "github.com/opentracing/opentracing-go/log"
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
func GetHealthCheck(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Health{"ok"})
}

// Encode count as json
func GetCount(w http.ResponseWriter, r *http.Request) {
    // Add Datadog Tracing
    span := opentracing.StartSpan("/getcount")
    defer span.Finish()
    
    // bump/return/log container count
    count++
    fmt.Println("Container Count: " + strconv.Itoa(count))
    json.NewEncoder(w).Encode(Counter{count})

    span.LogFields(logs.Int("count", count))
}

// main function
func main() {
    // create a Tracer configuration
    config := ddtrace.NewConfiguration()
    config.ServiceName = "alphadog"
    config.AgentHostname = "dd-agent"

    // initialize a Tracer and ensure a graceful shutdown
    // using the `closer.Close()`
    tracer, closer, err := ddtrace.NewTracer(config)
    if err != nil {
        // handle the configuration error
    }
    defer closer.Close()

    // set the Datadog tracer as a GlobalTracer
    opentracing.SetGlobalTracer(tracer)

    router := mux.NewRouter()
    router.HandleFunc("/getcount", GetCount).Methods("GET")
    router.HandleFunc("/healthcheck", GetHealthCheck).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", router))
}
