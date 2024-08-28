package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

var responses []Response

func web(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w)
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodDelete:
		handleDelete(w, r)
	case http.MethodPut:
		handlePut(w, r)
	}
}

func handleGet(w http.ResponseWriter) {
	json, _ := json.Marshal(responses)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var requestData Response

	json.NewDecoder(r.Body).Decode(&requestData)

	requestData.Id = len(responses)

	responses = append(responses, requestData)

	jsonResponse, _ := json.Marshal(responses)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	var requestData int
	json.NewDecoder(r.Body).Decode(&requestData)

	if requestData < 0 || requestData >= len(responses) {
		http.Error(w, "index out of range", http.StatusBadRequest)
	}

	for i := 0; i < len(responses); i++ {
		if requestData == i {
			responses = append(responses[:i], responses[(i+1):]...)
			i--
		}
	}
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	var requestData Response
	json.NewDecoder(r.Body).Decode(&requestData)
	if requestData.Id < len(responses) && requestData.Id >= 0 {
		responses[requestData.Id] = requestData
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responses[requestData.Id])
	}
}

func main() {
	fmt.Printf("Hello, world!")
	http.HandleFunc("/hello", web)
	http.ListenAndServe(":8000", nil)
}
