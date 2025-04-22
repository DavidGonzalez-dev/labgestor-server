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
	ObtenerEntradasProductos(c echo.Context) error
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

func (controller productoController) ObtenerEntradasProductos(c echo.Context) error {
	// Se obtiene el producto haciendo uso de la capa del repositorio
	productos, err := controller.Repo.ObtenerEntradasProductos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Hubo un error al obtener todos los usuarios", Error: err.Error()})
	}

	return c.JSON(http.StatusFound, response.Response{Data: productos})
}

func (controller productoController) CrearProducto(c echo.Context) error {
	// Se lee el cuerpo del request
	// Se lee la informacion del producto
	var requestBody struct {
		Producto struct {
			Numero_Registro   string `json:"numeroRegistro"`
			Nombre            string `json:"nombre"`
			Fecha_fabricacion string `json:"fechaFabricacion"`
			Fecha_vencimiento string `json:"fechaVencimiento"`
			Descripcion       string `json:"descripcion"`
			Compuesto_activo  string `json:"compuestoActivo"`
			Presentacion      string `json:"presentacion"`
			Cantidad          string `json:"cantidad"`
			Numero_lote       string `json:"numeroLote"`
			Tamano_lote       string `json:"tamanoLote"`
			IDCliente         int    `json:"idCliente"`
			IDFabricante      int    `json:"idFabricante"`
			IDTipo            int    `json:"idTipo"`
			IDEstado          int    `json:"idEstado"`
		} `json:"producto"`
		DetallesEntrada struct {
			PropositoAnalisis      string `json:"propositoAnalisis"`
			CondicionesAmbientales string `json:"condicionesAmbientales"`
			FechaRecepcion         string `json:"fechaRecepcion"`
			FechaInicioAnalisis    string `json:"fechaInicioAnalisis"`
			FechaFinalAnalisis     string `json:"fechaFinalAnalisis"`
			IDUsuario              string `json:"idUsuario"`
			NumeroRegistroProducto string `json:"numeroRegistroProducto"`
		} `json:"detallesEntrada"`
	}

	// Se lee la informacion de los detalles de la entrada del producto
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	// Crear una instancia del modelo del producto
	producto := models.Producto{
		NumeroRegistro:   requestBody.Producto.Numero_Registro,
		Nombre:           requestBody.Producto.Nombre,
		FechaFabricacion: requestBody.Producto.Fecha_fabricacion,
		FechaVencimiento: requestBody.Producto.Fecha_vencimiento,
		Descripcion:      requestBody.Producto.Descripcion,
		CompuestoActivo:  requestBody.Producto.Compuesto_activo,
		Presentacion:     requestBody.Producto.Presentacion,
		Cantidad:         requestBody.Producto.Cantidad,
		NumeroLote:       requestBody.Producto.Numero_lote,
		TamanoLote:       requestBody.Producto.Tamano_lote,
		IDCliente:        requestBody.Producto.IDCliente,
		IDFabricante:     requestBody.Producto.IDFabricante,
		IDTipo:           requestBody.Producto.IDTipo,
		IDEstado:         requestBody.Producto.IDEstado,
	}

	entradaProducto := models.EntradaProducto{
		PropositoAnalisis:      requestBody.DetallesEntrada.PropositoAnalisis,
		CondicionesAmbientales: requestBody.DetallesEntrada.CondicionesAmbientales,
		FechaRecepcion:         requestBody.DetallesEntrada.FechaRecepcion,
		FechaInicioAnalisis:    requestBody.DetallesEntrada.FechaInicioAnalisis,
		FechaFinalAnalisis:     requestBody.DetallesEntrada.FechaFinalAnalisis,
		IDUsuario:              requestBody.DetallesEntrada.IDUsuario,
		NumeroRegistroProducto: requestBody.DetallesEntrada.NumeroRegistroProducto,
	}

	if err := controller.Repo.CrearProducto(&producto, &entradaProducto); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al crear el producto", Error: err.Error()})
	}

	// Se retorna una respuesta exitosa
	return c.JSON(http.StatusCreated, response.Response{Message: "El Producto ha sido registrado con exito"})
}
