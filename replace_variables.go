package main

import (
	"regexp"
	"strings"
)

// ToDo: clarify the logic
func ReplaceVariables(variables map[string]string, input string) string {
	// {{          -> Coincide con las llaves de apertura literal
	// ([^{}]+)    -> Captura uno o más caracteres que NO sean llaves (el nombre de la variable)
	// }}          -> Coincide con las llaves de cierre literal
	re := regexp.MustCompile(`{{([^{}]+)}}`)

	result := re.ReplaceAllStringFunc(input, func(match string) string {
		// match será algo como "{{api_url}}"
		// Limpiamos las llaves para obtener solo la clave "api_url"
		key := strings.Trim(match, "{} ")

		// Buscamos en el mapa. Si no existe, dejamos el original o un string vacío
		if val, ok := variables[key]; ok {
			return val
		}
		return match // Si no la encuentra, devuelve {{variable}} tal cual
	})

	return result
}
