package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Interfaz que define los metodos del controlador
type FabricanteController interface {
	CrearFabricante(c echo.Context) error
	ActualizarFabricante(c echo.Context) error
	ObtenerFabricante(c echo.Context) error
	ObtenerFabricantes(c echo.Context) error
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
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	if requestBody.Nombre == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Nombre' es obligatorio"})
	}
	if requestBody.Direccion == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Direccion' es obligatorio"})
	}
	// Crear una instancia del modelo
	fabricante := models.Fabricante{
		Nombre:    requestBody.Nombre,
		Direccion: requestBody.Direccion,
	}

	// Se crea el fabricante haciendo uso de la capa del repositorio
	if err := controller.Repo.CrearFabricante(&fabricante); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al crear el fabricante", Error: err.Error()})
	}

	// Si todo sale bien se retorna un estado de 200
	return c.JSON(http.StatusOK, response.Response{Message: "Se registro el fabricante con exito"})
}

func (controller fabricanteController) ActualizarFabricante(c echo.Context) error {
	var requestBody struct {
		ID        int
		Nombre    string
		Direccion string
	}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo de request", Error: err.Error()})
	}

	// Realizar validaciones de campos
	if requestBody.Nombre == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Nombre' es obligatorio"})
	}
	if requestBody.Direccion == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El campo 'Direccion' es obligatorio"})
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
		return c.JSON(http.StatusInternalServerError, response.Response{
			Message: "No se pudo actualizar el fabricante",
			Error:   err.Error(),
		})
	}

	// Si todo salió bien, respondemos con un estado 200 y un mensaje de éxito
	return c.JSON(http.StatusOK, response.Response{Message: "Se actualizo el fabricante con exito"})
}

func (controller fabricanteController) ObtenerFabricante(c echo.Context) error {
	ID := c.Param("id")

	// Llamamos al repositorio para obtener el fabricante por ID
	fabricante, err := controller.Repo.ObtenerFabricante(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al obtener el fabricante", Error: err.Error()})
	}
	// Si todo salió bien, respondemos con un estado 200 y el cliente
	return c.JSON(http.StatusOK, response.Response{Data: fabricante})
}

func (controller fabricanteController) ObtenerFabricantes(c echo.Context) error {
	fabricantes, err := controller.Repo.ObtenerFabricantes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al obtener los farbicantes", Error: err.Error()})
	}
	// Si todo salió bien, respondemos con un estado 200 y el cliente
	return c.JSON(http.StatusOK, response.Response{Data: fabricantes})
}
