package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

type CajasBioburdenRepository interface {
	CrearCajaBioburden(cajaBioburden *models.CajasBioburden) error
	ObtenerCajasBioburdenID(id int) (*models.CajasBioburden, error)
	ObtenerCajasPorPruebaRecuento(idPruebaRecuento int) ([]models.CajasBioburden, error)
	ActualizarCajaBioburden(cajaBioburden *models.CajasBioburden) error
	EliminarCajaBioburden(id int) error
}

type cajasBioburdenRepository struct {
	DB *gorm.DB
}

func NewCajasBioburdenRepository(db *gorm.DB) CajasBioburdenRepository {
	return &cajasBioburdenRepository{DB: db}
}

// CrearCajaBioburden crea un nuevo registro de caja de bioburden en la base de datos
func (repo *cajasBioburdenRepository) CrearCajaBioburden(cajaBioburden *models.CajasBioburden) error {
	return repo.DB.Create(&cajaBioburden).Error
}

func (repo *cajasBioburdenRepository) ObtenerCajasBioburdenID(id int) (*models.CajasBioburden, error) {
	var cajaBioburden models.CajasBioburden
	// Se busca el registro de caja de bioburden por ID
	if err := repo.DB.First(&cajaBioburden, id).Error; err != nil {
		return nil, err
	}
	return &cajaBioburden, nil

}

func (repo *cajasBioburdenRepository) ObtenerCajasPorPruebaRecuento(idPruebaRecuento int) ([]models.CajasBioburden, error) {
	var cajas []models.CajasBioburden
	// Se busca el registro de caja de bioburden por ID de prueba de recuento
	if err := repo.DB.Select("id, tipo, metodo_siembra, resultado, medida_aritmetica").Where("id_prueba_recuento = ?", idPruebaRecuento).Find(&cajas).Error; err != nil {
		return nil, err
	}
	return cajas, nil
}

func (repo *cajasBioburdenRepository) ActualizarCajaBioburden(cajaBioburden *models.CajasBioburden) error {
	return repo.DB.Save(&cajaBioburden).Error
}

func (repo *cajasBioburdenRepository) EliminarCajaBioburden(id int) error {
	var cajaBioburden models.CajasBioburden
	if err := repo.DB.First(&cajaBioburden, "id = ?", id).Error; err != nil {
		return err
	}
	return repo.DB.Delete(&cajaBioburden).Error
}
