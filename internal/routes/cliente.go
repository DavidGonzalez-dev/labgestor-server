package routes

import (
	"labgestor-server/internal/controllers"
	"labgestor-server/internal/repository"
	"labgestor-server/middleware"

	"github.com/labstack/echo/v4"
)

// Estructura que usa a la capa de controllers
type clienteHandler struct {
	Controller controllers.ClienteController
}

// Funcion que instancia el hanlder
func NewClienteHandler(e *echo.Echo, controller controllers.ClienteController, userRepo repository.UsuarioRepository) {
	// Instanciar el handler
	handler := clienteHandler{Controller: controller}

	// Definir EndPoints (Puntos de entrada a la API)
	e.POST("/clientes", handler.Controller.CrearCliente, middleware.RequireAuth(userRepo, ""))          // Registro de cliente
	e.PUT("/clientes/:id", handler.Controller.ActualizarCliente, middleware.RequireAuth(userRepo, ""))  // Actualizar Cliente
	e.GET("/clientes/:id", handler.Controller.ObtenerCliente, middleware.RequireAuth(userRepo, ""))     // Obtener informacion de un cliente en especifico
	e.GET("/clientes", handler.Controller.ObtenerClientes, middleware.RequireAuth(userRepo, ""))        // Obtener informacion de todos los clientes
	e.DELETE("/clientes/:id", handler.Controller.EliminarCliente, middleware.RequireAuth(userRepo, "")) // Eliminar cliente
}
