package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

type ControlesNegativosRepository interface {
	CrearControlesNegativos(controlesNegativos *models.ControlesNegativosMedio) error
}

type controlesNegativosRepository struct {
	DB *gorm.DB
}

func NewControlesNegativosRepository(db *gorm.DB) ControlesNegativosRepository {
	return &controlesNegativosRepository{DB: db}
}

// Este metodo nos permite crear un registro de una prueba de recuento en la base de datos
func (repo *controlesNegativosRepository) CrearControlesNegativos(controlesNegativos *models.ControlesNegativosMedio) error {
	// Se crea el producto y se verifica que no hallan errores
	return repo.DB.Create(&controlesNegativos).Error
}
