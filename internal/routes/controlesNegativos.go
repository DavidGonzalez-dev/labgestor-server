package routes

import (
	"labgestor-server/internal/controllers"

	"github.com/labstack/echo/v4"
)

type controlesNegativosHandler struct {
	controllers.ControlesNegativosController
}

func NewControlesNegativosHandler(e *echo.Echo, controller controllers.ControlesNegativosController) {
	handler := controlesNegativosHandler{controller}

	//? -----------------------------------------------------
	//? Endpoints CRUD
	//? -----------------------------------------------------
	e.POST("/controlesNegativos", handler.CrearControlesNegativos)                 // Crear un nuevo registro de controles negativos
	e.GET("/controlesNegativos/:id", handler.ObtenerControlesNegativosID)          // Obtener un registro de controles negativos por ID
	e.GET("/controlesNegativos/producto/:id", handler.ObtenerControlesPorProducto) // Obtener todos los controles negativos por ID de producto
	e.PUT("/controlesNegativos/:id", handler.ActualizarControlesNegativos)         // Actualizar un registro de controles negativos por ID
	e.DELETE("/controlesNegativos/:id", handler.EliminarControlesNegativos)        // Eliminar un registro de controles negativos por ID
}
