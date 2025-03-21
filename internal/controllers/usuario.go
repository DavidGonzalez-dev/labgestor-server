package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Definicion de Constantes:
const SECRET string = "SECRET"
const PSWHASHLEVEL int = 10
var EXPTIME time.Time = time.Now().Add(time.Hour * 24)

// Interfaz que define los controladores usados como handlers en la API
type UsuarioController interface {
	RegistrarUsuario(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error 
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
		ID        string
		Nombres   string
		Apellidos string
		Correo    string
		RolID     int
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error al leer el cuerpo de la solicitud", "error": err.Error()})
	}

	//Creacion del usuario
	usuario := models.Usuario{
		ID:         requestBody.ID,
		Nombres:    requestBody.Nombres,
		Apellidos:  requestBody.Apellidos,
		Correo:     requestBody.Correo,
		Firma:      utils.GenerarFirmaUsuario(requestBody.Nombres, requestBody.Apellidos),
		Estado:     true,
		RolID:      requestBody.RolID,
	}

	if err := controller.Repo.CrearUsuario(&usuario); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error al crear el usuario", "error": err.Error()})
	}

	// Se retorna una respuesta exitosa
	return c.JSON(http.StatusOK, map[string]string{"message": "Usuario creado con exito"})
}

func (controller *usuarioController) Login(c echo.Context) error {
	// Obtener el Nombre de usuario y la contraseña
	var credenciales struct {
		ID         string
		Contrasena string
	}

	if err := c.Bind(&credenciales); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error al leer el cuerpo del request", "error": err.Error()})
	}
	// Verificar que el usuario este registrado en la base de datos
	usuario, err := controller.Repo.ObtenerUsuarioID(credenciales.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID o contrasena invalidos", "error": err.Error()})
	}

	//Comparar la contrasena enviada con la contrasena encriptada del usuario
	errContrasena := bcrypt.CompareHashAndPassword([]byte(usuario.Contrasena), []byte(credenciales.Contrasena))
	if errContrasena != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID o contrasena invalidos", "error": errContrasena.Error()})
	}

	// Generar una token JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": usuario.ID,
		"exp": EXPTIME.Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error al generar el token", "error": err.Error()})
	}

	// Enviar de vuelta la token
	c.SetCookie(&http.Cookie{Name: "sesionUsuario", Value: tokenString, Expires: EXPTIME, HttpOnly: true, Secure: false})
	return c.JSON(http.StatusOK, map[string]string{"message": "Se ha generado el token con exito"})
}

func (controller *usuarioController) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{Name: "sesionUsuario", Value: "", Expires: time.Now(), HttpOnly: true, Secure: false})
	// Delete the cookie
	return c.JSON(http.StatusOK, map[string]string{"message": "Se ha eliminado la sesion"})
}
