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

	// Definir EndPoints (Puntos de entrada a la API)
	e.POST("/fabricante/crearFabricante", handler.Controller.CrearFabricante)
	e.POST("/fabricante/actualizarFabricante", handler.Controller.ActualizarFabricante)
	e.GET("/fabricante/obtenerFabricante/:id", handler.Controller.ObtenerFabricante)
	e.GET("/fabricante/obtenerFabricantes", handler.Controller.ObtenerFabricantes)
}
