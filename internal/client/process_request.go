package client

import (
	"errors"
	"strings"
)

type Request struct {
	Method     string
	URI        string
	Body       string
	StatusCode int
}

func ProcessRequest(request string) (data Request, err error) {
	split := strings.Split(request, "\n")

	var method string
	var uri string
	var parts []string

	if len(split) < 2 {
		parts = strings.Split(request, " ")
	} else {
		parts = strings.Split(split[0], " ")
	}

	method = parts[0]
	uri = parts[1]

	if method == "GET" {
		return Get(uri), nil
	}

	if method == "DELETE" {
		return Request{}, errors.New("not implemented yet")
	}

	headers := strings.Split(split[1], " ")
	header := Headers{
		Type:  strings.ReplaceAll(headers[0], ":", ""),
		Value: headers[1],
	}

	body := strings.Join(split[2:], " ")
	body = strings.ReplaceAll(body, "\n", "")

	if method == "PUT" {
		return Put(uri, header, body), nil
	}

	return Request{}, errors.New("method not supported")
}
