package repository

import (
	"errors"
	"labgestor-server/internal/models"
	"time"

	"gorm.io/gorm"
)

// Interfaz que define los metodos que se emplean en la tabla de los productos en la base datos
type ProductoRepository interface {
	ObtenerProductoID(numeroRegistro string) (*models.RegistroEntradaProducto, error)
	ObtenerEntradasProductos() (*[]models.RegistroEntradaProducto, error)
	ObtenerEntradasProductosPorUsuario(idUsuario string) (*[]models.RegistroEntradaProducto, error)
	CrearProducto(producto *models.Producto, entradaProducto *models.RegistroEntradaProducto) error
	EliminarProducto(producto *models.Producto) error
	ActualizarProducto(producto *models.Producto) error
	ActualizarRegistroEntradaProducto(entradaProducto *models.RegistroEntradaProducto) error
	ObtenerInfoProducto(numeroRegitroProducto string) (*models.Producto, error)
	ObtenerInfoRegistroEntradaProducto(numeroRegistroProducto string) (*models.RegistroEntradaProducto, error)
	ActualizarEstadoProducto(newEstado int, numeroRegistroProducto string) error

	ObtenerProductosAnalizadosSemana() (map[string]any, error)
	ObtenerTipoProductosSemana() ([]map[string]any, error)
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
		Preload("Producto").              // Preload para que aparezca la informacion del producto
		Preload("Producto.Cliente").      // Preload para que aparezca la informacion del cliente del producto
		Preload("Producto.Fabricante").   // Preload para que aparezca la informacion del fabricante del producto
		Preload("Producto.TipoProducto"). // Preload para que aparezca la informacion del tipo del producto
		Preload("Producto.EstadoProducto").
		Preload("Usuario", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "nombres", "apellidos")
		}).                                                          // Preload para que aparezca la informacion del estado del producto
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
		Preload("Producto.TipoProducto").
		Preload("Producto.EstadoProducto").
		Preload("Usuario", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "firma")
		}).
		Find(&registrosEntradaProducto).Error; err != nil { // Se guardan los resultados en el slice declarado

		// En caso de haber ocurrido algun error devolvemos nil y el error
		return nil, err
	}

	// En caso de que todo halla salido bien se retorna el slice con los registro de entrada de productos y nil
	return &registrosEntradaProducto, nil
}

// Este metodo nos permite obtener solo los detalles del registro de la entrada del producto al area encargados de un usuario en especifico
func (repo *productoRepository) ObtenerEntradasProductosPorUsuario(idUsuario string) (*[]models.RegistroEntradaProducto, error) {
	var registrosEntradaProducto []models.RegistroEntradaProducto
	// Se hace un preload de la tablas anidadas para obtener infromacion adicional
	if err := repo.DB.
		Preload("Producto.TipoProducto").
		Preload("Producto.EstadoProducto").
		Where("id_usuario=?", idUsuario).
		Find(&registrosEntradaProducto).Error; err != nil { // Se guardan los resultados en el slice declarado

		// En caso de haber ocurrido algun error devolvemos nil y el error
		return nil, err
	}
	// En caso que todo halla salido bien se retorna el slice con los registros de entrada de productos y nil
	return &registrosEntradaProducto, nil
}

