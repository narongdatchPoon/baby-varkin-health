package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RequestBody struct {
	Msg string `json:"msg"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody

	// Parse the JSON body into the struct
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Use the parsed data
	fmt.Fprintf(w, "<h1>Received: %s</h1>", reqBody.Msg)

}

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler).Methods("POST")

}
