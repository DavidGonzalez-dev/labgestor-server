package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dns string = "host=localhost user=postgres password=password dbname=labgestor sslmode=disable"

func NewConexionDB() (*gorm.DB, error) {
	conexion, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	return conexion, err
}
