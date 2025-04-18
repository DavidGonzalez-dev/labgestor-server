package routes

import (
	"labgestor-server/internal/controllers"
	"labgestor-server/internal/repository"
	"labgestor-server/middleware"

	"github.com/labstack/echo/v4"
)

func NewUsuarioHanlder(e *echo.Echo, controller controllers.UsuarioController, repo repository.UsuarioRepository) {
	// -----------------------------------
	// Rutas Publicas Autenticacion
	// -----------------------------------
	e.POST("/login", controller.Login)                              // Iniciar Sesion de Usuario
	e.PATCH("/contrasena/actualizar", controller.CambiarContrasena) // Actualizar la contraseña
	e.GET("/validar-token", controller.ValidarToken)				// Validar el token

	// -----------------------------------
	// Rutas Privadas Autenticacion
	// -----------------------------------
	e.POST("/logout", controller.Logout) // Cerrar Sesion

	// -----------------------------------
	// Rutas CRUD
	// -----------------------------------
	e.POST("/usuarios/registrar", controller.RegistrarUsuario)                              // Registrar Usuario en el sistema
	e.PUT("/usuarios/actualizar", controller.ActualizarUsuario)                             // Actualizar la info de un usuario
	e.GET("/usuarios/:id", controller.ObtenerPerfil, middleware.RequireAuth(repo, "admin")) // Obtener la info de un usuario
	e.GET("/usuarios", controller.ObtenerUsuarios, middleware.RequireAuth(repo, "admin"))   // Obtener la lista de usuarios
	e.DELETE("/usuarios/:id", controller.DeshabilitarUsuario)                               // Dehabilitar el Ingreso de un Usuario
}
