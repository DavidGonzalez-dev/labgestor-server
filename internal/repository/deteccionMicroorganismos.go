package repository

import (
<<<<<<< HEAD
	"errors"
=======
>>>>>>> 74d06f274fce1dc6d8349036cda6921cc3a8c6b6
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
	
	if err := repo.DB.Create(&deteccion).Error ;err != nil {
		return err
	}
	
	return nil
}

// Este metodo nos permite obtener el registro de una deteccion de microorganismo basado en un id
func (repo *deteccionMicroorganismosRepository) ObtenerDeteccionMicroorganismosID(id int) (*models.DeteccionesMicroorganismos, error) {

	var deteccionMicroorganismos models.DeteccionesMicroorganismos

	if err := repo.DB.First(&deteccionMicroorganismos, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no existe un registro de deteccion de microorganismo con este id")
		}
		return nil, err
	}
	return &deteccionMicroorganismos, nil
}

// Este metodo nos permite actualizar el registro de una deteccion de microorganismo 
func (repo *deteccionMicroorganismosRepository) ActualizarDeteccionMicroorganismos(deteccion *models.DeteccionesMicroorganismos) error {
	return repo.DB.Save(&deteccion).Error
}

// Este metodo nos permite obtener informacion superficial acerca de todas las detecciones de microorganismos de un producto
func (repo *deteccionMicroorganismosRepository) ObtenerDeteccionMicroorganismosPorProducto(numeroRegistroProducto string) ([]models.DeteccionesMicroorganismos, error) {
	
	var detecciones []models.DeteccionesMicroorganismos

	result := repo.DB.Select("id", "nombre_microorganismo", "tratamiento", "estado").Where("numero_registro_producto = ?", numeroRegistroProducto).Find(&detecciones)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return detecciones, nil
}

// Este metodo nos permite eliminar un registro de deteccion de microorganismo segun su id
func (repo *deteccionMicroorganismosRepository) EliminarDeteccionMicroorganismos(id int) error {
	var deteccionMicroorganismos models.DeteccionesMicroorganismos
	if err := repo.DB.First(&deteccionMicroorganismos, "id = ?", id).Error; err != nil {
		return err
	}
	return repo.DB.Delete(&deteccionMicroorganismos).Error
}
