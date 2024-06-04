package queries

const ReservationReportQuery = `
    SELECT 
        TO_CHAR(reservation_time, 'Day') AS day_of_week,
        EXTRACT(HOUR FROM reservation_time) AS hour_of_day,
        num_diners,
        COUNT(*) AS reservation_count
    FROM 
        reservations
    GROUP BY 
        day_of_week, 
        hour_of_day, 
        num_diners
    ORDER BY 
        day_of_week, 
        hour_of_day, 
        num_diners;
`
