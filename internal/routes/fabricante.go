package routes

import (
	"labgestor-server/internal/controllers"

	"github.com/labstack/echo/v4"
)

// Estructura que usa a la capa de controllers
type fabricanteHandler struct {
	Controller controllers.FabricanteController
}

// Funcion que instancia el hanlder
func NewFabricanteHandler(e *echo.Echo, controller controllers.FabricanteController) {
	// Instanciar el handler
	handler := fabricanteHandler{Controller: controller}

	// ? ----------------------------------------------------------------------
	// ? Puntos de entrada a la API
	// ? ----------------------------------------------------------------------

	e.POST("/fabricantes/registrar", handler.Controller.CrearFabricante)      // Registrar un nuevo fabricante
	e.PUT("/fabricantes/actualizar", handler.Controller.ActualizarFabricante) // Actualizar la informacion de un fabricante en especifico
	e.GET("/fabricantes/:id", handler.Controller.ObtenerFabricante)           // Obtener informacion de un fabricante en especifico
	e.GET("/fabricantes/", handler.Controller.ObtenerFabricantes)             // Obtener la informacion de todos los fabricantes
	e.DELETE("/fabricantes/:id", handler.Controller.EliminarFabricante)       // Eliminar un fabricante en especifico
}
