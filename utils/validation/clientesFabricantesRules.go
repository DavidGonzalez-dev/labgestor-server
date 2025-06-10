package validation

import (
	"regexp"
)

// Se definene las reglas de validacion para los campos de clientes y fabricantes
var ClientesFabricantesRules = map[string]ValidationRule{
	"Nombre":    {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ\s]+$`), Message: "El nombre no puede contener numeros"},
	"Direccion": {Regex: regexp.MustCompile(`^(?i)(cra|cr|calle|cl|av|avenida|transversal|tv|diag|dg|manzana|mz|circular|circ)[a-z]*\.?\s*\d+[a-zA-Z]?\s*(#|n°|no\.?)\s*\d+[a-zA-Z]?(?:[-]\d+)?$`), Message: "Ingrese una direccion valida"},
}
