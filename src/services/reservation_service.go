package services

import (
	"database/sql"
	"log"
	"restaurant-reservation-management/src/models"
	"restaurant-reservation-management/src/queries"
)

type ReservationRepository interface {
	GetReservationReport() ([]models.ReservationReport, error)
}

type ReservationStorage struct {
	DB *sql.DB
}

func NewReservationService(db *sql.DB) *ReservationStorage {
	return &ReservationStorage{DB: db}
}

func (s *ReservationStorage) GetReservationReport() ([]models.ReservationReport, error) {
	rows, err := s.DB.Query(queries.ReservationReportQuery)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			log.Println(closeErr)
		}
	}()

	reservationsChan := make(chan models.ReservationReport)
	errorCh := make(chan error)

	go s.processReservations(rows, reservationsChan, errorCh)

	return s.validateChannels(reservationsChan, errorCh)
}

func (s *ReservationStorage) processReservations(rows *sql.Rows, reservationsChan chan models.ReservationReport, errorCh chan error) {
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
			return nil, err
		}
	}
}
