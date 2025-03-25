package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Interfaz que define los metodos del controlador
type ClienteController interface {
	CrearCliente(c echo.Context) error
}

// Structura que conecte con el repositorio
type clienteController struct {
	Repo repository.ClienteRepository
}

// Funcion de instancia de controlador
func NewClienteController(repo repository.ClienteRepository) ClienteController {
	return clienteController{Repo: repo}
}

// -------------------------------------
// Metodos del controlador
// -------------------------------------
func (controller clienteController) CrearCliente(c echo.Context) error {
	// Se lee el cuerpo del request
	var requestBody struct {
		Nombre   string
		Direccion string
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "No se pudo leer el cuerpo de request", "error": err.Error()})
	}

	// TODO: Realizar validaciones de campos

	// Crear una instancia del modelo
	cliente := models.Cliente{
		Nombre:   requestBody.Nombre,
		Direccion: requestBody.Direccion,
	}

	// Se crea el cliente haciendo uso de la capa del repositorio
	if err := controller.Repo.CrearCliente(&cliente); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "No se pudo crear el cliente", "error": err.Error()})
	}

	// Si todo sale bien se retorna un estado de 200
	return c.JSON(http.StatusOK, map[string]string{"message": "Se registro el cliente con exito"})
}
