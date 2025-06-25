package validation

import (
	"regexp"
)

var MonitoreosDeteccionesRules = map[string]ValidationRule{
	"VolumenMuestra":  {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ0-9\s]+$`), Message: "El volumen de muestra no debe de tener caracteres especiales", Requested: true},
	"NombreDiluyente": {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ\s]+$`), Message: "El nombre del diluyente no debe de tener caracteres especiales", Requested: true},
}
