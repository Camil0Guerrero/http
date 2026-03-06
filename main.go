package main

import (
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Uso: http <nombre_archivo>.http")
	}

	cnt, err := os.ReadFile(os.Args[1])
	check(err, "Error al leer el archivo")

	content := string(cnt)

	if hasVariables(content) {
		variables, trimmed := ExtractAndRemoveVariables(content)
		content = ReplaceVariables(variables, trimmed)
	}

	petitions := ExtractPetitions(content)

	var wg sync.WaitGroup

	for _, petition := range petitions {
		// 2. Indicar que hay una tarea pendiente
		wg.Add(1)

		go func(p string) {
			// 4. Indicar que la tarea terminó al salir de la función
			defer wg.Done()
			processPetition(p)
		}(petition) // Pasar la petición como argumento para evitar colisiones de memoria
	}

	// 3. Bloquear el main hasta que el contador de wg llegue a cero
	wg.Wait()
}

func hasVariables(body string) bool {
	parts := strings.SplitN(body, "\n\n", 2)

	return strings.Contains(parts[0], "@")
}

func check(e error, msg string) {
	if e != nil {
		log.Fatal(msg, ": ", e)
	}
}
