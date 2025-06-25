package routes

import (
	"labgestor-server/internal/controllers"

	"github.com/labstack/echo/v4"
)

type MonitoreosDeteccionesHandler struct {
	Controller controllers.MonitoreosDeteccionesController
}

func NewMonitoreosDeteccionesHandler(e *echo.Echo, controller controllers.MonitoreosDeteccionesController) {
	handler := MonitoreosDeteccionesHandler{Controller: controller}

	// ? ----------------------------------------------------------------------
	// ? Puntos de entrada a la API
	// ? ----------------------------------------------------------------------

	e.POST("/monitoreosDetecciones", handler.Controller.CrearMonitoreosDetecciones)                              // Registrar una nueva detección de monitoreo
	e.GET("/monitoreosDetecciones/detecciones/:id", handler.Controller.ObtenerMonitoreosDeteccionesPorDeteccion) // Obtener detecciones de monitoreo por ID de detección
	e.PUT("/monitoreosDetecciones/:id", handler.Controller.ActualizarMonitoreosDetecciones)                      // Actualizar una detección de monitoreo
	e.DELETE("/monitoreosDetecciones/:id", handler.Controller.EliminarMonitoreosDetecciones)                     // Eliminar una detección de monitoreo
}
