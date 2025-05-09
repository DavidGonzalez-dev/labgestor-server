package repository

import (
	"gorm.io/gorm"
	"labgestor-server/internal/models"
)

// Interfaz que define los metodos que se emplean en la tabla de los productos en la base datos
type ProductoRepository interface {
	ObtenerProductoID(numeroRegistro string) (*models.RegistroEntradaProducto, error)
	ObtenerEntradasProductos() (*[]models.RegistroEntradaProducto, error)
	CrearProducto(producto *models.Producto, entradaProducto *models.RegistroEntradaProducto) error
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
// Este metodo nos permite obtener todos los detalles del registro del producto incluyendo la entrada del mismo
func (repo *productoRepository) ObtenerProductoID(numeroRegistroProducto string) (*models.RegistroEntradaProducto, error) {

	//? ------------------------------------------------------------------
	//? Instanciamos y Precargamos toda la informacion de un producto
	//? ------------------------------------------------------------------
	// Realizamos la consulta utilizando el valor del numero de registro del producto como parametro
	var entradaProducto models.RegistroEntradaProducto
	if err := repo.DB.
		Preload("Producto").                                         // Preload para que aparezca la informacion del producto
		Preload("Producto.Cliente").                                 // Preload para que aparezca la informacion del cliente del producto
		Preload("Producto.Fabricante").                              // Preload para que aparezca la informacion del fabricante del producto
		Preload("Producto.TipoProducto").                            // Preload para que aparezca la informacion del tipo del producto
		Preload("Producto.EstadoProducto").                          // Preload para que aparezca la informacion del estado del producto
		Where("numero_registro_producto=?", numeroRegistroProducto). // Filtro para seleccionar solo el registro que coincida con el numero de registro de producto pasado.
		First(&entradaProducto).Error; err != nil {
		// En caso de un error al momento de hacer la precarga de los datos se retorna el error y nil
		return nil, err
	}

	//? ------------------------------------------------------------------
	//? Se retorna el objeto ya poblado con la informacion de la base de datos
	//? ------------------------------------------------------------------
	// Se retorna el objeto con la informacion
	return &entradaProducto, nil
}

// Este metodo nos permite obtener solo los detalles del registro de la entrada del producto al area.
func (repo *productoRepository) ObtenerEntradasProductos() (*[]models.RegistroEntradaProducto, error) {
	//? ---------------------------------------------------------------------
	//? Se crea un slice para guardar los registros de entrada de productos
	//? ---------------------------------------------------------------------
	var registrosEntradaProducto []models.RegistroEntradaProducto
	// Se hace un preload de la tablas anidadas para obtener infromacion adicional
	if err := repo.DB.
		Preload("Usuario", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "nombres", "apellidos") // Se obtiene el id, los nombres y los apellidos del usuario responsable
		}).
		Find(&registrosEntradaProducto).Error; err != nil { // Se guardan los resultados en el slice declarado

		// En caso de haber ocurrido algun error devolvemos nil y el error
		return nil, err
	}

	// En caso de que todo halla salido bien se retorna el slice con los registro de entrada de productos y nil
	return &registrosEntradaProducto, nil
}

// Este metodo nos permite crear un producto en la base de datos junto a su respectivo producto
func (repo *productoRepository) CrearProducto(producto *models.Producto, entradaProducto *models.RegistroEntradaProducto) error {

	//? -----------------------------------------------------------------
	//? Se crea el producto mediante una transaccion en la base de datos
	//? -----------------------------------------------------------------
	// Se hace uso de una transaccion para asegurarse de que ambos registros queden en la base de datos y no solo uno.
	err := repo.DB.Transaction(func(tx *gorm.DB) error {

		// Se crea el producto y se verifica que no hallan errores
		if err := tx.Create(producto).Error; err != nil {
			return err
		}

		// Se crea el registro de entrada del producto y se verifica que no hallan errores
		if err := tx.Create(entradaProducto).Error; err != nil {
			return err
		}
		// Si todo salio bien se retorna nil
		return nil
	})

	// Se retorna el valor de retorno de la transaccion
	return err
}
