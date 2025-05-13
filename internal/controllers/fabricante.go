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
type FabricanteController interface {
	CrearFabricante(c echo.Context) error
	ActualizarFabricante(c echo.Context) error
	ObtenerFabricante(c echo.Context) error
	ObtenerFabricantes(c echo.Context) error
	EliminarFabricante(c echo.Context) error
}

// Structura que conecte con el repositorio
type fabricanteController struct {
	Repo repository.FabricanteRepository
}

// Funcion de instancia de controlador
func NewFabricanteController(repo repository.FabricanteRepository) FabricanteController {
	return fabricanteController{Repo: repo}
}

// ? -------------------------------------
// ? Metodos del controlador
// ? -------------------------------------
// Este handler se encarga de crear un fabricante en la base de datos con base en la informacion pasada en el request.
func (controller fabricanteController) CrearFabricante(c echo.Context) error {
	//? -------------------------------------------------------------------
	//? Se lee el cuerpo de request
	//? -------------------------------------------------------------------
	var requestBody struct {
		Nombre    string `json:"nombre"`
		Direccion string `json:"direccion"`
	}
	// Se verifica si hubo un error al leer el cuerpo del request y en caso de haber un error se retorna un mensaje de error
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	//? -------------------------------------------------------------------
	//? Se crea una instancia del modelo del fabricante y se validan los
	//? campos del request.
	//? -------------------------------------------------------------------
	fabricante := models.Fabricante{
		Nombre:    requestBody.Nombre,
		Direccion: requestBody.Direccion,
	}
	// Se definen las reglas de validacion para el modelo del fabricante
	validationRules := map[string]validation.ValidationRule{
		"Nombre":    {Regex: regexp.MustCompile(`^[a-zA-Z\s]+$`), Message: "El nombre no puede contener numeros"},
		"Direccion": {Regex: regexp.MustCompile(`^(?i)(cra|cr|calle|cl|av|avenida|transversal|tv|diag|dg|manzana|mz|circular|circ)[a-z]*\.?\s*\d+[a-zA-Z]?\s*(#|n°|no\.?)\s*\d+[a-zA-Z]?(?:[-]\d+)?$`), Message: "Ingrese una direccion valida ejm:AV 45 n° 12, Calle 105 n° 8"},
	}
	// Se hace la validacion con base a las reglas establecidas y en caso de no cumplir alguna regla se decuelve un error informando acerca del error
	if err := validation.Validate(fabricante.ToMap(), validationRules); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	//? -------------------------------------------------------------------
	//? Se crea el fabricante en la base de datos
	//? -------------------------------------------------------------------
	// Se crea el fabricante haciendo uso de la capa del repositorio y en caso de presentarse un error se devuelve el mensaje de error
	if err := controller.Repo.CrearFabricante(&fabricante); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al crear el fabricante", Error: err.Error()})
	}
	// Si todo sale bien se retorna un estado de 200
	return c.JSON(http.StatusCreated, response.Response{Message: "Se registro el fabricante con exito"})
}

// Este handler se encarga de actualizar la informacion de un fabricante sercioandose que este exista en la base de datos
func (controller fabricanteController) ActualizarFabricante(c echo.Context) error {
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
	//? Buscamos y traemos el registro del fabricante desde la base de datos
	//? -------------------------------------------------------------------
	// Traemos el registro del fabricantre haciendo uso de la capa del repositorio
	fabricante, err := controller.Repo.ObtenerFabricanteID(requestBody.ID)

	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "El usuario que se quiere actualizar no existe"})
	}

	//? --------------------------------------------------------------------------------
	//? Actualizamos la informacion del fabricante y definimos las reglas de validacion
	//? --------------------------------------------------------------------------------
	// Actualizamos los datos del fabricante
	fabricante.Nombre = requestBody.Nombre
	fabricante.Direccion = requestBody.Direccion

	// Definimos las reglas de validacion
	validationRules := map[string]validation.ValidationRule{
		"Nombre":    {Regex: regexp.MustCompile(`^[a-zA-Z\s]+$`), Message: "El nombre no puede contener numeros"},
		"Direccion": {Regex: regexp.MustCompile(`^(?i)(cra|cr|calle|cl|av|avenida|transversal|tv|diag|dg|manzana|mz|circular|circ)[a-z]*\.?\s*\d+[a-zA-Z]?\s*(#|n°|no\.?)\s*\d+[a-zA-Z]?(?:[-]\d+)?$`), Message: "Ingrese una direccion valida ejm:AV 45 n° 12, Calle 105 n° 8"},
	}

	// Validamos los campos del fabricante y devolvemos en error en caso de no haber cumplido con alguna regla
	if err := validation.Validate(fabricante.ToMap(), validationRules); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	//? -------------------------------------------------------------------
	//? Actualizamos la informacion del fabricante en la base de datos
	//? -------------------------------------------------------------------
	// Llamamos al repositorio para actualizar el fabricante en la base de datos
	if err := controller.Repo.ActualizarFabricante(fabricante); err != nil {
		// Si hay un error al actualizar, lo retornamos con un mensaje adecuado
		return c.JSON(http.StatusNotModified, response.Response{Message: "No se pudo actualizar el fabricante", Error: err.Error()})
	}

	// Si todo salió bien, respondemos con un estado 200 y un mensaje de éxito
	return c.JSON(http.StatusOK, response.Response{Message: "Se actualizo el fabricante con exito"})
}

