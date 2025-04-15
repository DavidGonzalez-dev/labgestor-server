package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

// Interfaz que define los metodos que se emplean en la tabla de los productos en la base datos
type ProductoRepository interface {
	ObtenerProductoID(Numero_Registro string) (*models.Producto, error)
	ObtenerProductos() (*[]models.Producto, error)
}

// Structura que implementa la interfaz anteriormente definida
type productoRepository struct {
	DB *gorm.DB
}

// Funcion que devuelve una instancia de la estructura producto repository
func NewProductoRepository(db *gorm.DB) ProductoRepository {
	return &productoRepository{DB: db}
}

func (repo *productoRepository) ObtenerProductoID(Numero_Registro string) (*models.Producto, error) {
	var producto models.Producto

	//realizamos la consulta utilizando el valor del ID como parametro
	if err := repo.DB.First(&producto, Numero_Registro).Error; err != nil {
		return nil, err
	}

	return &producto, nil
}

func (repo *productoRepository) ObtenerProductos() (*[]models.Producto, error) {
	var productos []models.Producto
	if err := repo.DB.Find(&productos).Error; err != nil {
		return nil, err
	}
	return &productos, nil

}
