package main

import (
	"log"

	"github.com/mporracindie/go-elasticsearch/v6"
	"github.com/mporracindie/go-elasticsearch/v6/esapi"
	"github.com/mporracindie/go-elasticsearch/v6/esutil"
)

func init() {
	log.SetFlags(0)
}

func main() {
	var (
		res *esapi.Response
		err error
	)

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	doc := struct {
		Title string `json:"title"`
	}{Title: "Test"}

	res, err = es.Index("test", esutil.NewJSONReader(&doc), es.Index.WithRefresh("true"))
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}

	res, err = es.Search(
		es.Search.WithIndex("test"),
		es.Search.WithBody(esutil.NewJSONReader(&query)),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)
}
