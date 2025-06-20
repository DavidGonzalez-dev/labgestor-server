package routes

import (
	"labgestor-server/internal/controllers"

	"github.com/labstack/echo/v4"
)

type CajasBioburdenRoutes struct {
	// Controller es la capa de controladores que maneja las operaciones de cajas de bioburden
	Controller controllers.CajasBioburdenController
}

func NewCajasBioburdenHandler(e *echo.Echo, controller controllers.CajasBioburdenController) {
	handler := CajasBioburdenRoutes{Controller: controller}

	// Definir EndPoints (Puntos de entrada a la API)
	e.POST("/cajasBioburden", handler.Controller.CrearCajaBioburden)                              // Crear una nueva caja de bioburden
	e.GET("/cajasBioburden/:id", handler.Controller.ObtenerCajasBioburdenID)                      // Obtener una caja de bioburden por ID
	e.GET("/cajasBioburden/pruebaRecuento/:id", handler.Controller.ObtenerCajasPorPruebaRecuento) // Obtener cajas de bioburden por ID de prueba de recuento
	e.PUT("/cajasBioburden/:id", handler.Controller.ActualizarCajaBioburden)                      // Actualizar una caja de bioburden
	e.DELETE("/cajasBioburden/:id", handler.Controller.EliminarCajaBioburden)                     // Eliminar una caja de bioburden
}
