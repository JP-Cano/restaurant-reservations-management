package services

import (
	"github.com/jackc/pgx/v5"
	"log"
	"restaurant-reservation-management/src/database"
	"restaurant-reservation-management/src/models"
	"restaurant-reservation-management/src/queries"
)

type ReservationRepository interface {
	GetReservationReport() ([]models.ReservationReport, error)
}

type ReservationStorage struct {
	DB *pgx.Conn
}

func NewReservationService(db *pgx.Conn) *ReservationStorage {
	return &ReservationStorage{DB: db}
}

func (s *ReservationStorage) GetReservationReport() ([]models.ReservationReport, error) {
	rows, err := s.DB.Query(database.BCGContext, queries.ReservationReportQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reservationsChan := make(chan models.ReservationReport)
	errorCh := make(chan error)

	go s.processReservations(rows, reservationsChan, errorCh)

	return s.validateChannels(reservationsChan, errorCh)
}

func (s *ReservationStorage) processReservations(rows pgx.Rows, reservationsChan chan models.ReservationReport, errorCh chan error) {
	defer close(reservationsChan)
	for rows.Next() {
		var reservationReport models.ReservationReport
		if err := rows.Scan(&reservationReport.DayOfWeek, &reservationReport.HourOfDay, &reservationReport.NumDiners, &reservationReport.ReservationCount); err != nil {
			errorCh <- err
			return
		}
		reservationsChan <- reservationReport
	}
	if err := rows.Err(); err != nil {
		errorCh <- err
	}
}

func (s *ReservationStorage) validateChannels(reservationsChan chan models.ReservationReport, errorCh chan error) ([]models.ReservationReport, error) {
	var reservationReports []models.ReservationReport
	for {
		select {
		case report, ok := <-reservationsChan:
			if !ok {
				return reservationReports, nil
			}
			reservationReports = append(reservationReports, report)
		case err := <-errorCh:
			log.Println(err)
			close(errorCh)
			return nil, err
		}
	}
}
