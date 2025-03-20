package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"net/http"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Interfaz que define los controladores usados como handlers en la API
type UsuarioController interface {
	RegistrarUsuario(c echo.Context) error
}

// Structura que implementa a la interfaz definida arriba
type usuarioController struct {
	Repo repository.UsuarioRepository
}

// Funcion para instanciar la estructura usuarioController y acceder a los controladores del usuario
func NewUsuarioController(repo repository.UsuarioRepository) UsuarioController {
	return &usuarioController{Repo: repo}
}

// -------------------------------
// Definicion de los Controladores
// -------------------------------
func (controller *usuarioController) RegistrarUsuario(c echo.Context) error {

	// Obtener la informacion presente en el cuerpo del request
	var requestBody struct {
		ID         string
		Nombres    string
		Apellidos  string
		Correo     string
		Contrasena string
		Firma      string
		Estado     bool
		RolID      int
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}


	// Encriptar la contrasena
	hash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Contrasena), 10)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":"Error al encriptar la contrasena"})
	}

	// Generacion de la firma

	//Creacion del usuario
	usuario := models.Usuario{
		ID:         requestBody.ID,
		Nombres:    requestBody.Nombres,
		Apellidos:  requestBody.Apellidos,
		Correo:     requestBody.Correo,
		Contrasena: string(hash),
		Firma:      requestBody.Firma,
		Estado:     requestBody.Estado,
		RolID:      requestBody.RolID,
	}

	if controller.Repo.CrearUsuario(&usuario) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":"Error al crear el usuario"})
	}

	// Se retorna una respuesta exitosa
	return c.JSON(http.StatusOK, map[string]string{"message":"Usuario creado con exito"})
}


