package queries

const ReservationReportQuery = `
        SELECT 
            DAYNAME(reservation_time) AS day_of_week,
            HOUR(reservation_time) AS hour_of_day,
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
