package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

// Interfaz que define los metodos que se emplean en la tabla de los fabricante en la base datos
type FabricanteRepository interface {
	CrearFabricante(fabricante *models.Fabricante) error
	ActualizarFabricante(fabricante *models.Fabricante) error
	ObtenerFabricante(ID string) (*models.Fabricante, error)
	ObtenerFabricantes() (*[]models.Fabricante, error)
}

// Structura que implementa la interfaz anteriormente definida
type fabricanteRepository struct {
	DB *gorm.DB
}

// Funcion que devuelve una instancia de la estructura fabricnateepository
func NewFabricanterepository(db *gorm.DB) FabricanteRepository {
	return &fabricanteRepository{DB: db}
}

// ---------------------------
// Metodos de la estructura
// ---------------------------

func (repo *fabricanteRepository) CrearFabricante(fabricante *models.Fabricante) error {
	return repo.DB.Create(&fabricante).Error
}

func (repo *fabricanteRepository) ActualizarFabricante(fabricante *models.Fabricante) error {
	return repo.DB.Save(&fabricante).Error
}

func (repo *fabricanteRepository) ObtenerFabricante(ID string) (*models.Fabricante, error) {
	var fabricante models.Fabricante

	// Realizamos la consulta utilizando el valor del ID como parámetro
	if err := repo.DB.First(&fabricante, ID).Error; err != nil {
		return nil, err
	}

	return &fabricante, nil
}

func (repo *fabricanteRepository) ObtenerFabricantes() (*[]models.Fabricante, error) {
	var fabricantes []models.Fabricante
	if err := repo.DB.Find(&fabricantes).Error; err != nil {
		return nil, err
	}
	return &fabricantes, nil
}

// TODO: Implementar metodos para Modificar fabricante
