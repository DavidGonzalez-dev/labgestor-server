package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/mail"
	"labgestor-server/utils/mail/templates"
	"labgestor-server/utils/response"
	"labgestor-server/utils/verificationCodes"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Interfaz que define los metodos del controlador
type PasswordResetTokensController interface {
	SendEmailWithToken(c echo.Context) error
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
		UsuarioId string `json:"usuarioId"`
		Email     string `json:"correo"`
	}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Hubo un error al leer el cuerpo del request", Error: err.Error()})
	}

	// Segundo se valida que el usuario exista y que el correo este registrado
	usuario, err := controller.UsuarioRepo.ObtenerUsuarioID(requestBody.UsuarioId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error de parte del servidor al obtener el usuario", Error: err.Error()})
	} else if usuario == nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "El usuario no existe"})
	}
	if requestBody.Email != usuario.Correo {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El correo subministrado no corresponde a ningun usuario"})
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
		"created":          finalToken.CreatedTimestamp,
		"expires":          finalToken.ExpirationTimestamp,
		"verificationCode": verificationCode,
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
