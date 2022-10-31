package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kenfav/cong-reports-go/configs"
	"github.com/kenfav/cong-reports-go/handler"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Post("/", handler.Create)
	r.Put("/{id}", handler.Update)
	r.Delete("/{id}", handler.Delete)
	r.Get("/", handler.List)
  r.Get("/{id}/reports/{year}/{month}", handler.GetReportByMonth)
  r.Get("/{id}/reports", handler.ListAllReports)
	r.Get("/{id}", handler.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
