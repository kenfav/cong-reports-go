package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kenfav/cong-reports-go/models"
)

func ListAllReports(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	reports, err := models.ReportsGetAll(int64(id))
	if err != nil {
		log.Printf("Erro ao obter os registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
