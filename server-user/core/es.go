package core

import "github.com/olivere/elastic"

var (
	Es *elastic.Client
)

func SetEsClient(_es *elastic.Client) {
	Es = _es
}
