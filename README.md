# Restaurant Reservation Management System

This is a restaurant reservation management system designed to help restaurants manage their table reservations
efficiently.

## Features

- **Reservation Reporting**: Provides reports on reservation patterns, such as the frequency of reservations by day,
  hour, and party size, to help optimize table distribution and improve customer experience.
- **Concurrency**: Utilizes concurrency to process reservation data efficiently and handle multiple requests
  concurrently.
- **Database**: Stores reservation data in a PostgreSQL database to ensure data integrity and reliability.

## Installation

To install and run the reservation management system locally, follow these steps:

1. Clone the repository:

   ```bash
   git clone git@github.com:JP-Cano/restaurant-reservations-management.git
   ```

2. Install dependencies:

   ```bash
   cd restaurant-reservation-management
   go mod tidy
   ```

3. Set up the PostgreSQL database:

    - Install PostgreSQL and create a new database.
    - Configure the database connection in the `.env` file.

4. Start the server:

   ```bash
   go run main.go
   ```

5. Access the application
   at [https://logical-fort-425301-p0.uc.r.appspot.com/reservation-report](https://logical-fort-425301-p0.uc.r.appspot.com/reservation-report).

## Usage

1. **Generating Reports**: Retrieve reservation reports using the API endpoint to analyze reservation patterns and
   optimize table distribution.
