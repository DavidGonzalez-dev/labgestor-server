package routes

import (
	"labgestor-server/internal/controllers"

	"labgestor-server/internal/repository"
	"labgestor-server/middleware"

	"github.com/labstack/echo/v4"
)

type MonitoreosDeteccionesHandler struct {
	Controller controllers.MonitoreosDeteccionesController
}

func NewMonitoreosDeteccionesHandler(e *echo.Echo, controller controllers.MonitoreosDeteccionesController, userRepo repository.UsuarioRepository) {
	handler := MonitoreosDeteccionesHandler{Controller: controller}

	// ? ----------------------------------------------------------------------
	// ? Puntos de entrada a la API
	// ? ----------------------------------------------------------------------


	e.POST("/monitoreosDetecciones", handler.Controller.CrearMonitoreosDetecciones, middleware.RequireAuth(userRepo, ""))                              // Registrar una nueva detección de monitoreo
	e.GET("/monitoreosDetecciones/detecciones/:id", handler.Controller.ObtenerMonitoreosDeteccionesPorDeteccion, middleware.RequireAuth(userRepo, "")) // Obtener detecciones de monitoreo por ID de detección
	e.PUT("/monitoreosDetecciones/:id", handler.Controller.ActualizarMonitoreosDetecciones, middleware.RequireAuth(userRepo, ""))                      // Actualizar una detección de monitoreo
	e.DELETE("/monitoreosDetecciones/:id", handler.Controller.EliminarMonitoreosDetecciones, middleware.RequireAuth(userRepo, ""))                     // Eliminar una detección de monitoreo
}
