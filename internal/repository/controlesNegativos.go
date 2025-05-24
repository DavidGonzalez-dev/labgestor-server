package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

type ControlesNegativosRepository interface {
	CrearControlesNegativos(controlesNegativos *models.ControlesNegativosMedio) error
	ObtenerControlesNegativosID(id int) (*models.ControlesNegativosMedio, error)
	ObtenerControlesPorProducto(NumeroRegistroProducto string) ([]models.ControlesNegativosMedio, error)
	ActualizarControlesNegativos(controlesNegativos *models.ControlesNegativosMedio) error
	EliminarControlesNegativos(id int) error
}

type controlesNegativosRepository struct {
	DB *gorm.DB
}

func NewControlesNegativosRepository(db *gorm.DB) ControlesNegativosRepository {
	return &controlesNegativosRepository{DB: db}
}

// Este metodo nos permite crear un registro de los controles negativos en la base de datos
func (repo *controlesNegativosRepository) CrearControlesNegativos(controlesNegativos *models.ControlesNegativosMedio) error {
	// Se crea el producto y se verifica que no hallan errores
	return repo.DB.Create(&controlesNegativos).Error
}

// Este metodo nos permite obtener un registro de los controles negativos en la base de datos
func (repo *controlesNegativosRepository) ObtenerControlesNegativosID(NumeroRegistroProducto int) (*models.ControlesNegativosMedio, error) {
	// Se crea el producto y se verifica que no hallan errores
	var controlesNegativos models.ControlesNegativosMedio
	if err := repo.DB.First(&controlesNegativos, NumeroRegistroProducto).Error; err != nil {
		return nil, err
	}
	return &controlesNegativos, nil
}

// Este metodo nos permite obtener los registro de los controles negativos por producto en la base de datos
func (repo *controlesNegativosRepository) ObtenerControlesPorProducto(id string) ([]models.ControlesNegativosMedio, error) {

	var controlesNegativos []models.ControlesNegativosMedio
	if err := repo.DB.Where("numero_registro_producto = ?", id).Find(&controlesNegativos).Error; err != nil {
		return nil, err
	}
	return controlesNegativos, nil
}

// Este metodo nos permite actualizar un registro de los controles negativos en la base de datos
func (repo *controlesNegativosRepository) ActualizarControlesNegativos(controlesNegativos *models.ControlesNegativosMedio) error {
	// Se crea el producto y se verifica que no hallan errores
	return repo.DB.Save(&controlesNegativos).Error
}

func (repo *controlesNegativosRepository) EliminarControlesNegativos(id int) error {
	// Se crea el producto y se verifica que no hallan errores
	return repo.DB.Delete(&models.ControlesNegativosMedio{}, id).Error

}
