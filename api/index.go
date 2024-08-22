package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You SUCCESS DEPLOY GOLANG page 1</h1>")
}

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler).Methods("GET")

	// http.ListenAndServe(":8080", r)
}

// NOW I WANT TO CREATE NEW ROUTE / and /dua

// YOU MUST CREATE FILE FOR / for index.go and /dua for dua.go
