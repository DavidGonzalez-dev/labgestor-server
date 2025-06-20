package validation

import "regexp"

var CajasBioburdenRules = map[string]ValidationRule{
	"Tipo":             {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ\s]+$`), Message: "El tipo de caja debe contener solo letras y espacios"},
	"Resultado":        {Regex: regexp.MustCompile(`^[A-Za-z0-9\s]+$`), Message: "El resultado debe contener solo letras y  números"},
	"MetodoSiembra":    {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ\s]+$`), Message: "El método de siembra debe contener solo letras y espacios"},
	"MedidaAritmetica": {Requested: true},
	"FactorDisolucion": {Regex: regexp.MustCompile(`^[\w\s\p{P}\p{S}]+$`), Message: "El factor de disolución debe estar en minutos (1:10)", Requested: true},
}
