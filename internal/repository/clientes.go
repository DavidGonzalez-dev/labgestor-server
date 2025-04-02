package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

// Interfaz que define los metodos que se emplean en la tabla de los clientes en la base datos
type ClienteRepository interface {
	CrearCliente(cliente *models.Cliente) error
	ActualizarCliente(cliente *models.Cliente) error
}

// Structura que implementa la interfaz anteriormente definida
type clienteRepository struct {
	DB *gorm.DB
}

// Funcion que devuelve una instancia de la estructura clienteRepository
func NewClienterepository(db *gorm.DB) ClienteRepository {
	return &clienteRepository{DB: db}
}

// ---------------------------
// Metodos de la estructura
// ---------------------------

func (repo *clienteRepository) CrearCliente(cliente *models.Cliente) error {
	return repo.DB.Create(&cliente).Error
}

func (repo *clienteRepository) ActualizarCliente(cliente *models.Cliente) error {
	return repo.DB.Save(&cliente).Error
}

// TODO: Implementar metodos para Modificar Cliente
