package validation

import (
	"regexp"
)

var DetecccionMicroorganismosRules = map[string]ValidationRule{
	"NombreMicroorganismo":   {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ0-9\s()\/\-]+$`), Message: "El nombre del microorganismo solo debe contener letras y espacios.", Requested: true},
	"Especificacion":         {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ0-9\s()<>\-=\/.]+$`), Message: "La especificación contiene caracteres no permitidos.", Requested: true},
	"Tratamiento":            {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ0-9\s()\/\-]+$`), Message: "El tratamiento solo debe contener letras, espacios y guiones.", Requested: true},
	"MetodoUsado":            {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚÜÑáéíóúüñ]+(-[A-Za-zÁÉÍÓÚÜÑáéíóúüñ]+)*-\d+$`), Message: "El método usado solo debe contener letras, espacios y guiones.", Requested: true},
	"CantidadMuestra":        {Regex: regexp.MustCompile(`^[0-9]+[a-zA-Z]+$`), Message: "La cantidad debe incluir una unidad válida (ml, g, mg, µl). Ej: 10 mg", Requested: true},
	"VolumenDiluyente":       {Regex: regexp.MustCompile(`^[0-9]+[a-zA-Z]+$`), Message: "El volumen debe incluir una unidad válida (ml, g, mg, µl). Ej: 90 ml", Requested: true},
	"Resultado":              {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ0-9\s()\/\-]+$`), Message: "El resultado no debe de tener caracteres especiales", Requested: false},
	"NumeroRegistroProducto": {Regex: regexp.MustCompile(`^[A-Z]{4}-\d{4}-\d{4}$`), Message: "Error en el formato del numero de registro asegurate que el formato sea: AAAA-0000-0000", Requested: true},
}
