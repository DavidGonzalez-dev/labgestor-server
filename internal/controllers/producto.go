package controllers

import (
	"fmt"
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"labgestor-server/utils/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductoController interface {
	ObtenerProductoID(c echo.Context) error
	ObtenerRegistrosEntradaProductos(c echo.Context) error
	ObtenerRegistrosEntradaProductosPorUsuario(c echo.Context) error
	CrearProducto(c echo.Context) error
	EliminarProducto(c echo.Context) error
	ActualizarProducto(c echo.Context) error
	ActualizarRegistroEntradaProducto(c echo.Context) error
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

// Este handler nos devuelve con la informacion completa de un producto incluyendo los detalles del registro de la entrada al area
func (controller productoController) ObtenerProductoID(c echo.Context) error {
	//? -------------------------------------------------------------------
	//? Obtenemos el id del producto y se busca en la capa del repositorio
	//? -------------------------------------------------------------------
	// Se lee el cuerpo del request
	numeroRegistroProducto := c.Param("id")
	// Se obtiene el producto haciendo uso de la capa del repositorio y en caso de no encontrarlo se retorna un estado de notFound
	registroEntradaProducto, err := controller.Repo.ObtenerProductoID(numeroRegistroProducto)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Producto no encontrado", Error: err.Error()})
	}
	//  Se retorna el registro de la entrada del producto con todos los detalles del mismo
	return c.JSON(http.StatusFound, response.Response{Data: registroEntradaProducto})
}

// Este handler nos devuelve un array con los registros de entrada de los productos sin detalles.
func (controller productoController) ObtenerRegistrosEntradaProductos(c echo.Context) error {
	//? --------------------------------------------------------------
	//? Se Obtienen todos los registros de entrada de los productos
	//? --------------------------------------------------------------
	// Se obtiene el producto haciendo uso de la capa del repositorio y en dado caso de presentarse un error no se retorna informacion sino el error
	productos, err := controller.Repo.ObtenerEntradasProductos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Hubo un error al obtener todos los productos", Error: err.Error()})
	}

	// En caso de haber salido todo bien se retorna la informacion de los registros de entrada de los productos
	return c.JSON(http.StatusOK, response.Response{Data: productos})
}

// Este handler nos devuelve un array con los registros de entrada de los productos sin detalles.
func (controller productoController) ObtenerRegistrosEntradaProductosPorUsuario(c echo.Context) error {
	
	// Se obtiene el id del usuario
	id := c.Param("id")
	
	//? --------------------------------------------------------------
	//? Se Obtienen todos los registros de entrada de los productos
	//? --------------------------------------------------------------
	// Se obtiene el producto haciendo uso de la capa del repositorio y en dado caso de presentarse un error no se retorna informacion sino el error
	productos, err := controller.Repo.ObtenerEntradasProductosPorUsuario(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Hubo un error al obtener los productos del usuario", Error: err.Error()})
	}

	// En caso de haber salido todo bien se retorna la informacion de los registros de entrada de los productos
	return c.JSON(http.StatusOK, response.Response{Data: productos})
}

