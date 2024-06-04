package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"restaurant-reservation-management/src/database"
	"restaurant-reservation-management/src/server"
)

func main() {
	godotenv.Load(".env")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbUser, url.QueryEscape(dbPassword), dbHost, dbPort, dbName)

	db := database.NewPsqlStorage(dsn)
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s := server.New(fmt.Sprintf(":%s", port), db.Conn)
	if err := s.Serve(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