// Este metodo nos permite crear un producto en la base de datos junto a su respectivo producto
func (repo *productoRepository) CrearProducto(producto *models.Producto, entradaProducto *models.RegistroEntradaProducto) error {

	//? -----------------------------------------------------------------
	//? Se crea el producto mediante una transaccion en la base de datos
	//? -----------------------------------------------------------------
	// Se hace uso de una transaccion para asegurarse de que ambos registros queden en la base de datos y no solo uno.
	err := repo.DB.Transaction(func(tx *gorm.DB) error {

		// Maneja campos de fecha vacíos
		productoData := map[string]any{
			"numero_registro":  producto.NumeroRegistro,
			"nombre":           producto.Nombre,
			"descripcion":      producto.Descripcion,
			"compuesto_activo": producto.CompuestoActivo,
			"presentacion":     producto.Presentacion,
			"cantidad":         producto.Cantidad,
			"numero_lote":      producto.NumeroLote,
			"tamano_lote":      producto.TamanoLote,
			"id_cliente":       producto.IDCliente,
			"id_fabricante":    producto.IDFabricante,
			"id_tipo":          producto.IDTipo,
			"id_estado":        producto.IDEstado,
		}
		// Si las fechas no están vacías, agrégalas al mapa
		if producto.FechaFabricacion != "" {
			productoData["fecha_fabricacion"] = producto.FechaFabricacion
		}

		if producto.FechaVencimiento != "" {
			productoData["fecha_vencimiento"] = producto.FechaVencimiento
		}

		// Se crea el producto usando el mapa de datos y se verifica que no haya errores
		if err := tx.Model(&models.Producto{}).Create(productoData).Error; err != nil {
			return err
		}

		// Maneja campos de fecha vacíos para el registro de entrada
		entradaData := map[string]any{
			"proposito_analisis":       entradaProducto.PropositoAnalisis,
			"condiciones_ambientales":  entradaProducto.CondicionesAmbientales,
			"numero_registro_producto": entradaProducto.NumeroRegistroProducto,
			"id_usuario":               entradaProducto.IDUsuario,
		}
		// Si las fechas no están vacías, agrégalas al mapa
		if entradaProducto.FechaRecepcion != "" {
			entradaData["fecha_recepcion"] = entradaProducto.FechaRecepcion
		}

		if entradaProducto.FechaInicioAnalisis != "" {
			entradaData["fecha_inicio_analisis"] = entradaProducto.FechaInicioAnalisis
		}

		if entradaProducto.FechaFinalAnalisis != "" {
			entradaData["fecha_final_analisis"] = entradaProducto.FechaFinalAnalisis
		}
		// Se crea el registro de entrada del producto usando el mapa de datos y se verifica que no haya errores
		if err := tx.Model(&models.RegistroEntradaProducto{}).Create(entradaData).Error; err != nil {
			return err
		}
		// Si todo salio bien se retorna nil
		return nil
	})

	// Se retorna el valor de retorno de la transaccion
	return err
}

// Este metodo nos permie eliminar un registro de un producto en la base de datos
func (repo *productoRepository) EliminarProducto(producto *models.Producto) error {
	// Se elimina el producto y se verifica que no hallan errores
	if err := repo.DB.Delete(&producto).Error; err != nil {
		return err
	}
	// Si todo salio bien se retorna nil
	return nil
}

// Este metodo nos permite actualizar la informacion de un producto
func (repo *productoRepository) ActualizarProducto(producto *models.Producto) error {
	// Se actualiza el producto y se verifica que no hallan errores
	if err := repo.DB.Model(&models.Producto{}).Where("numero_registro = ?", producto.NumeroRegistro).Updates(map[string]any{
		"Nombre":           producto.Nombre,
		"FechaFabricacion": producto.FechaFabricacion,
		"FechaVencimiento": producto.FechaVencimiento,
		"Descripcion":      producto.Descripcion,
		"CompuestoActivo":  producto.CompuestoActivo,
		"Presentacion":     producto.Presentacion,
		"Cantidad":         producto.Cantidad,
		"NumeroLote":       producto.NumeroLote,
		"TamanoLote":       producto.TamanoLote,
		"IDCliente":        producto.IDCliente,
		"IDFabricante":     producto.IDFabricante,
		"IDTipo":           producto.IDTipo,
	}).Error; err != nil {
		// En caso de un error se retorna
		return err
	}
	return nil
}

