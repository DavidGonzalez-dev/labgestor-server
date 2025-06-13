package validation

import "regexp"

var PruebaRecuentoRules = map[string]ValidationRule{
	"MetodoUsado":            {Regex: regexp.MustCompile(`^[\p{L}\s\-]+$`), Message: "El método usado solo puede contener letras, espacios y guiones.", Requested: true},
	"Especificacion":         {Regex: regexp.MustCompile(`^[\p{L}0-9\s<>=\/.,°%:-]+$`), Message: "La especificación contiene caracteres no válidos.", Requested: true},
	"VolumenDiluyente":       {Regex: regexp.MustCompile(`^\d+\s?(ml|mL|ML)$`), Message: "Debe contener un número seguido de 'ml', por ejemplo: 5 ml o 10mL.", Requested: true},
	"TiempoDisolucion":       {Regex: regexp.MustCompile(`^\d+\s?(minutos|min|MIN)$`), Message: "Formato de tiempo no válido. Ejemplo: '5 minutos', '60min'.", Requested: true},
	"CantidadMuestra":        {Regex: regexp.MustCompile(`^\d+(\.\d+)?\s?(mg|g|ml|mL|MG|G)$`), Message: "Debe ser un número con unidad válida: mg, g, ml. Ejemplo: 0.5g, 200mg.", Requested: true},
	"Tratamiento":            {Regex: regexp.MustCompile(`^[\p{L}\s\-]+$`), Message: "El tratamiento solo puede contener letras, espacios y guiones.", Requested: true},
	"NombreRecuento":         {Regex: regexp.MustCompile(`^[\p{L}\s]+$`), Message: "Solo se permiten letras y espacios en el nombre del recuento.", Requested: true},
	"NumeroRegistroProducto": {Regex: regexp.MustCompile(`^[A-Z]{3,4}-\d{3,4}-\d{4}$|^[A-Z]{3}-\d{3}$`), Message: "Formato inválido. Usa AAA-000 o AAAA-0000-0000.", Requested: true},
}
