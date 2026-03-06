package main

import (
	"regexp"
	"strings"
)

func ReplaceVariables(variables map[string]string, input string) string {
	// {{       -> Matches literal opening braces
	// ([^{}]+) -> Captures one or more characters that are NOT braces (variable name)
	// }}       -> Matches literal closing braces
	re := regexp.MustCompile(`{{([^{}]+)}}`)

	result := re.ReplaceAllStringFunc(input, func(match string) string {
		// match will look like "{{api_url}}"
		// Trim braces and spaces to get only the key "api_url"
		key := strings.Trim(match, "{} ")

		// Search in the map. If not found, return the original placeholder
		if val, ok := variables[key]; ok {
			return val
		}
		return match // Variable not found, return as is.{{api_url}}
	})

	return result
}
