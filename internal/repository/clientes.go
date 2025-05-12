package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

// Interfaz que define los metodos que se emplean en la tabla de los clientes en la base datos
type ClienteRepository interface {
	CrearCliente(cliente *models.Cliente) error
	ActualizarCliente(cliente *models.Cliente) error
	ObtenerClienteID(ID int) (*models.Cliente, error)
	ObtenerClientes() (*[]models.Cliente, error)
	EliminarCliente(cliente *models.Cliente) error
}

// Structura que implementa la interfaz anteriormente definida
type clienteRepository struct {
	DB *gorm.DB
}

// Funcion que devuelve una instancia de la estructura clienteRepository
func NewClienterepository(db *gorm.DB) ClienteRepository {
	return &clienteRepository{DB: db}
}

// ? ---------------------------
// ? Metodos de la estructura
// ? ---------------------------
// Este metodo nos permite crear un registro de un cliente en la base de datos
func (repo *clienteRepository) CrearCliente(cliente *models.Cliente) error {
	return repo.DB.Create(&cliente).Error
}

// Este metodo nos permite actualizar el registro de un cliente en la base datos
func (repo *clienteRepository) ActualizarCliente(cliente *models.Cliente) error {
	return repo.DB.Save(&cliente).Error
}

// Este metodo nos permite obtener la informacion deun cliente en base a us ID
func (repo *clienteRepository) ObtenerClienteID(ID int) (*models.Cliente, error) {
	var cliente models.Cliente
	// Realizamos la consulta utilizando el valor del ID como parámetro
	if err := repo.DB.First(&cliente, ID).Error; err != nil {
		return nil, err
	}

	return &cliente, nil
}

// Este metodo nos permite obtener un slice con todos los registros de los clientes
func (repo *clienteRepository) ObtenerClientes() (*[]models.Cliente, error) {
	var clientes []models.Cliente
	if err := repo.DB.Find(&clientes).Error; err != nil {
		return nil, err
	}
	return &clientes, nil
}

// Este metodo nos permite eliminar un cliente de la base de datos
func (repo *clienteRepository) EliminarCliente(cliente *models.Cliente) error {
	return repo.DB.Delete(&cliente).Error
}
