package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Response struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Разрешаем только POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Парсим тело формы
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Собираем все параметры в текст
	var params []string
	for k, v := range r.Form {
		params = append(params, fmt.Sprintf("%s=%s", k, strings.Join(v, ",")))
	}

	resp := Response{
		ResponseType: "ephemeral",
		Text:         fmt.Sprintf("👋 Hello from Mattermost slash command! @qwerty \nReceived params:\n%s", strings.Join(params, "\n")),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Listening on :8083")
	http.ListenAndServe(":8083", nil)
}
