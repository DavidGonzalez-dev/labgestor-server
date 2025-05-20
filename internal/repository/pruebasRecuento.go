package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

type PruebaRecuentoRepository interface {
	CrearPruebaRecuento(pruebaRecuento *models.PruebaRecuento) error
	ObtenerPruebaRecuento(numeroRegistroProducto string) (*models.PruebaRecuento, error)
	ActualizarPruebaRecuento(pruebaRecuento *models.PruebaRecuento) error
}

type pruebaRecuentoRepository struct {
	DB *gorm.DB
}

func NewPruebaRecuentoRepository(db *gorm.DB) PruebaRecuentoRepository {
	return &pruebaRecuentoRepository{DB: db}
}

// Este metodo nos permite crear un registro de una prueba de recuento en la base de datos
func (repo *pruebaRecuentoRepository) CrearPruebaRecuento(pruebaRecuento *models.PruebaRecuento) error {
	// Se crea el producto y se verifica que no hallan errores
	return repo.DB.Create(&pruebaRecuento).Error
}

func (repo *pruebaRecuentoRepository) ObtenerPruebaRecuento(numeroRegistroProducto string) (*models.PruebaRecuento, error) {
	var pruebaRecuento models.PruebaRecuento
	if err := repo.DB.Where("numero_registro_producto = ?", numeroRegistroProducto).First(&pruebaRecuento).Error; err != nil {
		return nil, err
	}
	return &pruebaRecuento, nil
}

func (repo *pruebaRecuentoRepository) ActualizarPruebaRecuento(pruebaRecuento *models.PruebaRecuento) error {
	// Se actualiza el producto y se verifica que no hallan errores
	return repo.DB.Save(&pruebaRecuento).Error
}
