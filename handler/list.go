package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kenfav/cong-reports-go/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	publicadores, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter os registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publicadores)
}
