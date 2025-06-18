package rdbms

import (
	"database/sql"
	"fmt"

	"github.com/achmadnr21/cinema/config"

	_ "github.com/lib/pq" // Import driver PostgreSQL
)

var pgdatabase *sql.DB

// InitDB menginisialisasi koneksi database
func InitPG(Dbconf config.Database) error {
	var err error
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		Dbconf.Host, Dbconf.Port, Dbconf.Username, Dbconf.Password, Dbconf.Database, Dbconf.SSLMode,
	)

	pgdatabase, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	// Cek koneksi database
	if err = pgdatabase.Ping(); err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	return nil
}

// GetDB mengembalikan instance database
func GetPG() *sql.DB {
	return pgdatabase
}
