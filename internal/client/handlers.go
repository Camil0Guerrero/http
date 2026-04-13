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

func Get(uri string) Request {
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal("Error while Executing GET request", err)
	}
	defer res.Body.Close()

	resBody := res.Body
	body, err := io.ReadAll(resBody)
	if err != nil {
		log.Fatal("Error while reading response body", err)
	}

	printResponse(res)

	return Request{
		Method:     res.Request.Method,
		URI:        res.Request.URL.String(),
		Body:       string(body),
		StatusCode: res.StatusCode,
	}
}

type Headers struct {
	Type  string
	Value string
}

func Put(uri string, headers Headers, body string) Request {
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
	defer res.Body.Close()

	printResponse(res)
	return Request{
		Method:     res.Request.Method,
		URI:        res.Request.URL.String(),
		Body:       body,
		StatusCode: res.StatusCode,
	}
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
