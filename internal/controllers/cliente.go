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

// -------------------------------------
// Metodos del controlador
// -------------------------------------
func (controller clienteController) CrearCliente(c echo.Context) error {
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
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "El campo 'Nombre' es obligatorio"})
	}
	if requestBody.Direccion == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "El campo 'Direccion' es obligatorio"})
	}

	// Crear una instancia del modelo
	cliente := models.Cliente{
		Nombre:    requestBody.Nombre,
		Direccion: requestBody.Direccion,
	}

	// Se crea el cliente haciendo uso de la capa del repositorio
	if err := controller.Repo.CrearCliente(&cliente); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "No se pudo crear el cliente", "error": err.Error()})
	}

	// Si todo sale bien se retorna un estado de 200
	return c.JSON(http.StatusOK, map[string]string{"message": "Se registro el cliente con exito"})
}

func (controller clienteController) ActualizarCliente(c echo.Context) error {
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
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "El campo 'Nombre' es obligatorio",
		})
	}
	if requestBody.Direccion == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "El campo 'Direccion' es obligatorio",
		})
	}

	// Crear una instancia del modelo
	cliente := models.Cliente{
		ID:        requestBody.ID,
		Nombre:    requestBody.Nombre,
		Direccion: requestBody.Direccion,
	}

	// Llamamos al repositorio para actualizar el Cliente en la base de datos
	if err := controller.Repo.ActualizarCliente(&cliente); err != nil {
		// Si hay un error al actualizar, lo retornamos con un mensaje adecuado
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "No se pudo actualizar el cliente",
			"error":   err.Error(),
		})
	}

	// Si todo salió bien, respondemos con un estado 200 y un mensaje de éxito
	return c.JSON(http.StatusOK, map[string]string{"message": "Se actualizo el cliente con exito"})
}

func (controller clienteController) ObtenerCliente(c echo.Context) error {
	ID := c.Param("id")

	// Llamamos al repositorio para obtener el cliente por ID
	cliente, err := controller.Repo.ObtenerCliente(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "No se pudo obtener el cliente",
			"error": err.Error(),
		})
	}
	// Si todo salió bien, respondemos con un estado 200 y el cliente
	return c.JSON(http.StatusOK, cliente)
}

func (controller clienteController) ObtenerClientes(c echo.Context) error {
	// Llamamos al repositorio para obtener todos los clientes
	clientes, err := controller.Repo.ObtenerClientes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "No se pudo obtener los clientes",
			"error": err.Error(),
		})
	}

	// Si todo salió bien, respondemos con un estado 200 y la lista de clientes
	return c.JSON(http.StatusOK, clientes)
}
