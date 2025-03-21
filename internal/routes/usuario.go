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
	// Registrar Usuario
	e.POST("/registrarUsuario", handler.Controller.RegistrarUsuario)
	// Inciar Session
	e.POST("/login", handler.Controller.Login)
	// Cerrar Sesion
	e.POST("/logout", handler.Controller.Logout)
}

