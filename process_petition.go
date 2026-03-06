package main

import (
	"log"
	"strings"
)

func processPetition(petition string) {
	split := strings.Split(petition, "\n")

	var method string
	var uri string
	var parts []string

	if len(split) < 2 {
		parts = strings.Split(petition, " ")
	} else {
		parts = strings.Split(split[0], " ")
	}

	method = parts[0]
	uri = parts[1]

	if method == "GET" {
		Get(uri)
		return
	}

	if method == "DELETE" {
		log.Print("Not implemented yet")
		return
	}

	headers := strings.Split(split[1], " ")
	header := Headers{
		Type:  strings.ReplaceAll(headers[0], ":", ""),
		Value: headers[1],
	}

	body := strings.Join(split[2:], " ")
	body = strings.ReplaceAll(body, "\n", "")

	if method == "PUT" {
		Put(uri, header, body)
		return
	}

	log.Print("Not implemented yet")
}
