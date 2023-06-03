package _elastic

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"os"
)

// NewEs new elastic
func NewEs(url string) (*elastic.Client, error) {
	var err error
	logger := log.New(os.Stdout, "gin-frame", log.LstdFlags)
	esClient, err := elastic.NewClient(elastic.SetErrorLog(logger),
		elastic.SetSniff(false), elastic.SetURL(url))
	if err != nil {
		panic(err)
	}
	do, i, err := esClient.Ping(url).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Es result with code %d and version %v\n", i, do.Version))
	return esClient, err
}
