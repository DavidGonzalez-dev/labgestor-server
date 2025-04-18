package validation

type ValidationRule struct {
	Regex string
	Message string
}

func Validate(data map[string]any, validationRules []ValidationRule) error{
	// TODO: Validar si el usuario ingreso algun valor

	// TODO: Recorre las reglas de validacion y se verifica si los campos cumplen

	return nil
}