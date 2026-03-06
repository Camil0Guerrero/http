package main

import "strings"

func ExtractAndRemoveVariables(content string) (map[string]string, string) {
	var variables = make(map[string]string)

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

	return variables, strings.Join(parts[1:], "")
}
