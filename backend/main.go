package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Message struct {
    Text string `json:"text"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Message{Text: "Hello from Go backend!"})
}

func main() {
    http.HandleFunc("/api/hello", helloHandler)

    log.Println("✅ Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
