package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

type PruebaRecuentoRepository interface {
	CrearPruebaRecuento(pruebaRecuento *models.PruebaRecuento) error
	ObtenerPruebaRecuentoID(id int) (*models.PruebaRecuento, error)
	ActualizarPruebaRecuento(pruebaRecuento *models.PruebaRecuento) error
	ObtenerPruebasPorProducto(numeroRegistroProducto string) ([]models.PruebaRecuento, error)
	EliminarPruebaRecuento(id int) error
}

type pruebaRecuentoRepository struct {
	DB *gorm.DB
}


// Esta funcion nos permite instanciar el repositorio
// y recibir la base de datos como parametro
func NewPruebaRecuentoRepository(db *gorm.DB) PruebaRecuentoRepository {
	return &pruebaRecuentoRepository{DB: db}
}

// Este metodo nos permite crear un registro de una prueba de recuento en la base de datos
func (repo *pruebaRecuentoRepository) CrearPruebaRecuento(pruebaRecuento *models.PruebaRecuento) error {
	// Se crea el producto y se verifica que no hallan errores
	return repo.DB.Create(&pruebaRecuento).Error
}

func (repo *pruebaRecuentoRepository) ObtenerPruebaRecuentoID(id int) (*models.PruebaRecuento, error) {
	var pruebaRecuento models.PruebaRecuento
	if err := repo.DB.First(&pruebaRecuento, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &pruebaRecuento, nil
}

func (repo *pruebaRecuentoRepository) ActualizarPruebaRecuento(pruebaRecuento *models.PruebaRecuento) error {
	// Se actualiza el producto y se verifica que no hallan errores
	return repo.DB.Save(&pruebaRecuento).Error
}

func (repo *pruebaRecuentoRepository) ObtenerPruebasPorProducto(numeroRegistroProducto string) ([]models.PruebaRecuento, error) {
	var pruebas []models.PruebaRecuento
	if err := repo.DB.Select("nombre_recuento", "tratamiento", "estado").Where("numero_registro_producto = ?", numeroRegistroProducto).Find(&pruebas).Error; err != nil {
		return nil, err
	}
	return pruebas, nil
}

func (repo *pruebaRecuentoRepository) EliminarPruebaRecuento(id int) error {
	var pruebaRecuento models.PruebaRecuento
	if err := repo.DB.First(&pruebaRecuento, "id = ?", id).Error; err != nil {
		return err
	}
	return repo.DB.Delete(&pruebaRecuento).Error
}