package routes

import (
	"labgestor-server/internal/controllers"
	"labgestor-server/internal/repository"
	"labgestor-server/middleware"

	"github.com/labstack/echo/v4"
)

type pruebaRecuentoHandler struct {
	controllers.PruebaRecuentoController
}

func NewPruebaRecuentoHandler(e *echo.Echo, controller controllers.PruebaRecuentoController, userRepo repository.UsuarioRepository) {
	handler := pruebaRecuentoHandler{controller}

	//? -----------------------------------------------------
	//? Endpoints CRUD
	//? -----------------------------------------------------
	e.POST("/pruebasRecuento", handler.CrearPruebaRecuento,  middleware.RequireAuth(userRepo, ""))                               // Crear una nueva prueba de recuento
	e.GET("/pruebasRecuento/:id", handler.ObtenerPruebaRecuentoID,  middleware.RequireAuth(userRepo, ""))                        // Obtener una prueba de recuento por numero de registro de producto
	e.PUT("/pruebasRecuento/:id", handler.ActualizarPruebaRecuento,  middleware.RequireAuth(userRepo, ""))  
	e.PATCH("/actualizarEstadoPruebaRecuento/:id", handler.ActualizarEstadoPrueba,  middleware.RequireAuth(userRepo, ""))                     // Actualizar una prueba de recuento
	e.GET("/pruebasRecuento/producto/:numeroRegistro", handler.ObtenerPruebasPorProducto,  middleware.RequireAuth(userRepo, "")) // Obtener una prueba de recuento por id
	e.DELETE("/pruebasRecuento/:id", handler.EliminarPruebaRecuento,  middleware.RequireAuth(userRepo, ""))                      // Eliminar una prueba de recuento

}