// Este metodo nos permite actualizar el registro de entrada del producto
func (repo *productoRepository) ActualizarRegistroEntradaProducto(entradaProducto *models.RegistroEntradaProducto) error {
	// Preparar un mapa con los campos que siempre se actualizan
	updateData := map[string]any{
		"proposito_analisis":      entradaProducto.PropositoAnalisis,
		"condiciones_ambientales": entradaProducto.CondicionesAmbientales,
	}
	// Solo incluir fechas cuando no estén vacías
	if entradaProducto.FechaRecepcion != "" {
		updateData["fecha_recepcion"] = entradaProducto.FechaRecepcion
	}

	if entradaProducto.FechaInicioAnalisis != "" {
		updateData["fecha_inicio_analisis"] = entradaProducto.FechaInicioAnalisis
	}

	if entradaProducto.FechaFinalAnalisis != "" {
		updateData["fecha_final_analisis"] = entradaProducto.FechaFinalAnalisis
	}

	// Se actualiza el registro de entrada del producto con los campos no vacíos
	if err := repo.DB.Model(&models.RegistroEntradaProducto{}).Where("codigo_entrada = ?", entradaProducto.CodigoEntrada).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

// Este metodo Permite cambiar el estado de un producto
func (repo *productoRepository) ActualizarEstadoProducto(newEstado int, numeroRegistroProducto string) error {

	//Preparamos los datos que vamos a actualizar
	updateData := map[string]any{
		"id_estado": newEstado,
	}

	result := repo.DB.Model(&models.Producto{}).Where("numero_registro=?", numeroRegistroProducto).Updates(updateData)
	if result.Error != nil {
		return result.Error
	}
	// Revisamos que halla afectado a alguna fila
	if result.RowsAffected == 0 {
		return errors.New("el producto que estas intentando actualizar no existe")
	}
	return nil
}

// Este metodo nos trae las estadisticas de los productos ingresados en la semana
func (repo *productoRepository) ObtenerProductosAnalizadosSemana() (map[string]any, error) {
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	startOfWeek := now.AddDate(0, 0, -(weekday - 1)).Truncate(24 * time.Hour)

	result := make(map[string]any)
	diasSemana := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}

	// Iterar por cada día de la semana
	for i := 0; i < 7; i++ {
		currentDay := startOfWeek.AddDate(0, 0, i)
		nextDay := currentDay.AddDate(0, 0, 1)

		var count int64
		err := repo.DB.Model(&models.RegistroEntradaProducto{}).
			Where("fecha_recepcion >= ? AND fecha_recepcion < ?", currentDay, nextDay).
			Count(&count).Error

		if err != nil {
			return nil, err
		}

		result[diasSemana[i]] = count
	}

	return result, nil

}

// Este metodo nos permite obtener las estadisticas de cantidad de productos por tipo
func (repo *productoRepository) ObtenerTipoProductosSemana() ([]map[string]any, error) {

	var results []map[string]any

	repo.DB.Table("productos AS p").
		Select("ep.nombre_estado, COUNT(*) AS cantidad_productos").
		Joins("JOIN estado_productos ep ON p.id_estado = ep.id").
		Group("ep.nombre_estado").
		Scan(&results)

	return results, nil
}

// -------------------------------
// Metodos de Ayuda
// -------------------------------
// Esta funcion nos permite traer unicamente la informacion plana del producto, sin hacer preloads pesados. Esta funcion se crea para el uso interno de la apliacion como ayuda para el desarrollador
func (repo *productoRepository) ObtenerInfoProducto(numeroRegitroProducto string) (*models.Producto, error) {

	var producto models.Producto
	if err := repo.DB.Where("numero_registro = ?", numeroRegitroProducto).First(&producto).Error; err != nil {
		return nil, err
	}

	return &producto, nil
}

func (repo *productoRepository) ObtenerInfoRegistroEntradaProducto(numeroRegistroProducto string) (*models.RegistroEntradaProducto, error) {
	var registroEntradaProducto models.RegistroEntradaProducto
	if err := repo.DB.Where("numero_registro_producto = ?", numeroRegistroProducto).First(&registroEntradaProducto).Error; err != nil {
		return nil, err
	}
	return &registroEntradaProducto, nil
}
