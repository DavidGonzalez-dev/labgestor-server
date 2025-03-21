package utils

import (
	"fmt"
	"strings"
)

// Esta funcion es la encargada de generar las firmas de los usuarios
func GenerarFirmaUsuario(nombres string, apellidos string) string {

	// Se genera un slice para acceder unicamente al primer apellido
	primerApellido := strings.Split(strings.TrimSpace(apellidos), " ")[0]

	// Se retorna la firma con el formato especificado ej: David Gonzalez -> D. Gonzalez
	return fmt.Sprintf("%s. %s", string(nombres[0]), primerApellido)
}
