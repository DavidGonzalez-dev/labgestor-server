package repository

import (
	"gorm.io/gorm"
	"labgestor-server/internal/models"
)

// Interfaz que define los metodos que se emplean en la tabla de los productos en la base datos
type ProductoRepository interface {
	ObtenerProductoID(numeroRegistro string) (*models.Producto, error)
	ObtenerEntradasProductos() (*[]models.EntradaProducto, error)
	CrearProducto(producto *models.Producto, entradaProducto *models.EntradaProducto) error
}

// Structura que implementa la interfaz anteriormente definida
type productoRepository struct {
	DB *gorm.DB
}

// Funcion que devuelve una instancia de la estructura producto repository
func NewProductoRepository(db *gorm.DB) ProductoRepository {
	return &productoRepository{DB: db}
}

// -----------------------------------
// METODOS CRUD
// -----------------------------------
func (repo *productoRepository) ObtenerProductoID(numeroRegistro string) (*models.Producto, error) {
	var producto models.Producto

	//realizamos la consulta utilizando el valor del ID como parametro
	if err := repo.DB.Preload("Cliente").Preload("Fabricante").Preload("TipoProducto").Preload("EstadoProducto").First(&producto, numeroRegistro).Error; err != nil {
		return nil, err
	}

	return &producto, nil
}

func (repo *productoRepository) ObtenerEntradasProductos() (*[]models.EntradaProducto, error) {
	var entradasProductos []models.EntradaProducto
	if err := repo.DB.Preload("Usuario", func(db *gorm.DB)*gorm.DB {return db.Select("id", "nombres", "apellidos")}).Find(&entradasProductos).Error; err != nil {
		return nil, err
	}
	return &entradasProductos, nil

}

func (repo *productoRepository) CrearProducto(producto *models.Producto, entradaProducto *models.EntradaProducto) error {
	// Se hace uso de una transaccion para asegurarse dfe que ambos registros queden registrados de manera correcta

	err := repo.DB.Transaction(func (tx *gorm.DB) error {
		// Se crea el producto
		if err := tx.Create(producto).Error; err != nil {
			return err
		}
	
		// Se crea la entrada del producto en la base de datos
		if err := tx.Create(entradaProducto).Error; err != nil{
			return err
		}
		
		return nil
	})

	return err
}
