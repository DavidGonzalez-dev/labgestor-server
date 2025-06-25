package routes

import (
	"labgestor-server/internal/controllers"
	"labgestor-server/internal/repository"
	"labgestor-server/middleware"

	"github.com/labstack/echo/v4"
)

type productoHandler struct {
	controllers.ProductoController
}

func NewProductoHandler(e *echo.Echo, controller controllers.ProductoController, userRepo repository.UsuarioRepository) {
	handler := productoHandler{controller}

	//? -----------------------------------------------------
	//? Endpoints CRUD
	//? -----------------------------------------------------
	e.GET("/productos/:id", handler.ObtenerProductoID, middleware.RequireAuth(userRepo, ""))                                             // Obtener informacion de un producto en especifico
	e.GET("/registroEntradaProductos", handler.ObtenerRegistrosEntradaProductos, middleware.RequireAuth(userRepo, ""))                   // Obtener la informacion de todos los productos
	e.GET("registroEntradaProductos/user/:id", handler.ObtenerRegistrosEntradaProductosPorUsuario, middleware.RequireAuth(userRepo, "")) // Obtener la informacion de todos los productos de un usuario especifico
	e.POST("/productos", handler.CrearProducto, middleware.RequireAuth(userRepo, ""))                                                    // Crear un nuevo producto
	e.DELETE("/productos/:id", handler.EliminarProducto, middleware.RequireAuth(userRepo, ""))                                           // Eliminar un producto en especifico
	e.PUT("/productos/:id", handler.ActualizarProducto, middleware.RequireAuth(userRepo, ""))                                            // Actualizar un producto en especifico
	e.PUT("/registroEntradaProductos/:id", handler.ActualizarRegistroEntradaProducto, middleware.RequireAuth(userRepo, ""))              // Actualizar un registro de entrada de producto
	e.GET("/productos/:id/analisis", handler.ObtenerAnalisis, middleware.RequireAuth(userRepo, ""))                                      // Obtener los analisis de un producto
	e.PATCH("/actualizarEstadoProducto/:numeroRegistro", handler.ActualizarEstadoProducto, middleware.RequireAuth(userRepo, ""))         // Actualizar el estado de un producto
}
