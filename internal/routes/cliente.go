package routes

import (
	"labgestor-server/internal/controllers"

	"github.com/labstack/echo/v4"
)

// Estructura que usa a la capa de controllers
type clienteHandler struct {
	Controller controllers.ClienteController
}

// Funcion que instancia el hanlder
func NewClienteHandler(e *echo.Echo, controller controllers.ClienteController) {
	// Instanciar el handler
	hanlder := clienteHandler{Controller: controller}

	// Definir EndPoints (Puntos de entrada a la API)
	e.POST("/clientes/crearCliente", hanlder.Controller.CrearCliente)
}
