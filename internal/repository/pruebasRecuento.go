package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)


// En esta interfaz se definene los metodos del repositorio
type PruebaRecuentoRepository interface {
	CrearPruebaRecuento(pruebaRecuento *models.PruebaRecuento) error
}

type pruebaRecuentoRepository struct {
	DB *gorm.DB
}

// Esta funcion nos permite instanciar el repositorio
// y recibir la base de datos como parametro
func NewPruebaRecuentoRepository(db *gorm.DB) PruebaRecuentoRepository {
	return &pruebaRecuentoRepository{DB: db}
}


// ? ------------------------------------------------
// ? CRUD PRUEBAS DE RECUEENTO
// ? ------------------------------------------------

// Este metodo nos permite crear un registro de una prueba de recuento en la base de datos
func (repo *pruebaRecuentoRepository) CrearPruebaRecuento(pruebaRecuento *models.PruebaRecuento) error {
	// Se crea el producto y se retorna el error
	// en caso de que ocurra un error
	return repo.DB.Create(&pruebaRecuento).Error
}
