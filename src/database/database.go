package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

type PsqlStorage struct {
	Conn *pgx.Conn
}

var BCGContext = context.Background()

func NewPsqlStorage(url string) *PsqlStorage {
	conn, err := pgx.Connect(BCGContext, url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = conn.Ping(BCGContext)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Database connected successfully")
	return &PsqlStorage{Conn: conn}
}

func (p *PsqlStorage) Close() {
	err := p.Conn.Close(BCGContext)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error closing database connection: %v\n", err)
	}
}
