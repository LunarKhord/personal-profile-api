package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"net/http"
)



type MeResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Github string `json:"github"`
}

type Response struct {
	Message string `json:"message"`
}


func landingHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
            "message":"healthy",
        })

}


func healthHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
            "message":"healthy",
        })
}

func meHandle(w http.ResponseWriter, r *http.Request) {
	var profile MeResponse
	profile.Name = "Muhammad Hasim"
	profile.Email = "krazygenus@gmail.com"
	profile.Github = "https://github.com/LunarKhord"
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		fmt.Println(err)

	}
}



func main() {
    fmt.Println("Server started on :8080...")



        http.HandleFunc("GET /", landingHandle)
        http.HandleFunc("GET /health", healthHandle)
        http.HandleFunc("GET /me", meHandle)


        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }
        log.Fatal(http.ListenAndServe(":"+port, nil))
	}