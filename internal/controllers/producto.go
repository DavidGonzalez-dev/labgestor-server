package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
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
// Metodos del controlador
// -------------------------------------
func (controller productoController) ObtenerProductoID(c echo.Context) error {
	// Se lee el cuerpo del request
	Numero_Registro := c.Param("id")

	// Se obtiene el producto haciendo uso de la capa del repositorio
	producto, err := controller.Repo.ObtenerProductoID(Numero_Registro)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, producto)
}

func (controller productoController) ObtenerProductos(c echo.Context) error {
	// Se obtiene el producto haciendo uso de la capa del repositorio
	productos, err := controller.Repo.ObtenerProductos()
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, productos)
}

func (controller productoController) CrearProducto(c echo.Context) error {
	// Se lee el cuerpo del request
	var requestBody struct {
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

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	if requestBody.Numero_Registro == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Numero Registro' es obligatorio"})
	}
	if requestBody.Nombre == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Nombre' es obligatorio"})
	}
	if requestBody.Fecha_fabricacion == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Fecha Fabricacion' es obligatorio"})
	}
	if requestBody.Fecha_vencimiento == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Fecha Vencimiento' es obligatorio"})
	}
	if requestBody.Descripcion == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Descripcion' es obligatorio"})
	}
	if requestBody.Compuesto_activo == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Compuesto Activo' es obligatorio"})
	}
	if requestBody.Presentacion == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Presentacion' es obligatorio"})
	}
	if requestBody.Cantidad == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Cantidad' es obligatorio"})
	}
	if requestBody.Numero_lote == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Numero Lote' es obligatorio"})
	}
	if requestBody.Tamano_lote == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Tamano Lote' es obligatorio"})
	}

	//	if requestBody.Id_cliente == 0 {
	//		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Id CLiente' debe tener un cliente existente"})
	//	}
	//	if requestBody.Id_fabricante == 0 {
	//		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Id Fabricante' debe tener un fabricante existente"})
	//	}
	//	if requestBody.Id_tipo == 0 {
	//		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Id Tipo' debe tener un tipo existente"})
	//	}
	//	if requestBody.Id_fabricante == 0 {
	//		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Id Estado' debe tener un fabricante existente"})
	//	}

	// Crear una instancia del modelo
	producto := models.Producto{
		Numero_Registro:   requestBody.Numero_Registro,
		Nombre:            requestBody.Nombre,
		Fecha_fabricacion: requestBody.Fecha_fabricacion,
		Fecha_vencimiento: requestBody.Fecha_vencimiento,
		Descripcion:       requestBody.Descripcion,
		Compuesto_activo:  requestBody.Compuesto_activo,
		Presentacion:      requestBody.Presentacion,
		Cantidad:          requestBody.Cantidad,
		Numero_lote:       requestBody.Numero_lote,
		Tamano_lote:       requestBody.Tamano_lote,
		Id_cliente:        requestBody.Id_cliente,
		Id_fabricante:     requestBody.Id_fabricante,
		Id_tipo:           requestBody.Id_tipo,
		Id_estado:         requestBody.Id_estado,
	}
	controller.Repo.CrearProducto(&producto)

	// Se retorna una respuesta exitosa
	return c.JSON(http.StatusOK, response.Response{Message: "El Producto ha sido registrado con exito"})
}
