package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ExtractVariablesFromJSON(variables map[string]string, content []byte) {
	var rawData map[string]string

	err := json.Unmarshal(content, &rawData)
	check(err, "Error parsing JSON environment file")

	for key, value := range rawData {
		variables[key] = fmt.Sprintf("%v", strings.Trim(value, " "))
	}
}
