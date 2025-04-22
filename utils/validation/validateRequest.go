package validation

import (
	"errors"
	"fmt"
	"regexp"
)

type ValidationRule struct {
	Regex   *regexp.Regexp
	Message string
}

func Validate(data map[string]any, validationRules map[string]ValidationRule) error {
	
	//Recorremos las reglas de validacion
	for field, rule := range validationRules{
		
		// Verificamos que el dato pasado es una string
		value, ok := data[field].(string)
		if !ok {
			return fmt.Errorf("el campo %s debe ser un string", field)
		}

		// Verificamos que el campo exista en los datos y que cumpla con las expresiones regulares
		if data[field] == "" {
			return fmt.Errorf("el campo %s es obligatorio", field)
		} else if !rule.Regex.MatchString(value){
			return errors.New(rule.Message)
		}
	}
	
	return nil
}