// Este handler se encarga de obtener la informacion de un fabricante basandonos en el ID que se le pase como parametro al endpoint
func (controller fabricanteController) ObtenerFabricante(c echo.Context) error {
	// ? ----------------------------------------------------------------------
	// ? Obtenemos el ID y lo convertimos en un entero
	// ? ----------------------------------------------------------------------
	// Obtenemos el id desde el aprametro del endpoint
	idParam := c.Param("id")
	// Convertimos el id en un entero y nos cersioramos que nos hallan pasado un numero como id
	idFabricante, err := strconv.Atoi(idParam)
	// Verficamos que el usuario nos halla pasado un numero como parametro
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El id del fabricante tiene que ser un numero entero"})
	}

	// ? ----------------------------------------------------------------------
	// ? Obtenemos el registro del fabricante desde el repositorio
	// ? ----------------------------------------------------------------------
	// Llamamos al repositorio para obtener el fabricante por ID
	fabricante, err := controller.Repo.ObtenerFabricanteID(idFabricante)
	// Verificamos que el usuario exista
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Registro no encontrado", Error: err.Error()})
	}
	// Si todo salió bien, respondemos con un estado 200 y el cliente
	return c.JSON(http.StatusOK, response.Response{Data: fabricante})
}

// Este handler se encarga de obtener la informacion de todos los fabricantes.
func (controller fabricanteController) ObtenerFabricantes(c echo.Context) error {
	// ? ----------------------------------------------------------------------
	// ? Obtenemos todos los registros de los fabricantes
	// ? ----------------------------------------------------------------------
	// Obtenemos todos los registros de los fabricantes haciendo uso de la capa del repositorio
	fabricantes, err := controller.Repo.ObtenerFabricantes()
	// Verificamos que se hallan retornado de manera exitosa los registros de los fabricnates
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Error al obtener los farbicantes", Error: err.Error()})
	}
	// Si todo salió bien, respondemos con un estado 200 y el cliente
	return c.JSON(http.StatusOK, response.Response{Data: fabricantes})
}

// Este handler se encarga de eliminar un fabricante de la base de datos
func (controller fabricanteController) EliminarFabricante(c echo.Context) error {
	// ? ----------------------------------------------------------------------
	// ? Obtenemos el ID y lo convertimos en un entero
	// ? ----------------------------------------------------------------------
	// Obtenemos el id desde el aprametro del endpoint
	idParam := c.Param("id")
	// Convertimos el id en un entero y nos cersioramos que nos hallan pasado un numero como id
	idFabricante, err := strconv.Atoi(idParam)
	// Verficamos que el usuario nos halla pasado un numero como parametro
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El id del fabricante tiene que ser un numero entero"})
	}

	// ? ----------------------------------------------------------------------
	// ? Obtenemos el registro del fabricante desde el repositorio
	// ? ----------------------------------------------------------------------
	// Llamamos al repositorio para obtener el fabricante por ID
	fabricante, err := controller.Repo.ObtenerFabricanteID(idFabricante)
	// Verificamos que el usuario exista
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Registro no encontrado", Error: err.Error()})
	}
	
	// ? ----------------------------------------------------------------------
	// ? Eliminamos el registro del fabricante desde la base de datos
	// ? ----------------------------------------------------------------------
	if err := controller.Repo.EliminarFabricante(fabricante); err != nil {
		return c.JSON(http.StatusConflict, response.Response{Message: "Error al eliminar el fabricante", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Response{Message: "Se elimino"})
}
