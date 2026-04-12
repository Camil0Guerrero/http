package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Get(uri string) *http.Response {
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal("Error while Executing GET request", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Error while closing response body", err)
		}
	}(res.Body)

	printResponse(res)
	return res
}

type Headers struct {
	Type  string
	Value string
}

func Put(uri string, headers Headers, body string) *http.Response {
	req, err := http.NewRequest("PUT", uri, convertToBuffer(body))
	if err != nil {
		log.Fatal("Error Creating PUT request", err)
	}

	req.Header.Set(headers.Type, headers.Value)
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error while Executing PUT request", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Error while closing response body", err)
		}
	}(res.Body)

	printResponse(res)
	return res
}

func convertToBuffer(body string) *bytes.Buffer {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(body), &obj)
	if err != nil {
		log.Fatal("Error Parsing JSON", err)
	}

	jsonData, err := json.Marshal(obj)
	if err != nil {
		log.Fatal("Could not convert JSON", err)
	}

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
