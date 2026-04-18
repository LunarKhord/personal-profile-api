package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
)

type MeResponse struct {
    Name   string `json:"name"`
    Email  string `json:"email"`
    Github string `json:"github"`
}

func landingHandle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "API is running",
    })
}

func healthHandle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "healthy",
    })
}

func meHandle(w http.ResponseWriter, r *http.Request) {
    profile := MeResponse{
        Name:   "Muhammad Hasim",
        Email:  "krazygenus@gmail.com",
        Github: "https://github.com/LunarKhord/",    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(profile); err != nil {
        log.Printf("failed to encode response: %v", err)
    }
}

func main() {
    log.Println("Server started on :8080...")

    http.HandleFunc("GET /", landingHandle)
    http.HandleFunc("GET /health", healthHandle)
    http.HandleFunc("GET /me", meHandle)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Fatal(http.ListenAndServe(":"+port, nil))
}