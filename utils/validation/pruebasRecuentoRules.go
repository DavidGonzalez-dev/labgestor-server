package validation

import "regexp"

var PruebaRecuentoRules = map[string]ValidationRule{
	"MetodoUsado":            {Regex: regexp.MustCompile(`^[a-zA-Z0-9-]+$`), Message: "El codigo del metodo usado solo puede contener letras y numeros separados por guiones(-)", Requested: true},
	"Especificacion":         {Requested: true},
	"VolumenDiluyente":       {Regex: regexp.MustCompile(`^[0-9]+[a-zA-Z]+$`), Message: "El volumen diluyente solo puede contener letras, numeros y comas (,)", Requested: true},
	"TiempoDisolucion":       {Regex: regexp.MustCompile(`^[0-9]+[a-zA-Z]+$`), Message: "El tiempo de disolucion debe contener un número seguido de una unidad de tiempo (por ejemplo, 10min, 5h)", Requested: true},
	"CantidadMuestra":        {Regex: regexp.MustCompile(`^[0-9]+[a-zA-Z]+$`), Message: "La cantida de la muestra debe contener un número seguido de una unidad de masa (por ejemplo, 10mg, 5g)", Requested: true},
	"Tratamiento":            {Regex: regexp.MustCompile(`^[\p{L}\p{P}\s]+$`), Message: "El tratamiento solo puede contener letras y signos de puntuación", Requested: true},
	"NombreRecuento":         {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ0-9 ]+$`), Message: "El nombre de recuento solo puede contener letras", Requested: true},
	"NumeroRegistroProducto": {Regex: regexp.MustCompile(`^[A-Z]{4}-\d{4}-\d{4}$`), Message: "Error en el formato del numero de registro asegurate que el formato sea: AAAA-0000-0000", Requested: true},
}
