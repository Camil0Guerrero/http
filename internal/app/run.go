package app

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/Camil0Guerrero/http/internal/client"
	"github.com/Camil0Guerrero/http/internal/parser"
)

func Run(args []string) error {
	if len(args) < 2 {
		log.Fatal("Usage: ./http archive.http")
	}

	cnt, err := os.ReadFile(args[1])
	check(err, "Failed to read file")

	variables := make(map[string]string)

	if len(args) >= 3 {
		envCnt, err := os.ReadFile(args[2])
		check(err, "Failed to read environment file")

		parser.ExtractVariablesFromJSON(variables, envCnt)
	}

	content := string(cnt)

	if hasVariables(content) {
		content = parser.ExtractAndRemoveVariables(variables, content)
	}

	if len(variables) > 0 {
		content = parser.ReplaceVariables(variables, content)
	}

	petitions := parser.ExtractPetitions(content)

	var wg sync.WaitGroup

	for _, petition := range petitions {
		// Indicate a pending task
		wg.Add(1)

		go func(p string) {
			// Indicate task finished
			defer wg.Done()
			client.ProcessRequest(p)
		}(petition)
	}

	wg.Wait()
	return nil
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
