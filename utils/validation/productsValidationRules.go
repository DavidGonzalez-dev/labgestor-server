package validation

import "regexp"

// Se definene las reglas de validacion para los campos del producto
var ProductoRules = map[string]ValidationRule{
	"NumeroRegistro":   {Regex: regexp.MustCompile(`^[A-Z]{4}-\d{4}-\d{4}$`), Message: "Error en el formato del numero de registro asegurate que el formato sea: AAAA-0000-0000", Requested: true},
	"Nombre":           {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ\s]+$`), Message: "El nombre no puede contener numeros", Requested: true},
	"FechaFabricacion": {Regex: regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`), Message: "La fecha de fabricacion no es valida asegurese de que sea en el formato yyyy-mm-dd", Requested: true},
	"FechaVencimiento": {Regex: regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`), Message: "La fecha de vencimiento no es valida asegurese de que sea en el formato yyyy-mm-dd", Requested: true},
	"Descripcion":      {Regex: regexp.MustCompile(`^.+$`), Message: "La descripcion no puede estar vacia", Requested: true},
	"CompuestoActivo":  {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ\s]+$`), Message: "El compuesto activo no puede contener numeros", Requested: true},
	"Presentacion":     {Regex: regexp.MustCompile(`^[A-Za-zÁÉÍÓÚáéíóúÑñ\s]+$`), Message: "La presentacion no puede contener numeros", Requested: true},
	"Cantidad":         {Regex: regexp.MustCompile(`^[a-zA-Z0-9]+$`), Message: "La cantidad no puede contener caracteres especiales", Requested: true},
	"NumeroLote":       {Regex: regexp.MustCompile(`^[a-zA-Z0-9]+$`), Message: "El numero de lote no puede contener caracteres especiales", Requested: true},
	"TamanoLote":       {Regex: regexp.MustCompile(`^[a-zA-Z0-9]+$`), Message: "El tamano de lote no puede contener caracteres especiales"},
}

// Se definene las reglas de validacion para los campos de registro de entrada de productos
var RegistroEntradaRules = map[string]ValidationRule{
	"PropositoAnalisis":      {Regex: regexp.MustCompile(`^.+$`), Message: "El proposito de analisis no puede estar vacio", Requested: true},
	"CondicionesAmbientales": {Regex: regexp.MustCompile(`^.+$`), Message: "Las condiciones ambientales no pueden estar vacias", Requested: true},
	"FechaRecepcion":         {Regex: regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`), Message: "La fecha de recepcion no es valida asegurese de que sea en el formato yyyy-mm-dd", Requested: true},
	"FechaInicioAnalisis":    {Regex: regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`), Message: "La fecha de incio de analisis no es valida asegurese de que sea en el formato yyyy-mm-dd", Requested: true},
	"FechaFinalAnalisis":     {Regex: regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`), Message: "La fecha de final de analisis no es valida asegurese de que sea en el formato yyyy-mm-dd"},
	"IDUsuario":              {Regex: regexp.MustCompile(`^[0-9]+$`), Message: "El id de usuario solo puede contener numeros", Requested: true},
}
