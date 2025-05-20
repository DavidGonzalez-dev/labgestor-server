package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils"
	"labgestor-server/utils/response"
	"labgestor-server/utils/validation"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var EXPTIME time.Time = time.Now().Add((time.Hour * 24) * 1) // Configuracion de duracion del token

// Interfaz que define los controladores usados como handlers en la API
type UsuarioController interface {
	// Metodos CRUD
	ObtenerUsuarios(c echo.Context) error
	ActualizarUsuario(c echo.Context) error

	// Metodos de autenticacion
	RegistrarUsuario(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CambiarContrasena(c echo.Context) error
	ValidarToken(c echo.Context) error

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

// ? --------------------------------------------------
// ? Definicion de los Controladores para Autenticacion
// ? --------------------------------------------------
// Este handler se usa para registrar un usuario en el sistema
func (controller *usuarioController) RegistrarUsuario(c echo.Context) error {

	// ? ---------------------------------------------------------------
	// ? Se lee el cuerpo del request y se verfica que no hallan errores
	// ? ---------------------------------------------------------------
	// Obtener la informacion presente en el cuerpo del request
	var requestBody struct {
		ID        string `json:"documento"`
		Nombres   string `json:"nombres"`
		Apellidos string `json:"apellidos"`
		Correo    string `json:"correo"`
		RolID     int    `json:"rol"`
	}
	// En caso de error al leeer el cuerpo del request se devuelve un error
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	if usuario, _ := controller.Repo.ObtenerUsuarioID(requestBody.ID); usuario != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al crear el usuario", Error: "El usuario ya existe"})
	}

	// ? ---------------------------------------------------------------
	// ? Se instancia el modelo y se validan los campos
	// ? ---------------------------------------------------------------
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

	// Se definen las reglas de validacion
	validationRules := map[string]validation.ValidationRule{
		"ID":        {Regex: regexp.MustCompile(`^\d+$`), Message: "El id de usuario solo puede contener numeros"},
		"Nombres":   {Regex: regexp.MustCompile(`^[a-zA-Z\s]+$`), Message: "Los nombres solo pueden contener letras y espacios"},
		"Apellidos": {Regex: regexp.MustCompile(`^[a-zA-Z\s]+$`), Message: "Los apellidos solo pueden contener letras y espacios"},
		"Correo":    {Regex: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`), Message: "Ingrese una direccion de correo valida ej: ejemplo@gmail.com"},
	}

	// Se valida que los campos dados si cumplan con las reglas de validacion definidas
	if err := validation.Validate(usuario.ToMap(), validationRules); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	// Se valida que se halla pasado un rol valido para el usuario
	if usuario.RolID == 0 {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Rol invalido", Error: "El rol que esta intentando asignar no existe"})
	}

	// ? ---------------------------------------------------------------
	// ? Se crea el registro del usuario en la base de datos
	// ? ---------------------------------------------------------------
	// Se crea el usuario y se verifica que no hallan habido errores
	if err := controller.Repo.CrearUsuario(&usuario); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al crear el usuario", Error: err.Error()})
	}
	// Se retorna una respuesta exitosa
	return c.JSON(http.StatusCreated, response.Response{Message: "El usuario ha sido registrado con exito"})
}

// Este handler se usa para iniciar sesion con un token JWT, el token se guarda en una cookie segura el navegador
func (controller *usuarioController) Login(c echo.Context) error {

	// ? --------------------------------------------------
	// ? Leemos el cuerpo del request
	// ? --------------------------------------------------

	// Obtener el Nombre de usuario y la contraseña
	var credenciales struct {
		ID         string `json:"documento"`
		Contrasena string `json:"password"`
	}
	// Verificamos qyue no halla habido un error al momento de leer el cuerpo de request
	if err := c.Bind(&credenciales); err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	// ? --------------------------------------------------
	// ? Verificamos que el usuario exista
	// ? --------------------------------------------------

	// Verificar que el usuario este registrado en la base de datos
	usuario, err := controller.Repo.ObtenerUsuarioID(credenciales.ID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.Response{Message: "Acceso denegado", Error: "Ups! su id o contraseña son incorrectos, vuelva a intentarlo"})
	}

	// ? --------------------------------------------------
	// ? Hacemos las validaciones para el ingreso
	// ? --------------------------------------------------

	// Verificar que el usuario este habilitado para el ingreso
	if !usuario.Estado {
		return c.JSON(http.StatusUnauthorized, response.Response{Message: "Acceso Denegado", Error: "Este usuario ha sido deshabilitado por el administrador del sistema. Para poder ingresar de nuevo al sistema comuniquese con su empleador"})
	}

	//Comparar la contrasena enviada con la contrasena encriptada del usuario
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Contrasena), []byte(credenciales.Contrasena))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.Response{Message: "Acceso denegado", Error: "Ups! su id o contraseña son incorrectos, vuelva a intentarlo"})
	}

	// ? --------------------------------------------------
	// ? Generamos el JWT
	// ? --------------------------------------------------

	// Generar el JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": usuario.ID,
		"exp":    EXPTIME.Unix(),
		"rol":    usuario.Rol.NombreRol,
	})

	// Se codifica el token y se firma usando el SECRET
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al generar el token", Error: err.Error()})
	}

	// ? ---------------------------------------------------------
	// ? Se guarda el token en una cookie segura en el navegador
	// ? ---------------------------------------------------------

	// Enviar de vuelta la token
	c.SetCookie(&http.Cookie{Name: "authToken", Value: tokenString, Expires: EXPTIME, HttpOnly: true, Secure: false, SameSite: http.SameSiteLaxMode, Path: "/"})
	return c.JSON(http.StatusOK, response.Response{Message: "Se ha generado el token con exito", Data: map[string]string{"id": usuario.ID, "rol": usuario.Rol.NombreRol}})
}

// Este handler se usa para cerrar la sesion del usuario e inhabilitar el token JWT generado
func (controller *usuarioController) Logout(c echo.Context) error {

	// ? ---------------------------------------------------------
	// ? Eliminamos la cookie
	// ? ---------------------------------------------------------
	// Actualizamos la fecha de expiracion de la cookie para que expire ahora.
	c.SetCookie(&http.Cookie{Name: "authToken", Value: "", Expires: time.Now(), HttpOnly: true, Secure: false})
	return c.JSON(http.StatusOK, response.Response{Message: "Se ha cerrado la sesion con exito"})
}

// Este handler valida que el token sea valido y retorna una respuesta al usuario
func (controller *usuarioController) ValidarToken(c echo.Context) error {

	// ? ---------------------------------------------------------
	// ? Obtenemos las cookies del navegador
	// ? ---------------------------------------------------------

	// Obtenemos el token de autorizacion para inicio de sesion
	tokenCookie, err := c.Cookie("authToken")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.Response{Message: "No se encontro el token", Data: map[string]bool{"valid": false}})
	}

	// ? ---------------------------------------------------------
	// ? Se hace el decode del token
	// ? ---------------------------------------------------------
	// Hacemos el decode del token
	tokenString := tokenCookie.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	// Se verifica que el token sea valido
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.Response{Message: "Token invalido", Error: err.Error(), Data: map[string]bool{"valid": false}})
	}

	// ? -----------------------------------------------------------
	// ? Hacemos la validacion de la fecha de expiracion del token
	// ? -----------------------------------------------------------

	//Verificar el token
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		// Validamos la expiracion del token
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			return c.JSON(http.StatusUnauthorized, response.Response{Message: "Token Expirado", Data: map[string]bool{"valid": false}})
		}
	}
	// Si se llega este punto el token es valido
	return c.JSON(http.StatusOK, response.Response{Message: "El token es valido", Data: map[string]any{"valid": true, "rol": claims["rol"]}})
}

// Este handler se usa para actualizar la contraseña de un usuario ya creado. La contraseña se encripta antes de hacer el update en la base de datos.
func (controller *usuarioController) CambiarContrasena(c echo.Context) error {

	// ? ---------------------------------------------------------
	// ? Leemos el cuerpo del request
	// ? ---------------------------------------------------------
	// Obtenemos el cuerpo del request
	var requestBody struct {
		ID         string `json:"id"`
		Contrasena string `json:"contrasena"`
	}
	// Buscamos errores al momento de leer el cuerpo del request
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	// ? ---------------------------------------------------------
	// ? Verificamos que el ususario exista
	// ? ---------------------------------------------------------

	// Obtenemos el usuario y verificamos que exista
	usuario, err := controller.Repo.ObtenerUsuarioID(requestBody.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Error al actualizar la contrasena", Error: err.Error()})
	}

	// ? ---------------------------------------------------------
	// ? Se actualiza la contraseña en la base de datos
	// ? ---------------------------------------------------------
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

// ? --------------------------------------------------
// ? Definicion de los Controladores CRUD
// ? --------------------------------------------------
// Este handler retorna un objeto JSON con la informacion del usuario cuyo ID es pasado como parametro
func (controller *usuarioController) ObtenerPerfil(c echo.Context) error {

	// ? ---------------------------------------------------------
	// ? Se obtiene el id pasado por parametro
	// ? ---------------------------------------------------------
	id := c.Param("id")
	// Verificamos que se halla pasado una id valida
	if id == "" {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Id invalida"})
	}

	// ? ---------------------------------------------------------
	// ? Se busca el registro del usuario en base al id
	// ? ---------------------------------------------------------

	// Buscamos al usuario
	usuario, err := controller.Repo.ObtenerUsuarioID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Error al obtener el usuario", Error: err.Error()})
	}

	// Retornar el usuario
	return c.JSON(http.StatusOK, usuario)
}

// Este handler permite hacer un borrado logico del usuario, lo inhabilita
func (controller *usuarioController) DeshabilitarUsuario(c echo.Context) error {

	// ? ---------------------------------------------------------
	// ? Obtenemos al usuario con el id pasado como parametro
	// ? ---------------------------------------------------------

	id := c.Param("id")
	// Obtenemos el usuario
	usuario, err := controller.Repo.ObtenerUsuarioID(id)
	// Verificamos que el usuario exista
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Usuario no encontrado", Error: err.Error()})
	}

	// ? --------------------------------------------------------------------
	// ? Actualizamos el estado del usuario y actualizamos la base de datos
	// ? --------------------------------------------------------------------

	// Cambiamos el estado del usuario a false
	usuario.Estado = false
	controller.Repo.ActualizarUsuario(usuario)

	// Retornamos una respuesta correcta
	return c.JSON(http.StatusOK, response.Response{Message: "El usuario ha sido deshabilitado con exito"})

}

// Este hanldler permite obtener todos los registros de los usuarios
func (controller *usuarioController) ObtenerUsuarios(c echo.Context) error {

	// ? ------------------------------------------------------------------------
	// ? Obtenemos los registros de los usuarios en base a la cada de repositorio
	// ? ------------------------------------------------------------------------

	// Llamamos al repositorio para obtener todos los usuarios
	usuarios, err := controller.Repo.ObtenerUsuarios()
	// Verificamos que no halla ningun error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "No se pudo obtener los usuarios", Error: err.Error()})
	}

	// Si todo salió bien, respondemos con un estado 200 y los usuarios
	return c.JSON(http.StatusOK, response.Response{Data: usuarios})

}

// Este handler se encarga de actualizar un usuario siempre y cuando el usuario exista
func (controller *usuarioController) ActualizarUsuario(c echo.Context) error {

	// ? --------------------------------------------------------------------
	// ? Se lee el cuerpo del request
	// ? --------------------------------------------------------------------

	// Se lee el cuerpo del request
	var requestBody struct {
		ID        string `json:"id"`
		Nombres   string `json:"nombres"`
		Apellidos string `json:"apellidos"`
		Correo    string `json:"correo"`
		Estado    bool   `json:"estado"`
		RolID     int    `json:"rolId"`
	}
	// Se verifica que no hallan habido errores al leer el cuerpo del request
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo de request", Error: err.Error()})
	}

	// ? --------------------------------------------------------------------
	// ? Se verifica que el usuario exista
	// ? --------------------------------------------------------------------

	// Obtenemos el usuario de la base de datos
	usuario, err := controller.Repo.ObtenerUsuarioID(requestBody.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "No se pudo actualizar el usuario", Error: err.Error()})
	}

	// ? --------------------------------------------------------------------
	// ? Se actualiza la informacion del usuario y se validan los campos
	// ? --------------------------------------------------------------------
	usuario.Nombres = requestBody.Nombres
	usuario.Apellidos = requestBody.Apellidos
	usuario.Correo = requestBody.Correo
	usuario.Firma = utils.GenerarFirmaUsuario(requestBody.Nombres, requestBody.Apellidos)
	usuario.Estado = requestBody.Estado
	usuario.RolID = requestBody.RolID

	// Se crean las reglas de validacion
	validationRules := map[string]validation.ValidationRule{
		"Nombres":   {Regex: regexp.MustCompile(`^[a-zA-Z\s]+$`), Message: "El campo 'Nombres' solo puede contener letras y espacios (No puede estar vacio)"},
		"Apellidos": {Regex: regexp.MustCompile(`^[a-zA-Z\s]+$`), Message: "El campo 'Apellidos' solo puede contener letras y espacios (No puede estar vacio)"},
		"Correo":    {Regex: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`), Message: "El campo 'Correo' no es valido"},
	}
	// Se verifica que los campos cumplan con las reglas de validacion
	if err := validation.Validate(usuario.ToMap(), validationRules); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	// ? --------------------------------------------------------------------
	// ? Se actualiza el reguistro del usuario en la base de datos
	// ? --------------------------------------------------------------------
	// Se actualiza el usuario
	if err := controller.Repo.ActualizarUsuario(usuario); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "No se pudo actualizar el usuario", Error: err.Error()})
	}
	// Si todo salio bien se retorna una respuesta positiva
	return c.JSON(http.StatusOK, response.Response{Message: "Se actualizo el usuario con exito"})
}
