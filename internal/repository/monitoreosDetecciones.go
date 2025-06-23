package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

type MonitoreosDeteccionRepository interface {
	CrearMonitoreosDetecciones(monitoreosDeteccion *models.MonitoreosDeteccionesMicroorganismo) error
	ObtenerMonitoreosDeteccionesID(id int) (models.MonitoreosDeteccionesMicroorganismo, error)
	ObtenerMonitoreosDeteccionesPorDeteccion(idDeteccionMicroorganismos int) (*[]models.MonitoreosDeteccionesMicroorganismo, error)
	ActualizarMonitoreosDetecciones(monitoreosDeteccion *models.MonitoreosDeteccionesMicroorganismo) error
	EliminarMonitoreosDetecciones(monitoreosDeteccion *models.MonitoreosDeteccionesMicroorganismo) error
}

type monitoreosDeteccionRepository struct {
	DB *gorm.DB
}

func NewMonitoreosDeteccionesRepository(db *gorm.DB) MonitoreosDeteccionRepository {
	return &monitoreosDeteccionRepository{DB: db}
}

func (repo *monitoreosDeteccionRepository) CrearMonitoreosDetecciones(monitoreosDeteccion *models.MonitoreosDeteccionesMicroorganismo) error {
	return repo.DB.Create(&monitoreosDeteccion).Error
}

func (repo *monitoreosDeteccionRepository) ObtenerMonitoreosDeteccionesID(id int) (models.MonitoreosDeteccionesMicroorganismo, error) {
	var monitoreosDeteccion models.MonitoreosDeteccionesMicroorganismo
	if err := repo.DB.First(&monitoreosDeteccion, id).Error; err != nil {
		return models.MonitoreosDeteccionesMicroorganismo{}, err
	}
	return monitoreosDeteccion, nil

}

func (repo *monitoreosDeteccionRepository) ObtenerMonitoreosDeteccionesPorDeteccion(idDeteccionMicroorganismos int) (*[]models.MonitoreosDeteccionesMicroorganismo, error) {
	var monitoreosMicroorganismo *[]models.MonitoreosDeteccionesMicroorganismo
	if err := repo.DB.Where("id_deteccion_microorganismo = ?", idDeteccionMicroorganismos).Find(&monitoreosMicroorganismo).Error; err != nil {
		return nil, err
	}
	return monitoreosMicroorganismo, nil
}

func (repo *monitoreosDeteccionRepository) ActualizarMonitoreosDetecciones(monitoreosDeteccion *models.MonitoreosDeteccionesMicroorganismo) error {
	return repo.DB.Save(&monitoreosDeteccion).Error
}

func (repo *monitoreosDeteccionRepository) EliminarMonitoreosDetecciones(monitoreosDeteccion *models.MonitoreosDeteccionesMicroorganismo) error {
	return repo.DB.Delete(&monitoreosDeteccion).Error
}
