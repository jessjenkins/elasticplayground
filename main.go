package main

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	fmt.Println("Playing with elasticsearch")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	log.Println(es.Info())




}
