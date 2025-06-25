package routes

import (
	"labgestor-server/internal/controllers"
	"labgestor-server/internal/repository"
	"labgestor-server/middleware"

	"github.com/labstack/echo/v4"
)

type deteccionMicroorganismosHandler struct {
	controllers.DeteccionMicroorganismosController
}

func NewDeteccionMicroorganismosHandler(e *echo.Echo, controller controllers.DeteccionMicroorganismosController, userRepo repository.UsuarioRepository) {
	handler := deteccionMicroorganismosHandler{controller}

	//? -----------------------------------------------------
	//? Endpoints CRUD
	//? -----------------------------------------------------
	e.POST("/deteccionMicroorganismos", handler.CrearDeteccionMicroorganismos, middleware.RequireAuth(userRepo, ""))                          // Crear un nuevo registro de deteccion de microorganismos
	e.GET("/deteccionMicroorganismos/:id", handler.ObtenerDeteccionMicroorganismosID, middleware.RequireAuth(userRepo, ""))                   // Obtener un registro de deteccion de microorganismos por ID
	e.PUT("/deteccionMicroorganismos/:id", handler.ActualizarDeteccionMicroorganismos, middleware.RequireAuth(userRepo, ""))                  // Actualizar un registro de deteccion de microorganismos por ID
	e.GET("/deteccionMicroorganismos/producto/:id", handler.ObtenerDeteccionMicroorganismosPorProducto, middleware.RequireAuth(userRepo, "")) // Obtener todos los registros de deteccion de microorganismos por numero de registro de producto
	e.PATCH("/terminarDeteccionMicroorganismos/:id", handler.TerminarDeteccionMicroorganismos, middleware.RequireAuth(userRepo, ""))
	e.DELETE("/deteccionMicroorganismos/:id", handler.EliminarDeteccionMicroorganismos, middleware.RequireAuth(userRepo, ""))                 // Eliminar un registro de deteccion de microorganismos por ID
}
