package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"labgestor-server/utils/validation"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

type ProductoController interface {
	ObtenerProductoID(c echo.Context) error
	ObtenerRegistrosEntradaProductos(c echo.Context) error
	CrearProducto(c echo.Context) error
	EliminarProducto(c echo.Context) error
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
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Hubo un error al obtener todos los usuarios", Error: err.Error()})
	}

	// En caso de haber salido todo bien se retorna la informacion de los registros de entrada de los productos
	return c.JSON(http.StatusFound, response.Response{Data: productos})
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
			IDEstado         int    `json:"idEstado"`
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
		IDEstado:         requestBody.Producto.IDEstado,
	}
	// Se definene las regals de validacion para los campos del producto
	validationRules := map[string]validation.ValidationRule{
		"NumeroRegistro":   {Regex: regexp.MustCompile(`^[A-Z]{4}-\d{4}-\d{4}$`), Message: "Error en el formato del numero de registro asegurate que el formato sea: AAAA-0000-0000"},
		"Nombre":           {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ\s]+$`), Message: "El nombre no puede contener numeros"},
		"FechaFabricacion": {Regex: regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`), Message: "La fecha de fabricacion no es valida asegurese de que sea en el formato yyyy-mm-dd"},
		"FechaVencimiento": {Regex: regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`), Message: "La fecha de vencimiento no es valida asegurese de que sea en el formato yyyy-mm-dd"},
		"Descripcion":      {Regex: regexp.MustCompile(`^.+$`), Message: "La descripcion no puede estar vacia"},
		"CompuestoActivo":  {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ\s]+$`), Message: "El compuesto activo no puede contener numeros"},
		"Presentacion":     {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ\s]+$`), Message: "La presentacion no puede contener numeros"},
		"Cantidad":         {Regex: regexp.MustCompile(`^[a-zA-Z0-9]+$`), Message: "La cantidad no puede contener caracteres especiales"},
		"NumeroLote":       {Regex: regexp.MustCompile(`^[a-zA-Z0-9]+$`), Message: "El numero de lote no puede contener caracteres especiales"},
		"TamanoLote":       {Regex: regexp.MustCompile(`^[a-zA-Z0-9]+$`), Message: "El tamano de lote no puede contener caracteres especiales"},
	}
	// Se hace la validacion de los campos y en caso de no cumplir con alguna regla se devuelve un error con la informacion del campo que no cumple las reglas de validacion
	if err := validation.Validate(producto.ToMap(), validationRules); err != nil {
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
	// Se definene las reglas de validacion para los campos de registro de entrada de productos
	validationRulesEntrada := map[string]validation.ValidationRule{
		"PropositoAnalisis":      {Regex: regexp.MustCompile(`^.+$`), Message: "El proposito de analisis no puede estar vacio"},
		"CondicionesAmbientales": {Regex: regexp.MustCompile(`^.+$`), Message: "Las condiciones ambientales no pueden estar vacias"},
		"FechaRecepcion":         {Regex: regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`), Message: "La fecha de recepcion no es valida asegurese de que sea en el formato yyyy-mm-dd"},
		"FechaInicioAnalisis":    {Regex: regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`), Message: "La fecha de incio de analisis no es valida asegurese de que sea en el formato yyyy-mm-dd"},
		"FechaFinalAnalisis":     {Regex: regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`), Message: "La fecha de final de analisis no es valida asegurese de que sea en el formato yyyy-mm-dd"},
		"IDUsuario":              {Regex: regexp.MustCompile(`^[0-9]+$`), Message: "El id de usuario solo puede contener numeros"},
	}
	// Se hace la validacion de los campos y en caso de no cumplir con alguna regla se devuelve un error con la informacion del campo que no cumple las reglas de validacion
	if err := validation.Validate(entradaProducto.ToMap(), validationRulesEntrada); err != nil {
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
