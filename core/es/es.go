package es

import (
	_elastic "github.com/jeffcail/ginframe/pkg/elastic"
	"github.com/olivere/elastic"
)

var (
	EsClient *elastic.Client
	err      error
)

func InitEs() {
	EsClient, err = _elastic.NewEs()
	if err != nil {
		panic(err)
	}
}
