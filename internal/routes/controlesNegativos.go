package routes

import (
	"labgestor-server/internal/controllers"
	"labgestor-server/internal/repository"
	"labgestor-server/middleware"

	"github.com/labstack/echo/v4"
)

type controlesNegativosHandler struct {
	controllers.ControlesNegativosController
}

func NewControlesNegativosHandler(e *echo.Echo, controller controllers.ControlesNegativosController, userRepo repository.UsuarioRepository) {
	handler := controlesNegativosHandler{controller}

	//? -----------------------------------------------------
	//? Endpoints CRUD
	//? -----------------------------------------------------

	e.POST("/controlesNegativos", handler.CrearControlesNegativos, middleware.RequireAuth(userRepo, ""))                 // Crear un nuevo registro de controles negativos
	e.GET("/controlesNegativos/:id", handler.ObtenerControlesNegativosID, middleware.RequireAuth(userRepo, ""))          // Obtener un registro de controles negativos por ID
	e.GET("/controlesNegativos/producto/:id", handler.ObtenerControlesPorProducto, middleware.RequireAuth(userRepo, "")) // Obtener todos los controles negativos por ID de producto
	e.PUT("/controlesNegativos/:id", handler.ActualizarControlesNegativos, middleware.RequireAuth(userRepo, ""))         // Actualizar un registro de controles negativos por ID
	e.DELETE("/controlesNegativos/:id", handler.EliminarControlesNegativos, middleware.RequireAuth(userRepo, ""))        // Eliminar un registro de controles negativos por ID
}
