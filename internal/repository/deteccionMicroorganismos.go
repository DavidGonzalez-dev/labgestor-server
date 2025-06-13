package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

// Interfaz que define los metodos que debe implementar el repositorio de deteccion de microorganismos
type DeteccionMicroorganismosRepository interface {
	CrearDeteccionMicroorganismos(deteccion *models.DeteccionesMicroorganismos) error
	ObtenerDeteccionMicroorganismosID(id int) (*models.DeteccionesMicroorganismos, error)
	ActualizarDeteccionMicroorganismos(deteccion *models.DeteccionesMicroorganismos) error
	ObtenerDeteccionMicroorganismosPorProducto(numeroRegistroProducto string) ([]models.DeteccionesMicroorganismos, error)
	EliminarDeteccionMicroorganismos(id int) error
}

type deteccionMicroorganismosRepository struct {
	DB *gorm.DB
}

func NewDeteccionMicroorganismosRepository(db *gorm.DB) DeteccionMicroorganismosRepository {
	return &deteccionMicroorganismosRepository{DB: db}
}

// Este metodo nos permite crear un registro de deteccion de microorganismos en la base de datos
func (repo *deteccionMicroorganismosRepository) CrearDeteccionMicroorganismos(deteccion *models.DeteccionesMicroorganismos) error {
	// Se crea el registro y se verifica que no hallan errores
	return repo.DB.Create(&deteccion).Error
}

func (repo *deteccionMicroorganismosRepository) ObtenerDeteccionMicroorganismosID(id int) (*models.DeteccionesMicroorganismos, error) {
	// Se crea el registro y se verifica que no hallan errores
	var deteccionMicroorganismos models.DeteccionesMicroorganismos
	if err := repo.DB.First(&deteccionMicroorganismos, id).Error; err != nil {
		return nil, err
	}
	return &deteccionMicroorganismos, nil
}

func (repo *deteccionMicroorganismosRepository) ActualizarDeteccionMicroorganismos(deteccion *models.DeteccionesMicroorganismos) error {
	// Se actualiza el registro y se verifica que no hallan errores
	return repo.DB.Save(&deteccion).Error
}

func (repo *deteccionMicroorganismosRepository) ObtenerDeteccionMicroorganismosPorProducto(numeroRegistroProducto string) ([]models.DeteccionesMicroorganismos, error) {
	var detecciones []models.DeteccionesMicroorganismos
	if err := repo.DB.Select("id", "nombre_microorganismo", "tratamiento", "estado").Where("numero_registro_producto = ?", numeroRegistroProducto).Find(&detecciones).Error; err != nil {
		return nil, err
	}
	return detecciones, nil
}

func (repo *deteccionMicroorganismosRepository) EliminarDeteccionMicroorganismos(id int) error {
	var deteccionMicroorganismos models.DeteccionesMicroorganismos
	if err := repo.DB.First(&deteccionMicroorganismos, "id = ?", id).Error; err != nil {
		return err
	}
	return repo.DB.Delete(&deteccionMicroorganismos).Error
}
