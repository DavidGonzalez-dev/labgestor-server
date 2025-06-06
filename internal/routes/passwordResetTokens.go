package routes

import (
	"labgestor-server/internal/controllers"

	"github.com/labstack/echo/v4"
)

type passwordResetTokenHandler struct {
	Controller controllers.PasswordResetTokensController
}

// Funcion para instanciar el manejador de rutas
func NewPasswordResetTokensHandler(e *echo.Echo,controller controllers.PasswordResetTokensController) {
	handler := passwordResetTokenHandler{Controller: controller}

	// ? ----------------------------------------------------------------------
	// ? Puntos de entrada a la API
	// ? ----------------------------------------------------------------------
	e.POST("/passwordResetMail", handler.Controller.SendEmailWithToken)
}