package parser

import "strings"

func ExtractAndRemoveVariables(variables map[string]string, content string) string {
	// Separate the header from the content
	parts := strings.SplitN(content, "\n\n", 2)

	for variable := range strings.SplitSeq(parts[0], "@") {
		split := strings.Split(variable, "=")

		if len(split) != 2 {
			continue
		}

		contentSerialized := strings.Split(split[1], " ")

		// Example: api_url: https://api.example.com
		variables[strings.Trim(split[0], " ")] = strings.Trim(contentSerialized[1], "\n")
	}

	return strings.Join(parts[1:], "")
}
