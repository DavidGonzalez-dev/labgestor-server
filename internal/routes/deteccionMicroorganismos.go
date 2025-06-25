package routes

import (
	"labgestor-server/internal/controllers"

	"github.com/labstack/echo/v4"
)

type deteccionMicroorganismosHandler struct {
	controllers.DeteccionMicroorganismosController
}

func NewDeteccionMicroorganismosHandler(e *echo.Echo, controller controllers.DeteccionMicroorganismosController) {
	handler := deteccionMicroorganismosHandler{controller}

	//? -----------------------------------------------------
	//? Endpoints CRUD
	//? -----------------------------------------------------
	e.POST("/deteccionMicroorganismos", handler.CrearDeteccionMicroorganismos)                          // Crear un nuevo registro de deteccion de microorganismos
	e.GET("/deteccionMicroorganismos/:id", handler.ObtenerDeteccionMicroorganismosID)                   // Obtener un registro de deteccion de microorganismos por ID
	e.PUT("/deteccionMicroorganismos/:id", handler.ActualizarDeteccionMicroorganismos)                  // Actualizar un registro de deteccion de microorganismos por ID
	e.GET("/deteccionMicroorganismos/producto/:id", handler.ObtenerDeteccionMicroorganismosPorProducto) // Obtener todos los registros de deteccion de microorganismos por numero de registro de producto
	e.PATCH("/terminarDeteccionMicroorganismos/:id", handler.TerminarDeteccionMicroorganismos)
	e.DELETE("/deteccionMicroorganismos/:id", handler.EliminarDeteccionMicroorganismos)                 // Eliminar un registro de deteccion de microorganismos por ID
}
