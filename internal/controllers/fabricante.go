package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response struct to standardize API responses
type Response struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// Interfaz que define los metodos del controlador
type FabricanteController interface {
	CrearFabricante(c echo.Context) error
	ActualizarFabricante(c echo.Context) error
}

// Structura que conecte con el repositorio
type fabricanteController struct {
	Repo repository.FabricanteRepository
}

// Funcion de instancia de controlador
func NewFabricanteController(repo repository.FabricanteRepository) FabricanteController {
	return fabricanteController{Repo: repo}
}

// -------------------------------------
// Metodos del controlador
// -------------------------------------
func (controller fabricanteController) CrearFabricante(c echo.Context) error {
	// Se lee el cuerpo del request
	var requestBody struct {
		Nombre    string
		Direccion string
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "No se pudo leer el cuerpo de request", "error": err.Error()})
	}

	// TODO: Realizar validaciones de campos
	if requestBody.Nombre == "" {
		return c.JSON(http.StatusBadRequest, Response{
			Message: "El campo 'Nombre' es obligatorio",
		})
	}
	if requestBody.Direccion == "" {
		return c.JSON(http.StatusBadRequest, Response{
			Message: "El campo 'Direccion' es obligatorio",
		})
	}
	// Crear una instancia del modelo
	fabricante := models.Fabricante{
		Nombre:    requestBody.Nombre,
		Direccion: requestBody.Direccion,
	}

	// Se crea el fabricante haciendo uso de la capa del repositorio
	if err := controller.Repo.CrearFabricante(&fabricante); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "No se pudo crear el fabricante", "error": err.Error()})
	}

	// Si todo sale bien se retorna un estado de 200
	return c.JSON(http.StatusOK, map[string]string{"message": "Se registro el fabricante con exito"})
}

func (controller fabricanteController) ActualizarFabricante(c echo.Context) error {
	var requestBody struct {
		ID        int
		Nombre    string
		Direccion string
	}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "No se pudo leer el cuerpo del request", "error": err.Error()})
	}

	// Realizar validaciones de campos
	if requestBody.Nombre == "" {
		return c.JSON(http.StatusBadRequest, Response{
			Message: "El campo 'Nombre' es obligatorio",
		})
	}
	if requestBody.Direccion == "" {
		return c.JSON(http.StatusBadRequest, Response{
			Message: "El campo 'Direccion' es obligatorio",
		})
	}

	// Crear una instancia del modelo
	fabricante := models.Fabricante{
		ID:        requestBody.ID,
		Nombre:    requestBody.Nombre,
		Direccion: requestBody.Direccion,
	}

	// Llamamos al repositorio para actualizar el fabricante en la base de datos
	if err := controller.Repo.ActualizarFabricante(&fabricante); err != nil {
		// Si hay un error al actualizar, lo retornamos con un mensaje adecuado
		return c.JSON(http.StatusInternalServerError, Response{
			Message: "No se pudo actualizar el fabricante",
			Error:   err.Error(),
		})
	}

	// Si todo salió bien, respondemos con un estado 200 y un mensaje de éxito
	return c.JSON(http.StatusOK, map[string]string{"message": "Se actualizo el fabricante con exito"})
}
