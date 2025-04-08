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
	e.POST("/login", controller.Login) // Iniciar Sesion de Usuario 
	e.PUT("/contrasena/actualizar", controller.CambiarContrasena) // Actualizar la contraseña
	
	// -----------------------------------
	// Rutas Privadas Autenticacion
	// -----------------------------------
	e.POST("/logout", controller.Logout) // Cerrar Sesion

	// -----------------------------------
	// Rutas CRUD
	// -----------------------------------
	e.POST("/usuarios/registrar", controller.RegistrarUsuario) // Registrar Usuario en el sistema
	e.GET("/usuarios/:id", controller.ObtenerPerfil, middleware.RequireAuth(repo, "admin")) // Obtener la info de un usuario
	e.DELETE("/usuarios/:id", controller.DeshabilitarUsuario, middleware.RequireAuth(repo, "admin")) // Dehabilitar el Ingreso de un Usuario
}	

