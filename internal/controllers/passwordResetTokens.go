package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/mail"
	"labgestor-server/utils/mail/templates"
	"labgestor-server/utils/response"
	"labgestor-server/utils/verificationCodes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Interfaz que define los metodos del controlador
type PasswordResetTokensController interface {
	SendEmailWithToken(c echo.Context) error
	VerifySendToken(c echo.Context) error
}

type passwordController struct {
	Repo        repository.PasswordResetTokenRepository
	UsuarioRepo repository.UsuarioRepository
}

// Funcion para instanciar el controlador
func NewPasswordResetTokensController(repo repository.PasswordResetTokenRepository, userRepository repository.UsuarioRepository) PasswordResetTokensController {
	return &passwordController{Repo: repo, UsuarioRepo: userRepository}
}

// Este metodo nos permite enviar un correo con el codigo de verificacion para el cambio de contraseña
func (controller *passwordController) SendEmailWithToken(c echo.Context) error {

	// Primero se lee el cuerpo del request
	var requestBody struct {
		Email     string `json:"correo"`
	}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Hubo un error al leer el cuerpo del request", Error: err.Error()})
	}

	// Segundo se valida que el usuario exista y que el correo este registrado
	usuario, err := controller.UsuarioRepo.ObtenerUsuarioCorreo(requestBody.Email)

	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Hubo un error al enviar el correo", Error: err.Error()})
	}

	// Se genera el codigo de verificacion y el token
	verificationCode := verificationCodes.GenerarCodigoVerificacion()

	finalToken := models.PasswordResetToken{
		CreatedTimestamp:    time.Now().UTC(),
		ExpirationTimestamp: time.Now().Add(time.Minute * 15).UTC(),
		Used:                false,
		IdUsuario:           usuario.ID,
	}

	// Generamos el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":           finalToken.IdUsuario,
		"verificationCode": verificationCode,
		"iat":              finalToken.CreatedTimestamp.Unix(),
		"exp":              finalToken.ExpirationTimestamp.Unix(),
		"used":             finalToken.Used,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al generar el token", Error: err.Error()})
	}

	finalToken.Token = tokenString

	// Se guarda en la Base de datos
	if err := controller.Repo.Create(&finalToken); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Hubo un error al guardar el token", Error: err.Error()})
	}

	// Se envia el correo
	mailBody := mailTemplates.GetPasswordResetMessage(verificationCode)
	if err := mail.SendEmail(requestBody.Email, "Recuperacion de Contraseña", mailBody); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Hubo un error al enviar el email", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Correo enviado con exito"})
}

// Este metodo se usa para validar que el token suministrado por el usuario sea valido con el de la base de datos y en dado caso se setea una cookie con el token de autorizacion
func (controller *passwordController) VerifySendToken(c echo.Context) error {

	// Se leen los datos enviado (codigoVerificacion, userID)
	var requestBody struct {
		UserId           string `json:"usuarioId"`
		VerificationCode string `json:"codigoVerificacion"`
	}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	// Validamos que el usuario exista
	if usuario, _ := controller.UsuarioRepo.ObtenerUsuarioID(requestBody.UserId); usuario == nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Este usuario no existe"})
	}

	// Buscamos el token mas reciente y valido para ese usuario
	userRecentToken, err := controller.Repo.GetMostRecentTokenByUserID(requestBody.UserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Hubo un error al recuperar el token", Error: err.Error()})
	}
	if userRecentToken == nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "El usuario no tiene tokens"})
	}

	// Decodificamos el token
	token, err := jwt.Parse(userRecentToken.Token, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Println("Error parsing token in VerifySendToken controller:", err)
		return c.JSON(http.StatusUnauthorized, response.Response{Message: "Token inválido", Error: err.Error()})
	}

	// Verficamos el token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al verificar el token"})
	}

	// Verificamos que el codigo de verificacion concuerde
	if claims["verificationCode"] != requestBody.VerificationCode {
		return c.JSON(http.StatusUnauthorized, response.Response{Message: "El codigo de verificacion es incorrecto"})
	}

	// Verificamos que el token no halla sido usado
	if userRecentToken.Used {
		return c.JSON(http.StatusGone, response.Response{Message: "El token ya fue usado"})
	}

	// Marcamos el token como usado
	if err := controller.Repo.MarkAsUsed(userRecentToken.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al actualizar el token"})
	}

	// Seteamos una cookie segura en el navegador para que el front pueda usar el endopoint de cambio de contraseña
	c.SetCookie(&http.Cookie{Name: "resetPasswordToken", Value: userRecentToken.Token, Expires: userRecentToken.ExpirationTimestamp, HttpOnly: true, SameSite: http.SameSiteLaxMode})

	return c.JSON(http.StatusOK, response.Response{Message: "El token es valido"})
}
