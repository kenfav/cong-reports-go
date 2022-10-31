package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kenfav/cong-reports-go/models"
)

func GetReportByMonth(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	year, err := strconv.Atoi(chi.URLParam(r, "year"))
	month, err := strconv.Atoi(chi.URLParam(r, "month"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.GetReportByMonth(int64(id), int64(year), int64(month))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rows)

}
