package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		ResponseType: "ephemeral",
		Text:         "👋 Hello from Mattermost slash command! @qwerty",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Listening on :8083")
	http.ListenAndServe(":8080", nil)
}
