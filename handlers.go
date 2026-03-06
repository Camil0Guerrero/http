package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Get(uri string) {
	res, err := http.Get(uri)
	check(err, "Error al realizar la petición GET")

	defer res.Body.Close()

	printResponse(res)
}

type Headers struct {
	Type  string
	Value string
}

func Put(uri string, headers Headers, body string) {
	req, err := http.NewRequest("PUT", uri, convertToBuffer(body))
	check(err, "Error al crear la petición PUT")

	req.Header.Set(headers.Type, headers.Value)
	client := &http.Client{}

	res, err := client.Do(req)
	check(err, "Error al realizar la petición PUT")

	defer res.Body.Close()

	printResponse(res)
}

func convertToBuffer(body string) *bytes.Buffer {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(body), &obj)
	check(err, "Error al convertir el body a JSON")

	jsonData, err := json.Marshal(obj)
	check(err, "Error al convertir el body a JSON")

	return bytes.NewBuffer(jsonData)
}

func printResponse(res *http.Response) {
	const yellow = "\033[0;33m"
	const green = "\033[92m"
	const blue = "\033[96m"
	const colorNone = "\033[0m"

	fmt.Print("============================================== \n")
	fmt.Fprintf(os.Stdout, "%s %v %s -> %s %v %s \n", yellow, res.Request.Method, colorNone, yellow, res.Request.URL, colorNone)
	fmt.Print("============================================== \n")
	fmt.Fprintf(os.Stdout, "Status code:%s %v %s \n", green, res.StatusCode, colorNone)

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Fprintf(os.Stdout, "Body: %s%v %s \n", blue, string(bodyBytes), colorNone)
}
