package parser

import (
	"bufio"
	"log"
	"strings"
)

// ExtractPetitions Separate the petitions by the title (###)
// Remove the title (###)
func ExtractPetitions(content string) []string {
	var petitions []string
	var currentPetition strings.Builder // Is used to efficiently build a string using Builder.Write methods

	reader := strings.NewReader(content)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "###") {
			// Add the current petition and initialize a new one
			if currentPetition.Len() > 0 {
				petitions = append(petitions, strings.TrimSpace(currentPetition.String()))
				currentPetition.Reset()
			}

			// Don't add the title
			continue
		}

		// Add the line to the current petition
		currentPetition.WriteString(line)
		// Keep the format of the petition
		currentPetition.WriteString("\n")
	}

	if currentPetition.Len() > 0 {
		petitions = append(petitions, strings.TrimSpace(currentPetition.String()))
	}

	log.Printf("Parser: Find %d petitions", len(petitions))
	return petitions
}
