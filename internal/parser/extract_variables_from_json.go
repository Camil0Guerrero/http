package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func ExtractVariablesFromJSON(variables map[string]string, content []byte) {
	var rawData map[string]string

	err := json.Unmarshal(content, &rawData)
	if err != nil {
		log.Fatal("Error parsing JSON environment file", err)
	}

	for key, value := range rawData {
		variables[key] = fmt.Sprintf("%v", strings.Trim(value, " "))
	}
}
