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
	handler := clienteHandler{Controller: controller}

	// Definir EndPoints (Puntos de entrada a la API)
	e.POST("/clientes/registrar", handler.Controller.CrearCliente)      // Registro de cliente
	e.PUT("/clientes/actualizar", handler.Controller.ActualizarCliente) // Actualizar Cliente
	e.GET("/clientes/:id", handler.Controller.ObtenerCliente)           // Obtener informacion de un cliente en especifico
	e.GET("/clientes", handler.Controller.ObtenerClientes)              // Obtener informacion de todos los clientes
	e.DELETE("/clientes/:id", handler.Controller.EliminarCliente)       // Eliminar cliente
}
