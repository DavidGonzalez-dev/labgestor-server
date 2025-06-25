package routes

import (
	"labgestor-server/internal/controllers"
	"labgestor-server/internal/repository"
	"labgestor-server/middleware"

	"github.com/labstack/echo/v4"
)

// Estructura que usa a la capa de controllers
type fabricanteHandler struct {
	Controller controllers.FabricanteController
}

// Funcion que instancia el hanlder
func NewFabricanteHandler(e *echo.Echo, controller controllers.FabricanteController, userRepo repository.UsuarioRepository) {
	// Instanciar el handler
	handler := fabricanteHandler{Controller: controller}

	// ? ----------------------------------------------------------------------
	// ? Puntos de entrada a la API
	// ? ----------------------------------------------------------------------

	e.POST("/fabricantes", handler.Controller.CrearFabricante, middleware.RequireAuth(userRepo, ""))          // Registrar un nuevo fabricante
	e.PUT("/fabricantes/:id", handler.Controller.ActualizarFabricante, middleware.RequireAuth(userRepo, ""))  // Actualizar la informacion de un fabricante en especifico
	e.GET("/fabricantes/:id", handler.Controller.ObtenerFabricante, middleware.RequireAuth(userRepo, ""))     // Obtener informacion de un fabricante en especifico
	e.GET("/fabricantes", handler.Controller.ObtenerFabricantes, middleware.RequireAuth(userRepo, ""))        // Obtener la informacion de todos los fabricantes
	e.DELETE("/fabricantes/:id", handler.Controller.EliminarFabricante, middleware.RequireAuth(userRepo, "")) // Eliminar un fabricante en especifico
}
