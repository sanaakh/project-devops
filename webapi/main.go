package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	TeamName string `json:"team_name"`
}

func whoAmIHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{TeamName: "Team DevOps : Ilyes, Maxime, Lydia et Sanaa :) "}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/whoami", whoAmIHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