// Este handler nos permite crear un producto en la base de datos con su respectivo registro de entrada al area
func (controller productoController) CrearProducto(c echo.Context) error {

	//? --------------------------------------------------------------------------
	//? Bind de la informacion del request
	//? --------------------------------------------------------------------------
	// Se lee el cuerpo del request y en caso de haber algun error se devuelve un estado de peticion erronea
	var requestBody struct {
		Producto struct {
			NumeroRegistro   string `json:"numeroRegistro"`
			Nombre           string `json:"nombre"`
			FechaFabricacion string `json:"fechaFabricacion"`
			FechaVencimiento string `json:"fechaVencimiento"`
			Descripcion      string `json:"descripcion"`
			CompuestoActivo  string `json:"compuestoActivo"`
			Presentacion     string `json:"presentacion"`
			Cantidad         string `json:"cantidad"`
			NumeroLote       string `json:"numeroLote"`
			TamanoLote       string `json:"tamanoLote"`
			IDCliente        int    `json:"idCliente"`
			IDFabricante     int    `json:"idFabricante"`
			IDTipo           int    `json:"idTipo"`
		} `json:"producto"`
		DetallesEntrada struct {
			PropositoAnalisis      string `json:"propositoAnalisis"`
			CondicionesAmbientales string `json:"condicionesAmbientales"`
			FechaRecepcion         string `json:"fechaRecepcion"`
			FechaInicioAnalisis    string `json:"fechaInicioAnalisis"`
			FechaFinalAnalisis     string `json:"fechaFinalAnalisis"`
			IDUsuario              string `json:"idUsuario"`
		} `json:"detallesEntrada"`
	}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	// Se verifica que el producto no exista
	if producto, _ := controller.Repo.ObtenerInfoProducto(requestBody.Producto.NumeroRegistro); producto != nil {
		return c.JSON(http.StatusConflict, response.Response{Message: "Errro al crear el producto", Error: fmt.Sprintf("Ya existe un producto con el numero de registro: %s", producto.NumeroRegistro)})
	}

	//? --------------------------------------------------------------------------
	//? Lectura y validacion de los atributos del producto
	//? --------------------------------------------------------------------------
	// Crear una instancia del modelo del producto
	producto := models.Producto{
		NumeroRegistro:   requestBody.Producto.NumeroRegistro,
		Nombre:           requestBody.Producto.Nombre,
		FechaFabricacion: requestBody.Producto.FechaFabricacion,
		FechaVencimiento: requestBody.Producto.FechaVencimiento,
		Descripcion:      requestBody.Producto.Descripcion,
		CompuestoActivo:  requestBody.Producto.CompuestoActivo,
		Presentacion:     requestBody.Producto.Presentacion,
		Cantidad:         requestBody.Producto.Cantidad,
		NumeroLote:       requestBody.Producto.NumeroLote,
		TamanoLote:       requestBody.Producto.TamanoLote,
		IDCliente:        requestBody.Producto.IDCliente,
		IDFabricante:     requestBody.Producto.IDFabricante,
		IDTipo:           requestBody.Producto.IDTipo,
		IDEstado:         1,
	}

	// Se hace la validacion de los campos y en caso de no cumplir con alguna regla se devuelve un error con la informacion del campo que no cumple las reglas de validacion
	if err := validation.Validate(producto.ToMap(), validation.ProductoRules); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	//? --------------------------------------------------------------------------
	//? Lectura y validacion de los atributos del registro de entrada del producto
	//? --------------------------------------------------------------------------
	//Se crea una instancia de la entrada del producto
	entradaProducto := models.RegistroEntradaProducto{
		PropositoAnalisis:      requestBody.DetallesEntrada.PropositoAnalisis,
		CondicionesAmbientales: requestBody.DetallesEntrada.CondicionesAmbientales,
		FechaRecepcion:         requestBody.DetallesEntrada.FechaRecepcion,
		FechaInicioAnalisis:    requestBody.DetallesEntrada.FechaInicioAnalisis,
		FechaFinalAnalisis:     requestBody.DetallesEntrada.FechaFinalAnalisis,
		IDUsuario:              requestBody.DetallesEntrada.IDUsuario,
		NumeroRegistroProducto: requestBody.Producto.NumeroRegistro,
	}
	// Se hace la validacion de los campos y en caso de no cumplir con alguna regla se devuelve un error con la informacion del campo que no cumple las reglas de validacion
	if err := validation.Validate(entradaProducto.ToMap(), validation.RegistroEntradaRules); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	//? --------------------------------------------------------------------------
	//? Creacion del producto en la base de datos
	//? --------------------------------------------------------------------------
	// Si todas las reglas de validacion pasaron se crea el producto en la base de datos ,si hay algun error mientras se crea se retorna el error
	if err := controller.Repo.CrearProducto(&producto, &entradaProducto); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al crear el producto", Error: err.Error()})
	}

	// Si todo salio bien se retorna una respuesta exitosa que indique que el producto se ha registrado con exito
	return c.JSON(http.StatusCreated, response.Response{Message: "El Producto ha sido registrado con exito"})
}

