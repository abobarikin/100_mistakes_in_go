package main

import (
	"log"
	"net/http"
)

func main() {
	createClient()
}

func createClient() error {
	var client *http.Client
	var err error
	tracing := true
	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}
	if err != nil {
		return err
	}
	log.Println(client)

	return nil
}

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}