package parser

import (
	"strings"
)

func ReplaceVariables(variables map[string]string, input string) string {
	var args []string

	for key, value := range variables {
		target := "{{" + key + "}}"
		args = append(args, target, value)
	}

	// Replacer expect a list of pairs: old, new, old2, new2, ...
	replacer := strings.NewReplacer(args...)

	return replacer.Replace(input)
}
