package infrastructure

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const dns string = "host=localhost user=postgres password=Davidfelipe2017 dbname=labgestor sslmode=disable"

func NewConexionDB() (*gorm.DB, error) {
    conexion, err := gorm.Open(postgres.Open(dns), &gorm.Config{
        Logger: logger.New(
            log.New(os.Stdout, "\r\n[GORM] ", log.LstdFlags),
            logger.Config{ 
                SlowThreshold:             time.Second,  // SQL que demoran más de 1s
                LogLevel:                  logger.Info,   // Nivel de log (Error, Warn, Info)
                IgnoreRecordNotFoundError: true,          // Ignorar "record not found"
                Colorful:                  true,          // Colores en terminal
            },
        ),
    })
    return conexion, err
}