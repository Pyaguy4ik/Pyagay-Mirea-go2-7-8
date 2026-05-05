package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("TASKS_PORT")
    if port == "" {
        port = "8082"
    }

    mux := http.NewServeMux()

    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        json.NewEncoder(w).Encode(map[string]string{
            "status":  "ok",
            "service": "tasks",
        })
    })

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Tasks service is running",
            "version": "1.0",
        })
    })

    addr := ":" + port
    log.Println("tasks service started on", addr)

    if err := http.ListenAndServe(addr, mux); err != nil {
        log.Fatal(err)
    }
}
