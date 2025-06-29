package routes

import (
	"labgestor-server/internal/controllers"
	"labgestor-server/internal/repository"
	"labgestor-server/middleware"

	"github.com/labstack/echo/v4"
)

func NewUsuarioHanlder(e *echo.Echo, controller controllers.UsuarioController, repo repository.UsuarioRepository) {
	//? -----------------------------------
	//? Rutas Publicas Autenticacion
	//? -----------------------------------
	e.POST("/login", controller.Login)                       // Iniciar Sesion de Usuario
	e.PATCH("/passwordReset", controller.CambiarContrasena, middleware.RequireResetPasswordToken()) // Actualizar la contraseña
	e.GET("/validar-token", controller.ValidarToken)         // Validar el token

	//? -----------------------------------
	//? Rutas Privadas Autenticacion
	//? -----------------------------------
	e.POST("/logout", controller.Logout) // Cerrar Sesion

	//? -----------------------------------
	//? Rutas CRUD
	//? -----------------------------------
	e.POST("/usuarios", controller.RegistrarUsuario,  middleware.RequireAuth(repo, "admin"))                                 // Registrar Usuario en el sistema
	e.PUT("/usuarios/:id", controller.ActualizarUsuario, middleware.RequireAuth(repo, ""))                             // Actualizar la info de un usuario
	e.GET("/usuarios/:id", controller.ObtenerPerfil, middleware.RequireAuth(repo, ""))                                 // Obtener la info de un usuario
	e.GET("/usuarios", controller.ObtenerUsuarios, middleware.RequireAuth(repo, "admin")) // Obtener la lista de usuarios
	e.DELETE("/usuarios/:id", controller.DeshabilitarUsuario, middleware.RequireAuth(repo, ""))                        // Dehabilitar el Ingreso de un Usuario
}
