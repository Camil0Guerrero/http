package main

import (
	"bufio"
	"log"
	"strings"
)

// ExtractPetitions toma todo el texto del archivo y devuelve un slice
// con cada bloque de petición (método, url, body, headers) limpio de títulos ###.
func ExtractPetitions(content string) []string {
	var petitions []string
	var currentPetition strings.Builder // Builder es más eficiente que concatenar strings en bucles

	// Convertimos el string gigante en algo que podemos leer línea por línea
	reader := strings.NewReader(content)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		// 1. Detectamos el separador de bloques
		if strings.HasPrefix(line, "###") {
			// Si ya estábamos construyendo una petición, la guardamos y empezamos una nueva
			if currentPetition.Len() > 0 {
				petitions = append(petitions, strings.TrimSpace(currentPetition.String()))
				currentPetition.Reset()
			}

			// Como es un separador ### (título), NO lo añadimos a currentPetition.
			continue
		}

		// 2. Si no es un título, es parte de la petición. Lo añadimos.
		currentPetition.WriteString(line)
		currentPetition.WriteString("\n") // Restauramos el salto de línea que scanner.Text() quita
	}

	// 3. Al terminar el bucle, capturamos la última petición si existe
	if currentPetition.Len() > 0 {
		petitions = append(petitions, strings.TrimSpace(currentPetition.String()))
	}

	log.Printf("Parser: Encontradas %d peticiones limpias.", len(petitions))
	return petitions
}
