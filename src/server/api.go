package server

import (
	"database/sql"
	"log"
	"net/http"
	"restaurant-reservation-management/src/controllers"
	"restaurant-reservation-management/src/middlewares"
	"restaurant-reservation-management/src/services"
)

type Server struct {
	Addr string
	DB   *sql.DB
}

func New(addr string, db *sql.DB) *Server {
	return &Server{
		Addr: addr,
		DB:   db,
	}
}

func (s *Server) Serve() error {
	router := http.NewServeMux()
	s.registerRoutes(router)

	server := &http.Server{
		Addr:    s.Addr,
		Handler: middlewares.RequestLoggerMiddleware(router),
	}

	log.Printf("Server running on port %s", s.Addr)
	return server.ListenAndServe()
}

func (s *Server) registerRoutes(router *http.ServeMux) {
	reservationController := s.registerDependencies()
	router.HandleFunc("/reservation-report", reservationController.GetReservationReport)
}

func (s *Server) registerDependencies() *controllers.ReservationController {
	reservationService := services.NewReservationService(s.DB)
	reservationController := controllers.NewReservationController(reservationService)
	return reservationController
}
