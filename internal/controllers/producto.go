package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"labgestor-server/utils/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductoController interface {
	ObtenerProductoID(c echo.Context) error
	ObtenerProductos(c echo.Context) error
	CrearProducto(c echo.Context) error
}

type productoController struct {
	Repo repository.ProductoRepository
}

func NewProductoController(repo repository.ProductoRepository) ProductoController {
	return productoController{Repo: repo}
}

// -------------------------------------
// CONTROLADORES CURD
// -------------------------------------
func (controller productoController) ObtenerProductoID(c echo.Context) error {
	// Se lee el cuerpo del request
	numeroRegistro := c.Param("id")

	// Se obtiene el producto haciendo uso de la capa del repositorio
	producto, err := controller.Repo.ObtenerProductoID(numeroRegistro)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al obtener el producto", Error: err.Error()})
	}

	// Verificar si existe el producto
	if producto.NumeroRegistro == "" {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Producto no encontrado"})
	}

	return c.JSON(http.StatusFound, response.Response{Data: producto})
}

// TODO: Modificar de donde sale la informacion de los productos. Sale de la tabla entradas de productos no de la de productos
func (controller productoController) ObtenerProductos(c echo.Context) error {
	// Se obtiene el producto haciendo uso de la capa del repositorio
	productos, err := controller.Repo.ObtenerProductos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Hubo un error al obtener todos los usuarios", Error: err.Error()})
	}

	return c.JSON(http.StatusFound, response.Response{Data: productos})
}

func (controller productoController) CrearProducto(c echo.Context) error {
	// Se lee el cuerpo del request
	// Se lee la informacion del producto
	var requestBody struct {
		producto struct {
			Numero_Registro   string
			Nombre            string
			Fecha_fabricacion string
			Fecha_vencimiento string
			Descripcion       string
			Compuesto_activo  string
			Presentacion      string
			Cantidad          string
			Numero_lote       string
			Tamano_lote       string
			Id_cliente        int
			Id_fabricante     int
			Id_tipo           int
			Id_estado         int
		}
		detallesEntrada struct {
			// TODO: Valores de la tabla de entradas de productos
		}
	}

	// Se lee la informacion de los detalles de la entrada del producto
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	// Hacer validaciones de campo
	validationRules := []validation.ValidationRule{}
	err := validation.Validate(map[string]any{}, validationRules)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{})
	}

	// Crear una instancia del modelo del producto
	producto := models.Producto{
		NumeroRegistro:   requestBody.producto.Numero_Registro,
		Nombre:           requestBody.producto.Nombre,
		FechaFabricacion: requestBody.producto.Fecha_fabricacion,
		FechaVencimiento: requestBody.producto.Fecha_vencimiento,
		Descripcion:      requestBody.producto.Descripcion,
		CompuestoActivo:  requestBody.producto.Compuesto_activo,
		Presentacion:     requestBody.producto.Presentacion,
		Cantidad:         requestBody.producto.Cantidad,
		NumeroLote:       requestBody.producto.Numero_lote,
		TamanoLote:       requestBody.producto.Tamano_lote,
		ClienteID:        requestBody.producto.Id_cliente,
		FabricanteID:     requestBody.producto.Id_fabricante,
		TipoID:           requestBody.producto.Id_tipo,
		EstadoID:         requestBody.producto.Id_estado,
	}

	// TODO: Se crea una instancia del registro de la entrada del producto

	controller.Repo.CrearProducto(&producto)

	// Se retorna una respuesta exitosa
	return c.JSON(http.StatusCreated, response.Response{Message: "El Producto ha sido registrado con exito"})
}
