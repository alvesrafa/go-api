package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"alvesrafa.dev/study/models"
)

func List(w http.ResponseWriter, r *http.Request) {

	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Error: erro ao obter todos os registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}
