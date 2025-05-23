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
	e.POST("/controlesNegativos", handler.CrearControlesNegativos) // Crear un nuevo registro de controles negativos
}
