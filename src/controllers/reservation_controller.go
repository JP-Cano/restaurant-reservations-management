package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurant-reservation-management/src/services"
)

type ReservationController struct {
	Service services.ReservationRepository
}

func NewReservationController(service services.ReservationRepository) *ReservationController {
	return &ReservationController{Service: service}
}

func (c *ReservationController) GetReservationReport(w http.ResponseWriter, r *http.Request) {
	reports, err := c.Service.GetReservationReport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(reports)
	if err != nil {
		log.Fatal(err)
	}
}
