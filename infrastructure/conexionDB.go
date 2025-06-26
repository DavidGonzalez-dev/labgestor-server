package infrastructure

import (
	// "fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConexionDB() (*gorm.DB, error) {
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_DATABASE")

	// dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbName)
	dns := "host=localhost user=postgres password=Davidfelipe2017 dbname=labgestor sslmode=disable"

	conexion, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n[GORM] ", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second, // SQL que demoran más de 1s
				LogLevel:                  logger.Info, // Nivel de log (Error, Warn, Info)
				IgnoreRecordNotFoundError: true,        // Ignorar "record not found"
				Colorful:                  true,        // Colores en terminal
			},
		),
	})
	return conexion, err
}
