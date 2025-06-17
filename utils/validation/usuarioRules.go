package validation

import (
	"regexp"
)

var UsuarioRules = map[string]ValidationRule{
	"ID":        {Regex: regexp.MustCompile(`^\d+$`), Message: "El id de usuario solo puede contener numeros"},
	"Nombres":   {Regex: regexp.MustCompile(`^[a-zA-Z\s]+$`), Message: "Los nombres solo pueden contener letras y espacios"},
	"Apellidos": {Regex: regexp.MustCompile(`^[a-zA-Z\s]+$`), Message: "Los apellidos solo pueden contener letras y espacios"},
	"Correo":    {Regex: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`), Message: "Ingrese una direccion de correo valida ej: ejemplo@gmail.com"},
}
