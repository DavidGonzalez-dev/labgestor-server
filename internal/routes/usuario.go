package routes

import (
	"labgestor-server/internal/controllers"

	"github.com/labstack/echo/v4"
)

type usuarioHandler struct {
	Controller controllers.UsuarioController
}

func NewUsuarioHanlder(e *echo.Echo, controller controllers.UsuarioController) {
	handler := usuarioHandler{Controller: controller}
	e.POST("/registraUsuario", handler.Controller.RegistrarUsuario)
}
