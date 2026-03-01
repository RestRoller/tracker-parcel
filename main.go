package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    // Простое HTTP-приложение для демонстрации
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Parcel Tracker App - Version %s\n", getVersion())
        fmt.Fprintf(w, "Service is running!\n")
    })

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "OK")
    })

    port := getPort()
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getVersion() string {
    return "1.0.2"
}

func getPort() string {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    return port
}
