package validation

import (
	"regexp"
)

var ControlesNegativosRules = map[string]ValidationRule{
	"MedioCultivo":           {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ\s]+$`), Message: "El medio de cultivo no puede contener numeros", Requested: true},
	"Resultado":              {Requested: true},
	"NumeroRegistroProducto": {Regex: regexp.MustCompile(`^[A-Z]{4}-\d{4}-\d{4}$`), Message: "Error en el formato del numero de registro asegurate que el formato sea: AAAA-0000-0000", Requested: true},
}
