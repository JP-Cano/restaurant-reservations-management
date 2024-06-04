package models

type ReservationReport struct {
	DayOfWeek        string `json:"day_of_week"`
	HourOfDay        int    `json:"hour_of_day"`
	NumDiners        int    `json:"num_diners"`
	ReservationCount int    `json:"reservation_count"`
}
