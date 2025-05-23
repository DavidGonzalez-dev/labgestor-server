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
	e.POST("/pruebasRecuento", handler.CrearPruebaRecuento)                               // Crear una nueva prueba de recuento
	e.GET("/pruebasRecuento/:id", handler.ObtenerPruebaRecuentoID)                        // Obtener una prueba de recuento por numero de registro de producto
	e.PUT("/pruebasRecuento/:id", handler.ActualizarPruebaRecuento)                       // Actualizar una prueba de recuento
	e.GET("/pruebasRecuento/producto/:numeroRegistro", handler.ObtenerPruebasPorProducto) // Obtener una prueba de recuento por id
	e.DELETE("/pruebasRecuento/:id", handler.EliminarPruebaRecuento)                      // Eliminar una prueba de recuento

}
