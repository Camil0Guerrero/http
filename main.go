package main

import (
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./http archive.http")
	}

	cnt, err := os.ReadFile(os.Args[1])
	check(err, "Failed to read file")

	variables := make(map[string]string)

	if len(os.Args) >= 2 {
		envCnt, err := os.ReadFile(os.Args[2])
		check(err, "Failed to read environment file")

		ExtractVariablesFromJSON(variables, envCnt)
	}

	content := string(cnt)

	if hasVariables(content) {
		content = ExtractAndRemoveVariables(variables, content)
	}

	if len(variables) > 0 {
		content = ReplaceVariables(variables, content)
	}

	petitions := ExtractPetitions(content)

	var wg sync.WaitGroup

	for _, petition := range petitions {
		// Indicate a pending task
		wg.Add(1)

		go func(p string) {
			// Indicate task finished
			defer wg.Done()
			processPetition(p)
		}(petition)
	}

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
