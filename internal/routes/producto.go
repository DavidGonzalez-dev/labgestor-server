package routes

import (
	"labgestor-server/internal/controllers"

	"github.com/labstack/echo/v4"
)

type productoHandler struct {
	controllers.ProductoController
}

func NewProductoHandler(e *echo.Echo, controller controllers.ProductoController) {
	handler := productoHandler{controller}
	// Definir EndPoints (Puntos de entrada a la API)
	e.GET("/productos/:id", handler.ObtenerProductoID)              // Obtener informacion de un producto en especifico
	e.GET("/entradas-productos/", handler.ObtenerEntradasProductos) // Obtener la informacion de todos los productos
	e.POST("/productos/crear", handler.CrearProducto)               // Crear un nuevo producto
}
