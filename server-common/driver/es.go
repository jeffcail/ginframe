package driver

import (
	_elastic "github.com/jeffcail/ginframe/server-common/pkg/elastic"
	"github.com/olivere/elastic"
)

func InitEs(url string) (*elastic.Client, error) {
	es, err := _elastic.NewEs(url)
	if err != nil {
		return nil, err
	}
	return es, nil
}
