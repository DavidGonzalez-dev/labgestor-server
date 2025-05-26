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
<<<<<<< HEAD
	e.POST("/pruebasRecuento/crear", handler.CrearPruebaRecuento)                         // Crear una nueva prueba de recuento
	e.GET("/pruebasRecuento/:id", handler.ObtenerPruebaRecuentoID)                        // Obtener una prueba de recuento por numero de registro de producto
	e.PUT("/pruebasRecuento/actualizar", handler.ActualizarPruebaRecuento)                // Actualizar una prueba de recuento
	e.GET("/pruebasRecuento/producto/:numeroRegistro", handler.ObtenerPruebasPorProducto) // Obtener una prueba de recuento por id
	e.DELETE("/pruebasRecuento/eliminar/:id", handler.EliminarPruebaRecuento)             // Eliminar una prueba de recuento
=======
	e.POST("/pruebasRecuento", handler.CrearPruebaRecuento)                               // Crear una nueva prueba de recuento
	e.GET("/pruebasRecuento/:id", handler.ObtenerPruebaRecuentoID)                        // Obtener una prueba de recuento por numero de registro de producto
	e.PUT("/pruebasRecuento/:id", handler.ActualizarPruebaRecuento)                       // Actualizar una prueba de recuento
	e.GET("/pruebasRecuento/producto/:numeroRegistro", handler.ObtenerPruebasPorProducto) // Obtener una prueba de recuento por id
	e.DELETE("/pruebasRecuento/:id", handler.EliminarPruebaRecuento)                      // Eliminar una prueba de recuento

>>>>>>> 0fee36fb23a0a4d5dd82e7dbec779d4f64aafe56
}
