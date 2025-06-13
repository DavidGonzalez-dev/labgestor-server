package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConexionDB() (*gorm.DB, error) {

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", dbHost, dbUser, dbPassword, dbName, dbPort)

	conexion, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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
