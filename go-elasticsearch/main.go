package main

import (
	"fmt"
	es "github.com/elastic/go-elasticsearch/v8"
	"log"
)

var (
	client *es.Client
)

func init() {
	var err error
	client, err = es.NewClient(es.Config{
		Addresses: []string{"http://43.154.141.194:9200"},
		Username:  "elastic",
		Password:  "changeme",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println(client.Info())
}
