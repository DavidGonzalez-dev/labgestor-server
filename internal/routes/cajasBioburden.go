package routes

import (
	"labgestor-server/internal/controllers"

	"labgestor-server/internal/repository"
	"labgestor-server/middleware"

	"github.com/labstack/echo/v4"
)

type CajasBioburdenRoutes struct {
	// Controller es la capa de controladores que maneja las operaciones de cajas de bioburden
	Controller controllers.CajasBioburdenController
}

func NewCajasBioburdenHandler(e *echo.Echo, controller controllers.CajasBioburdenController, userRepo repository.UsuarioRepository) {
	handler := CajasBioburdenRoutes{Controller: controller}

	// Definir EndPoints (Puntos de entrada a la API)
	e.POST("/cajasBioburden", handler.Controller.CrearCajaBioburden, middleware.RequireAuth(userRepo, ""))                              // Crear una nueva caja de bioburden
	e.GET("/cajasBioburden/:id", handler.Controller.ObtenerCajasBioburdenID, middleware.RequireAuth(userRepo, ""))                      // Obtener una caja de bioburden por ID
	e.GET("/cajasBioburden/pruebaRecuento/:id", handler.Controller.ObtenerCajasPorPruebaRecuento, middleware.RequireAuth(userRepo, "")) // Obtener cajas de bioburden por ID de prueba de recuento
	e.PUT("/cajasBioburden/:id", handler.Controller.ActualizarCajaBioburden, middleware.RequireAuth(userRepo, ""))                      // Actualizar una caja de bioburden
	e.DELETE("/cajasBioburden/:id", handler.Controller.EliminarCajaBioburden, middleware.RequireAuth(userRepo, ""))                     // Eliminar una caja de bioburden
}
