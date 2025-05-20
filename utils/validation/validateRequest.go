package validation

import (
	"errors"
	"fmt"
	"regexp"
)

type ValidationRule struct {
	Regex     *regexp.Regexp
	Message   string
	Requested bool
}

func Validate(data map[string]any, validationRules map[string]ValidationRule) error {
	
	// Recorremos las reglas de validación
	for field, rule := range validationRules {
		// Verificamos si el campo existe en los datos
		value, exists := data[field]
		
		// Si el campo no existe o es nil
		if !exists || value == nil {
			// Si es obligatorio, retornamos error
			if rule.Requested {
				return fmt.Errorf("el campo %s es obligatorio", field)
			}
			// Si no es obligatorio, continuamos con el siguiente campo
			continue
		}
		
		// Verificamos que el dato sea una string
		strValue, ok := value.(string)
		if !ok {
			return fmt.Errorf("el campo %s debe ser un string", field)
		}
		
		// Verificamos si el campo está vacío
		if strValue == "" {
			// Si es obligatorio, retornamos error
			if rule.Requested {
				return fmt.Errorf("el campo %s es obligatorio", field)
			}
			// Si no es obligatorio, continuamos con el siguiente campo
			continue
		}
		
		// Si el campo tiene valor y hay un regex definido, validamos
		if rule.Regex != nil {
			if !rule.Regex.MatchString(strValue) {
				return errors.New(rule.Message)
			}
		}
	}

	return nil
}
