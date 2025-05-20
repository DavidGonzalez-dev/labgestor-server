package routes

import (
	"labgestor-server/internal/controllers"

	"github.com/labstack/echo/v4"
)

type pruebaRecuentoHandler struct {
	controllers.PruebaRecuentoController
}

func NewPruebaRecuentoHandler(e *echo.Echo, controller controllers.PruebaRecuentoController) {
	handler := pruebaRecuentoHandler{controller}

	//? -----------------------------------------------------
	//? Endpoints CRUD
	//? -----------------------------------------------------
	e.POST("/pruebasRecuento/crear", handler.CrearPruebaRecuento)                       // Crear una nueva prueba de recuento
	e.GET("/pruebasRecuento/producto/:id", handler.ObtenerPruebaRecuento)               // Obtener una prueba de recuento por numero de registro de producto
	e.PUT("/pruebasRecuento/actualizar/producto/:id", handler.ActualizarPruebaRecuento) // Actualizar una prueba de recuento
}
