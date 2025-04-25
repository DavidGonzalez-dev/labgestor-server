package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"labgestor-server/utils/validation"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Interfaz que define los metodos del controlador
type ClienteController interface {
	CrearCliente(c echo.Context) error
	ActualizarCliente(c echo.Context) error
	ObtenerCliente(c echo.Context) error
	ObtenerClientes(c echo.Context) error
}

// Structura que conecte con el repositorio
type clienteController struct {
	Repo repository.ClienteRepository
}

// Funcion de instancia de controlador
func NewClienteController(repo repository.ClienteRepository) ClienteController {
	return clienteController{Repo: repo}
}

// ? -------------------------------------
// ? Metodos del controlador
// ? -------------------------------------
// Este handler permite crear un cliente en la base de datos.
func (controller clienteController) CrearCliente(c echo.Context) error {

	//? --------------------------------------------------------------------
	//? Leemos el cuerpo del request
	//? --------------------------------------------------------------------
	// Se lee el cuerpo del request
	var requestBody struct {
		Nombre    string `json:"nombre"`
		Direccion string `json:"direccion"`
	}
	// Verificamos si hubo un error leyendo el cuerpo del request
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "No se pudo leer el cuerpo del request", Error: err.Error()})
	}

	//? --------------------------------------------------------------------
	//? Se valida la informacion del request
	//? --------------------------------------------------------------------
	// Crear una instancia del modelo
	cliente := models.Cliente{
		Nombre:    requestBody.Nombre,
		Direccion: requestBody.Direccion,
	}
	// Se crean las reglas de validacion para el registro de un cliente
	validationRules := map[string]validation.ValidationRule{
		"Nombre":    {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ\s]+$`), Message: "El nombre no puede contener numeros"},
		"Direccion": {Regex: regexp.MustCompile(`^(?i)(cra|cr|calle|cl|av|avenida|transversal|tv|diag|dg|manzana|mz|circular|circ)[a-z]*\.?\s*\d+[a-zA-Z]?\s*(#|n°|no\.?)\s*\d+[a-zA-Z]?(?:[-]\d+)?$`), Message: "Ingrese una direccion valida"},
	}
	// Se verifica si las validaciones de campos fueron exitosas y en caso contrario se retorna un mensaje con el error en el campo respectivo
	if err := validation.Validate(cliente.ToMap(), validationRules); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	//? --------------------------------------------------------------------
	//? Se crea el registro en la base de datos haciendo uso del repositorio
	//? --------------------------------------------------------------------
	// Se crea el cliente haciendo uso de la capa del repositorio
	if err := controller.Repo.CrearCliente(&cliente); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al crear el cliente", Error: err.Error()})
	}
	// Si todo sale bien se retorna un estado de 200
	return c.JSON(http.StatusCreated, response.Response{Message: "Cliente creado con exito"})
}

// Este handler permite actualizar un registro de un cliente en la base de datos siempre y cuando el id del registro exista.
func (controller clienteController) ActualizarCliente(c echo.Context) error {
	//? -------------------------------------------------------------------
	//? Se lee el cuerpo de request
	//? -------------------------------------------------------------------
	var requestBody struct {
		ID        int    `json:"id"`
		Nombre    string `json:"nombre"`
		Direccion string `json:"direccion"`
	}
	// Se verifica si hubo alguin error al leer el cuerpo del request.
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Error al leer el cuerpo de request", Error: err.Error()})
	}

	//? -------------------------------------------------------------------
	//? Buscamos y traemos el registro del cliente desde la base de datos
	//? -------------------------------------------------------------------
	// Traemos el registro del fabricantre haciendo uso de la capa del repositorio
	cliente, err := controller.Repo.ObtenerClienteID(requestBody.ID)
	// Verificamos que ID que se paso corresponde a un cliente
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "El usuario que se quiere actualizar no existe"})
	}

	//? --------------------------------------------------------------------------------
	//? Actualizamos la informacion del cliente y definimos las reglas de validacion
	//? --------------------------------------------------------------------------------
	// Actualizamos los datos del cliente
	cliente.Nombre = requestBody.Nombre
	cliente.Direccion = requestBody.Direccion

	// Definimos las reglas de validacion
	validationRules := map[string]validation.ValidationRule{
		"Nombre":    {Regex: regexp.MustCompile(`^[a-zA-Z\s]+$`), Message: "El nombre solo puede contener letras"},
		"Direccion": {Regex: regexp.MustCompile(`^(?i)(cra|cr|calle|cl|av|avenida|transversal|tv|diag|dg|manzana|mz|circular|circ)[a-z]*\.?\s*\d+[a-zA-Z]?\s*(#|n°|no\.?)\s*\d+[a-zA-Z]?(?:[-]\d+)?$`), Message: "Ingrese una direccion valida ejm:AV 45 n° 12, Calle 105 n° 8"},
	}
	// Validamos los campos del cliente y devolvemos en error en caso de no haber cumplido con alguna regla
	if err := validation.Validate(cliente.ToMap(), validationRules); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	//? -------------------------------------------------------------------
	//? Actualizamos la informacion del cliente en la base de datos
	//? -------------------------------------------------------------------
	// Llamamos al repositorio para actualizar el cliente en la base de datos
	if err := controller.Repo.ActualizarCliente(cliente); err != nil {
		// Si hay un error al actualizar, lo retornamos con un mensaje adecuado
		return c.JSON(http.StatusNotModified, response.Response{Message: "No se pudo actualizar el cliente", Error: err.Error()})
	}
	// Si todo salió bien, respondemos con un estado 200 y un mensaje de éxito
	return c.JSON(http.StatusOK, response.Response{Message: "Se actualizo el cliente con exito"})
}

// Este handler permite obtener la informacion de un cliente dependiendo del id proporcionado
func (controller clienteController) ObtenerCliente(c echo.Context) error {
	// ? ----------------------------------------------------------------------
	// ? Obtenemos el ID y lo convertimos en un entero
	// ? ----------------------------------------------------------------------
	// Obtenemos el id desde el aprametro del endpoint
	idParam := c.Param("id")
	// Convertimos el id en un entero y nos cersioramos que nos hallan pasado un numero como id
	idCliente, err := strconv.Atoi(idParam)
	// Verficamos que el usuario nos halla pasado un numero como parametro
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El id del cliente tiene que ser un numero entero"})
	}

	// ? ----------------------------------------------------------------------
	// ? Obtenemos el registro del cliente desde el repositorio
	// ? ----------------------------------------------------------------------
	// Llamamos al repositorio para obtener el cliente por ID
	cliente, err := controller.Repo.ObtenerClienteID(idCliente)
	// Verificamos que el usuario exista
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Registro no encontrado", Error: err.Error()})
	}
	// Si todo salió bien, respondemos con un estado 200 y el cliente
	return c.JSON(http.StatusOK, response.Response{Data: cliente})
}

// Este hanlder permite obtener un slice con todos los registros de los clientes
func (controller clienteController) ObtenerClientes(c echo.Context) error {
	//? ---------------------------------------------------------------------------
	//? Obtenemos todos los registros de los clientes haciendo uso de la capa del repositorio
	//? ---------------------------------------------------------------------------
	// Llamamos al repositorio para obtener todos los clientes
	clientes, err := controller.Repo.ObtenerClientes()
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Error al obtener la informacion", Error: err.Error()})
	}
	// Si todo salió bien, respondemos con un estado 200 y la lista de clientes
	return c.JSON(http.StatusOK, response.Response{Data: clientes})
}
