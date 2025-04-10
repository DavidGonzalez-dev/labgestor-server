package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils"
	"labgestor-server/utils/response"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var EXPTIME time.Time = time.Now().Add((time.Hour * 24) * 1)

// Interfaz que define los controladores usados como handlers en la API
type UsuarioController interface {
	// Metodos CRUD
	ObtenerUsuarios(c echo.Context) error

	// Metodos de autenticacion
	RegistrarUsuario(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CambiarContrasena(c echo.Context) error

	// Metodos Modulo de Usuarios
	ObtenerPerfil(c echo.Context) error
	DeshabilitarUsuario(c echo.Context) error
}

// Structura que implementa a la interfaz definida arriba
type usuarioController struct {
	Repo repository.UsuarioRepository
}

// Funcion para instanciar la estructura usuarioController y acceder a los controladores del usuario
func NewUsuarioController(repo repository.UsuarioRepository) UsuarioController {
	return &usuarioController{Repo: repo}
}

// --------------------------------------------------
// Definicion de los Controladores para Autenticacion
// --------------------------------------------------
// Este handler se usa para registrar un usuario en el sistema
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
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	//Creacion del usuario
	usuario := models.Usuario{
		ID:        requestBody.ID,
		Nombres:   requestBody.Nombres,
		Apellidos: requestBody.Apellidos,
		Correo:    requestBody.Correo,
		Firma:     utils.GenerarFirmaUsuario(requestBody.Nombres, requestBody.Apellidos),
		Estado:    true,
		RolID:     requestBody.RolID,
	}
	controller.Repo.CrearUsuario(&usuario)

	// Se retorna una respuesta exitosa
	return c.JSON(http.StatusOK, response.Response{Message: "El usuario ha sido registrado con exito"})
}

// Este handler se usa para iniciar sesion con un token JWT, el token se guarda en una cookie segura el navegador
func (controller *usuarioController) Login(c echo.Context) error {
	// Obtener el Nombre de usuario y la contraseña
	var credenciales struct {
		ID         string
		Contrasena string
	}
	if err := c.Bind(&credenciales); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	// Verificar que el usuario este registrado en la base de datos
	usuario, err := controller.Repo.ObtenerUsuarioID(credenciales.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al hacer la validacion del usuario", Error: err.Error()})
	}
	if usuario.ID == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Id o contrasena Invalidos"})
	}

	//Comparar la contrasena enviada con la contrasena encriptada del usuario
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Contrasena), []byte(credenciales.Contrasena))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID o Contrasena Invalidos", Error: err.Error()})
	}

	// Verificar que el usuario este habilitado para el ingreso
	if !usuario.Estado {
		return c.JSON(http.StatusUnauthorized, response.Response{Message: "Acceso Denegado"})
	}

	// Generar una token JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": usuario.ID,
		"exp":    EXPTIME.Unix(),
		"rol":    usuario.Rol.NombreRol,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Errror al generar el token", Error: err.Error()})
	}

	// Enviar de vuelta la token
	c.SetCookie(&http.Cookie{Name: "sesionUsuario", Value: tokenString, Expires: EXPTIME, HttpOnly: true, Secure: false})
	return c.JSON(http.StatusInternalServerError, response.Response{Message: "Se ha generado el token con exito"})
}

// Este handler se usa para cerrar la sesion del usuario e inhabilitar el token JWT generado
func (controller *usuarioController) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{Name: "sesionUsuario", Value: "", Expires: time.Now(), HttpOnly: true, Secure: false})
	// Delete the cookie
	return c.JSON(http.StatusOK, response.Response{Message: "Se ha cerrado la sesion con exito"})
}

// Este handler se usa para actualizar la contraseña de un usuario ya creado. La contraseña se encripta antes de hacer el update en la base de datos.
func (controller *usuarioController) CambiarContrasena(c echo.Context) error {

	// Obtenemos el cuerpo del request
	var requestBody struct {
		ID         string
		Contrasena string
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	// Obtenemos el usuario y verificamos que exista
	usuario, err := controller.Repo.ObtenerUsuarioID(requestBody.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al cambiar la contraseña", Error: err.Error()})
	}
	if usuario.ID == "0" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Usuario no encontrado"})
	}

	// Hasheamos la contraseña
	passwordLevel, _ := strconv.Atoi(os.Getenv("PSWHASHLEVEL"))
	hash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Contrasena), passwordLevel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al encriptar la contraseña", Error: err.Error()})
	}

	//Actualizamos la informacion del usuario
	usuario.Contrasena = string(hash)
	controller.Repo.ActualizarUsuario(usuario)

	return c.JSON(http.StatusOK, response.Response{Message: "Se actualizo la contrasena con exito"})
}

// --------------------------------------------------
// Definicion de los Controladores CRUD
// --------------------------------------------------
// Este handler retorna un objeto JSON con la informacion del usuario cuyo ID es pasado como parametro
func (controller *usuarioController) ObtenerPerfil(c echo.Context) error {
	id := c.Param("id")

	// Verificamos que se halla pasado una id valida
	if id == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Id invalida"})
	}

	// Buscamos al usuario
	usuario, err := controller.Repo.ObtenerUsuarioID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al obtener el usuario", Error: err.Error()})
	}
	if usuario.ID == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Usuario no econtrado"})
	}

	// Retornar el usuario
	return c.JSON(http.StatusOK, usuario)
}

// Este handler permite hacer un borrado logico del usuario, lo inhabilita
func (controller *usuarioController) DeshabilitarUsuario(c echo.Context) error {
	// Obtenemos el usuario segun el parametro id
	id := c.Param("id")
	usuario, err := controller.Repo.ObtenerUsuarioID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al deshabilitar el usuario", Error: err.Error()})
	}
	if usuario.ID == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Usuario no encontrado"})
	}

	// Cambiamos el estado del usuario a false
	usuario.Estado = false
	controller.Repo.ActualizarUsuario(usuario)

	// Retornamos una respuesta correcta
	return c.JSON(http.StatusOK, response.Response{Message: "El usuario ha sido deshabilitado con exito"})

}

func (controller *usuarioController) ObtenerUsuarios(c echo.Context) error {
	// Llamamos al repositorio para obtener todos los usuarios
	usuarios, err := controller.Repo.ObtenerUsuarios()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "No se pudo obtener los usuarios", "error": err.Error()})
	}

	// Si todo salió bien, respondemos con un estado 200 y los usuarios
	return c.JSON(http.StatusOK, usuarios)

}