func (controller productoController) ActualizarProducto(c echo.Context) error {
	//? --------------------------------------------------------------------------
	//? Bind de la informacion del request
	//? --------------------------------------------------------------------------
	// Se lee el cuerpo del request y en caso de haber algun error se devuelve un estado de peticion erronea
	var requestBody struct {
		NumeroRegistro   string `json:"numeroRegistro"`
		Nombre           string `json:"nombre"`
		FechaFabricacion string `json:"fechaFabricacion"`
		FechaVencimiento string `json:"fechaVencimiento"`
		Descripcion      string `json:"descripcion"`
		CompuestoActivo  string `json:"compuestoActivo"`
		Presentacion     string `json:"presentacion"`
		Cantidad         string `json:"cantidad"`
		NumeroLote       string `json:"numeroLote"`
		TamanoLote       string `json:"tamanoLote"`
		IDCliente        int    `json:"idCliente"`
		IDFabricante     int    `json:"idFabricante"`
		IDTipo           int    `json:"idTipo"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}
	//? --------------------------------------------------------------------------
	//? Lectura y validacion de los atributos del producto
	//? --------------------------------------------------------------------------

	//Obtenemos el producto de la base de datos
	producto, err := controller.Repo.ObtenerInfoProducto(requestBody.NumeroRegistro)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Producto no encontrado", Error: err.Error()})
	}

	// ?-----------------------------------------------------------
	// ? Se actualizan los campos del producto
	// ?-----------------------------------------------------------\
	producto.Nombre = requestBody.Nombre
	producto.FechaFabricacion = requestBody.FechaFabricacion
	producto.FechaVencimiento = requestBody.FechaVencimiento
	producto.Descripcion = requestBody.Descripcion
	producto.CompuestoActivo = requestBody.CompuestoActivo
	producto.Presentacion = requestBody.Presentacion
	producto.Cantidad = requestBody.Cantidad
	producto.NumeroLote = requestBody.NumeroLote
	producto.TamanoLote = requestBody.TamanoLote
	producto.IDCliente = requestBody.IDCliente
	producto.IDFabricante = requestBody.IDFabricante
	producto.IDTipo = requestBody.IDTipo
	


	if err := validation.Validate(producto.ToMap(), validation.ProductoRules); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}
	//? --------------------------------------------------------------------------
	//? Actualizacion del producto en la base de datos
	//? --------------------------------------------------------------------------

	if err := controller.Repo.ActualizarProducto(producto); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al actualizar el producto", Error: err.Error()})
	}
	// Si todo salio bien se retorna una respuesta exitosa que indique que el producto se ha registrado con exito
	return c.JSON(http.StatusOK, response.Response{Message: "El Producto ha sido actualizado con exito"})
}

// Este handler nos permite actualizar el registro de entrada del producto
func (controller productoController) ActualizarRegistroEntradaProducto(c echo.Context) error {
	//? --------------------------------------------------------------------------
	//? Bind de la informacion del request
	//? --------------------------------------------------------------------------
	
	// Se lee el cuerpo del request y en caso de haber algun error se devuelve un estado de peticion erronea
	var requestBody struct {
		NumeroRegistroProducto string `json:"numeroRegistroProducto"`
		PropositoAnalisis      string `json:"propositoAnalisis"`
		CondicionesAmbientales string `json:"condicionesAmbientales"`
		FechaRecepcion         string `json:"fechaRecepcion"`
		FechaInicioAnalisis    string `json:"fechaInicioAnalisis"`
		FechaFinalAnalisis     string `json:"fechaFinalAnalisis"`
	}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	//? --------------------------------------------------------------------------
	//? Lectura y validacion de los atributos del registro de entrada del producto
	//? --------------------------------------------------------------------------

	registroEntradaProducto, err := controller.Repo.ObtenerInfoRegistroEntradaProducto(requestBody.NumeroRegistroProducto)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Registro de entrada no encontrado", Error: err.Error()})
	}

	// Se hace la actualizacion de los valores del registro de la entrada del producto
	registroEntradaProducto.PropositoAnalisis = requestBody.PropositoAnalisis
	registroEntradaProducto.CondicionesAmbientales = requestBody.CondicionesAmbientales
	registroEntradaProducto.FechaRecepcion = requestBody.FechaRecepcion
	registroEntradaProducto.FechaInicioAnalisis = requestBody.FechaInicioAnalisis
	registroEntradaProducto.FechaFinalAnalisis = requestBody.FechaFinalAnalisis

	// Se hace la validacion de los campos
	validationRules := validation.RegistroEntradaRules
	delete(validationRules, "IDUsuario") // Se elimina la regla de validacion para el campo IDUsuario

	// Se hace la validacion de los campos y en caso de no cumplir con alguna regla se devuelve un error con la informacion del campo que no cumple las reglas de validacion
	if err := validation.Validate(registroEntradaProducto.ToMap(), validationRules); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	//? --------------------------------------------------------------------------
	//? Actualizacion del registro de entrada del producto en la base de datos
	//? --------------------------------------------------------------------------
	if err := controller.Repo.ActualizarRegistroEntradaProducto(registroEntradaProducto); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al actualizar el registro de entrada del producto", Error: err.Error()})
	}
	// Si todo salio bien se retorna una respuesta exitosa que indique que el producto se ha registrado con exito
	return c.JSON(http.StatusOK, response.Response{Message: "El registro de entrada del producto ha sido actualizado con exito"})
}

// Este handler nos permite eliminar un producto y todos los registros relacionados
func (controller productoController) EliminarProducto(c echo.Context) error {

	// Obtenemos el parametro del endpoint
	numeroRegistroProducto := c.Param("id")
	println(numeroRegistroProducto)

	// Se trae el registro del producto para verificar que si exista
	producto, err := controller.Repo.ObtenerInfoProducto(numeroRegistroProducto)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Este producto no existe", Error: err.Error()})
	}

	// Se elimina el producto
	if err := controller.Repo.EliminarProducto(producto); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al eliminar el producto", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "El producto se elimino con exito"})
}
