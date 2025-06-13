package mail

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

// Funcion para enviar Email
func SendEmail(recepient, subject, body string) error {

	// Obtener las variables de entorno
	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")
	fromEmail := os.Getenv("MAILGUN_FROM_EMAIL")

	// Verificamos que la informacion sea valida
	if domain == "" || apiKey == "" || fromEmail == "" {
		return fmt.Errorf("faltan configuracion de MailGun revisa las variables de entorno")
	}

	// Instanciamos cliente de Mailgun
	mgClient := mailgun.NewMailgun(domain, apiKey)

	// Creamos el mensaje
	message := mailgun.NewMessage(
		fromEmail,
		subject,
		"",
		recepient,
	)
	// Se setea el html en el cuerpo del correo
	message.SetHTML(body)

	// Establecemos tiempo limite para el envio del mensaje
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// Enviamos el mensahe
	resp, id, err := mgClient.Send(ctx, message)
	if err != nil {
		return err
	}
	fmt.Printf("Email enviado exitosamente, id del mensaje: %s respuesta del servidor: %s", id, resp)
	return nil
}