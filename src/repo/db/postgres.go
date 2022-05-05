package postgres

import (
	"fmt"
	"os"

	"github/robotxt/iie-app/src/logging"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// PostgresCon database connection
type PostgresCon struct {
	DB *gorm.DB
}

// Connect to mysql db
func (db *PostgresCon) Connect() (*gorm.DB, error) {

	dbinfo := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	logging.Info(dbinfo)
	conn, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		logging.Fatal(err)
		logging.Fatal("Could not connect database")
	}
	db.DB = conn
	return conn, err
}
